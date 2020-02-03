## Golang shared-c lib with IPC

### Dependency
- golang >= 1.10 (which support generate dll in windows)
- gcc

### Build

#### For windows

- Install ZeroMQ in `msys2-mingw64`
```shell script
pacman -S mingw-w64-x86_64-zeromq
```

- Download `pkgconfiglite` due to `pkg-config` has [issue](https://github.com/rust-lang/pkg-config-rs/issues/51) on Windows:
```shell script
https://sourceforge.net/projects/pkgconfiglite/
```

- Copy `pkgconfiglite` to the PATH of `msys2-mingw64` and make sure that PATH *DON'T* contains any blank(like `C:/Program Files/msys2/mingw64/bin` will cause problems in make build).


- Start build
```shell script
make build
```

#### For Unix like platform (Ubuntu for example)

- Install pkg-config:
```shell script
sudo apt-get install pkg-config
```

- Install libzmq:
```shell script
https://zeromq.org/download/#linux
```

- Modify `Makefile` line 5 as the libzmq version you just install:
```
For 4.x: zmq_4_x
For 3.x: zmq_3_x
For 2.1: zmq_2_1
For 2.2.x: (empty)
```

- Find libzmq's pkg-config file `libzmq.pc`:
```shell script
find / | grep libzmq
```

- Modify `Makefile` line 8 as the dir that `libzmq.pc` exist
- Start build
```shell script
make build
```
----
### Run
Firstly, run `program-c` in one console window(and keep it open, also you can open as much `program-c` as you like):
```shell script
make run-program-c
```

Secondly, run `queue-go` in one console window(and keep it open):
```shell script
make run-queue-go
```

Finally, run `collector-go` in one console window(and keep it open):
```shell script
make run-collector-go
```

Input message in `program-c` which will go through `queue-go` and come to `collector-go`.

