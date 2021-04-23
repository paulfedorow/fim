all: build/fim build/bench

build/fim:
	mkdir -p build
	go build -o build/fim fim/fim.go

build/bench:
	mkdir -p build
	go build -o build/bench bench/bench.go

clean:
	rm -r build
