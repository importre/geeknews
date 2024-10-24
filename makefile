all: clean build

build:
	@go build -o gn
	@shasum -a 256 gn

release: clean build
	gh release create $(VERSION) \
		--title $(VERSION) \
		--generate-notes \
		gn

clean:
	@rm -f gn
