package main

import (
	"net"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/miketheprogrammer/go-thrust"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	address := listener.Addr().String()

	thrust.InitLogger()
	thrust.Start()

	thrustWindow := thrust.NewWindow("http://"+address+"/browser.html", nil)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	http.Handle("/", http.FileServer(&assetfs.AssetFS{
		Asset: Asset, AssetDir: AssetDir,
	}))
	http.ListenAndServe(address, nil)
}
