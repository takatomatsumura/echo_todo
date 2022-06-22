OUTDIR ?= out
TARGET ?= server
SRCS := $(shell find . -name cmd -prune -o -type f -name '*.go' -print)

build: $(OUTDIR)/$(TARGET)

clean:
	$(RM) -r $(OUTDIR)

run: $(OUTDIR)/$(TARGET)
	exec $<

start: $(OUTDIR)/liveserver
	exec $<

update:
	go get -d -u
	go mod tidy

.PHONY: build clean run start update

$(OUTDIR)/$(TARGET): $(SRCS) go.mod go.sum
	go build -o $@

$(OUTDIR)/liveserver: $(wildcard cmd/liveserver/*.go)
	go build -o $@ ./cmd/liveserver

ent/ent.go: $(wildcard ent/schema/*.go)
	go generate ./ent
