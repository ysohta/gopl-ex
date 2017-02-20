package bzip

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type writer struct {
	w io.Writer // underlying output stream
	f *os.File
}

func NewWriter(out io.Writer) io.WriteCloser {
	// create temporary file
	tmpfile, err := ioutil.TempFile("", ".tmp")
	if err != nil {
		log.Fatal(err)
	}

	w := &writer{w: out, f: tmpfile}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	return w.f.Write(data)
}

func (w *writer) Close() error {
	// close file
	w.f.Close()

	// execute bzip2 to stdout
	cmd := exec.Command("bzip2", "-c", w.f.Name())
	cmd.Stdout = w.w
	return cmd.Run()
}
