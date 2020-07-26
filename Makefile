
.PHONY: all
all: gitsubmod build
	echo done

.PHONY: build
build: 
	pushd src && go build -o ../build/ .

.PHONY: clean
clean:
	rm -rf build

antlr:
	antlr4 grammars-v4/c/C.g4 -o build/antlr4/java
	antlr4 grammars-v4/c/C.g4 -Dlanguage=Go -o build/antlr4/golang

gitsubmod:
	git submodule update --init --recursive
	
.PHONY: run
run: build
	./build/CCodeHash
