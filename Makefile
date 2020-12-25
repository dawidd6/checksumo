PREFIX = /usr/local
DESTDIR =
PROGRAM = checksumo

GTK_TAG = gtk_3_22

build:
	go build -tags $(GOTK_TAG) -o $(PROGRAM)

test:
	go test -tags $(GOTK_TAG) -v -count=1 ./...

install: build
	install -d $(DESTDIR)$(PREFIX)
	install $(PROGRAM) $(DESTDIR)$(PREFIX)
