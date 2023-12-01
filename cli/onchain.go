package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
)

func (c *Cli) ListRefundables() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.ListRefundables()
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) PrepareRefund(swapAddress, toAddress string, satPerVByte uint32) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.PrepareRefund(breez_sdk.PrepareRefundRequest{
		SwapAddress: swapAddress,
		ToAddress:   toAddress,
		SatPerVbyte: satPerVByte,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) PrepareSweep(toAddress string, satPerVByte uint64) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.PrepareSweep(breez_sdk.PrepareSweepRequest{
		ToAddress:   toAddress,
		SatPerVbyte: satPerVByte,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) RecommendedFees() error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.RecommendedFees()
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) Refund(swapAddress, toAddress string, satPerVByte uint32) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.Refund(breez_sdk.RefundRequest{
		SwapAddress: swapAddress,
		ToAddress:   toAddress,
		SatPerVbyte: satPerVByte,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) Sweep(toAddress string, satPerVByte uint32) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.Sweep(breez_sdk.SweepRequest{
		ToAddress:   toAddress,
		SatPerVbyte: satPerVByte,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}
