package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func (c *Cli) AuthLnurl(lnurl string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	inputType, err := breez_sdk.ParseInput(lnurl)
	if err != nil {
		return err
	}

	switch lnurlInput := inputType.(type) {
	case breez_sdk.InputTypeLnUrlAuth:
		response, err := c.sdk.LnurlAuth(lnurlInput.Data)
		if err != nil {
			return err
		}

		switch response.(type) {
		case breez_sdk.LnUrlCallbackStatusOk:
			c.PrintSuccess("Authenticated")
		}

		c.PrettyPrint(response)
	default:
		return fmt.Errorf("not a LNURL-auth")
	}

	return nil
}

func (c *Cli) PayLnurl(lnurl string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	inputType, err := breez_sdk.ParseInput(lnurl)
	if err != nil {
		return err
	}

	switch lnurlInput := inputType.(type) {
	case breez_sdk.InputTypeLnUrlPay:
		amountMsat, err := c.PromptUint64("Amount to pay in millisatoshi (min %v, max %v): ", lnurlInput.Data.MinSendable, lnurlInput.Data.MaxSendable)
		if err != nil {
			return err
		}

		comment, err := c.PromptNil("Comment (optional): ")
		if err != nil {
			return err
		}

		response, err := c.sdk.PayLnurl(breez_sdk.LnUrlPayRequest{
			Data:       lnurlInput.Data,
			AmountMsat: amountMsat,
			Comment:    comment,
		})
		if err != nil {
			return err
		}

		switch response.(type) {
		case breez_sdk.LnUrlPayResultEndpointSuccess:
			c.PrintSuccess("Payment sent")
		}

		c.PrettyPrint(response)
	default:
		return fmt.Errorf("not a LNURL-pay")
	}

	return nil
}

func (c *Cli) WithdrawLnurl(lnurl string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	inputType, err := breez_sdk.ParseInput(lnurl)
	if err != nil {
		return err
	}

	switch lnurlInput := inputType.(type) {
	case breez_sdk.InputTypeLnUrlWithdraw:
		amountMsat, err := c.PromptUint64("Amount to withdraw in millisatoshi (min %v, max %v): ", lnurlInput.Data.MinWithdrawable, lnurlInput.Data.MaxWithdrawable)
		if err != nil {
			return err
		}

		description, err := c.PromptNil("Description (optional): ")
		if err != nil {
			return err
		}

		response, err := c.sdk.WithdrawLnurl(breez_sdk.LnUrlWithdrawRequest{
			Data:        lnurlInput.Data,
			AmountMsat:  amountMsat,
			Description: description,
		})
		if err != nil {
			return err
		}

		switch response.(type) {
		case breez_sdk.LnUrlWithdrawResultOk:
			c.PrintSuccess("Payment received")
		}

		c.PrettyPrint(response)
	default:
		return fmt.Errorf("not a LNURL-withdraw")
	}

	return nil
}
