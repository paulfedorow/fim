common_go_files := $(shell find {dataset,ints,mine} -name "*.go")

all: build/fim build/bench

build/fim: $(common_go_files) $(shell find fim -name "*.go")
	mkdir -p build
	go build -o build/fim fim/fim.go

build/bench: $(common_go_files) $(shell find bench -name "*.go")
	mkdir -p build
	go build -o build/bench bench/bench.go

clean:
	rm -rf build
