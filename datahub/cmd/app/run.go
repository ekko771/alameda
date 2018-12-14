package app

import (
	"errors"
	"strings"

	"github.com/containers-ai/alameda/datahub"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envVarPrefix = "ALAMEDA_DATAHUB"
)

var (
	scope  *log.Scope
	config datahub.Config

	configurationFilePath string

	RunCmd = &cobra.Command{
		Use:   "run",
		Short: "start alameda datahub server",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			var (
				err error

				server *datahub.Server
			)

			initConfig()
			initLogger()

			server, err = datahub.NewServer(config)
			if err != nil {
				panic(err)
			}

			if err = server.Run(); err != nil {
				server.Stop()
				panic(err)
			}
		},
	}
)

func init() {
	parseFlag()
}

func parseFlag() {
	RunCmd.Flags().StringVar(&configurationFilePath, "config", "", "The path to datahub configuration file.")
}

func initConfig() {

	config = datahub.NewDefaultConfig()

	initViperSetting()
	mergeConfigFileValueWithDefaultConfigValue()
}

func initViperSetting() {

	viper.SetEnvPrefix(envVarPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

func mergeConfigFileValueWithDefaultConfigValue() {

	if configurationFilePath == "" {

	} else {

		viper.SetConfigFile(configurationFilePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(errors.New("Read configuration file failed: " + err.Error()))
		}
		err = viper.Unmarshal(&config)
		if err != nil {
			panic(errors.New("Unmarshal configuration failed: " + err.Error()))
		}
	}
}

func initLogger() {

	scope = log.RegisterScope("datahub", "datahub server log", 0)
}
