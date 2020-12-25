PREFIX = /usr
DESTDIR =
PROGRAM = checksumo

GOTK_TAG = gtk_3_22

build:
	glib-compile-resources --target=resources.h --generate-source data/data.gresource.xml
	go build -mod=vendor -trimpath -tags $(GOTK_TAG) -v -o $(PROGRAM)

test:
	go test -mod=vendor -tags $(GOTK_TAG) -v -count=1 ./...

install: build
	install -d $(DESTDIR)$(PREFIX)/bin
	install $(PROGRAM) $(DESTDIR)$(PREFIX)/bin
	install -d $(DESTDIR)$(PREFIX)/share/applications
	install data/$(PROGRAM).desktop $(DESTDIR)$(PREFIX)/share/applications
	install -d $(DESTDIR)$(PREFIX)/share/icons/hicolor/scalable/apps
	#install data/$(PROGRAM).svg $(DESTDIR)$(PREFIX)/share/icons/hicolor/scalable/apps
