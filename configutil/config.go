package configutil

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configParam default is "config.yml", empty string will not load a file
func Setup(cmd *cobra.Command, configParam string) (*viper.Viper, error) {
	cfg := viper.GetViper()

	cfg.AutomaticEnv()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(
		"-", "_",
		".", "__", // TODO: switch to . -> _
	))

	if cmd != nil {
		cfg.BindPFlags(cmd.Flags())
	}

	if cmd != nil && configParam == "" {
		configParam, _ = cmd.Flags().GetString("config")
	}

	hasFileConfig := configParam != ""

	if hasFileConfig {
		fileName := configParam
		configType := ""

		dir := filepath.Dir(configParam)
		fileName = filepath.Base(configParam)

		configType = filepath.Ext(configParam)
		configType = strings.TrimLeft(configType, ".")
		fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) // Must be without extension

		if configType == "" {
			configType = "yml"
		}

		cfg.SetConfigName(fileName)
		cfg.SetConfigType(configType)

		cfg.AddConfigPath(dir)
		logrus.WithField("dir", dir).WithField("file", fileName+"."+configType).WithField("type", configType).Info("Loading config")
		// Add binary dir
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		cfg.AddConfigPath(dir)

		dir, err = filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), "/config"))
		if err != nil {
			panic(err)
		}
		cfg.AddConfigPath(dir)

		err = cfg.ReadInConfig()
		if err != nil {
			return cfg, nil
		}
	}

	return cfg, nil
}
