package acquire

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
)

type Game struct{}

var fonts = []struct {
	draw2d.FontData
	path string
}{
	{
		FontData: draw2d.FontData{
			Name:   "Noto Sans",
			Family: draw2d.FontFamilySans,
			Style:  draw2d.FontStyleNormal,
		},
		path: "NotoSans-Regular.ttf",
	},
	{
		FontData: draw2d.FontData{
			Name:   "Noto Sans",
			Family: draw2d.FontFamilySans,
			Style:  draw2d.FontStyleNormal,
		},
		path: "NotoSans-Regular.ttf",
	},
}

func init() {
	for _, f := range fonts {
		data, err := ioutil.ReadFile(f.path)
		if err != nil {
			panic(fmt.Sprintf("failed to read font '%s', %s", f.path, err))
		}
		font, err := truetype.Parse(data)
		if err != nil {
			panic(fmt.Sprintf("failed to parse font '%s', %s", f.path, err))
		}
		draw2d.RegisterFont(f.FontData, font)
	}
}

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
	gc.SetFontData(draw2d.FontData{
		Name:   "Noto Sans",
		Family: draw2d.FontFamilySans,
		Style:  draw2d.FontStyleNormal,
	})
	gc.SetFillColor(color.RGBA{0, 0, 0, 255})
	gc.SetFontSize(20)
	gc.Rotate(0.1)
	gc.FillStringAt("BLAH!", 50, 50)
	return nil
}
