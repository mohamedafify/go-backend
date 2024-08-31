GC = go
BIN = bin/server
SRC = $(shell find . -name '*.go')

all: ${BIN}

${BIN}: ${SRC}
	${GC} build -o ${BIN} . 

clean:
	rm -f ${BIN}

.PHONY: all clean
