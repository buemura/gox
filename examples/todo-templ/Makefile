.PHONY: generate build start dev clean

generate:
	templ generate

build: generate
	go build -o todo .

start: build
	./todo

dev: generate
	templ generate --watch & \
	go run . & \
	wait

clean:
	rm -f todo views/*_templ.go
