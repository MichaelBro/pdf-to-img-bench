package main

import (
	"github.com/disintegration/imaging"
	"github.com/gen2brain/go-fitz"
	"image"
	"log"
)

var fistPage = 0

func GetPreviewImage(file []byte) (image.Image, error) {
	doc, err := fitz.NewFromMemory(file)
	if err != nil {
		return nil, err
	}

	defer doc.Close()

	srcImage, err := doc.ImageDPI(fistPage, 300)

	if err != nil {
		return nil, err
	}

	return fit(&srcImage, imaging.Box), err
}

func renderPageByFitz(file []byte) {
	img, err := GetPreviewImage(file)
	if err != nil {
		panic(err)
	}

	saveToJpeg(img)
}

func fit(image *image.Image, filter imaging.ResampleFilter) *image.NRGBA {
	return imaging.Fit(*image, MAX_SIZE, MAX_SIZE, filter)
}

func saveToJpeg(img image.Image) {
	err := imaging.Save(img, "fitz.jpg", imaging.JPEGQuality(75))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
