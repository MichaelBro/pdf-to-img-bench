//go:build cgo

package main

import "os"

var MAX_SIZE = 921

func main() {

	pdfBytes, err := os.ReadFile("../files/big.pdf")
	if err != nil {
		panic(err)
	}

	err = renderPageByPdfium(pdfBytes)
	if err != nil {
		panic(err)
	}

	renderPageByFitz(pdfBytes)

	r, err := os.Open("../files/big.pdf")
	if err != nil {
		panic(err)
	}
	pdfBoxToImage(r)
	defer r.Close()
}
