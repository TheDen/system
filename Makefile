BINARY=system
SRC=cmd/system/*.go
PLATFORMS := linux
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p bin/$(os)
	GOOS=$(os) GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o bin/$(os)/$(BINARY) $(SRC)

.PHONY: sync
sync:
	govendor sync

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: clean
clean:
	rm -rfv bin/*
