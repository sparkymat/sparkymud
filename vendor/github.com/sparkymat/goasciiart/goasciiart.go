// by mo2zie

package asciiart

import (
	"github.com/nfnt/resize"

	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func scaleImageToAsciiWidth(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	h := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	return img, w, h
}

func Convert2AsciiOfWidth(img image.Image, width int) []byte {
	scaledImage, w, h := scaleImageToAsciiWidth(img, width)

	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(scaledImage.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}
	return buf.Bytes()
}
