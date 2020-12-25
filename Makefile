PREFIX = /usr/local
DESTDIR =
PROGRAM = ghashverifier

GTK_HEADER_FILE = /usr/include/gtk-3.0/gtk/gtkversion.h
GTK_MAJOR = 3
GTK_MINOR = $(shell grep '\#define GTK_MINOR_VERSION' $(GTK_HEADER_FILE) | grep -Po '[0-9]+')
GTK_GO_TAG = gtk_$(GTK_MAJOR)_$(GTK_MINOR)

build:
	go build -tags $(GTK_GO_TAG) -o $(PROGRAM)

install: build
	install -d $(DESTDIR)$(PREFIX)
	install $(PROGRAM) $(DESTDIR)$(PREFIX)
