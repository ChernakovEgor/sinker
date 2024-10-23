run:
	go build -C ./cmd/ -o sinker && ./cmd/sinker

clean:
	go -C ./cmd/ clean
