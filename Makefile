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

###########################################################
# Load external lib build tasks
###########################################################
include script/*.mk

# Install dependencies
dev-install: zig-install

# Build the library
build: zig-build

clean:
	rm -rf dist/*
	rm -rf tmp

