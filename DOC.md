## Golang shared-c with IPC

### Project Constitute

```
./
├── DOC.md
├── Makefile
├── README.md
├── go.mod
├── go.sum
└── src
    ├── c
    │   └── main.c
    └── go
        ├── collector
        │   └── main.go
        ├── libinterop
        │   └── main.go
        └── queue
            └── main.go
```

----

### Dependency
- Golang >= 1.10
- GCC
- libzmq
- [gozmq](github.com/alecthomas/gozmq)

----
### API of libinterop

#### func [Init()](https://github.com/Lqlsoftware/GolangDLL/blob/master/src/go/libinterop/main.go#L16);
Initialize ZeroMQ.

#### func [Send(GoString parameter)](https://github.com/Lqlsoftware/GolangDLL/blob/master/src/go/libinterop/main.go#L37);
Send a GoString to queue by ZeroMQ

----
### Windows Platform

In the purpose of using `cgo` and having unified style of `Makefile`, the build operation is based on `gcc` not `vc`.
Here we use `msys2` as an Linux-style environment, and `mingw-w64-x86_64-gcc` as compiler.

#### Steps for building

- Install msys2:
```
https://www.msys2.org/
```

- Install `mingw64/mingw-w64-x86_64-gcc` and `mingw-w64-x86_64-zeromq` in `msys2-mingw64`:
```shell script
pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-zeromq
```

- Download `pkgconfiglite` due to `pkg-config` has an [issue](https://github.com/rust-lang/pkg-config-rs/issues/51) of 
compatibility on Windows:
```shell script
https://sourceforge.net/projects/pkgconfiglite/
```

- Copy `pkgconfiglite` to the PATH of `msys2-mingw64` and make sure that PATH *DON'T* contains any blank(like `C:/Program Files/msys2/mingw64/bin` will cause problems in make build).

- Install golang and export `go` in `$PATH`:
```shell script
export PATH=/c/go/bin:$PATH
```

- Start build:
```shell script
make build 
```

- Distribute to folder `dist`:
```shell script
make install 
```

----
### Linux Platform
It is easier to build this project on Linux or macOS as the environment can be set up in few commands.
Here we give example by using Ubuntu.

- Install `golang` and `pkg-config`:
```shell script
sudo apt-get install golang pkg-config
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

- Distribute to folder `dist` (may need root privilege):
```shell script
sudo make install 
```