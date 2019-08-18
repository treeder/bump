build: 
	go build -o bump ./cmd

install: build
	sudo cp bump /usr/local/bin/bump

docker: 
	docker build -t treeder/bump:latest .

push: docker
	# todo: version these, or auto push this using CI
	# docker push treeder/bump:latest

release: docker
	./release.sh

test:
	go test ./...

.PHONY: install test build docker release
