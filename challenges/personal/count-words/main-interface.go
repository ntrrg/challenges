package main

import (
	"bufio"
	"context"
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

	ctx := context.Background()
	ht := NewMurmurV3_32HT[WordCount](BUF_SIZE, 0)
	wc := NewWordCounter(&ht)
	buf := make([]byte, BUF_SIZE)

	if _, err := io.CopyBuffer(&wc, os.Stdin, buf); err != nil {
		fmt.Fprintf(os.Stderr, "cannot copy data from stdin: %v\n", err)
		os.Exit(1)
	}

	if err := wc.Flush(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "cannot flush word counter: %v\n", err)
		os.Exit(1)
	}

	words, errWCF := wc.Freq(ctx)
	if errWCF != nil {
		fmt.Fprintf(os.Stderr, "cannot get word frequency: %v\n", errWCF)
		os.Exit(1)
	}

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

		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "cannot profile memory usage: %v\n", err)
			os.Exit(1)
		}
	}
}

//////////////////
// Word Counter //
//////////////////

type WordCount struct {
	Word  string
	Count int
}

type WordCountService interface {
	WordFreqCounter
	WordParser
}

type WordCountFullService interface {
	WordCountService

	WordCounter
	WordLister
	WordUniqueCounter
}

type WordCounter interface {
	Total(context.Context) (int, error)
}

type WordFreqCounter interface {
	Freq(context.Context) ([]WordCount, error)
}

type WordLister interface {
	Words(context.Context) ([]string, error)
}

type WordParser interface {
	Write(context.Context, []byte) (int, error)
	Flush(context.Context) error
}

type WordUniqueCounter interface {
	Len(context.Context) (int, error)
}

type WordCounterImpl struct {
	buf []byte
	ht  HashTable[WordCount]
}

func NewWordCounter(ht HashTable[WordCount]) WordCounterImpl {
	return WordCounterImpl{
		buf: make([]byte, 0, 30),
		ht:  ht,
	}
}

func (wc *WordCounterImpl) Flush(ctx context.Context) error {
	if len(wc.buf) == 0 {
		return nil
	}

	c, ok, errHTG := wc.ht.Get(ctx, wc.buf)
	if errHTG != nil {
		return fmt.Errorf("cannot get value from hash table: %w", errHTG)
	}

	if !ok {
		c.Word = string(wc.buf)
	} else if len(c.Word) != len(wc.buf) || !CompareSlice(c.Word, wc.buf) {
		return fmt.Errorf("hash table collision: %s", string(wc.buf))
	}

	c.Count++

	if err := wc.ht.Set(ctx, wc.buf, c); err != nil {
		return fmt.Errorf("cannot set value on hash table: %w", err)
	}

	wc.buf = wc.buf[:0]

	return nil
}

func (wc WordCounterImpl) Freq(ctx context.Context) ([]WordCount, error) {
	words, errHTVL := wc.ht.Values(ctx)
	if errHTVL != nil {
		return nil, fmt.Errorf("cannot get values from hash table: %w", errHTVL)
	}

	return words, nil
}

func (wc WordCounterImpl) Words(ctx context.Context) ([]string, error) {
	l, errHTL := wc.ht.Len(ctx)
	if errHTL != nil {
		return nil, fmt.Errorf("cannot get length from hash table: %w", errHTL)
	}

	vals, errHTVL := wc.ht.Values(ctx)
	if errHTVL != nil {
		return nil, fmt.Errorf("cannot get values from hash table: %w", errHTL)
	}

	words := make([]string, 0, l)
	for _, c := range vals {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context done: %w", ctx.Err())
		default:
			words = append(words, c.Word)
		}
	}

	return words, nil
}

func (wc *WordCounterImpl) Write(ctx context.Context, p []byte) (int, error) {
	for _, b := range p {
		if b == ' ' || b == '\n' {
			if err := wc.Flush(context.Background()); err != nil {
				return 0, err
			}

			continue
		}

		// To lower case.
		if b >= 'A' && b <= 'Z' {
			b += 'a' - 'A'
		}

		wc.buf = append(wc.buf, b)
	}

	return len(p), nil
}

////////////////
// Hash table //
////////////////

type HashTableEntry[K, V any] struct {
	Key K
	Val V
}

type HashTable[K, V any] interface {
	Get(K) (V, error)
	Len() (int, error)
	Keys() ([]K, error)
	Values() ([]V, error)
	Entries() ([]HashTableEntry[K, V], error)

	Set(K, V) error
	Delete(K) error
}

type HashTableCtx[K, V any] interface {
	GetCtx(context.Context, K) (V, error)
	LenCtx(context.Context) (int, error)
	KeysCtx(context.Context) ([]K, error)
	ValuesCtx(context.Context) ([]V, error)
	EntriesCtx(context.Context) ([]HashTableEntry[K, V], error)

	SetCtx(context.Context, K, V) error
	DeleteCtx(context.Context, K) error
}

type HashType interface{~uint | ~uint32 | ~uint64}

type HashingFunc func[T any, HT HashType](context.Context, T) (HT, error)

type HashTableImpl[K, V any, HT HashType] struct {
	h  HashingFunc[K, HT]
	ht map[HT]HashTableEntry[K, V]
}

func NewHashTable[K, V any, HT HashType](size int, hf HashingFunc[K, HT]) HashTableImpl[K, V, HT] {
	return HashTableImpl[K, V, HT]{
		h:  hf,
		ht: make(map[HT]HashTableEntry[K, V], size),
	}
}

func (ht *HashTableImpl[K, V any]) Get(k K) (V, error) {
	return ht.GetCtx(context.Background(), k)
}

func (ht *HashTableImpl[T]) GetCtx(ctx context.Context, k []byte) (T, error) {
	var v T

	if err := ctx.Err(); err != nil {
		return v, err
	}

	h, errH := ht.h(ctx, k)
	if errH != nil {
		return v, fmt.Errorf("cannot hash key: %w", errH)
	}

	v, ok := ht.ht[h]
	if !ok {
		return v, fmt.Errorf("key not found (k: %v; h: %v)", k, h)
	}

	return e.Val, ok, nil
}

func (ht HashTableImpl[T]) Len(_ context.Context) (int, error) {
	return len(ht.ht), nil
}

func (ht HashTableImpl[T]) Keys(ctx context.Context) ([][]byte, error) {
	keys := make([][]byte, 0, len(ht.ht))
	for _, e := range ht.ht {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context done: %w", ctx.Err())
		default:
			keys = append(keys, e.Key)
		}
	}

	return keys, nil
}

func (ht HashTableImpl[T]) Entries(ctx context.Context) ([]HashTableEntry[T], error) {
	entries := make([]HashTableEntry[T], 0, len(ht.ht))
	for _, e := range ht.ht {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context done: %w", ctx.Err())
		default:
			entries = append(entries, e)
		}
	}

	return entries, nil
}

func (ht HashTableImpl[T]) Values(ctx context.Context) ([]T, error) {
	vals := make([]T, 0, len(ht.ht))
	for _, e := range ht.ht {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context done: %w", ctx.Err())
		default:
			vals = append(vals, e.Val)
		}
	}

	return vals, nil
}

func (ht *HashTableImpl[T]) Set(_ context.Context, k []byte, v T) error {
	nk := make([]byte, 0, len(k))
	copy(nk, k)

	h := MurmurV3x32(k, ht.seed)
	ht.ht[h] = HashTableEntry[T]{Key: nk, Val: v}

	return nil
}

func (ht *HashTableImpl[T]) Delete(_ context.Context, k []byte) error {
	h := MurmurV3x32(k, ht.seed)
	delete(ht.ht, h)
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
