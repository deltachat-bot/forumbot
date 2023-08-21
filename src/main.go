package main

import (
	"github.com/deltachat-bot/deltabot-cli-go/botcli"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/spf13/cobra"
)

var cli = botcli.New("forumbot")

func main() {
	cli.OnBotInit(func(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
		bot.On(deltachat.EventSecurejoinInviterProgress{}, onEvent)
		bot.On(deltachat.EventWebxdcStatusUpdate{}, onEvent)
		bot.OnNewMsg(onNewMsg)
	})

	if err := cli.Start(); err != nil {
		cli.Logger.Error(err)
	}
}
