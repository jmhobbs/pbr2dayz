package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunSavesConvertedImagesBesideCO(t *testing.T) {
	tempDir := t.TempDir()

	coPath := filepath.Join(tempDir, "shirt_co.png")
	nohqPath := filepath.Join(tempDir, "shirt_nohq.png")
	asPath := filepath.Join(tempDir, "shirt_as.png")
	smdiPath := filepath.Join(tempDir, "shirt_smdi.png")

	writePNG(t, coPath, rgbaImage(color.RGBA{R: 10, G: 20, B: 30, A: 255}))
	writePNG(t, nohqPath, rgbaImage(color.RGBA{R: 40, G: 50, B: 60, A: 255}))
	writePNG(t, asPath, rgbaImage(color.RGBA{R: 255, G: 70, B: 255, A: 255}))
	writePNG(t, smdiPath, rgbaImage(color.RGBA{R: 255, G: 80, B: 165, A: 255}))

	err := run([]string{coPath, nohqPath, asPath, smdiPath})

	assert.NoError(t, err)
	assert.Equal(t, color.RGBA{R: 10, G: 20, B: 30, A: 255}, readRGBA(t, filepath.Join(tempDir, "shirt_basecolor.png")))
	assert.Equal(t, color.RGBA{R: 40, G: 50, B: 60, A: 255}, readRGBA(t, filepath.Join(tempDir, "shirt_normal.png")))
	assert.Equal(t, uint8(70), readGray(t, filepath.Join(tempDir, "shirt_ao.png")))
	assert.Equal(t, uint8(80), readGray(t, filepath.Join(tempDir, "shirt_metallic.png")))
	assert.Equal(t, uint8(90), readGray(t, filepath.Join(tempDir, "shirt_roughness.png")))
}

func TestRunRejectsWrongArgumentCount(t *testing.T) {
	err := run([]string{"only-one"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "usage:")
}

func rgbaImage(pixel color.RGBA) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.SetRGBA(0, 0, pixel)
	return img
}

func writePNG(t *testing.T, path string, img image.Image) {
	t.Helper()

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("create %s: %v", path, err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		t.Fatalf("encode %s: %v", path, err)
	}
}

func readRGBA(t *testing.T, path string) color.RGBA {
	t.Helper()

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}

	return color.RGBAModel.Convert(img.At(0, 0)).(color.RGBA)
}

func readGray(t *testing.T, path string) uint8 {
	t.Helper()

	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}

	return color.GrayModel.Convert(img.At(0, 0)).(color.Gray).Y
}
