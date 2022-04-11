# phony rules
.PHONY: build
build:
	@docker run --rm -it -v E:/Data-Structures-in-Golang/binary_search/:/go/src test-go bash -c "go build -o go-test binary_search.go"

.PHONY: build-image
build-image:
	@docker build -t=test-go . 

.PHONY: run
run:
	@docker run --rm -it -v E:/Data-Structures-in-Golang/binary_search/:/go/src test-go bash -c "go run binary_search.go"
