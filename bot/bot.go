package bot

import (
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wol-tg-bot/util"
)

type WolBotConfig struct {
	Token             string
	ValidUserNameList []string
	MacAddr           string
	IpAddr            string
	Port              int
}

func Start(config *WolBotConfig) error {
	// init wol bot
	settings := telebot.Settings{Token: config.Token}
	bot, err := telebot.NewBot(settings)
	if err != nil {
		log.Panicf("wol bot init error: %v", err)
		return err
	}

	// register route logic
	bot.Handle("/power", func(ctx telebot.Context) error {
		log.Println("wol bot receive /power command")

		// check user permission
		if isUserInvalid(config.ValidUserNameList, ctx.Sender().Username) {
			log.Printf("wol bot detect /power command from invalid userName: %s", ctx.Sender().Username)
			return nil
		}

		// try wake target machine and reply message
		err := util.DoWake(config.MacAddr, config.IpAddr, config.Port)
		if err != nil {
			log.Printf("wol bot do wake macAddr: %s, ipAddr: %s, port: %d, error: %v",
				config.MacAddr,
				config.IpAddr,
				config.Port,
				err)
			_, _ = bot.Reply(ctx.Message(), "Error waking up the machine please retry later")
			return nil
		}

		_, _ = bot.Reply(ctx.Message(), "Done")
		log.Println("wol bot done /power command")
		return nil
	})

	// start in goroutine
	go bot.Start()
	log.Println("wol bot start")

	// wait close signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("shutting down wol bot...")
	bot.Stop()
	log.Println("wol bot closed")

	return nil
}

func isUserInvalid(validUserNameList []string, userName string) bool {
	for _, validUser := range validUserNameList {
		if userName == validUser {
			return false
		}
	}
	return true
}
