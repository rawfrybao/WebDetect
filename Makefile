# Go compiler
GO := go

# Output directory
OUT_DIR := bin

# Name of the executable
FETCHEXECUTABLE := xpathfetch
WEBHOOKEXECUTABLE := webhook
DELPREFIX := rm -rf

# Detect operating system
ifeq ($(OS),Windows_NT)
    FETCHEXECUTABLE := $(FETCHEXECUTABLE).exe
	WEBHOOKEXECUTABLE := $(WEBHOOKEXECUTABLE).exe
	DELPREFIX := rmdir /S /Q
	MKDIR = if not exist $(OUT_DIR) mkdir $(OUT_DIR)
else
    ifeq ($(OS),MINGW32_NT-6.2)
        FETCHEXECUTABLE := $(FETCHEXECUTABLE).exe
		WEBHOOKEXECUTABLE := $(WEBHOOKEXECUTABLE).exe
		DELPREFIX := rmdir /S /Q
		MKDIR = if not exist $(OUT_DIR) mkdir $(OUT_DIR)
	else
		MKDIR = mkdir -p $(OUT_DIR)
    endif
endif

.PHONY: all build clean generate

all: build

build:
	@$(MKDIR)
	$(GO) build -o $(OUT_DIR)/$(FETCHEXECUTABLE) ./cmd/xpathfetch/main.go
	$(GO) build -o $(OUT_DIR)/$(WEBHOOKEXECUTABLE) ./cmd/webhook/main.go

clean:
	@$(DELPREFIX) $(OUT_DIR)

generate:
	@go mod download
	@go generate ./...