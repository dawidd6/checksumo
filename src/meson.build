go = find_program('go', required: true)
go_gtk_tag = 'gtk_3_22'
go_env = 'CGO_CFLAGS=-I' + meson.current_build_dir()
go_args = [
    '-v', '-mod=vendor', '-tags', go_gtk_tag
]

custom_target(
    meson.project_name(),
    build_by_default: true,
    command: [
        go_env, go, 'build', go_args, '-o', 'checksumo', meson.current_source_dir()
    ],
    console: true,
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
