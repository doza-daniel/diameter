all: smart naive pointgen

smart: bin
	go build -o bin/smart cmd/smart/main.go

naive: bin
	go build -o bin/naive cmd/naive/main.go

pointgen: bin
	go build -o bin/pointgen cmd/pointgen/main.go

benchmark: smart naive pointgen
	benchmarks.sh 10000 200

pdf:
	cd docs; \
	pdflatex dijametar.tex

bin:
	mkdir -p bin

.PHONY: clean

clean:
	rm -rf bin/ *.dat *.png docs/dijametar.{aux,out,log} docs/*.pdf
	
