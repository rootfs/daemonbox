PKG = github.com/codegangsta/negroni\
	  github.com/gorilla/mux\
	  github.com/unrolled/render

godeps:
	GOPATH=`godep path` godep save ./...
	GOPATH=`godep path` go get ${PKG}

build: main.go
	godep go build -o daemonbox main.go

clean:
	rm -rf Godeps daemonbox

all: godeps build
