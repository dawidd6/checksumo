PREFIX = /usr
APP = checksumo
APP_ID = com.github.dawidd6.checksumo
GOTK_TAG = gtk_3_22
GO_FLAGS = -v -mod=vendor -tags=$(GOTK_TAG) -ldflags="-s -w -X main.appName=$(APP) -X main.appID=$(APP_ID) -X main.localeDir=$(DESTDIR)$(PREFIX)/share/locale -X main.uiResource=/data/$(APP).ui"
POTFILES = \
	data/checksumo.ui
LANGUAGES = \
	pl

build:
	glib-compile-resources --target=src/resources.h --generate-source data/$(APP).gresource.xml
	go build $(GO_FLAGS) -o $(APP) ./src

test:
	go test $(GO_FLAGS) ./...

fmt:
	gofmt -w $(shell find . -name '*.go' -not -path './vendor/*')

po-extract:
	xgettext --output po/default.pot $(POTFILES)

po-init:
	$(foreach LANG,$(LANGUAGES),msginit --no-translator --input po/default.pot --output po/$(LANG).po --locale $(LANG))

po-update:
	$(foreach LANG,$(LANGUAGES),msgmerge --backup off --update po/$(LANG).po po/default.pot)

po-build:
	$(foreach LANG,$(LANGUAGES),msgfmt --output-file po/$(LANG).mo po/$(LANG).po)

install: po-build
	install -D -m755 $(APP) $(DESTDIR)$(PREFIX)/bin/$(APP)
	install -D -m644 data/$(APP).desktop $(DESTDIR)$(PREFIX)/share/applications/$(APP_ID).desktop
	install -D -m644 data/$(APP).svg $(DESTDIR)$(PREFIX)/share/icons/$(APP_ID).svg
	install -D -m644 data/$(APP).gschema.xml $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/$(APP_ID).gschema.xml
	install -D -m644 data/$(APP).appdata.xml $(DESTDIR)$(PREFIX)/share/metainfo/$(APP_ID).appdata.xml
	install -D -m644 po/pl.mo $(DESTDIR)$(PREFIX)/share/locale/pl/LC_MESSAGES/$(APP).mo
	$(foreach LANG,$(LANGUAGES),install -D -m644 po/$(LANG).mo $(DESTDIR)$(PREFIX)/share/locale/$(LANG)/LC_MESSAGES/$(APP).mo)
	glib-compile-schemas $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(APP)
	rm -f $(DESTDIR)$(PREFIX)/share/applications/$(APP_ID).desktop
	rm -f $(DESTDIR)$(PREFIX)/share/icons/$(APP_ID).svg
	rm -f $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/$(APP_ID).gschema.xml
	rm -f $(DESTDIR)$(PREFIX)/share/metainfo/$(APP_ID).appdata.xml
	$(foreach LANG,$(LANGUAGES),rm -f $(DESTDIR)$(PREFIX)/share/locale/$(LANG)/LC_MESSAGES/$(APP).mo)
