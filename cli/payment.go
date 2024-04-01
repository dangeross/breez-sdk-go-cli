package cli

import (
	"fmt"

	"github.com/breez/breez-sdk-go/breez_sdk"
	"github.com/dangeross/breez-sdk-go-cli/internal/util"
)

func (c *Cli) ListPayments(filters *[]string, fromTimestamp, toTimestamp *int64, includeFailures *bool, offset, limit *uint32) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	paymentTypeFilter, err := util.AsPaymentTypeFilterList(filters)
	if err != nil {
		return err
	}

	response, err := c.sdk.ListPayments(breez_sdk.ListPaymentsRequest{
		Filters:         paymentTypeFilter,
		FromTimestamp:   fromTimestamp,
		ToTimestamp:     toTimestamp,
		IncludeFailures: includeFailures,
		Offset:          offset,
		Limit:           limit,
	})
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) PaymentByHash(paymentHash string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	response, err := c.sdk.PaymentByHash(paymentHash)
	if err != nil {
		return err
	}

	c.PrettyPrint(response)
	return nil
}

func (c *Cli) SetPaymentMetadata(paymentHash, metadata string) error {
	if c.sdk == nil {
		return fmt.Errorf("SDK is not initialized. Try 'connect'")
	}

	err := c.sdk.SetPaymentMetadata(paymentHash, metadata)
	if err != nil {
		return err
	}

	return nil
}
