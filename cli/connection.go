package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func (c *Cli) Connect(inviteCode *string) error {
	if c.sdk != nil {
		return fmt.Errorf("SDK is already initialized")
	}

	mnemonic, err := c.getMnemonic()
	if err != nil {
		return err
	}

	seed, err := breez_sdk.MnemonicToSeed(mnemonic)
	if err != nil {
		return err
	}

	nodeConfig := breez_sdk.NodeConfigGreenlight{
		Config: breez_sdk.GreenlightNodeConfig{
			InviteCode: inviteCode,
		},
	}

	config := breez_sdk.DefaultConfig(breez_sdk.EnvironmentTypeProduction, c.config.ApiKey, nodeConfig)
	config.WorkingDir = c.dataDir

	if c.sdk, err = breez_sdk.Connect(config, seed, c); err != nil {
		return err
	}

	c.Println("Node was registered successfully")
	return nil
}

func (c *Cli) Disconnect() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is already disconnected")
	}

	err := c.sdk.Disconnect()
	if err != nil {
		return err
	}

	c.sdk.Destroy()
	c.sdk = nil

	c.Println("Node was stopped successfully")
	return nil
}

func (c *Cli) SetApiKey(apiKey string) error {
	c.config.ApiKey = apiKey
	c.writeConfig(*c.config)

	return nil
}
