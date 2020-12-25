PREFIX ?= /usr/local
DESTDIR =
PROGRAM = ghashverifier

build:
	go build -o $(PROGRAM)

install: build
	install -d $(DESTDIR)$(PREFIX)
	install $(PROGRAM) $(DESTDIR)$(PREFIX)
