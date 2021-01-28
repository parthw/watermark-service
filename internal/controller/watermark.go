package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watermark-service/internal/logger"
	"github.com/watermark-service/internal/service"
)

// PutWaterMark to save files to directory
func PutWaterMark(ctx *gin.Context) error {
	bgFile, _ := ctx.FormFile("bg")
	logoFile, _ := ctx.FormFile("logo")

	logger.Log.Infof("bgFile - %v, logoFile -%v", bgFile.Filename, logoFile.Filename)

	dst := "/Users/parth/Documents/personal-git/watermark-service/files/"
	bgImgPath := dst + bgFile.Filename
	markImgName := dst + logoFile.Filename
	err := ctx.SaveUploadedFile(bgFile, bgImgPath)
	if err != nil {
		logger.Log.Errorf("Failed to save file with err - ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	err = ctx.SaveUploadedFile(logoFile, markImgName)
	if err != nil {
		logger.Log.Error("Failed to save file with err - ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	outImgName := dst + ctx.Request.Form.Get("outname")
	err = service.AddWaterMark(bgImgPath, markImgName, outImgName)
	if err != nil {
		logger.Log.Error("Failed to create watermark image with err - ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.File(outImgName)
	return nil
}
