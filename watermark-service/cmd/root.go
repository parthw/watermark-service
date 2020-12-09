package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "watermark-service",
	Short: "To watermark images",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To watermark images")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

// InitializeCobraViper to init cobra cli and viper config
func InitializeCobraViper() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.watermark-service.yaml)")

	initStartCmd()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// setting up default config, in case no config file is present
	// if any map parameter is used, then whole map will get overwrite
	// that's why in case of log, log.file is used which viper does not overwrite
	// log.file create a internal map in result it joins it
	defaultConfig()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
		}

		// Search config in home directory with name ".watermark-service" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".watermark-service")
	}

	//viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func defaultConfig() {
	viper.SetDefault("env", "dev")
	viper.SetDefault("log.file", "watermark-service.log")
	viper.SetDefault("log.file.maxsize", "10") //megabytes
	viper.SetDefault("log.file.maxbackups", "7")
	viper.SetDefault("log.file.maxage", "7") //days
	viper.SetDefault("log.level", "info")
}
