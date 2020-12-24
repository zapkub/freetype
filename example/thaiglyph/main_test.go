package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func ExampleThaiGlyph() {
	wd, _ := os.Getwd()
	promptttfbytes, _ := ioutil.ReadFile(path.Join(wd, "/prompt-regular.ttf"))
	promptttf, err := truetype.Parse(promptttfbytes)
	if err != nil {
		log.Fatal(err)
	}
	var tmpdir = path.Join(wd, "tmp")
	os.MkdirAll(tmpdir, 0755)

	file, err := os.Create(path.Join(tmpdir, "drawtext.png"))
	if err != nil {
		panic(err)
	}
	img := image.NewRGBA(image.Rect(0, 0, 1024, 800))
	c := freetype.NewContext()
	c.SetFont(promptttf)
	c.SetDPI(144)
	c.SetFontSize(18)
	c.SetDst(img)
	c.SetSrc(image.White)
	c.SetClip(image.Rect(0, 0, 800, 400))
	pt := freetype.Pt(0, 0+int(c.PointToFixed(18)>>6))
	c.DrawString("มั้ย น้ำ ฟัน ดั้นด้น", pt)
	png.Encode(file, img)
	fmt.Println("done")
	// output: done
}
