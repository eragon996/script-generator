# make the echo stop for each build
#MAKEFLAGS += -s

GO_MAIN_FILE = sgen.go
GO_EXEC_FILE = sgen

clean:
	rm -rf sgen

build : clean
	-go build ${GO_MAIN_FILE}

help : build
	-./${GO_EXEC_FILE} help

t : build
	-./${GO_EXEC_FILE} sh