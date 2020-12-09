package watermark

import (
	"image"
	"math"

	"github.com/disintegration/imaging"
	"github.com/watermark-services/watermark-service/internal/utils"
)

// AddWaterMark to add watermark to image
func AddWaterMark(bgImgName, markImgName, outImgName string) error {

	utils.Log.Infof("Initiating add watermark process with backgroud image - %v and watermark image - %v",
		bgImgName, markImgName)
	bgImg, err := openImage(bgImgName)
	if err != nil {
		utils.Log.Error(err)
		return err
	}

	markImg, err := resizeImage(markImgName, 200, 200)
	if err != nil {
		utils.Log.Error(err)
		return err
	}

	bgDim := bgImg.Bounds().Max
	markDim := markImg.Bounds().Max

	bgAspectRatio := math.Round(float64(bgDim.X) / float64(bgDim.Y))
	bgXPos, bgYPos := calculateMarkPos(bgDim, markDim, bgAspectRatio)

	dst := imaging.Paste(bgImg, markImg, image.Pt(bgXPos, bgYPos))

	err = imaging.Save(dst, outImgName)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	return nil

}

func openImage(name string) (image.Image, error) {
	src, err := imaging.Open(name)
	if err != nil {
		return nil, err
	}
	return src, nil
}

func resizeImage(imageName string, width, height int) (image.Image, error) {
	img, err := openImage(imageName)
	if err != nil {
		return nil, err
	}
	return imaging.Fit(img, width, height, imaging.Lanczos), nil
}

func calculateMarkPos(bgDim, markDim image.Point, aspectRatio float64) (int, int) {
	bgX := bgDim.X
	bgY := bgDim.Y
	markX := markDim.X
	markY := markDim.Y

	padding := 20 * int(aspectRatio)
	return bgX - markX - padding, bgY - markY - padding
}
