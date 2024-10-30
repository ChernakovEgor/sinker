run:
	go build -o sinker ./cmd/main.go && ./sinker

clean:
	rm sinker
