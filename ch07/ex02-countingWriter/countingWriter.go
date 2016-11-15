package countingWriter

import "io"

type CountWriter struct {
	io.Writer
	cnt int64
}

func (w *CountWriter) Write(p []byte) (n int, err error) {
	n, err = w.Writer.Write(p)
	w.cnt += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{w, 0}
	return &cw, &cw.cnt
}
