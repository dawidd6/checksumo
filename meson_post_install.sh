#!/usr/bin/env bash

set -euo pipefail

SCHEMAS_DIR="$MESON_INSTALL_DESTDIR_PREFIX/share/glib-2.0/schemas"

echo Compiling gsettings schemas...
glib-compile-schemas "$SCHEMAS_DIR"
