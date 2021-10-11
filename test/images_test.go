package test

import (
	"fmt"
	"go-file-system"
	"os"
	"strconv"
	"testing"
)

func Test_Clip(t *testing.T) {
	src := "../_file/7/0/05d014c303e86ad1c7ae1de9876efedd"
	dst := src + "_150x150"
	fmt.Println("src=", src)
	fmt.Println("dst=", dst)
	fIn, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer fIn.Close()
	fOut, _ := os.Create(dst)
	defer fOut.Close()
	err = main.Clip(fIn, fOut, 0, 0, 150, 150, 100)
	if err != nil {
		panic(err)
	}
}

func Test_Scale(t *testing.T) {
	width := 240
	height := 240
	quality := 50
	src := "../_file/7/0/05d014c303e86ad1c7ae1de9876efedd"
	dst := src + "_scalex" + strconv.Itoa(width) + "x" + strconv.Itoa(height) + "quality=" + strconv.Itoa(quality)
	fmt.Println("src=", src)
	fmt.Println("dst=", dst)
	fIn, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer fIn.Close()
	fOut, _ := os.Create(dst)
	defer fOut.Close()
	err = main.Scale(fIn, fOut, width, height, quality)
	if err != nil {
		panic(err)
	}
}
