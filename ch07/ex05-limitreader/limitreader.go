package limitreader

import "io"

type LimitedReader struct {
	io.Reader
	n      int64
	remain int
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.remain < 1 {
		return 0, io.EOF
	}
	if r.remain > len(p) {
		r.remain -= len(p)
		return r.Reader.Read(p)
	}

	n, err = r.Reader.Read(p)
	if n > r.remain {
		n = r.remain
	}
	r.remain -= n
	p = p[:n]
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	var lr io.Reader
	lr = &LimitedReader{r, n, int(n)}
	return lr
}
