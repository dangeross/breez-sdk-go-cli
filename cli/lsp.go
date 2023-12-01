package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func (c *Cli) CloseLspChannels() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.CloseLspChannels()
	if err != nil {
		return err
	}

	c.PrintSuccess("Lsp channels closing")
	return nil
}

func (c *Cli) ConnectLsp(lspId string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.ConnectLsp(lspId)
	if err != nil {
		return err
	}

	c.PrintSuccess("Lsp connected")
	return nil
}

func (c *Cli) LspInfo() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	lspInformation, err := c.sdk.LspInfo()
	if err != nil {
		return err
	}

	c.PrettyPrint(lspInformation)
	return nil
}

func (c *Cli) ListLsps() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.ListLsps()
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) OpenChannelFee(amountMsat uint64) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.OpenChannelFee(breez_sdk.OpenChannelFeeRequest{
		AmountMsat: amountMsat,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}
