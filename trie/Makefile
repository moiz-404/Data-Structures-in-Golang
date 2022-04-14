# phony rules
.PHONY: build
build:
	@docker run --rm -it -v E:/Data-Structures-in-Golang/trie/:/go/src test-go bash -c "go build -o go-test trie.go"

.PHONY: build-image
build-image:
	@docker build -t=test-go . 

.PHONY: run
run:
	@docker run --rm -it -v E:/Data-Structures-in-Golang/trie/:/go/src test-go bash -c "go run trie.go"
