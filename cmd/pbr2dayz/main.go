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
	if len(args) != 5 {
		return fmt.Errorf("usage: pbr2dayz <basecolor> <normal> <ao> <metallic> <roughness>")
	}

	baseColorPath := args[0]
	baseColor, err := loadImage(baseColorPath)
	if err != nil {
		return err
	}

	normal, err := loadImage(args[1])
	if err != nil {
		return err
	}

	ao, err := loadImage(args[2])
	if err != nil {
		return err
	}

	metallic, err := loadImage(args[3])
	if err != nil {
		return err
	}

	roughness, err := loadImage(args[4])
	if err != nil {
		return err
	}

	dayz := pbr2dayz.PBR{
		BaseColor: baseColor,
		Normal:    normal,
		AO:        ao,
		Metallic:  metallic,
		Roughness: roughness,
	}.Convert()

	outputDir := filepath.Dir(baseColorPath)
	baseName := baseNameFromPath(baseColorPath)

	if err := savePNG(filepath.Join(outputDir, baseName+"_co.png"), dayz.CO); err != nil {
		return err
	}
	if err := savePNG(filepath.Join(outputDir, baseName+"_nohq.png"), dayz.NOHQ); err != nil {
		return err
	}
	if err := savePNG(filepath.Join(outputDir, baseName+"_as.png"), dayz.AS); err != nil {
		return err
	}
	if err := savePNG(filepath.Join(outputDir, baseName+"_smdi.png"), dayz.SMDI); err != nil {
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
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
