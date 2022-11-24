mkdir "result"

go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=result/cpu.out -memprofile=result/mem.out -trace=result/trace.out | tee result/bench.txt

go tool pprof -pdf bench.test result/cpu.out > result/cpu.pdf
go tool pprof -pdf --alloc_space bench.test result/mem.out > result/alloc_space.pdf
go tool pprof -pdf --alloc_objects bench.test result/mem.out > result/alloc_objects.pdf
go tool pprof -pdf --inuse_space bench.test result/mem.out > result/inuse_space.pdf
go tool pprof -pdf --inuse_objects bench.test result/mem.out > result/inuse_objects.pdf