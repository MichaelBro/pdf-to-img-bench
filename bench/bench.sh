mkdir "result"

run_bench() {
  go test -run=. -bench=$2 -benchtime=5s -count 5 -benchmem -cpuprofile=$1/cpu.out -memprofile=$1/mem.out -trace=$1/trace.out | tee $1/bench.txt
}

to_pdf() {
  go tool pprof -pdf bench.test $1/cpu.out > $1/cpu.pdf
  go tool pprof -pdf --alloc_space bench.test $1/mem.out > $1/alloc_space.pdf
  go tool pprof -pdf --alloc_objects bench.test $1/mem.out > $1/alloc_objects.pdf
  go tool pprof -pdf --inuse_space bench.test $1/mem.out > $1/inuse_space.pdf
  go tool pprof -pdf --inuse_objects bench.test $1/mem.out > $1/inuse_objects.pdf
}

process() {
  mkdir $1
  run_bench $1 $2
  to_pdf $1
}

process "result/pdfium" BenchmarkRenderPdfium
process "result/fitz" BenchmarkRenderFitz
process "result/pdfbox" BenchmarkRenderPdfbox
