package watermark

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
)

// PlaceImage to place image
func PlaceImage(outName, bgImg, markImg, markDimensions, locationDimensions string) {
	//lX, lY := ParseCoordinates(locationDimensions, "x")

	//src := OpenImage(bgImg)

}

// OpenImage to open the image
func OpenImage(name string) image.Image {
	src, err := imaging.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	return src
}
