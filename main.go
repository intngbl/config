package config

import (
	"flag"
	"github.com/revel/config"
	"log"
	"os"
	"strings"
)

const (
	pathSeparator  = string(os.PathSeparator)
	configBaseName = `app.cfg`
)

var Config *config.Config

var configFile string

func init() {
	var err error

	flagConfigFile := flag.String("c", "", "Path to app.conf file. If not provided will look for parents")

	if *flagConfigFile == "" {
		// User didn't provide a settings file path.
		if configFile, err = lookForConfigFile(); err != nil {
			// No app.conf found in this or any of the parents
			panic(`No settings file found in this or any of the parent directories: ` + err.Error())
		}
	} else {
		configFile = *flagConfigFile
	}

	// Attempt to read config file.
	if Config, err = config.ReadDefault(configFile); err != nil {
		panic(`Could not read settings file "` + configFile + `": ` + err.Error())
	}

	log.Printf("Loaded configuration file at: " + configFile)
}

func lookForConfigFile() (fileName string, err error) {

	if configFile == "" {
		var pwd string
		var testFileName string

		if pwd, err = os.Getwd(); err != nil {
			return
		}

		// Splitting path by using slashes.
		chunks := strings.Split(pwd, pathSeparator)

		// Looking for file in parent directories.
		for i := len(chunks); i > 0; i-- {
			testFileName = strings.Join(chunks[0:i], pathSeparator) + pathSeparator + configBaseName
			// File exists.
			_, err = os.Stat(testFileName)
			if err == nil {
				return testFileName, nil
			}
		}

		return "", ErrConfigFileNotFound
	}

	return configFile, nil
}
