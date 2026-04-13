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

func TestRunSavesConvertedImagesBesideBaseColor(t *testing.T) {
	tempDir := t.TempDir()

	baseColorPath := filepath.Join(tempDir, "shirt_basecolor.png")
	normalPath := filepath.Join(tempDir, "shirt_normal.png")
	aoPath := filepath.Join(tempDir, "shirt_ao.png")
	metallicPath := filepath.Join(tempDir, "shirt_metallic.png")
	roughnessPath := filepath.Join(tempDir, "shirt_roughness.png")

	writePNG(t, baseColorPath, rgbaImage(color.RGBA{R: 10, G: 20, B: 30, A: 255}))
	writePNG(t, normalPath, rgbaImage(color.RGBA{R: 40, G: 50, B: 60, A: 255}))
	writePNG(t, aoPath, grayImage(70))
	writePNG(t, metallicPath, grayImage(80))
	writePNG(t, roughnessPath, grayImage(90))

	err := run([]string{baseColorPath, normalPath, aoPath, metallicPath, roughnessPath})

	assert.NoError(t, err)
	assert.Equal(t, color.RGBA{R: 10, G: 20, B: 30, A: 255}, readRGBA(t, filepath.Join(tempDir, "shirt_basecolor_co.png")))
	assert.Equal(t, color.RGBA{R: 40, G: 50, B: 60, A: 255}, readRGBA(t, filepath.Join(tempDir, "shirt_basecolor_nohq.png")))
	assert.Equal(t, color.RGBA{R: 255, G: 70, B: 255, A: 255}, readRGBA(t, filepath.Join(tempDir, "shirt_basecolor_as.png")))
	assert.Equal(t, color.RGBA{R: 255, G: 80, B: 165, A: 255}, readRGBA(t, filepath.Join(tempDir, "shirt_basecolor_smdi.png")))
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

func grayImage(value uint8) image.Image {
	img := image.NewGray(image.Rect(0, 0, 1, 1))
	img.SetGray(0, 0, color.Gray{Y: value})
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
