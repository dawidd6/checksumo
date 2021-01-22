#!/usr/bin/env bash

set -euo pipefail

datadir="$MESON_INSTALL_PREFIX/share"

# Package managers set this so we don't need to run
if test -z "${DESTDIR-}"; then
    echo Updating icon cache...
    gtk-update-icon-cache -qtf "$datadir/icons/hicolor"

    echo Updating desktop database...
    update-desktop-database -q "$datadir/applications"

    echo Compiling GSettings schemas...
    glib-compile-schemas "$datadir/glib-2.0/schemas"
fi
