package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dangeross/breez-sdk-go-cli/internal/util"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ApiKey   string `json:"api_key"`
	LogLevel string `json:"log_level"`
}

func (c *Cli) readConfig() error {
	configFilePath := filepath.Join(c.dataDir, "config.json")
	config := Config{
		LogLevel: "debug",
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		b, err := json.Marshal(config)
		if err != nil {
			return err
		}

		if err = os.WriteFile(configFilePath, b, 0644); err != nil {
			return fmt.Errorf("could not write config file")
		}

		c.config = &config
		return nil
	}

	b, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &config); err != nil {
		return err
	}

	c.config = &config
	return nil
}

func (c *Cli) getMnemonic() (string, error) {
	phraseFilePath := filepath.Join(c.dataDir, "phrase")
	if _, err := os.Stat(phraseFilePath); os.IsNotExist(err) {
		mnemonic, err := util.GenerateMnemonic()
		if err != nil {
			return "", err
		}

		if err = os.WriteFile(phraseFilePath, []byte(mnemonic), 0644); err != nil {
			return "", fmt.Errorf("could not write phrase file")
		}

		return mnemonic, nil
	}

	if b, err := os.ReadFile(phraseFilePath); err == nil {
		return string(b[:]), nil
	}

	return "", fmt.Errorf("could not read phrase file")
}

func (c *Cli) initLog() error {
	logFilePath := filepath.Join(c.dataDir, "sdk.log")

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	c.log = logrus.New()
	c.log.Out = file

	if level, err := logrus.ParseLevel(c.config.LogLevel); err != nil {
		c.log.SetLevel(level)
	}

	return nil
}