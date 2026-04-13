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

func TestDayZConvert(t *testing.T) {
	co := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	co.SetNRGBA(0, 0, color.NRGBA{R: 10, G: 20, B: 30, A: 40})
	co.SetNRGBA(1, 0, color.NRGBA{R: 40, G: 50, B: 60, A: 70})
	co.SetNRGBA(0, 1, color.NRGBA{R: 70, G: 80, B: 90, A: 100})
	co.SetNRGBA(1, 1, color.NRGBA{R: 100, G: 110, B: 120, A: 130})

	nohq := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	nohq.SetNRGBA(0, 0, color.NRGBA{R: 130, G: 120, B: 110, A: 100})
	nohq.SetNRGBA(1, 0, color.NRGBA{R: 90, G: 80, B: 70, A: 60})
	nohq.SetNRGBA(0, 1, color.NRGBA{R: 50, G: 40, B: 30, A: 20})
	nohq.SetNRGBA(1, 1, color.NRGBA{R: 10, G: 20, B: 30, A: 40})

	as := image.NewRGBA(image.Rect(0, 0, 1, 1))
	as.SetRGBA(0, 0, color.RGBA{R: 255, G: 80, B: 255, A: 255})

	smdi := image.NewRGBA(image.Rect(0, 0, 1, 1))
	smdi.SetRGBA(0, 0, color.RGBA{R: 255, G: 120, B: 225, A: 255})

	pbr := DayZ{
		CO:   co,
		NOHQ: nohq,
		AS:   as,
		SMDI: smdi,
	}.Convert()

	assert.Equal(t, image.Rect(0, 0, 2, 2), pbr.BaseColor.Bounds())
	assert.Equal(t, image.Rect(0, 0, 2, 2), pbr.Normal.Bounds())
	assert.Equal(t, image.Rect(0, 0, 2, 2), pbr.AO.Bounds())
	assert.Equal(t, image.Rect(0, 0, 2, 2), pbr.Metallic.Bounds())
	assert.Equal(t, image.Rect(0, 0, 2, 2), pbr.Roughness.Bounds())

	assert.Equal(t, color.RGBA{R: 10, G: 20, B: 30, A: 255}, rgbaAt(pbr.BaseColor, 0, 0))
	assert.Equal(t, color.RGBA{R: 100, G: 110, B: 120, A: 255}, rgbaAt(pbr.BaseColor, 1, 1))

	assert.Equal(t, color.RGBA{R: 130, G: 120, B: 110, A: 255}, rgbaAt(pbr.Normal, 0, 0))
	assert.Equal(t, color.RGBA{R: 10, G: 20, B: 30, A: 255}, rgbaAt(pbr.Normal, 1, 1))

	assert.Equal(t, uint8(80), grayAt(pbr.AO, 0, 0))
	assert.Equal(t, uint8(80), grayAt(pbr.AO, 1, 1))

	assert.Equal(t, uint8(120), grayAt(pbr.Metallic, 0, 0))
	assert.Equal(t, uint8(120), grayAt(pbr.Metallic, 1, 1))

	assert.Equal(t, uint8(30), grayAt(pbr.Roughness, 0, 0))
	assert.Equal(t, uint8(30), grayAt(pbr.Roughness, 1, 1))
}

func TestPBRConvertRoundTrip(t *testing.T) {
	original := PBR{
		BaseColor: newNRGBAImage(image.Rect(0, 0, 2, 2), []color.NRGBA{
			{R: 1, G: 2, B: 3, A: 255},
			{R: 4, G: 5, B: 6, A: 255},
			{R: 7, G: 8, B: 9, A: 255},
			{R: 10, G: 11, B: 12, A: 255},
		}),
		Normal: newNRGBAImage(image.Rect(0, 0, 2, 2), []color.NRGBA{
			{R: 200, G: 190, B: 180, A: 255},
			{R: 170, G: 160, B: 150, A: 255},
			{R: 140, G: 130, B: 120, A: 255},
			{R: 110, G: 100, B: 90, A: 255},
		}),
		AO: newGrayImage(image.Rect(0, 0, 2, 2), []uint8{
			13, 26,
			39, 52,
		}),
		Metallic: newGrayImage(image.Rect(0, 0, 2, 2), []uint8{
			64, 96,
			128, 160,
		}),
		Roughness: newGrayImage(image.Rect(0, 0, 2, 2), []uint8{
			5, 25,
			125, 250,
		}),
	}

	roundTripped := original.Convert().Convert()

	assertEqualRGBAImage(t, original.BaseColor, roundTripped.BaseColor)
	assertEqualRGBAImage(t, original.Normal, roundTripped.Normal)
	assertEqualGrayImage(t, original.AO, roundTripped.AO)
	assertEqualGrayImage(t, original.Metallic, roundTripped.Metallic)
	assertEqualGrayImage(t, original.Roughness, roundTripped.Roughness)
}

func rgbaAt(img image.Image, x int, y int) color.RGBA {
	return color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
}

func grayAt(img image.Image, x int, y int) uint8 {
	return color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y
}

func newNRGBAImage(bounds image.Rectangle, pixels []color.NRGBA) image.Image {
	img := image.NewNRGBA(bounds)
	index := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.SetNRGBA(x, y, pixels[index])
			index++
		}
	}

	return img
}

func newGrayImage(bounds image.Rectangle, pixels []uint8) image.Image {
	img := image.NewGray(bounds)
	index := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.SetGray(x, y, color.Gray{Y: pixels[index]})
			index++
		}
	}

	return img
}

func assertEqualRGBAImage(t *testing.T, expected image.Image, actual image.Image) {
	t.Helper()
	assert.Equal(t, expected.Bounds(), actual.Bounds())

	for y := expected.Bounds().Min.Y; y < expected.Bounds().Max.Y; y++ {
		for x := expected.Bounds().Min.X; x < expected.Bounds().Max.X; x++ {
			assert.Equal(t, rgbaAt(expected, x, y), rgbaAt(actual, x, y))
		}
	}
}

func assertEqualGrayImage(t *testing.T, expected image.Image, actual image.Image) {
	t.Helper()
	assert.Equal(t, expected.Bounds(), actual.Bounds())

	for y := expected.Bounds().Min.Y; y < expected.Bounds().Max.Y; y++ {
		for x := expected.Bounds().Min.X; x < expected.Bounds().Max.X; x++ {
			assert.Equal(t, grayAt(expected, x, y), grayAt(actual, x, y))
		}
	}
}
