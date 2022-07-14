all: build
build:
	GOARCH=arm64 go build -o ups_monitor

.PHONY: clean
clean:
	rm -rf ups_monitor
