package config

import (
	"os"
)

type draw struct {
	FontPath string  // 字体路径
	FontDpi  float64 // 分辨率
	FontSize float64 // 字体大小
}

var Draw draw

func init() {
	Draw.FontPath = os.Getenv("SG_DRAW_FONT_PATH")
	Draw.FontDpi = 72
	Draw.FontSize = 25
}
