package main

import (
	"github.com/spf13/viper"
	"github.com/watermark-services/watermark-service/cmd"
	"github.com/watermark-services/watermark-service/internal/utils"
)

func main() {
	cmd.Execute()
	utils.Log.Info(viper.AllSettings())
}
