package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func (c *Cli) ParseInput(s string) error {
	response, err := breez_sdk.ParseInput(s)
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) ServiceCheckCheck() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.ServiceHealthCheck()
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) ReportPaymentFailure(paymentHash string, comment *string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.ReportIssue(breez_sdk.ReportIssueRequestPaymentFailure{
		Data: breez_sdk.ReportPaymentFailureDetails{
			PaymentHash: paymentHash,
			Comment:     comment,
		},
	})
	if err != nil {
		return err
	}

	c.PrintSuccess("Report sent")
	return nil
}

func (c *Cli) ExecuteDevCommand(command string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.ExecuteDevCommand(command)
	if err != nil {
		return err
	}

	c.Println(response)
	return nil
}
