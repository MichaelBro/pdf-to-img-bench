package main

import (
	"os"
	"testing"
)

var filesBytes = pdfBytesForBenchmark()

var toPdfPath = "../files"

func pdfBytesForBenchmark() [][]byte {

	files, err := os.ReadDir(toPdfPath)
	if err != nil {
		panic(err)
	}

	res := make([][]byte, len(files))
	for i, file := range files {
		res[i], err = os.ReadFile(toPdfPath + "/" + file.Name())
		if err != nil {
			panic(err)
		}
	}

	return res
}

func BenchmarkRenderPdfium(b *testing.B) {
	b.ResetTimer()

	for _, pdfBytes := range filesBytes {
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			b.StartTimer()
			_, err := renderImageByPixels(pdfBytes)
			b.StopTimer()
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkRenderFitz(b *testing.B) {
	b.ResetTimer()

	for _, pdfBytes := range filesBytes {
		b.StopTimer()

		for i := 0; i < b.N; i++ {
			b.StartTimer()
			_, err := GetPreviewImage(pdfBytes)
			b.StopTimer()
			if err != nil {
				panic(err)
			}
		}
	}
}

func BenchmarkRenderPdfbox(b *testing.B) {
	b.ResetTimer()

	b.StopTimer()

	files, err := os.ReadDir(toPdfPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		for i := 0; i < b.N; i++ {
			r, err := os.Open(toPdfPath + "/" + file.Name())
			if err != nil {
				panic(err)
			}

			b.StartTimer()
			pdfBoxToImage(r)
			b.StopTimer()

			defer r.Close()
		}
	}
}
