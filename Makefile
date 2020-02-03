.PHONY: clean demo run build
default: build

# ZeroMQ version
ZEROMQ_TAG := -tags zmq_4_x

# pkg-config PATH
PKGCONFIG_PATH ?= /usr/lib/x86_64-linux-gnu/pkgconfig
export PKG_CONFIG_PATH=$(PKGCONFIG_PATH)

# Arch of env
ARCH 	= $(shell uname -s)
LINUX 	= Linux
DARWIN 	= Darwin

# Project, source, and build paths
PROJECT_ROOT 		:= $(shell pwd)
BUILD_DIR 		:= $(PROJECT_ROOT)/bin
DIST_DIR		:= $(PROJECT_ROOT)/dist
C_SRC_DIR 		:= $(PROJECT_ROOT)/src/c
QUEUE_SRC_DIR 		:= $(PROJECT_ROOT)/src/go/queue
LIBINTEROP_SRC_DIR 	:= $(PROJECT_ROOT)/src/go/libinterop
COLLECTOR_SRC_DIR	:= $(PROJECT_ROOT)/src/go/collector

# Golang lib names
LIBINTEROP_HEADER	 	:= libinterop.h
LIBINTEROP_WINDOWS		:= libinterop.dll
LIBINTEROP_LINUX 		:= libinterop.so
LIBINTEROP_DARWIN 		:= libinterop.dylib

# Executable names
PROGRAM_C 	:= program_c
QUEUE_GO 	:= queue_go
QUEUE_WIN	:= queue_go.exe
COLLECTOR_GO	:= collector_go
COLLECTOR_WIN	:= collector_go.exe

# Complier
CC 		?= gcc
GO 		?= go		# Change to your go executable file path
CC_FLAGS 	:= -g -O2



# Golang lib
libinterop-windows:
	GOHOSTOS=windows GOHOSTARCH=amd64 CGO_ENABLED=1 $(GO) build $(ZEROMQ_TAG) -i -x -v -ldflags "-s -w" -buildmode=c-shared -o $(BUILD_DIR)/$(LIBINTEROP_WINDOWS) $(LIBINTEROP_SRC_DIR)/*.go

libinterop-linux:
	GOHOSTOS=linux GOHOSTARCH=amd64 CGO_ENABLED=1 $(GO) build $(ZEROMQ_TAG) -i -x -v -ldflags "-s -w" -buildmode=c-shared -o $(BUILD_DIR)/$(LIBINTEROP_LINUX) $(LIBINTEROP_SRC_DIR)/*.go

libinterop-darwin:
	GOHOSTOS=darwin GOHOSTARCH=amd64 CGO_ENABLED=1 $(GO) build $(ZEROMQ_TAG) -i -x -v -ldflags "-s -w" -buildmode=c-shared -o $(BUILD_DIR)/$(LIBINTEROP_DARWIN) $(LIBINTEROP_SRC_DIR)/*.go

libinterop: libinterop-windows \
	libinterop-linux \
	libinterop-darwin

libinterop-no-unix: libinterop-windows \

libinterop-no-windows: libinterop-linux \
	libinterop-darwin

# Golang queue
queue-go:
	$(GO) build $(ZEROMQ_TAG) -i -x -v -o $(BUILD_DIR)/$(QUEUE_GO) $(QUEUE_SRC_DIR)/*.go; \
	if [ $(ARCH) != $(DARWIN) -a $(ARCH) != $(LINUX) ]; then \
		mv $(BUILD_DIR)/$(QUEUE_GO) $(BUILD_DIR)/$(QUEUE_WIN); \
	fi

collector-go:
	$(GO) build $(ZEROMQ_TAG) -i -x -v -o $(BUILD_DIR)/$(COLLECTOR_GO) $(COLLECTOR_SRC_DIR)/*.go; \
	if [ $(ARCH) != $(DARWIN) -a $(ARCH) != $(LINUX) ]; then \
		mv $(BUILD_DIR)/$(COLLECTOR_GO) $(BUILD_DIR)/$(COLLECTOR_WIN); \
	fi

# C-executable
program-c:
	$(CC) $(CC_FLAGS) $(C_SRC_DIR)/main.c -o $(BUILD_DIR)/$(PROGRAM_C) -I $(BUILD_DIR) -L $(BUILD_DIR) -linterop

run-program-c:
	cd $(BUILD_DIR) && env LD_LIBRARY_PATH=$(BUILD_DIR)  $(BUILD_DIR)/$(PROGRAM_C)

run-queue-go:
	cd $(BUILD_DIR) && $(BUILD_DIR)/$(QUEUE_GO)

run-collector-go:
	cd $(BUILD_DIR) && $(BUILD_DIR)/$(COLLECTOR_GO)

all: build

build: libinterop \
	program-c \
	queue-go \
	collector-go

install:
	# Make dir for distribute
	if [ ! -d $(DIST_DIR) ]; then \
		 mkdir $(DIST_DIR); \
	fi
	cp $(BUILD_DIR)/$(PROGRAM_C) $(DIST_DIR)
	cp $(BUILD_DIR)/$(QUEUE_GO)* $(DIST_DIR)
	cp $(BUILD_DIR)/$(COLLECTOR_GO)* $(DIST_DIR) 

	# Copying 'libinterop.so' to '/usr/lib/', may need a root privilege
	if [ $(ARCH) = $(LINUX) ]; then \
		cp $(BUILD_DIR)/$(LIBINTEROP_LINUX) /usr/lib/$(LIBINTEROP_LINUX); \
	elif [ $(ARCH) = $(DARWIN) ]; then \
		cp $(BUILD_DIR)/$(LIBINTEROP_DARWIN) /usr/local/lib/$(LIBINTEROP_DARWIN); \
	else \
		cp $(BUILD_DIR)/$(LIBINTEROP_WINDOWS) $(DIST_DIR); \
		cp /mingw64/bin/libzmq.dll $(DIST_DIR); \
		cp /mingw64/bin/libwinpthread-1.dll $(DIST_DIR); \
		cp /mingw64/bin/libgcc_s_seh-1.dll $(DIST_DIR); \
		cp /mingw64/bin/libstdc++-6.dll $(DIST_DIR); \
		cp /mingw64/bin/libsodium-23.dll $(DIST_DIR); \
	fi
	

# Clean
clean:
	# dll files
	rm -f $(BUILD_DIR)/$(LIBINTEROP_WINDOWS)
	rm -f $(BUILD_DIR)/$(LIBINTEROP_LINUX)
	rm -f $(BUILD_DIR)/$(LIBINTEROP_DARWIN)

	# header files
	rm -f $(BUILD_DIR)/$(LIBINTEROP_HEADER)

	# executable files
	rm -Rf $(BUILD_DIR)/$(PROGRAM_C)*
	rm -Rf $(BUILD_DIR)/$(QUEUE_GO)*
	rm -Rf $(BUILD_DIR)/$(COLLECTOR_GO)*
