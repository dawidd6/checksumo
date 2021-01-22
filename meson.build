project('com.github.dawidd6.checksumo', 'c')

gnome = import('gnome')
i18n = import('i18n')

dependency('gtk+-3.0', version : '>=3.22')
dependency('glib-2.0')
dependency('gio-2.0')

subdir('data')
subdir('po')
subdir('ui')
subdir('src')

meson.add_install_script('meson_post_install.sh')