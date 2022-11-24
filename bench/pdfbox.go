package main

import (
	"context"
	"github.com/sfomuseum/go-pdfbox"
	"image"
	"image/jpeg"
	"io"
	"os"
)

func pdfBoxToImage(r *os.File) {
	ctx := context.Background()
	//defer r.Close()
	pdfBox, err := pdfbox.New(ctx, "pdfbox://")
	if err != nil {
		panic(err)
	}

	cb := func(ctx context.Context, path string, r io.Reader) error {
		_, _, err := image.Decode(r)
		if err != nil {
			panic(err)
		}
		return nil
	}

	err = pdfbox.PDFToImage(ctx, pdfBox, r, 1, 1, cb)
	if err != nil {
		panic(err)
	}
}

func saveJpg(img *image.Image) error {
	// Write the output to a file.
	f, err := os.Create("pdfbox.jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	return jpeg.Encode(f, *img, nil)
}
