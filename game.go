package acquire

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d"
)

type Game struct{}

func (g *Game) PlayerRender(
	player int,
	gc draw2d.GraphicContext,
	bounds image.Rectangle,
) error {
	gc.SetFillColor(color.RGBA{0, 255, 0, 255})
	gc.SetStrokeColor(color.RGBA{255, 0, 0, 255})
	gc.SetLineWidth(5)
	gc.MoveTo(10, 10)
	gc.LineTo(float64(bounds.Max.X-50), 10)
	gc.QuadCurveTo(float64(bounds.Max.X-10), 10, float64(bounds.Max.X-10), 50)
	gc.LineTo(float64(bounds.Max.X-10), float64(bounds.Max.Y-10))
	gc.LineTo(10, float64(bounds.Max.Y-10))
	gc.Close()
	gc.FillStroke()
	return nil
}
