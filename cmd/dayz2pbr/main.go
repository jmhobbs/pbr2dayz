package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	_ "image/gif"
	_ "image/jpeg"

	"pbr2dayz"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 4 {
		return fmt.Errorf("usage: dayz2pbr <co> <nohq> <as> <smdi>")
	}

	coPath := args[0]
	co, err := loadImage(coPath)
	if err != nil {
		return err
	}

	nohq, err := loadImage(args[1])
	if err != nil {
		return err
	}

	as, err := loadImage(args[2])
	if err != nil {
		return err
	}

	smdi, err := loadImage(args[3])
	if err != nil {
		return err
	}

	pbr := pbr2dayz.DayZ{
		CO:   co,
		NOHQ: nohq,
		AS:   as,
		SMDI: smdi,
	}.Convert()

	outputDir := filepath.Dir(coPath)
	baseName := baseNameFromPath(coPath)

	if err := savePNG(filepath.Join(outputDir, baseName+"_basecolor.png"), pbr.BaseColor); err != nil {
		return err
	}
	if err := savePNG(filepath.Join(outputDir, baseName+"_normal.png"), pbr.Normal); err != nil {
		return err
	}
	if err := savePNG(filepath.Join(outputDir, baseName+"_ao.png"), pbr.AO); err != nil {
		return err
	}
	if err := savePNG(filepath.Join(outputDir, baseName+"_metallic.png"), pbr.Metallic); err != nil {
		return err
	}
	if err := savePNG(filepath.Join(outputDir, baseName+"_roughness.png"), pbr.Roughness); err != nil {
		return err
	}

	return nil
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}

	return img, nil
}

func savePNG(path string, img image.Image) error {
	if img == nil {
		return fmt.Errorf("save %s: %w", path, errors.New("image is nil"))
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("encode %s: %w", path, err)
	}

	return nil
}

func baseNameFromPath(path string) string {
	fileName := filepath.Base(path)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	for _, suffix := range []string{"_co", "_nohq", "_as", "_smdi"} {
		if strings.HasSuffix(baseName, suffix) {
			return strings.TrimSuffix(baseName, suffix)
		}
	}

	return baseName
}
