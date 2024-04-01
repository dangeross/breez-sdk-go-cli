package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func (c *Cli) Backup() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.Backup()
	if err != nil {
		return err
	}

	c.PrintSuccess("Backup started")
	return nil
}

func (c *Cli) CheckMessage(message, pubkey, signature string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.CheckMessage(breez_sdk.CheckMessageRequest{
		Message:   message,
		Pubkey:    pubkey,
		Signature: signature,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) ConfigureNode(closeToAddress *string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.ConfigureNode(breez_sdk.ConfigureNodeRequest{
		CloseToAddress: closeToAddress,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Cli) NodeCredentials() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	nodeCredentials, err := c.sdk.NodeCredentials()
	if err != nil {
		return err
	}

	c.PrettyPrint(nodeCredentials)
	return nil
}

func (c *Cli) NodeInfo() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	nodeState, err := c.sdk.NodeInfo()
	if err != nil {
		return err
	}

	c.PrettyPrint(nodeState)
	return nil
}

func (c *Cli) RegisterWebhook(url string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.RegisterWebhook(url)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cli) SignMessage(message string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.SignMessage(breez_sdk.SignMessageRequest{
		Message: message,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) StaticBackup() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := breez_sdk.StaticBackup(breez_sdk.StaticBackupRequest{
		WorkingDir: c.dataDir,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) Sync() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.Sync()
	if err != nil {
		return err
	}

	c.PrintSuccess("Synced")
	return nil
}
