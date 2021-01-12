project('com.github.dawidd6.checksumo', 'c')

gnome = import('gnome')
i18n = import('i18n')

go = find_program('go', required: true)
go_gtk_tag = 'gtk_3_22'
go_env = 'CGO_CFLAGS=-I@0@'.format(meson.current_build_dir())
go_args = '-v -mod=vendor -tags=@0@'.format(go_gtk_tag)

dependency('gtk+-3.0', version : '>=3.22')
dependency('glib-2.0')
dependency('gio-2.0')

subdir('data')
subdir('po')

custom_target(
    meson.project_name(),
    build_by_default: true,
    build_always: true,
    command: '@0@ @1@ build @2@ -o @3@ @4@'.format(
        go_env, go.path(), go_args, meson.project_name(), meson.current_source_dir()
    ).split(),
    install: true,
    install_dir: get_option('bindir'),
    output: meson.project_name(),
    depends: resources
)

test(
    'test',
    go,
    args: 'test @0@ @1@'.format(
        go_args, join_paths(meson.current_source_dir(), '...')
    ).split(),
    env: go_env,
)

meson.add_install_script('meson/post_install.sh')

