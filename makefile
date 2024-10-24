all: clean build

build:
	go build -o gn

release: clean build
	gh release create $(VERSION) \
		--title $(VERSION) \
		--generate-notes \
		gn

clean:
	rm -f gn
