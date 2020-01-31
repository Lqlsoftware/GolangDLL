## Golang shared-c lib with IPC

### Dependency
- golang >= 1.10 (which support generate dll in windows)
- gcc

### Build
Build lib for all arch of host (Windows, Linux, OSX), and all programs. 
```shell script
make build
```

### Run
Firstly, run `program-c` in one console window(and keep it open):
```shell script
make run-program-c
```

Secondly, run `queue-go` in one console window(and keep it open):
```shell script
make run-queue-go
```

Finally, run `collector-go` in one console window(and keep it open):
```shell script
make run-queue-go
```

Input message in `program-c` which will go through `queue-go` and come to `collector-go`.

