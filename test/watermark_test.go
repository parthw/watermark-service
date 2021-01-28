package watermark_test

import (
	"os"
	"testing"

	"github.com/watermark-service/internal/logger"
	"github.com/watermark-service/internal/service"
)

func TestWatermarkPackageFunctions(t *testing.T) {
	logger.InitializeLogger()
	outputFile := "testdata/test-final-image.png"
	err := service.AddWaterMark("testdata/test-bg.png", "testdata/test-logo.png", outputFile)
	if err != nil {
		t.Errorf("Failed to create image with error %v - ", err)
	}
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Output image %v does not exist", outputFile)
	}
	os.Remove(outputFile)
}
