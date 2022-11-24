package main

import (
	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/single_threaded"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"
)

// Be sure to close pools/instances when you're done with them.
var pool pdfium.Pool
var instance pdfium.Pdfium

// init rename to runInit for off auto run
func init() {
	// Init the PDFium library and return the instance to open documents.
	pool = single_threaded.Init(single_threaded.Config{})

	var err error
	instance, err = pool.GetInstance(time.Second * 30)
	if err != nil {
		log.Fatal(err)
	}
}

func renderPageByPdfium(pdfBytes []byte) error {
	img, err := renderImageByPixels(pdfBytes)
	if err != nil {
		return err
	}

	return saveAsJpg(img)
}

func saveAsJpg(img *image.RGBA) error {
	// Write the output to a file.
	f, err := os.Create("pdfium.jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	return jpeg.Encode(f, img, nil)
}

func renderImage(pdfBytes *[]byte, dpi int) (*image.RGBA, error) {
	// Open the PDF using PDFium (and claim a worker)
	doc, err := instance.OpenDocument(&requests.OpenDocument{
		File: pdfBytes,
	})
	if err != nil {
		return nil, err
	}

	// Always close the document, this will release its resources.
	defer func(instance pdfium.Pdfium, request *requests.FPDF_CloseDocument) {
		_, err := instance.FPDF_CloseDocument(request)
		if err != nil {

		}
	}(instance, &requests.FPDF_CloseDocument{
		Document: doc.Document,
	})
	// Render the page in DPI 200.
	pageRender, err := instance.RenderPageInDPI(&requests.RenderPageInDPI{
		DPI: dpi, // The DPI to render the page in.
		Page: requests.Page{
			ByIndex: &requests.PageByIndex{
				Document: doc.Document,
				Index:    0,
			},
		}, // The page to render, 0-indexed.
	})
	if err != nil {
		return nil, err
	}

	return pageRender.Result.Image, nil
}

func renderImageByPixels(pdfBytes []byte) (*image.RGBA, error) {
	// Open the PDF using PDFium (and claim a worker)
	//runInit()
	doc, err := instance.OpenDocument(&requests.OpenDocument{
		File: &pdfBytes,
	})
	if err != nil {
		return nil, err
	}

	// Always close the document, this will release its resources.
	defer func(instance pdfium.Pdfium, request *requests.FPDF_CloseDocument) {
		_, err := instance.FPDF_CloseDocument(request)
		if err != nil {

		}
	}(instance, &requests.FPDF_CloseDocument{
		Document: doc.Document,
	})
	// Render the page in DPI 200.
	pageRender, err := instance.RenderPageInPixels(&requests.RenderPageInPixels{
		Page: requests.Page{
			ByIndex: &requests.PageByIndex{
				Document: doc.Document,
				Index:    0,
			},
		}, // The page to render, 0-indexed.
		Width:  MAX_SIZE,
		Height: MAX_SIZE,
	})
	if err != nil {
		return nil, err
	}

	return pageRender.Result.Image, nil
}
