package cli

import (
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
		Name: "node_info",
		Help: "get the latest node state",
		Run: func(ctx *grumble.Context) (err error) {
			return c.NodeInfo()
		},
	})

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

	c.AddCommand(&grumble.Command{
		Name: "send_payment",
		Help: "send a lightning payment",
		Args: func(a *grumble.Args) {
			a.String("bolt11", "bolt11 lightning invoice")
			a.Uint64("amount_msat", "amount to send in millisatoshis", grumble.Default(0))
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

	return nil
}
