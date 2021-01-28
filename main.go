package main

import (
	"github.com/spf13/viper"
	"github.com/watermark-service/cmd"
	"github.com/watermark-service/internal/logger"
)

func main() {
	cmd.Execute()
	logger.Log.Info(viper.AllSettings())
}
