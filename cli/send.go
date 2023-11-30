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

func (c *Cli) SendSpontaneousPayment(nodeId string, amountMsat uint64) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized")
	}

	request := breez_sdk.SendSpontaneousPaymentRequest{
		NodeId:     nodeId,
		AmountMsat: amountMsat,
	}

	response, err := c.sdk.SendSpontaneousPayment(request)

	if err != nil {
		return err
	}

	c.PrintSuccess("Payment sent")
	c.PrettyPrint(response.Payment)
	return nil
}

func (c *Cli) SendOnchain(amountSat uint64, address string, satPerVByte uint32) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized")
	}

	pairInfo, err := c.sdk.FetchReverseSwapFees(breez_sdk.ReverseSwapFeesRequest{
		SendAmountSat: &amountSat,
	})

	if err != nil {
		return err
	}

	response, err := c.sdk.SendOnchain(breez_sdk.SendOnchainRequest{
		AmountSat:               amountSat,
		OnchainRecipientAddress: address,
		PairHash:                pairInfo.FeesHash,
		SatPerVbyte:             satPerVByte,
	})

	if err != nil {
		return err
	}

	c.PrettyPrint(response.ReverseSwapInfo)
	return nil
}
