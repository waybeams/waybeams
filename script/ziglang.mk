###########################################
# Ziglang download and build task
###########################################

# Example URL
# https://ziglang.org/download/0.12.0/zig-linux-x86_64-0.12.0.tar.xz

# Ziglang version
ZIG_VERSION?=0.12.0
ZIG_PLATFORM?=$(shell uname | tr '[:upper:]' '[:lower:]')
ZIG_ARCH?=$(shell uname -m)
ZIG_PATH=lib/zig
ZIG_BIN=$(ZIG_PATH)/bin/zig
ZIG_TMP=tmp/zig
ZIG_FULLPATH=$(PROJECT_ROOT)/$(ZIG_PATH)
ZIG_TAR?=zig-$(ZIG_PLATFORM)-$(ZIG_ARCH)-$(ZIG_VERSION).tar.xz
ZIG_URL=https://ziglang.org/download/$(ZIG_VERSION)/$(ZIG_TAR)

$(ZIG_BIN):
	mkdir -p $(ZIG_PATH)
	mkdir -p $(ZIG_TMP)
	wget -O $(ZIG_TMP)/$(ZIG_TAR) $(ZIG_URL)
	touch $(ZIG_TMP)/$(ZIG_TAR)
	tar -xf $(ZIG_TMP)/$(ZIG_TAR) -C $(ZIG_PATH) --strip-components=1
	mkdir -p $(ZIG_PATH)/bin
	mv $(ZIG_PATH)/zig $(ZIG_PATH)/bin/zig
	rm -rf tmp; true

zig-install: $(ZIG_BIN)

zig-build:
	$(ZIG_BIN) build

zig-test:
	$(ZIG_BIN) test
