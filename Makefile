PREFIX = /usr
APP = checksumo
APP_ID = com.github.dawidd6.checksumo
GOTK_TAG = gtk_3_22
GO_FLAGS = -v -mod=vendor -tags=$(GOTK_TAG)

build:
	glib-compile-resources --target=resources.h --generate-source data/$(APP).gresource.xml
	go build $(GO_FLAGS) -o $(APP)

test:
	go test $(GO_FLAGS) ./...

fmt:
	gofmt -w $(shell find . -name '*.go' -not -path './vendor/*')

install:
	install -D -m755 $(APP) $(DESTDIR)$(PREFIX)/bin/$(APP)
	install -D -m644 data/$(APP).desktop $(DESTDIR)$(PREFIX)/share/applications/$(APP_ID).desktop
	install -D -m644 data/$(APP).svg $(DESTDIR)$(PREFIX)/share/icons/$(APP_ID).svg
	install -D -m644 data/$(APP).gschema.xml $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/$(APP_ID).gschema.xml
	install -D -m644 data/$(APP).appdata.xml $(DESTDIR)$(PREFIX)/share/metainfo/$(APP_ID).appdata.xml
	glib-compile-schemas $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(APP)
	rm -f $(DESTDIR)$(PREFIX)/share/applications/$(APP_ID).desktop
	rm -f $(DESTDIR)$(PREFIX)/share/icons/$(APP_ID).svg
	rm -f $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/$(APP_ID).gschema.xml
	rm -f $(DESTDIR)$(PREFIX)/share/metainfo/$(APP_ID).appdata.xml
