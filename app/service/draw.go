package service

import (
	"bytes"
	"encoding/base64"
	"github.com/golang/freetype/truetype"
	"github.com/samego-ai/draw_roi/app/helper"
	"github.com/samego-ai/draw_roi/config"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

type DrawService struct {
}

// NewDrawService
func NewDrawService() *DrawService {
	return &DrawService{}
}

// Rectangle 矩形结构体 /**
type Rectangle struct {
	X int
	Y int
	W int
	H int
}

// Brush 笔刷 /**
type Brush struct {
	FontType  *truetype.Font
	FontSize  float64
	FontColor *image.Uniform
	FontDPI   float64
}

// NewBrush 实例化笔刷 /**
func NewBrush(FontPath string, FontSize float64, FontColor *image.Uniform, fontDPI float64) (*Brush, error) {
	fontFile, err := ioutil.ReadFile(FontPath)
	if err != nil {
		return nil, err
	}
	fontType, err := truetype.Parse(fontFile)
	if err != nil {
		return nil, err
	}
	return &Brush{FontType: fontType, FontSize: FontSize, FontColor: FontColor, FontDPI: fontDPI}, nil
}

// DrawMultiRectangle 图片上画多个框
func (service *DrawService) DrawMultiRectangle(base64Image string, rectangles []Rectangle, title string) (string, error) {
	// 1.加载图片作为背景
	background, err := jpeg.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Image)))

	// 2.新建笔刷
	brush, _ := NewBrush(config.Draw.FontPath, config.Draw.FontSize, image.Black, config.Draw.FontDpi)

	// 3.图片转成RGBA
	des := helper.Image2RGBA(background)

	// 4.绘制框
	for _, rectangle := range rectangles {
		var c *freetype.Context
		c = freetype.NewContext()
		c.SetDPI(brush.FontDPI)
		c.SetFont(brush.FontType)
		c.SetHinting(font.HintingFull)
		c.SetFontSize(brush.FontSize)
		c.SetClip(des.Bounds())
		c.SetDst(des)
		cGreen := image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255})

		c.SetSrc(cGreen)
		drawStringContent := "."
		for i := rectangle.X; i < (rectangle.X + rectangle.W); i++ {
			_, _ = c.DrawString(drawStringContent, freetype.Pt(i, rectangle.Y))
			_, _ = c.DrawString(drawStringContent, freetype.Pt(i, rectangle.Y+rectangle.H))
		}
		for j := rectangle.Y; j < (rectangle.Y + rectangle.H); j++ {
			_, _ = c.DrawString(drawStringContent, freetype.Pt(rectangle.X, j))
			_, _ = c.DrawString(drawStringContent, freetype.Pt(rectangle.X+rectangle.W, j))
		}
	}

	// 5.写入文字
	c := freetype.NewContext()
	c.SetDPI(brush.FontDPI)
	c.SetFont(brush.FontType)
	c.SetHinting(font.HintingNone)
	c.SetFontSize(brush.FontSize)
	c.SetClip(des.Bounds())
	c.SetDst(des)
	c.SetSrc(brush.FontColor)
	_, _ = c.DrawString(title, freetype.Pt(30, 30))

	// 6.image 转 base64
	emptyBuff := bytes.NewBuffer(nil)                  // 开辟一个新的空buff
	_ = jpeg.Encode(emptyBuff, des, nil)               // img写入到buff
	dist := make([]byte, 1048576*10)                   // 开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) // buff转成base64

	return string(dist[0:bytes.IndexByte(dist, 0)]), err
}
