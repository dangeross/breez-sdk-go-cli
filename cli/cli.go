package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/breez/breez-sdk-go/breez_sdk"
	"github.com/dangeross/breez-sdk-go-cli/internal/util"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type Cli struct {
	*grumble.App

	log *logrus.Logger
	sdk *breez_sdk.BlockingBreezServices

	dataDir string
	config  *Config
}

func Run() {
	c := &Cli{
		App: grumble.New(&grumble.Config{
			Name:                  "sdk",
			Description:           "A Breez SDK Go CLI",
			HistoryFile:           "history.txt",
			Prompt:                "sdk> ",
			PromptColor:           color.New(color.FgHiCyan, color.Bold),
			HelpHeadlineColor:     color.New(color.FgHiCyan),
			HelpHeadlineUnderline: true,
			HelpSubCommands:       true,

			Flags: func(f *grumble.Flags) {
				f.String("d", "data_dir", "", "optional data dir, default to current directory")
			},
		}),
	}

	c.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("__________                                  _________________   ____  __.")
		a.Println("\\______   \\_______   ____   ____ ________  /   _____/\\______ \\ |    |/ _|")
		a.Println(" |    |  _/\\_  __ \\_/ __ \\_/ __ \\\\___   /  \\_____  \\  |    |  \\|      <  ")
		a.Println(" |    |   \\ |  | \\/\\  ___/\\  ___/ /    /   /        \\ |    `   \\    |  \\ ")
		a.Println(" |______  / |__|    \\___  >\\___  >_____ \\ /_______  //_______  /____|__ \\")
		a.Println("        \\/              \\/     \\/      \\/         \\/         \\/        \\/")
		a.Println()
	})

	c.OnInit(func(app *grumble.App, flags grumble.FlagMap) (err error) {
		c.dataDir = flags.String("data_dir")

		if len(c.dataDir) == 0 {
			if c.dataDir, err = os.Getwd(); err != nil {
				return fmt.Errorf("failed to get the current working directory: %v", err)
			}
		}

		if c.dataDir, err = filepath.Abs(c.dataDir); err != nil {
			return err
		}

		if err = os.MkdirAll(c.dataDir, os.ModePerm); err != nil {
			return err
		}

		app.Config().HistoryFile = filepath.Join(c.dataDir, "history.txt")

		// Init config
		if err := c.readConfig(); err != nil {
			panic(err)
		}

		// Init log
		c.initLog()
		breez_sdk.SetLogStream(c)

		return c.load()
	})

	c.OnClose(func() (err error) {
		if c.sdk != nil {
			c.sdk.Disconnect()
			c.sdk.Destroy()
			c.sdk = nil
		}

		return
	})

	grumble.Main(c.App)
}

func (c *Cli) PrettyPrint(i interface{}) {
	if b, err := json.MarshalIndent(i, "", "  "); err == nil {
		c.Println(string(b))
	}
}

func (c *Cli) PrintSuccess(str string) {
	color.New(color.FgGreen, color.Bold).Fprintln(c.App, str)
}

func (c *Cli) load() error {
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
		Name: "send_payment",
		Help: "send a lightning payment",
		Args: func(a *grumble.Args) {
			a.String("bolt11", "bolt11 lightning invoice")
		},
		Flags: func(f *grumble.Flags) {
			f.Uint64("a", "amount_msat", 0, "amount to send in millisatoshis")
		},
		Run: func(ctx *grumble.Context) (err error) {
			bolt11 := ctx.Args.String("bolt11")
			amountMsat := util.NilUint64(ctx.Flags.Uint64("amount_msat"))

			return c.SendPayment(bolt11, amountMsat)
		},
	})

	return nil
}
