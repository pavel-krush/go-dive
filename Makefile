PROTOC_IMAGE=namely/protoc-all:1.51_2
.PHONY: gen
gen:
	docker run --rm -it --entrypoint=tools/protogen.bash -v $$PWD:/defs $(PROTOC_IMAGE)
