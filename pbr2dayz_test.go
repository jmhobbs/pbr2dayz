package pbr2dayz

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPBRConvert(t *testing.T) {
	baseColor := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	baseColor.SetNRGBA(0, 0, color.NRGBA{R: 10, G: 20, B: 30, A: 40})
	baseColor.SetNRGBA(1, 0, color.NRGBA{R: 40, G: 50, B: 60, A: 70})
	baseColor.SetNRGBA(0, 1, color.NRGBA{R: 70, G: 80, B: 90, A: 100})
	baseColor.SetNRGBA(1, 1, color.NRGBA{R: 100, G: 110, B: 120, A: 130})

	normal := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	normal.SetNRGBA(0, 0, color.NRGBA{R: 130, G: 120, B: 110, A: 100})
	normal.SetNRGBA(1, 0, color.NRGBA{R: 90, G: 80, B: 70, A: 60})
	normal.SetNRGBA(0, 1, color.NRGBA{R: 50, G: 40, B: 30, A: 20})
	normal.SetNRGBA(1, 1, color.NRGBA{R: 10, G: 20, B: 30, A: 40})

	ao := image.NewGray(image.Rect(0, 0, 1, 1))
	ao.SetGray(0, 0, color.Gray{Y: 80})

	metallic := image.NewGray(image.Rect(0, 0, 1, 1))
	metallic.SetGray(0, 0, color.Gray{Y: 120})

	roughness := image.NewGray(image.Rect(0, 0, 1, 1))
	roughness.SetGray(0, 0, color.Gray{Y: 30})

	dayz := PBR{
		BaseColor: baseColor,
		Normal:    normal,
		AO:        ao,
		Metallic:  metallic,
		Roughness: roughness,
	}.Convert()

	assert.Equal(t, image.Rect(0, 0, 2, 2), dayz.CO.Bounds())
	assert.Equal(t, image.Rect(0, 0, 2, 2), dayz.NOHQ.Bounds())
	assert.Equal(t, image.Rect(0, 0, 2, 2), dayz.AS.Bounds())
	assert.Equal(t, image.Rect(0, 0, 2, 2), dayz.SMDI.Bounds())

	assert.Equal(t, color.RGBA{R: 10, G: 20, B: 30, A: 255}, rgbaAt(dayz.CO, 0, 0))
	assert.Equal(t, color.RGBA{R: 100, G: 110, B: 120, A: 255}, rgbaAt(dayz.CO, 1, 1))

	assert.Equal(t, color.RGBA{R: 130, G: 120, B: 110, A: 255}, rgbaAt(dayz.NOHQ, 0, 0))
	assert.Equal(t, color.RGBA{R: 10, G: 20, B: 30, A: 255}, rgbaAt(dayz.NOHQ, 1, 1))

	assert.Equal(t, color.RGBA{R: 255, G: 80, B: 255, A: 255}, rgbaAt(dayz.AS, 0, 0))
	assert.Equal(t, color.RGBA{R: 255, G: 80, B: 255, A: 255}, rgbaAt(dayz.AS, 1, 1))

	assert.Equal(t, color.RGBA{R: 255, G: 120, B: 225, A: 255}, rgbaAt(dayz.SMDI, 0, 0))
	assert.Equal(t, color.RGBA{R: 255, G: 120, B: 225, A: 255}, rgbaAt(dayz.SMDI, 1, 1))
}

func rgbaAt(img image.Image, x int, y int) color.RGBA {
	return color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
}
