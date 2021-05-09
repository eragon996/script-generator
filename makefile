# make the echo stop for each build
#MAKEFLAGS += -s

GO_EXEC_FILE = sgen

clean:
	rm -rf sgen

build : clean
	-go build -o ${GO_EXEC_FILE}

help : build
	-./${GO_EXEC_FILE} help

t : build
	-./${GO_EXEC_FILE} sh

install: build
	-cp ${GO_EXEC_FILE} ~/bin/sgen
