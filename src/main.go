package main

import (
	"path/filepath"

	"github.com/deltachat-bot/deltabot-cli-go/botcli"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/spf13/cobra"
)

var cli = botcli.New("forumbot")

func init() {
	cli.OnBotInit(func(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
		onBotInit(bot, cli.AppDir)
	})
}

func onBotInit(bot *deltachat.Bot, appDir string) {
	initDB(filepath.Join(appDir, "bot.db"))
	bot.OnUnhandledEvent(onEvent)
	bot.OnNewMsg(onNewMsg)
}

func main() {
	if err := cli.Start(); err != nil {
		cli.Logger.Error(err)
	}
}
