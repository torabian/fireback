package workspaces

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// Helper function to decode image
func decodeImage(file string) (image.Image, string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	var img image.Image
	var format string

	switch {
	case strings.HasSuffix(file, ".jpg"), strings.HasSuffix(file, ".jpeg"):
		img, err = jpeg.Decode(f)
		format = "jpeg"
	case strings.HasSuffix(file, ".gif"):
		img, err = gif.Decode(f)
		format = "gif"
	case strings.HasSuffix(file, ".bmp"):
		img, err = bmp.Decode(f)
		format = "bmp"
	case strings.HasSuffix(file, ".tiff"), strings.HasSuffix(file, ".tif"):
		img, err = tiff.Decode(f)
		format = "tiff"
	default:
		img, err = png.Decode(f)
		format = "png"
	}
	return img, format, err
}

// Helper function to save webp image
func saveWebP(img image.Image, file string, quality uint) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Use webp.Encode to write the image to file with specified quality
	return webp.Encode(f, img, &webp.Options{Lossless: false, Quality: float32(quality)})
}

// Helper function to resize image
func resizeImage(img image.Image, width uint, height uint) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

func ConvertToWebp(input string, distPath string, size ImageCropSize) error {

	// Decode the image
	img, format, err := decodeImage(input)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return err
	}

	fmt.Printf("Decoded %s image\n", format)

	// Convert and save images in different sizes
	resizedImg := resizeImage(img, uint(size.Width), uint(size.Height))
	quality := uint(90)
	if size.Quality > 0 {
		quality = size.Quality
	}
	err = saveWebP(resizedImg, distPath, quality)
	if err != nil {
		fmt.Println("Error saving webp image:", err)
	}

	return nil
}
