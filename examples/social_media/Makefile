GOX = go run ../../cmd/gox

.PHONY: generate build start dev clean

generate:
	$(GOX) generate -dir .

build: generate
	go build -o social_media .

start: build
	./social_media

dev: generate
	$(GOX) watch -dir . & \
	go run . & \
	wait

clean:
	rm -f social_media views/*_gox.go
