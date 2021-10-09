package helper

import "image"
import "image/draw"

// Image2RGBA 图片转RGBA格式 /**
func Image2RGBA(img image.Image) *image.RGBA {
	baseSrcBounds := img.Bounds().Max
	des := image.NewRGBA(image.Rect(0, 0, baseSrcBounds.X, baseSrcBounds.Y))
	draw.Draw(des, des.Bounds(), img, img.Bounds().Min, draw.Over)
	return des
}
