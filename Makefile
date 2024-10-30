run:
	go build -C ./cmd/ -o sinker && ./cmd/sinker

tests:

clean:
	go -C ./cmd/ clean
