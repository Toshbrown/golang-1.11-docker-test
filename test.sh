#/bin/bash

echo "Running test on"
uname -a

echo "With docker version"
docker --version

echo "Testing golang:1.10-alpine3.8"
docker run --rm -v $(pwd):/usr/src/myapp -w /usr/src/myapp golang:1.10-alpine3.8 go run /usr/src/myapp/test.go

echo "Testing golang:1.11-alpine3.8"
docker run --rm -v $(pwd):/usr/src/myapp -w /usr/src/myapp golang:1.11-alpine3.8 go run /usr/src/myapp/test.go

echo "Testing golang:golang:1.11"
docker run --rm -v $(pwd):/usr/src/myapp -w /usr/src/myapp golang:1.11 go run /usr/src/myapp/test.go

echo "testing on local version"
go version
go run ./test.go