
.PHONY: all
all: gitsubmod build
	echo done

.PHONY: build
build: antlr
	pushd src && go build -o ../build/ .

.PHONY: clean
clean:
	rm -rf build

antlr:
	antlr4 grammars-v4/c/C.g4 -o build/antlr4/java
	antlr4 grammars-v4/c/C.g4 -Dlanguage=Go -o build/antlr4/golang
	ln -s $PWD/build/antlr4/golang/grammars-v4/c $PWD/build/antlr4/golang/parser

gitsubmod:
	git submodule update --init --recursive
	
.PHONY: run
run: build
	./build/CCodeHash
