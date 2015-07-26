package gfxBot

import (
	"os"
)

type Image struct {
	Data	[]byte
	Width	int
	Height	int
	Caption	string
	Ext	string
}

func NewImage(ext string, caption string) (*Image, error) {
	return &Image{Ext: ext, Caption: caption}, nil
}

func (img *Image) Save(path string) (error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(img.Data)
	if err != nil {
		return err
	}

	return nil
}

func (img *Image) Write(p []byte) (n int, err error) {
	if img.Data == nil {
		img.Data = make([]byte, len(p))
		copy(img.Data, p)
	} else {
		img.Data = append(img.Data, p...)
	}
	return len(p), nil
}

func (img *Image) Close() (err error) {
	return
}
