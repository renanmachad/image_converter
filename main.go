package main

import (
	"nuxui.org/nuxui/nux"
	_ "nuxui.org/nuxui/ui"
)

func main() {
	nux.Run(nux.NewWindow(nux.Attr{
		"width":  "500px",
		"height": "1:1",
		"title":  "Image converter",
	}))
}
