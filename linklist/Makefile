# phony rules
.PHONY: build
build:
	@docker run --rm -it -v E:/Data-Structures-in-Golang/linklist/:/go/src test-go bash -c "go build -o go-test linklist.go"

.PHONY: build-image
build-image:
	@docker build -t=test-go . 

.PHONY: run
run:
	@docker run --rm -it -v E:/Data-Structures-in-Golang/linklist/:/go/src test-go bash -c "go run linklist.go"
