project('com.github.dawidd6.checksumo', 'c')

gnome = import('gnome')
i18n = import('i18n')

go = find_program('go', required: true)
go_gtk_tag = 'gtk_3_22'
go_env = 'CGO_CFLAGS=-I' + meson.current_build_dir()
go_args = [
    '-v', '-mod=vendor', '-tags', go_gtk_tag
]

dependency('gtk+-3.0', version : '>=3.22')
dependency('glib-2.0')
dependency('gio-2.0')

subdir('data')
subdir('po')
subdir('ui')

custom_target(
    meson.project_name(),
    build_by_default: true,
    command: [
        go_env, go, 'build', go_args, '-o', 'checksumo', meson.current_source_dir() / 'src'
    ],
    install: true,
    install_dir: get_option('bindir'),
    output: 'checksumo',
    depends: resources
)

test(
    'test',
    go,
    args: [
        'test', go_args, meson.current_source_dir() / '...'
    ],
    env: go_env,
)

meson.add_install_script('meson_post_install.sh')