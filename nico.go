package main

import (
	"time"

	nskiplist "github.com/awesome-cap/skiplist"
)

func nInserts(n int) {
	list := nskiplist.New(12)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}
}

func nWorstInserts(n int) {
	list := nskiplist.New(12)
	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(float64(i), testByteString)
	}
}

func nAvgSearch(n int) {
	list := nskiplist.New(12)

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_,_ = list.Get(float64(i))
	}
}

func nSearchEnd(n int) {
	list := nskiplist.New(12)

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_,_ = list.Get(float64(n))
	}
}

func nDelete(n int) {
	list := nskiplist.New(12)

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Del(float64(i))
	}
}

func nWorstDelete(n int) {
	list := nskiplist.New(12)

	for i := 0; i < n; i++ {
		list.Set(float64(n-i), testByteString)
	}

	defer timeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Del(float64(n - i))
	}
}

