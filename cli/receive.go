package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
	"github.com/dangeross/breez-sdk-go-cli/internal/util"
	qrcode "github.com/mdp/qrterminal/v3"
)

func (c *Cli) BuyBitcoin(providerText string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	provider, err := util.AsProvider(providerText)
	if err != nil {
		return err
	}

	response, err := c.sdk.BuyBitcoin(breez_sdk.BuyBitcoinRequest{
		Provider: provider,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) ReceivePayment(amountMsat uint64, description string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	request := breez_sdk.ReceivePaymentRequest{
		AmountMsat:  amountMsat,
		Description: description,
	}

	response, err := c.sdk.ReceivePayment(request)

	if err != nil {
		return err
	}

	qrcode.GenerateHalfBlock(response.LnInvoice.Bolt11, qrcode.L, c.App)
	return nil
}

func (c *Cli) ReceiveOnchain() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.ReceiveOnchain(breez_sdk.ReceiveOnchainRequest{})

	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}
