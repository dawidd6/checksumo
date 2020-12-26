PREFIX = /usr/local
DESTDIR =
PROGRAM = checksumo

GOTK_TAG = gtk_3_22

build:
	glib-compile-resources --target=resources.h --generate-source data/data.gresource.xml
	go build -mod=vendor -trimpath -tags $(GOTK_TAG) -v -o $(PROGRAM)

test:
	go test -mod=vendor -tags $(GOTK_TAG) -v -count=1 ./...

install:
	install -D -m755 $(PROGRAM) $(DESTDIR)$(PREFIX)/bin/$(PROGRAM)
	install -D -m644 data/$(PROGRAM).desktop $(DESTDIR)$(PREFIX)/share/applications/$(PROGRAM).desktop
	install -D -m644  data/$(PROGRAM).svg $(DESTDIR)$(PREFIX)/share/icons/$(PROGRAM).svg

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(PROGRAM)
	rm -f $(DESTDIR)$(PREFIX)/share/applications/$(PROGRAM).desktop
	rm -f $(DESTDIR)$(PREFIX)/share/icons/$(PROGRAM).svg
