package configutil

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Setup(cmd *cobra.Command, requireFile bool) *viper.Viper {
	cfg := viper.GetViper()
	configName, err := cmd.Flags().GetString("config")
	if err != nil {
		logrus.WithError(err).Debug("Failed to read config flag")
	} else {
		cfg.SetConfigName(configName)
	}
	cfg.SetConfigType("json")
	// Add working dir
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg.AddConfigPath(dir)
	// Add binary dir
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	cfg.AddConfigPath(dir)
	cfg.BindPFlags(cmd.Flags())

	err = cfg.ReadInConfig()
	if err != nil && requireFile {
		logrus.Fatal(errors.New(err.Error() + ": Failed to read config file"))
	}
	return cfg
}
