project('com.github.dawidd6.checksumo', 'c')

gnome = import('gnome')
i18n = import('i18n')

gtk = dependency('gtk+-3.0', version : '>=3.22')
glib = dependency('glib-2.0')
gio = dependency('gio-2.0')

gnome.compile_resources(
    'resources', join_paths('data', meson.project_name() + '.gresource.xml'),
    source_dir: 'data'
)

gnome.compile_schemas(
    depend_files: join_paths('data', meson.project_name() + '.gschema.xml')
)

go = find_program('go', required: true)
gobuild = [
    'CGO_CFLAGS=I' + meson.current_build_dir(),
    go, 'build',
    '-mod, 'vendor',
    '-tags', 'gtk_3_22',
    '-o', meson.project_name(),
    meson.current_source_dir()
]

custom_target(
    meson.project_name(),
    build_by_default: true,
    command: gobuild,
    install: true,
    install_dir: get_option('bindir'),
    input: ['main.go'],
    output: meson.project_name()
)

install_data(
    join_paths('data', meson.project_name() + '.svg'),
    install_dir: join_paths(get_option('datadir'), 'icons')
)

install_data(
    join_paths('data', meson.project_name() + '.gschema.xml'),
    install_dir: join_paths(get_option('datadir'), 'glib-2.0', 'schemas')
)

i18n.gettext(meson.project_name(),
  args: [
    '--directory=' + meson.source_root(),
    '--from-code=UTF-8'
  ]
)

i18n.merge_file(
    input: join_paths('data', meson.project_name() + '.desktop.in'),
    output: meson.project_name() + '.desktop',
    po_dir: join_paths(meson.source_root(), 'po'),
    type: 'desktop',
    install: true,
    install_dir: join_paths(get_option('datadir'), 'applications')
)

i18n.merge_file(
    input: join_paths('data', meson.project_name() + '.appdata.xml.in'),
    output: meson.project_name() + '.appdata.xml',
    po_dir: join_paths(meson.source_root(), 'po'),
    install: true,
    install_dir: join_paths(get_option('datadir'), 'metainfo')
)

meson.add_install_script('meson/post_install.sh')
