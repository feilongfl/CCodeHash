
WORKDIR := ${PWD}

.PHONY: all
all: gitsubmod build
	echo done

.PHONY: build
build: ./src/parser
	pushd src && go build -o ../build/CCodeHash .

.PHONY: clean
clean:
	rm -rf build
	rm -rf ./src/parser

./src/parser:
	# pushd grammars-v4/c/ && antlr4 C.g4 -o ${WORKDIR}/build/antlr4/java
	pushd grammars-v4/c/ && antlr4 C.g4 -Dlanguage=Go -o ${WORKDIR}/src/parser

gitsubmod:
	git submodule update --init --recursive
	
.PHONY: run
test: build
	@echo '========== TEST START ================='
	./build/CCodeHash ./test/bt.c

build/antlr4/java:
	pushd grammars-v4/c/ && antlr4 C.g4 -o ${WORKDIR}/build/antlr4/java

.PHONY: bttest
bttest:build/antlr4/java
	pushd build/antlr4/java && env CLASSPATH=/usr/share/java/antlr-complete.jar javac *.java
	pushd build/antlr4/java && grun C compilationUnit -gui ${WORKDIR}/test/bt.c
	# pushd build/antlr4/java && grun C compilationUnit -ps ${WORKDIR}/build/test.ps ${WORKDIR}/test/bt.c

