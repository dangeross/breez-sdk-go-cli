package cli

import (
	"strings"

	"github.com/dangeross/breez-sdk-go-cli/internal/util"
	"github.com/desertbit/grumble"
)

func (c *Cli) load() error {
	c.AddCommand(&grumble.Command{
		Name: "set_api_key",
		Help: "configure the Breez API key",
		Args: func(a *grumble.Args) {
			a.String("api_key", "api key")
		},
		Run: func(ctx *grumble.Context) (err error) {
			apiKey := ctx.Args.String("api_key")

			return c.SetApiKey(apiKey)
		},
	})

	/**
	 * Connection
	 */

	c.AddCommand(&grumble.Command{
		Name: "connect",
		Help: "initialize an SDK instance",
		Flags: func(f *grumble.Flags) {
			f.String("i", "invite_code", "", "optional greenlight invite code")
		},
		Run: func(ctx *grumble.Context) (err error) {
			inviteCode := util.NilString(ctx.Flags.String("invite_code"))

			return c.Connect(inviteCode)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "disconnect",
		Help: "stop the SDK instance",
		Run: func(ctx *grumble.Context) (err error) {
			return c.Disconnect()
		},
	})

	/**
	 * Node
	 */

	c.AddCommand(&grumble.Command{
		Name: "backup",
		Help: "start a backup of the local data",
		Run: func(ctx *grumble.Context) (err error) {
			return c.Backup()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "check_message",
		Help: "verify a signed message",
		Args: func(a *grumble.Args) {
			a.String("message", "signed message")
			a.String("pubkey", "pubkey of the signer")
			a.String("signature", "message signature")
		},
		Run: func(ctx *grumble.Context) (err error) {
			message := ctx.Args.String("message")
			pubkey := ctx.Args.String("pubkey")
			signature := ctx.Args.String("signature")

			return c.CheckMessage(message, pubkey, signature)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "configure_node",
		Help: "configure the node",
		Flags: func(f *grumble.Flags) {
			f.StringL("close_to_address", "", "address to send funds to during a mutual channel close")
		},
		Run: func(ctx *grumble.Context) (err error) {
			closeToAddress := util.NilString(ctx.Flags.String("close_to_address"))

			return c.ConfigureNode(closeToAddress)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "node_credentials",
		Help: "get the node credentials",
		Run: func(ctx *grumble.Context) (err error) {
			return c.NodeCredentials()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "node_info",
		Help: "get the latest node state",
		Run: func(ctx *grumble.Context) (err error) {
			return c.NodeInfo()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "register_webhook",
		Help: "register a webhook URL that will be called on specific events",
		Args: func(a *grumble.Args) {
			a.String("url", "webhook URL")
		},
		Run: func(ctx *grumble.Context) (err error) {
			url := ctx.Args.String("url")

			return c.RegisterWebhook(url)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "sign_message",
		Help: "sign a message",
		Args: func(a *grumble.Args) {
			a.String("message", "message to sign")
		},
		Run: func(ctx *grumble.Context) (err error) {
			message := ctx.Args.String("message")

			return c.SignMessage(message)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "static_backup",
		Help: "fetch the static backup data",
		Run: func(ctx *grumble.Context) (err error) {
			return c.StaticBackup()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "sync",
		Help: "sync local data with remote node",
		Run: func(ctx *grumble.Context) (err error) {
			return c.Sync()
		},
	})

	/**
	 * LNURL
	 */

	c.AddCommand(&grumble.Command{
		Name: "lnurl_auth",
		Help: "authenticate using LNURL-auth",
		Args: func(a *grumble.Args) {
			a.String("lnurl", "LNURL-auth request")
		},
		Run: func(ctx *grumble.Context) (err error) {
			lnurl := ctx.Args.String("lnurl")

			return c.AuthLnurl(lnurl)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "lnurl_pay",
		Help: "send payment using LNURL-pay",
		Args: func(a *grumble.Args) {
			a.String("lnurl", "LNURL-pay request")
		},
		Run: func(ctx *grumble.Context) (err error) {
			lnurl := ctx.Args.String("lnurl")

			return c.PayLnurl(lnurl)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "lnurl_withdraw",
		Help: "receive payment using LNURL-withdraw",
		Args: func(a *grumble.Args) {
			a.String("lnurl", "LNURL-withdraw request")
		},
		Run: func(ctx *grumble.Context) (err error) {
			lnurl := ctx.Args.String("lnurl")

			return c.WithdrawLnurl(lnurl)
		},
	})

	/**
	 * Receive
	 */

	c.AddCommand(&grumble.Command{
		Name: "buy_bitcoin",
		Help: "generate a URL to buy bitcoin",
		Args: func(a *grumble.Args) {
			a.String("provider", "provider name", grumble.Default("moonpay"))
		},
		Run: func(ctx *grumble.Context) (err error) {
			provider := ctx.Args.String("provider")

			return c.BuyBitcoin(provider)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "receive_payment",
		Help: "generate a bolt11 invoice",
		Args: func(a *grumble.Args) {
			a.Uint64("amount_msat", "amount to receive in millisatoshis")
			a.String("description", "payment description", grumble.Default(""))
		},
		Run: func(ctx *grumble.Context) (err error) {
			amountMsat := ctx.Args.Uint64("amount_msat")
			description := ctx.Args.String("description")

			return c.ReceivePayment(amountMsat, description)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "receive_onchain",
		Help: "receive on-chain using a swap",
		Run: func(ctx *grumble.Context) (err error) {
			return c.ReceiveOnchain()
		},
	})

	/**
	 * Send
	 */

	c.AddCommand(&grumble.Command{
		Name: "send_payment",
		Help: "send a lightning payment",
		Args: func(a *grumble.Args) {
			a.String("bolt11", "bolt11 lightning invoice")
			a.Uint64("amount_msat", "amount to send in millisatoshis", grumble.Default(uint64(0)))
		},
		Run: func(ctx *grumble.Context) (err error) {
			bolt11 := ctx.Args.String("bolt11")
			amountMsat := util.NilUint64(ctx.Args.Uint64("amount_msat"))

			return c.SendPayment(bolt11, amountMsat)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "send_spontaneous_payment",
		Help: "send a spontaneous (keysend) payment",
		Args: func(a *grumble.Args) {
			a.String("node_id", "receiving node id")
			a.Uint64("amount_msat", "amount to send in millisatoshis")
		},
		Run: func(ctx *grumble.Context) (err error) {
			nodeId := ctx.Args.String("node_id")
			amountMsat := ctx.Args.Uint64("amount_msat")

			return c.SendSpontaneousPayment(nodeId, amountMsat)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "send_onchain",
		Help: "send on-chain using a reverse swap",
		Args: func(a *grumble.Args) {
			a.Uint64("amount_sat", "amount to send in satoshis")
			a.String("address", "receiving on-chain address")
			a.Uint("sat_per_vbyte", "fee rate for the claim transaction")
		},
		Run: func(ctx *grumble.Context) (err error) {
			amountSat := ctx.Args.Uint64("amount_sat")
			address := ctx.Args.String("address")
			satPerVByte := uint32(ctx.Args.Uint("sat_per_vbyte"))

			return c.SendOnchain(amountSat, address, satPerVByte)
		},
	})

	/**
	 * Payment
	 */

	c.AddCommand(&grumble.Command{
		Name: "list_payments",
		Help: "list and filter payments",
		Flags: func(f *grumble.Flags) {
			f.StringL("filters", "", "filter payment types (comma-separated)")
			f.Int64("f", "from_timestamp", 0, "from unix timestamp")
			f.Int64("t", "to_timestamp", 0, "to unix timestamp")
			f.Bool("i", "include_failures", false, "include failed payments")
			f.Uint("o", "offset", 0, "offset of filtered payments")
			f.Uint("l", "limit", 0, "limit payments list size")
		},
		Run: func(ctx *grumble.Context) (err error) {
			filters := util.NilStringArray(strings.Split(ctx.Flags.String("filters"), ","))
			fromTimestamp := util.NilInt64(ctx.Flags.Int64("from_timestamp"))
			toTimestamp := util.NilInt64(ctx.Flags.Int64("to_timestamp"))
			includeFailures := ctx.Flags.Bool("include_failures")
			offset := util.NilUint32(uint32(ctx.Flags.Uint("offset")))
			limit := util.NilUint32(uint32(ctx.Flags.Uint("limit")))

			return c.ListPayments(filters, fromTimestamp, toTimestamp, &includeFailures, offset, limit)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "payment_by_hash",
		Help: "fetch a payment by its hash",
		Args: func(a *grumble.Args) {
			a.String("payment_hash", "hash of the payment")
		},
		Run: func(ctx *grumble.Context) (err error) {
			paymentHash := ctx.Args.String("payment_hash")

			return c.PaymentByHash(paymentHash)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "set_payment_metadata",
		Help: "set the metadata for a given payment",
		Args: func(a *grumble.Args) {
			a.String("payment_hash", "hash of the payment")
			a.String("metadata", "JSON encoded metadata")
		},
		Run: func(ctx *grumble.Context) (err error) {
			paymentHash := ctx.Args.String("payment_hash")
			metadata := ctx.Args.String("metadata")

			return c.SetPaymentMetadata(paymentHash, metadata)
		},
	})

	/**
	 * LSP
	 */

	c.AddCommand(&grumble.Command{
		Name: "close_lsp_channels",
		Help: "close all LSP channels",
		Run: func(ctx *grumble.Context) (err error) {
			return c.CloseLspChannels()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "connect_lsp",
		Help: "connect to another LSP",
		Args: func(a *grumble.Args) {
			a.String("lsp_id", "LSP id to connect to")
		},
		Run: func(ctx *grumble.Context) (err error) {
			lspId := ctx.Args.String("lsp_id")

			return c.ConnectLsp(lspId)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "list_lsps",
		Help: "list available LSPs",
		Run: func(ctx *grumble.Context) (err error) {
			return c.ListLsps()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "lsp_info",
		Help: "get the latest LSP information",
		Run: func(ctx *grumble.Context) (err error) {
			return c.LspInfo()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "open_channel_fee",
		Help: "calculate the open channel fee",
		Args: func(a *grumble.Args) {
			a.Uint64("amount_msat", "amount to receive in millisatoshis", grumble.Default(uint64(0)))
		},
		Flags: func(f *grumble.Flags) {
			f.Int("e", "expiry", 0, "expiry in seconds")
		},
		Run: func(ctx *grumble.Context) (err error) {
			amountMsat := util.NilUint64(ctx.Flags.Uint64("amount_msat"))
			expiry := util.NilUint32(uint32(ctx.Flags.Uint("expiry")))

			return c.OpenChannelFee(amountMsat, expiry)
		},
	})

	/**
	 * Onchain
	 */

	c.AddCommand(&grumble.Command{
		Name: "list_refundables",
		Help: "list refundable swap addresses",
		Run: func(ctx *grumble.Context) (err error) {
			return c.ListRefundables()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "prepare_redeem_onchain_funds",
		Help: "calculate the fee for a potential transaction",
		Args: func(a *grumble.Args) {
			a.String("to_address", "address to send redeem to")
			a.Uint64("sat_per_vbyte", "fee rate for the redeem transaction")
		},
		Run: func(ctx *grumble.Context) (err error) {
			toAddress := ctx.Args.String("to_address")
			satPerVByte := uint32(ctx.Args.Uint("sat_per_vbyte"))

			return c.PrepareRedeemOnchainFunds(toAddress, satPerVByte)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "prepare_refund",
		Help: "prepare a refund transaction for an incomplete swap",
		Args: func(a *grumble.Args) {
			a.String("swap_address", "address of the incomplete swap")
			a.String("to_address", "address to send refund to")
			a.Uint("sat_per_vbyte", "fee rate for the refund transaction")
		},
		Run: func(ctx *grumble.Context) (err error) {
			swapAddress := ctx.Args.String("swap_address")
			toAddress := ctx.Args.String("to_address")
			satPerVByte := uint32(ctx.Args.Uint("sat_per_vbyte"))

			return c.PrepareRefund(swapAddress, toAddress, satPerVByte)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "recommended_fees",
		Help: "list recommended fees based on the mempool",
		Run: func(ctx *grumble.Context) (err error) {
			return c.RecommendedFees()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "redeem_onchain_funds",
		Help: "send on-chain funds to an external address",
		Args: func(a *grumble.Args) {
			a.String("to_address", "address to send redeem to")
			a.Uint("sat_per_vbyte", "fee rate for the redeem transaction")
		},
		Run: func(ctx *grumble.Context) (err error) {
			toAddress := ctx.Args.String("to_address")
			satPerVByte := uint32(ctx.Args.Uint("sat_per_vbyte"))

			return c.RedeemOnchainFunds(toAddress, satPerVByte)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "refund",
		Help: "broadcast a refund transaction for an incomplete swap",
		Args: func(a *grumble.Args) {
			a.String("swap_address", "address of the incomplete swap")
			a.String("to_address", "address to send refund to")
			a.Uint("sat_per_vbyte", "fee rate for the refund transaction")
		},
		Run: func(ctx *grumble.Context) (err error) {
			swapAddress := ctx.Args.String("swap_address")
			toAddress := ctx.Args.String("to_address")
			satPerVByte := uint32(ctx.Args.Uint("sat_per_vbyte"))

			return c.Refund(swapAddress, toAddress, satPerVByte)
		},
	})

	/**
	 * Swap
	 */

	c.AddCommand(&grumble.Command{
		Name: "fetch_onchain_fees",
		Help: "fetch the current fees for a reverse swap",
		Args: func(a *grumble.Args) {
			a.Uint64("send_amount_sat", "amount to send in satoshis")
		},
		Run: func(ctx *grumble.Context) (err error) {
			sendAmountSat := util.NilUint64(ctx.Args.Uint64("send_amount_sat"))

			return c.FetchReverseSwapFees(sendAmountSat)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "in_progress_reverse_swaps",
		Help: "fetch the in-progress reverse swaps",
		Run: func(ctx *grumble.Context) (err error) {
			return c.InProgressReverseSwaps()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "in_progress_swap",
		Help: "fetch the in-progress swap",
		Run: func(ctx *grumble.Context) (err error) {
			return c.InProgressSwap()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "max_reverse_swap_amount",
		Help: "calculate the maximum amount for a reverse swap",
		Run: func(ctx *grumble.Context) (err error) {
			return c.MaxReverseSwapAmount()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "rescan_swaps",
		Help: "rescan all swaps",
		Run: func(ctx *grumble.Context) (err error) {
			return c.RescanSwaps()
		},
	})

	/**
	 * Support
	 */

	c.AddCommand(&grumble.Command{
		Name: "parse",
		Help: "parse text to get its type and relevant metadata",
		Args: func(a *grumble.Args) {
			a.String("input", "input text to parse")
		},
		Run: func(ctx *grumble.Context) (err error) {
			input := ctx.Args.String("input")

			return c.ParseInput(input)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "service_health_check",
		Help: "fetch the latest service health check",
		Run: func(ctx *grumble.Context) (err error) {
			return c.ServiceCheckCheck()
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "report_payment_failure",
		Help: "send a payment failure report",
		Args: func(a *grumble.Args) {
			a.String("payment_hash", "hash of the failed payment")
		},
		Flags: func(f *grumble.Flags) {
			f.String("c", "comment", "", "comment about the failed payment")
		},
		Run: func(ctx *grumble.Context) (err error) {
			paymentHash := ctx.Args.String("payment_hash")
			comment := util.NilString(ctx.Flags.String("comment"))

			return c.ReportPaymentFailure(paymentHash, comment)
		},
	})

	c.AddCommand(&grumble.Command{
		Name: "execute_dev_command",
		Help: "execute a low level node command (used for debugging)",
		Args: func(a *grumble.Args) {
			a.String("command", "command to send")
		},
		Run: func(ctx *grumble.Context) (err error) {
			command := ctx.Args.String("command")

			return c.ExecuteDevCommand(command)
		},
	})

	return nil
}
