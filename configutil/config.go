package configutil

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Setup(cmd *cobra.Command) (*viper.Viper, error) {
	cfg := viper.GetViper()
	configParam := "config.yaml"
	if cmd != nil {
		configParam, _ = cmd.Flags().GetString("config")
	}
	fileName := configParam
	configType := ""

	dir := filepath.Dir(configParam)
	fileName = filepath.Base(configParam)

	configType = filepath.Ext(configParam)
	configType = strings.TrimLeft(configType, ".")
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) // Must be without extension

	if configType == "" {
		configType = "json"
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

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(
		"-", "_",
		".", "__",
	))

	// Read in config

	if cmd != nil {
		cfg.BindPFlags(cmd.Flags())
	}

	err = cfg.ReadInConfig()

	return cfg, err
}
