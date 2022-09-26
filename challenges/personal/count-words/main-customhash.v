module main

import math.bits
import os

fn main() {
	stdin := os.stdin()
	mut buf := []u8{len: 1 << 16}
	mut wc := new_word_counter()

	for {
		n := stdin.read(mut buf) or {
			match err {
				none {
					break
				}
				else {
					eprintln('cannot read input: $err')
					exit(1)
				}
			}
		}

		wc.write(mut buf[..n])?
	}

	wc.flush()?

	mut words := wc.freq()
	words.sort(b.count < a.count)

	for c in words {
		println('$c.word $c.count')
	}
}

struct WordCount {
mut:
	word  string
	count int
}

[unsafe]
fn (mut wc WordCount) free() {
	unsafe {
		wc.word.free()
	}
}

struct WordCounter {
mut:
	buf []u8
	ht  map[u32]WordCount
}

fn new_word_counter() WordCounter {
	return WordCounter{
		buf: []u8{cap: 30}
		ht: map[u32]WordCount{}
	}
}

fn (mut wc WordCounter) flush() ? {
	if wc.buf.len == 0 {
		return
	}

	h := murmur3_32(wc.buf, 0)
	mut c := wc.ht[h]

	if c.count == 0 {
		c.word = wc.buf.bytestr()
	} else if c.word.len != wc.buf.len || !compare_slice(c.word, wc.buf) {
		return error('hash table collision: $c.word == $wc.buf.bytestr()')
	}

	c.count++
	wc.ht[h] = c
	wc.buf.clear()
}

[unsafe]
fn (mut wc WordCounter) free() {
	unsafe {
		wc.buf.free()
		wc.ht.free()
	}
}

fn (wc WordCounter) freq() []WordCount {
	mut words := []WordCount{cap: wc.ht.len}

	for _, c in wc.ht {
		words << c
	}

	return words
}

fn (wc WordCounter) len() int {
	return wc.ht.len
}

fn (wc WordCounter) words() []string {
	mut words := []string{cap: wc.ht.len}

	for _, c in wc.ht {
		words << c.word
	}

	return words
}

fn (mut wc WordCounter) write(mut buf []u8) ?int {
	mut i, mut j := 0, 0

	for ; j < buf.len; j++ {
		if buf[j] == ` ` || buf[j] == `\n` {
			wc.buf << buf[i..j]
			wc.flush()?
			i = j + 1
			continue
		}

		if buf[j] >= `A` && buf[j] <= `Z` {
			buf[j] += `a` - `A`
		}
	}

	if i < j {
		wc.buf << buf[i..j]
	}

	return buf.len
}

fn compare_slice(str string, s []u8) bool {
	if str.len != s.len {
		return false
	}

	for i, _ in s {
		if str[i] != s[i] {
			return false
		}
	}

	return true
}

const (
	v3x32_c1 = 0xcc9e2d51
	v3x32_c2 = 0x1b873593
	v3x32_r1 = 15
	v3x32_r2 = 13
	v3x32_m  = 5
	v3x32_n  = 0xe6546b64
)

fn murmur3_32(data []u8, seed u32) u32 {
	mut h := seed

	l := data.len

	for i := 0; i + 4 <= l; i += 4 {
		k := unsafe { *&u32(usize(data.data) + usize(i)) }
		h ^= murmur3_32_scramble(k)
		h = bits.rotate_left_32(h, v3x32_r2)
		h = h * v3x32_m + v3x32_n
	}

	rem := l % 4

	if rem > 0 {
		mut k := u32(0)

		for j := 1; j <= rem; j++ {
			k <<= 8
			k |= u32(data[l - j])
		}

		h ^= murmur3_32_scramble(k)
	}

	h ^= u32(l)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
}

fn murmur3_32_scramble(k u32) u32 {
	mut k2 := k
	k2 *= v3x32_c1
	k2 = bits.rotate_left_32(k2, v3x32_r1)
	k2 *= v3x32_c2

	return k2
}
