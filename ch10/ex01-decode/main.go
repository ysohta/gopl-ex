package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	_ "image/gif" // register GIF decoder
	"image/jpeg"
	_ "image/jpeg" // register JPEG decoder
	"image/png"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

var (
	format = flag.String("format", "jpeg", "output image format")
)

func main() {
	flag.Parse()

	if err := convert(*format, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", *format, err)
		os.Exit(1)
	}
}

func convert(format string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "convert %q -> %q\n",kind,format)
	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{
			NumColors: 256,
			Drawer:    draw.FloydSteinberg,
		})
	default:
		return fmt.Errorf("invalid format:%q", format)
	}
}
