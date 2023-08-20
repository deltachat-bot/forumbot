package main

import (
	"fmt"

	"github.com/deltachat-bot/deltabot-cli-go/botcli"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/spf13/cobra"
)

// echo back received text
func echo(bot *deltachat.Bot, accId deltachat.AccountId, msgId deltachat.MsgId) {
	msg, _ := bot.Rpc.GetMessage(accId, msgId)
	if msg.FromId > deltachat.ContactLastSpecial && msg.Text != "" {
		bot.Rpc.MiscSendTextMessage(accId, msg.ChatId, msg.Text) //nolint:errcheck
	}
}

func onInfoCmd(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
	sysinfo, _ := bot.Rpc.GetSystemInfo()
	for key, val := range sysinfo {
		fmt.Printf("%v=%#v\n", key, val)
	}
}

func onBotInit(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
	bot.OnNewMsg(echo)
}

func onBotStart(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
	cli.Logger.Info("OnBotStart event triggered: bot is about to start!")
}

func main() {
	cli := botcli.New("echobot")

	// add an "info" CLI subcommand as example
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "display information about the Delta Chat core running in this system",
		Args:  cobra.ExactArgs(0),
	}
	cli.AddCommand(infoCmd, onInfoCmd)

	cli.OnBotInit(onBotInit)
	cli.OnBotStart(onBotStart)

	if err := cli.Start(); err != nil {
		cli.Logger.Error(err)
	}
}
