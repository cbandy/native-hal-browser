
hal-browser.go: hal-browser/
	go-bindata -o $@ -prefix=$< -ignore=[.]git -ignore=[.]md $<...

clean:
	rm -f hal-browser.go

setup:
	go get github.com/jteeuwen/go-bindata/...
	go get -d github.com/elazarl/go-bindata-assetfs
	go get -d github.com/miketheprogrammer/go-thrust

.PHONY: clean setup
