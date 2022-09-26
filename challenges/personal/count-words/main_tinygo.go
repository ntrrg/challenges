package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"os"
	"reflect"
	"runtime/pprof"
	"sort"
	"unsafe"
)

const (
	BUF_SIZE        = 1 << 16
	HASH_TABLE_SIZE = 1 << 16
)

func main() {
	if cpup := os.Getenv("CPUPROFILE"); cpup != "" {
		f, err := os.Create(cpup)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot create CPU profile file: %v\n", err)
			os.Exit(1)
		}

		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "cannot profile CPU usage: %v\n", err)
			os.Exit(1)
		}

		defer pprof.StopCPUProfile()
	}

	buf := make([]byte, BUF_SIZE)
	wc := NewWordCounter(HASH_TABLE_SIZE)

	if _, err := io.CopyBuffer(wc, os.Stdin, buf); err != nil {
		fmt.Fprintf(os.Stderr, "cannot copy data from stdin: %v\n", err)
		os.Exit(1)
	}

	if err := wc.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot flush word counter: %v\n", err)
		os.Exit(1)
	}

	words := wc.Freq()

	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	w := bufio.NewWriterSize(os.Stdout, 1<<16)
	defer w.Flush()

	for _, word := range words {
		fmt.Fprintln(w, word.Word, word.Count)
	}

	if memp := os.Getenv("MEMPROFILE"); memp != "" {
		f, err := os.Create(memp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot create memory profile file: %v\n", err)
			os.Exit(1)
		}

		defer f.Close()

		//if err := pprof.WriteHeapProfile(f); err != nil {
		//	fmt.Fprintf(os.Stderr, "cannot profile memory usage: %v\n", err)
		//	os.Exit(1)
		//}
	}
}

//////////////////
// Word Counter //
//////////////////

type WordCount struct {
	Word  string
	Count int
}

type WordCounter struct {
	len  int
	exp  int
	buf  []byte
	data []*WordCount
}

func NewWordCounter(size int) *WordCounter {
	return &WordCounter{
		exp:  GetPowerOf2(size),
		data: make([]*WordCount, size),
	}
}

func (wc *WordCounter) Flush() error {
	if len(wc.buf) == 0 {
		return nil
	}

	if err := wc.countWord(wc.buf); err != nil {
		return err
	}

	wc.buf = wc.buf[:0]
	return nil
}

func (wc *WordCounter) Freq() []*WordCount {
	words := make([]*WordCount, 0, wc.len)

	for _, c := range wc.data {
		if c == nil {
			continue
		}

		words = append(words, c)
	}

	return words
}

func (wc *WordCounter) Words() []string {
	words := make([]string, 0, wc.len)

	for _, c := range wc.data {
		if c == nil {
			continue
		}

		words = append(words, c.Word)
	}

	return words
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	// This is >1s slower.
	//for _, b := range p {
	//	if b == ' ' || b == '\n' {
	//		if err := wc.countWord(wc.buf); err != nil {
	//			return 0, err
	//		}

	//		wc.buf = wc.buf[:0]
	//		continue
	//	}

	//	if b >= 'A' && b <= 'Z' {
	//		b += 'a' - 'A'
	//	}

	//	wc.buf = append(wc.buf, b)
	//}

	l := len(p)
	wl := 0

	for i, b := range p {
		if b == ' ' || b == '\n' {
			word := p[i-wl : i]

			if len(wc.buf) > 0 {
				word = append(wc.buf, word...)
				wc.buf = wc.buf[:0]
			}

			if err := wc.countWord(word); err != nil {
				return 0, err
			}

			wl = 0
			continue
		}

		if b >= 'A' && b <= 'Z' {
			p[i] += 'a' - 'A'
		}

		wl++
	}

	if wl > 0 {
		wc.buf = append(wc.buf, p[l-wl:]...)
	}

	return len(p), nil
}

func (wc *WordCounter) countWord(word []byte) error {
	if len(word) == 0 {
		return nil
	}

	l := len(wc.data)
	h := MurmurV3x32(word, 0)
	i := int(h)

lookUp:
	for {
		i = LookupIndex32(h, wc.exp, i)
		c := wc.data[i]

		switch {
		case c != nil && len(c.Word) == len(word) && CompareSlice(c.Word, word):
			c.Count++
			break lookUp
		case c == nil:
			// This is actually faster than preallocating.
			wc.data[i] = &WordCount{Word: string(word), Count: 1}
			wc.len++
			break lookUp
		case wc.len == l:
			return fmt.Errorf("hash table is full")
		}
	}

	return nil
}

////////////
// Murmur //
////////////

const (
	v3x32_c1 = 0xcc9e2d51
	v3x32_c2 = 0x1b873593
	v3x32_r1 = 15
	v3x32_r2 = 13
	v3x32_m  = 5
	v3x32_n  = 0xe6546b64
)

func MurmurV3x32(data []byte, seed uint32) uint32 {
	h := seed

	l := len(data)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	for i := 0; i+4 <= l; i += 4 {
		k := *(*uint32)(unsafe.Pointer(sh.Data + uintptr(i)))
		h ^= murmurv3x32_Scramble(k)
		h = bits.RotateLeft32(h, v3x32_r2)
		h = h*v3x32_m + v3x32_n
	}

	if rem := l % 4; rem > 0 {
		var k uint32

		for j := 1; j <= rem; j++ {
			k <<= 8
			k |= uint32(data[l-j])
		}

		h ^= murmurv3x32_Scramble(k)
	}

	h ^= uint32(l)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

func murmurv3x32_Scramble(k uint32) uint32 {
	k *= v3x32_c1
	k = bits.RotateLeft32(k, v3x32_r1)
	k *= v3x32_c2
	return k
}

///////////
// Utils //
///////////

func CompareSlice(str string, s []byte) bool {
	if len(str) != len(s) {
		return false
	}

	for i := range s {
		if str[i] != s[i] {
			return false
		}
	}

	return true
}

func GetPowerOf2(x int) int {
	exp := 0

	for y := x >> 1; y != 0; y >>= 1 {
		exp++
	}

	return exp
}

// LookupIndex32 is a hash table iterator function.
//
// Mask, step and index. See https://nullprogram.com/blog/2022/08/08/.
func LookupIndex32(h uint32, exp, i int) int {
	mask := 1<<exp - 1
	step := (h >> (32 - exp)) | 1
	return (i + int(step)) & mask
}
