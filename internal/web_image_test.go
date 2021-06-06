package internal

import (
	"os"
	"testing"
)

func TestDownWebImage(t *testing.T) {
	url := `https://upload-images.jianshu.io/upload_images/7535793-f95a964979b61050.png`
	dir := `./`
	imagePath, err := DownWebImage(url, dir)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = os.Stat(imagePath)
	if err != nil {
		t.Error(err.Error())
	}

	os.Remove(imagePath)
}

func TestDownWebImageWithoutExt(t *testing.T) {
	url := `https://upload-images.jianshu.io/upload_images/24630328-df0561b8c1d131a1.image`
	dir := `./`
	imagePath, err := DownWebImage(url, dir)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = os.Stat(imagePath)
	if err != nil {
		t.Error(err.Error())
	}

	os.Remove(imagePath)
}
