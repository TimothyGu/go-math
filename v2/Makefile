GENERATED_PKGS := i16math i32math i64math i8math imath u16math u32math u64math u8math umath
GENERATED := $(addsuffix /const.gen.go,$(GENERATED_PKGS))
TEMPLATES := $(wildcard generate/*.tmpl)

.PHONY: all
all: $(GENERATED)
	go build ./...

.PHONY: clean
clean:
	rm -rf $(GENERATED_PKGS)

$(GENERATED): generate/generate.go $(TEMPLATES)
	go generate

.PHONY: generate
generate: $(GENERATED)

.PHONY: test
test: $(GENERATED)
	go test ./...
