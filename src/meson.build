go = find_program('go', required: true)
go_env = 'CGO_CFLAGS=-I' + meson.current_build_dir() / '../ui'
go_args = ['-v', '-mod=vendor']

custom_target(
    meson.project_name(),
    build_by_default: false,
    build_always_stale: true,
    command: [
        go_env, go, 'build', go_args, '-o', '@OUTPUT@', meson.current_source_dir()
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
