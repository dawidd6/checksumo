go = find_program('go', required: true)
go_env = ['CGO_CFLAGS=-I' + meson.current_build_dir() / '../ui']

custom_target(
    'build',
    build_by_default: true,
    build_always_stale: true,
    command: [
        go_env, go, 'build', '-v', '-mod=vendor', '-o', '@OUTPUT@', meson.current_source_dir()
    ],
    console: true,
    install: true,
    install_dir: get_option('bindir'),
    output: 'checksumo',
    depends: resources
)
