package main

import (
	"image"

	"github.com/brdgme/acquire"
	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	a := &acquire.Game{}
	bounds := image.Rect(0, 0, 300, 300)
	dest := image.NewRGBA(bounds)
	gc := draw2dimg.NewGraphicContext(dest)
	if err := a.PlayerRender(1, gc, bounds); err != nil {
		panic(err)
	}
	if err := draw2dimg.SaveToPngFile("/tmp/egg.png", dest); err != nil {
		panic(err)
	}
}
