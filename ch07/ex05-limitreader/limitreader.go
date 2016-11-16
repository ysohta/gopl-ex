package limitreader

import "io"

type LimitedReader struct {
	io.Reader
	n      int64
	remain int
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	// nothing remains
	if r.remain < 1 {
		return 0, io.EOF
	}

	n, err = r.Reader.Read(p)

	// over limit
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
