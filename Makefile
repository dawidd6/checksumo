PREFIX = /usr/local
DESTDIR =
PROGRAM = checksumo
VERSION = $(shell git describe --tags 2>/dev/null || git rev-parse HEAD)

GOTK_TAG = gtk_3_22

build:
	glib-compile-resources --target=resources.h --generate-source data/data.gresource.xml
	go build -mod=vendor -tags $(GOTK_TAG) -v -o $(PROGRAM) -ldflags "-s -w -X main.Version=$(VERSION)"

test:
	go test -mod=vendor -tags $(GOTK_TAG) -v -count=1 ./...

install:
	install -D -m755 $(PROGRAM) $(DESTDIR)$(PREFIX)/bin/$(PROGRAM)
	install -D -m644 data/$(PROGRAM).desktop $(DESTDIR)$(PREFIX)/share/applications/$(PROGRAM).desktop
	install -D -m644 data/$(PROGRAM).svg $(DESTDIR)$(PREFIX)/share/icons/$(PROGRAM).svg

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(PROGRAM)
	rm -f $(DESTDIR)$(PREFIX)/share/applications/$(PROGRAM).desktop
	rm -f $(DESTDIR)$(PREFIX)/share/icons/$(PROGRAM).svg
