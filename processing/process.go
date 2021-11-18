package processing

import (
	"image"
	"image/jpeg"
	"os"
)

func StripMetadata(src, dest string) error {
	var failed error

	origFile, err := os.Open(src)
	if err != nil {
		failed = err
	}
	defer origFile.Close()

	img, _, err := image.Decode(origFile)
	if err != nil {
		failed = err
	}

	outfile, err := os.Create(dest)
	if err != nil {
		failed = err
	}
	defer outfile.Close()

	if err := jpeg.Encode(outfile, img, nil); err != nil {
		failed = err
	}

	return failed
}
