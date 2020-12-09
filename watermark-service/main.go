package main

import (
	"github.com/watermark-services/watermark-service/cmd"
	"github.com/watermark-services/watermark-service/internal/utils"
)

func main() {

	cmd.InitializeCobraViper()
	utils.InitializeLogger()
	utils.Log.Info("Application logger initialized, starting application")
	cmd.Execute()
}
