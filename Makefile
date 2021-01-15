PREFIX = /usr
APP = checksumo
APP_ID = com.github.dawidd6.checksumo
GOTK_TAG = gtk_3_22
GO_SET_VARS = \
	-X main.appName=$(APP) \
	-X main.appID=$(APP_ID) \
	-X main.localeDomain=$(APP_ID) \
	-X main.localeDir=$(DESTDIR)$(PREFIX)/share/locale \
	-X main.uiResource=/data/$(APP).ui
GO_FLAGS = -v -mod=vendor -tags=$(GOTK_TAG) -ldflags="-s -w $(GO_SET_VARS)"
POTFILES = \
	data/checksumo.ui
LANGUAGES = \
	pl

build: build-resources
	go build $(GO_FLAGS) -o $(APP) ./src

build-resources:
	glib-compile-resources --target=src/resources.h --generate-source data/$(APP).gresource.xml

build-schemas:
	glib-compile-schemas $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas

build-po:
	$(foreach LANG,$(LANGUAGES),msgfmt --output-file po/$(LANG).mo po/$(LANG).po)

test:
	go test $(GO_FLAGS) ./...

fmt:
	gofmt -w $(shell find . -name '*.go' -not -path './vendor/*')

extract-pot:
	xgettext --output po/default.pot $(POTFILES)

init-po:
	$(foreach LANG,$(LANGUAGES),msginit --no-translator --input po/default.pot --output po/$(LANG).po --locale $(LANG))

update-po:
	$(foreach LANG,$(LANGUAGES),msgmerge --backup off --update po/$(LANG).po po/default.pot)

install-bin:
	install -D -m755 $(APP) $(DESTDIR)$(PREFIX)/bin/$(APP)

install-data:
	install -D -m644 data/$(APP).desktop $(DESTDIR)$(PREFIX)/share/applications/$(APP_ID).desktop
	install -D -m644 data/$(APP).svg $(DESTDIR)$(PREFIX)/share/icons/$(APP_ID).svg
	install -D -m644 data/$(APP).gschema.xml $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/$(APP_ID).gschema.xml
	install -D -m644 data/$(APP).appdata.xml $(DESTDIR)$(PREFIX)/share/metainfo/$(APP_ID).appdata.xml

install-po:
	$(foreach LANG,$(LANGUAGES),install -D -m644 po/$(LANG).mo $(DESTDIR)$(PREFIX)/share/locale/$(LANG)/LC_MESSAGES/$(APP_ID).mo)

install: build build-po install-bin install-data install-po build-schemas

uninstall-bin:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(APP)

uninstall-data:
	rm -f $(DESTDIR)$(PREFIX)/share/applications/$(APP_ID).desktop
	rm -f $(DESTDIR)$(PREFIX)/share/icons/$(APP_ID).svg
	rm -f $(DESTDIR)$(PREFIX)/share/glib-2.0/schemas/$(APP_ID).gschema.xml
	rm -f $(DESTDIR)$(PREFIX)/share/metainfo/$(APP_ID).appdata.xml

uninstall-po:
	$(foreach LANG,$(LANGUAGES),rm -f $(DESTDIR)$(PREFIX)/share/locale/$(LANG)/LC_MESSAGES/$(APP_ID).mo)

uninstall: uninstall-bin uninstall-data uninstall-po