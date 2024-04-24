###########################################################
# Build Script
###########################################################

# Conditionally set SHELL to zsh if /usr/bin/zsh exists
ifeq ($(shell test -e /usr/bin/zsh && echo true),true)
  SHELL=/usr/bin/zsh
endif

# Find Operation System details (darwin or linux)
PLATFORM:=$(shell uname | tr A-Z a-z)

# uname -m gives x86_64, we just want the left side (x86)
ARCH=$(shell uname -m | cut -d'_' -f1)

# Set the root folder for the project (git repo)
PROJECT_ROOT=$(shell git rev-parse --show-toplevel)

BUILD_ENV:=development

###########################################################
# Configure Zig
###########################################################

ZIG_VERSION=0.12.0

.PHONY: build clean run test

# Build the library
build:
	zig build

###########################################################
# Load external lib build tasks
###########################################################
include script/*.mk

# Install dependencies
dev-install: zig-install

# Build and run the example
run: build
	zig run examples/todo/main.zig

test:
	zig test ./examples/todo/main.zig

clean:
	rm -rf tmp zig-cache zig-out; true

