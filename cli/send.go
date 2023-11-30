package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func (c *Cli) SendPayment(bolt11 string, amountMsat *uint64) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized")
	}

	request := breez_sdk.SendPaymentRequest{
		Bolt11:     bolt11,
		AmountMsat: amountMsat,
	}

	response, err := c.sdk.SendPayment(request)

	if err != nil {
		return err
	}

	c.PrintSuccess("Payment sent")
	c.PrettyPrint(response.Payment)
	return nil
}
