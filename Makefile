default:
	go build github.com/andern/lin
test:
	go test ./...
bench:
	go test -bench=. ./...
deps:
	go get -u -v github.com/andern/lin
