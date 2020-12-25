PREFIX = /usr/local
DESTDIR =
PROGRAM = checksumo

build:
	go build -tags gtk_3_18 -o $(PROGRAM)

test:
	go test -tags gtk_3_18 -v -count=1 ./...

install: build
	install -d $(DESTDIR)$(PREFIX)
	install $(PROGRAM) $(DESTDIR)$(PREFIX)
