package pbr2dayz

import (
	"image"
	"image/color"
)

type PBR struct {
	BaseColor image.Image
	Normal    image.Image
	AO        image.Image
	Metallic  image.Image
	Roughness image.Image
}

type DayZ struct {
	CO   image.Image
	NOHQ image.Image
	AS   image.Image
	SMDI image.Image
}

// Convert transforms a PBR texture set into the DayZ texture layout.
func (p PBR) Convert() DayZ {
	targetBounds := firstAvailableBounds(p.BaseColor, p.Normal, p.AO, p.Metallic, p.Roughness)

	return DayZ{
		CO:   convertToOpaqueRGB(p.BaseColor, targetBounds),
		NOHQ: convertToOpaqueRGB(p.Normal, targetBounds),
		AS:   packAS(p.AO, targetBounds),
		SMDI: packSMDI(p.Metallic, p.Roughness, targetBounds),
	}
}

// firstAvailableBounds returns the first non-empty image bounds normalized to the origin.
func firstAvailableBounds(images ...image.Image) image.Rectangle {
	for _, src := range images {
		if src == nil {
			continue
		}

		bounds := src.Bounds()
		if bounds.Empty() {
			continue
		}

		return image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	}

	return image.Rectangle{}
}

// convertToOpaqueRGB copies RGB data into a fully opaque RGBA image.
func convertToOpaqueRGB(src image.Image, targetBounds image.Rectangle) image.Image {
	if src == nil || targetBounds.Empty() {
		return nil
	}

	dst := image.NewRGBA(targetBounds)
	for y := targetBounds.Min.Y; y < targetBounds.Max.Y; y++ {
		for x := targetBounds.Min.X; x < targetBounds.Max.X; x++ {
			sampled := sampleRGBA(src, x, y, targetBounds)
			dst.SetRGBA(x, y, color.RGBA{R: sampled.R, G: sampled.G, B: sampled.B, A: 0xff})
		}
	}

	return dst
}

// packAS builds the AS texture with ambient occlusion stored in the green channel.
func packAS(ao image.Image, targetBounds image.Rectangle) image.Image {
	if ao == nil || targetBounds.Empty() {
		return nil
	}

	dst := image.NewRGBA(targetBounds)
	for y := targetBounds.Min.Y; y < targetBounds.Max.Y; y++ {
		for x := targetBounds.Min.X; x < targetBounds.Max.X; x++ {
			ambientOcclusion := sampleGray(ao, x, y, targetBounds)
			dst.SetRGBA(x, y, color.RGBA{R: 0xff, G: ambientOcclusion, B: 0xff, A: 0xff})
		}
	}

	return dst
}

// packSMDI builds the SMDI texture from metallic and roughness inputs.
func packSMDI(metallic image.Image, roughness image.Image, targetBounds image.Rectangle) image.Image {
	if metallic == nil || roughness == nil || targetBounds.Empty() {
		return nil
	}

	dst := image.NewRGBA(targetBounds)
	for y := targetBounds.Min.Y; y < targetBounds.Max.Y; y++ {
		for x := targetBounds.Min.X; x < targetBounds.Max.X; x++ {
			metalness := sampleGray(metallic, x, y, targetBounds)
			gloss := 0xff - sampleGray(roughness, x, y, targetBounds)
			dst.SetRGBA(x, y, color.RGBA{R: 0xff, G: metalness, B: gloss, A: 0xff})
		}
	}

	return dst
}

// sampleRGBA reads a color from the source image mapped to the destination coordinates.
func sampleRGBA(src image.Image, dstX int, dstY int, targetBounds image.Rectangle) color.NRGBA {
	srcX, srcY := mapDestinationToSource(src.Bounds(), dstX, dstY, targetBounds)
	return color.NRGBAModel.Convert(src.At(srcX, srcY)).(color.NRGBA)
}

// sampleGray reads a grayscale value from the source image mapped to the destination coordinates.
func sampleGray(src image.Image, dstX int, dstY int, targetBounds image.Rectangle) uint8 {
	srcX, srcY := mapDestinationToSource(src.Bounds(), dstX, dstY, targetBounds)
	return color.GrayModel.Convert(src.At(srcX, srcY)).(color.Gray).Y
}

// mapDestinationToSource scales destination coordinates into source image coordinates.
func mapDestinationToSource(srcBounds image.Rectangle, dstX int, dstY int, dstBounds image.Rectangle) (int, int) {
	srcX := scaleCoordinate(dstX-dstBounds.Min.X, dstBounds.Dx(), srcBounds.Min.X, srcBounds.Dx())
	srcY := scaleCoordinate(dstY-dstBounds.Min.Y, dstBounds.Dy(), srcBounds.Min.Y, srcBounds.Dy())
	return srcX, srcY
}

// scaleCoordinate maps a destination axis position to the nearest source axis position.
func scaleCoordinate(dstPos int, dstSize int, srcMin int, srcSize int) int {
	if dstSize <= 1 || srcSize <= 1 {
		return srcMin
	}

	scaled := ((2*dstPos + 1) * srcSize) / (2 * dstSize)
	if scaled >= srcSize {
		scaled = srcSize - 1
	}

	return srcMin + scaled
}
