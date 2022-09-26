module main

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

	wc.flush()

	mut words := wc.freq()
	words.sort(b.count < a.count)

	for c in words {
		println('$c.word $c.count')
	}
}

struct WordCount {
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
	ht  map[string]int
}

fn new_word_counter() WordCounter {
	return WordCounter{
		ht: map[string]int{}
	}
}

fn (mut wc WordCounter) flush() {
	if wc.buf.len == 0 {
		return
	}

	word := wc.buf.bytestr()
	wc.buf.clear()
	wc.ht[word]++
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

	for k, v in wc.ht {
		words << WordCount{
			word: k
			count: v
		}
	}

	return words
}

fn (wc WordCounter) len() int {
	return wc.ht.len
}

fn (wc WordCounter) words() []string {
	mut words := []string{cap: wc.ht.len}

	for k, _ in wc.ht {
		words << k
	}

	return words
}

fn (mut wc WordCounter) write(mut buf []u8) ?int {
	mut i, mut j := 0, 0

	for ; j < buf.len; j++ {
		if buf[j] == ` ` || buf[j] == `\n` {
			wc.buf << buf[i..j]
			wc.flush()
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
