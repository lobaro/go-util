package configutil

import (
	"fmt"
	"github.com/lobaro/go-util/fileutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configFile default is "config.yml", empty string will try to get the configFile from "config" cmd flag
func Setup(cmd *cobra.Command, configFile string) (*viper.Viper, error) {
	cfg := viper.GetViper()

	cfg.AutomaticEnv()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(
		"-", "_",
		".", "__", // TODO: switch to . -> _
	))

	if cmd != nil {
		cfg.BindPFlags(cmd.Flags())
	}

	if cmd != nil && configFile == "" {
		configFile, _ = cmd.Flags().GetString("config")
	}

	hasFileConfig := configFile != ""

	if hasFileConfig {
		fileName := configFile
		configType := ""

		dir := filepath.Dir(configFile)
		fileName = filepath.Base(configFile)

		configType = filepath.Ext(configFile)
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

		if cfg.ConfigFileUsed() != "" &&  !fileutil.MustExists(cfg.ConfigFileUsed()) {
			return cfg, fmt.Errorf("config file does not exist: %s", cfg.ConfigFileUsed())
		}

		err = cfg.ReadInConfig()
		if err != nil {
			return cfg, nil
		}
	}

	return cfg, nil
}
