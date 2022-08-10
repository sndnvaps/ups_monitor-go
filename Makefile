DESTDIR?=/usr
PREFIX?=/local
GOARCH?=arm64

all: build
build:
	GOARCH=$(GOARCH) go build -o ups_monitor_go

.PHONY: install
install:
	$Q echo "[Install]"
	$Q mkdir -p		$(DESTDIR)$(PREFIX)/bin
	$Q cp ups_monitor_go	$(DESTDIR)$(PREFIX)/bin

.PHONY: uninstall
uninstall:
	rm -rf $(DESTDIR)$(PREFIX)/bin/ups_monitor_go

.PHONY: clean
clean:
	rm -rf ups_monitor_go
