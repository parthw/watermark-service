package watermark_test

import (
	"os"
	"testing"

	watermark "github.com/watermark-services/watermark-service/internal"
	"github.com/watermark-services/watermark-service/internal/utils"
)

func TestWatermarkPackageFunctions(t *testing.T) {
	utils.InitializeLogger()
	outputFile := "testdata/test-final-image.png"
	err := watermark.AddWaterMark("testdata/test-bg.png", "testdata/test-logo.png", outputFile)
	if err != nil {
		t.Errorf("Failed to create image with error %v - ", err)
	}
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Output image %v does not exist", outputFile)
	}
	os.Remove(outputFile)
}
