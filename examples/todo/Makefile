GOX = go run ../../cmd/gox

.PHONY: generate build start dev clean

generate:
	$(GOX) generate -dir .

build: generate
	go build -o todo .

start: build
	./todo

dev: generate
	$(GOX) watch -dir . & \
	go run . & \
	wait

clean:
	rm -f todo views/views_gox.go
