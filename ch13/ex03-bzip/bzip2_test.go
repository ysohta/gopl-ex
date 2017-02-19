package bzip_test

import (
	"bytes"
	"compress/bzip2" // reader
	"io"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/ysohta/gopl-ex/ch13/ex03-bzip" // writer
)

func TestBzip2Async(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	var nErrors, nGoRoutine, n int32
	nGoRoutine = 10

	var wg sync.WaitGroup
	for n = 0; n < nGoRoutine; n++ {
		wg.Add(1)
		go func() {
			tee := io.MultiWriter(w, &uncompressed)
			for i := 0; i < 10; i++ {
				io.WriteString(tee, "hello")
			}

			// count error on Close()
			if err := w.Close(); err != nil {
				atomic.AddInt32(&nErrors, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	want := nGoRoutine - 1
	if nErrors != want {
		t.Errorf("actual:%d want:%d", nErrors, want)
	}
}

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}
