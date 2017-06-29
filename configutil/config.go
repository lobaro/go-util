package configutil

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Setup(cmd *cobra.Command, requireFile bool) *viper.Viper {
	cfg := viper.GetViper()
	configParam, err := cmd.Flags().GetString("config")
	fileName := configParam
	configType := ""

	dir := filepath.Dir(configParam)
	fileName = filepath.Base(configParam)

	configType = filepath.Ext(configParam)
	configType = strings.TrimLeft(configType, ".")
	fileName = strings.TrimSuffix(fileName, configType) // Must be without extenstion

	if configType == "" {
		configType = "json"
	}

	if err != nil {
		logrus.WithError(err).Debug("Failed to read config flag")
	} else {
		cfg.SetConfigName(fileName)
	}
	cfg.SetConfigType(configType)

	cfg.AddConfigPath(dir)
	logrus.WithField("dir", dir).WithField("file", fileName+"."+configType).WithField("type", configType).Info("Loading config")
	// Add binary dir
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	cfg.AddConfigPath(dir)
	cfg.BindPFlags(cmd.Flags())

	err = cfg.ReadInConfig()
	if err != nil {
		if requireFile {
			logrus.WithError(err).Fatal("Failed to read config file")
		} else {
			logrus.WithError(err).Info("Failed to read config file")
		}
	}
	return cfg
}
