package main

import (
	"testing"
)

func TestMurmurHashV3x32(t *testing.T) {
	cases := []struct {
		data []byte
		want uint32
	}{
		{data: []byte(""), want: 0},
		{data: []byte("h"), want: 3565335251},
		{data: []byte("he"), want: 2020321539},
		{data: []byte("hel"), want: 4121571398},
		{data: []byte("hell"), want: 2707938291},
		{data: []byte("hello"), want: 613153351},
		{data: []byte("hello,"), want: 4091689770},
		{data: []byte("hello, "), want: 317845336},
		{data: []byte("hello, world"), want: 345750399},
	}

	for _, c := range cases {
		got := MurmurHashV3x32(c.data, 0)

		if got != c.want {
			t.Errorf("%q; got: %d, want: %d", c.data, got, c.want)
		}
	}
}
