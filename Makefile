.PHONY: clean demo run build
default: build

# Arch of env
ARCH 	= $(shell uname -s)
LINUX 	= Linux
DARWIN 	= Darwin

# Project, source, and build paths
PROJECT_ROOT 		:= $(shell pwd)
BUILD_DIR 			:= $(PROJECT_ROOT)/bin
C_SRC_DIR 			:= $(PROJECT_ROOT)/src/c
QUEUE_SRC_DIR 		:= $(PROJECT_ROOT)/src/go/queue
LIBINTEROP_SRC_DIR 	:= $(PROJECT_ROOT)/src/go/libinterop
COLLECTOR_SRC_DIR	:= $(PROJECT_ROOT)/src/go/collector

# Golang lib names
LIBINTEROP_WINDOWS 			:= libinterop_windows.dll
LIBINTEROP_WINDOWS_HEADER 	:= libinterop_windows.h
LIBINTEROP_LINUX 			:= libinterop_linux.so
LIBINTEROP_LINUX_HEADER		:= libinterop_linux.h
LIBINTEROP_DARWIN 			:= libinterop_darwin.dylib
LIBINTEROP_DARWIN_HEADER 	:= libinterop_darwin.h

# Executable names
PROGRAM_C 	:= program_c
QUEUE_GO 	:= queue_go
COLLECTOR_GO:= collector_go

# Complier
CC 			?= gcc
GO 			?= go		# Change to your go executable file path
CC_FLAGS 	:= -g -O2



# Golang lib
libinterop-windows:
	GOHOSTOS=windows GOHOSTARCH=amd64 CGO_ENABLED=1 $(GO) build -i -x -v -ldflags "-s -w" -buildmode=c-shared -o $(BUILD_DIR)/$(LIBINTEROP_WINDOWS) $(LIBINTEROP_SRC_DIR)/*.go

libinterop-linux:
	GOHOSTOS=linux GOHOSTARCH=amd64 CGO_ENABLED=1 $(GO) build -i -x -v -ldflags "-s -w" -buildmode=c-shared -o $(BUILD_DIR)/$(LIBINTEROP_LINUX) $(LIBINTEROP_SRC_DIR)/*.go

libinterop-darwin:
	GOHOSTOS=darwin GOHOSTARCH=amd64 CGO_ENABLED=1 $(GO) build -i -x -v -ldflags "-s -w" -buildmode=c-shared -o $(BUILD_DIR)/$(LIBINTEROP_DARWIN) $(LIBINTEROP_SRC_DIR)/*.go

libinterop: libinterop-windows \
	libinterop-linux \
	libinterop-darwin

libinterop-no-unix: libinterop-windows \

libinterop-no-windows: libinterop-linux \
	libinterop-darwin \

# Golang queue
queue-go:
	$(GO) build -i -x -v -o $(BUILD_DIR)/$(QUEUE_GO) $(QUEUE_SRC_DIR)/*.go

collector-go:
	$(GO) build -i -x -v -o $(BUILD_DIR)/$(COLLECTOR_GO) $(COLLECTOR_SRC_DIR)/*.go

# C-executable
program-c:
	@if [ $(ARCH) = $(DARWIN) ]; \
	then \
		$(CC) $(CC_FLAGS) $(C_SRC_DIR)/main.c -o $(BUILD_DIR)/$(PROGRAM_C) -I $(BUILD_DIR) -L $(BUILD_DIR) -linterop_darwin; \
	elif [ $(ARCH) = $(LINUX) ]; \
	then \
		$(CC) $(CC_FLAGS) $(C_SRC_DIR)/main.c -o $(BUILD_DIR)/$(PROGRAM_C) -I $(BUILD_DIR) -L $(BUILD_DIR) -linterop_linux; \
	else \
		$(CC) $(CC_FLAGS) $(C_SRC_DIR)/main.c -o $(BUILD_DIR)/$(PROGRAM_C) -I $(BUILD_DIR) -L $(BUILD_DIR) -linterop_windows; \
	fi

run-program-c:
	cd $(BUILD_DIR) && env LD_LIBRARY_PATH=$(BUILD_DIR)  $(BUILD_DIR)/$(PROGRAM_C) &

run-queue-go:
	cd $(BUILD_DIR) && $(BUILD_DIR)/$(QUEUE_GO) &

run-collector-go:
	cd $(BUILD_DIR) && $(BUILD_DIR)/$(COLLECTOR_GO) &

all: build

build: libinterop \
	program-c \
	queue-go \
	collector-go

# Clean
clean:
	# dll files
	rm -f $(BUILD_DIR)/$(LIBINTEROP_WINDOWS)
	rm -f $(BUILD_DIR)/$(LIBINTEROP_LINUX)
	rm -f $(BUILD_DIR)/$(LIBINTEROP_DARWIN)

	# header files
	rm -f $(BUILD_DIR)/$(LIBINTEROP_WINDOWS_HEADER)
	rm -f $(BUILD_DIR)/$(LIBINTEROP_LINUX_HEADER)
	rm -f $(BUILD_DIR)/$(LIBINTEROP_DARWIN_HEADER)

	# executable files
	rm -Rf $(BUILD_DIR)/$(PROGRAM_C)*
	rm -Rf $(BUILD_DIR)/$(QUEUE_GO)*