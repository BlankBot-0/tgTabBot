package telegram

import (
	"flag"
	"fmt"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
	"github.com/sinshu/go-meltysynth/meltysynth"
	"log"
	"os"
	"strings"
	"tgScoreBot/src"
	"tgScoreBot/src/tabs"
	"time"
)

type TgProcessor struct {
	tg *telego.Bot
	sf *meltysynth.SoundFont
}

func NewTgProcessor(cfg Config, sf *meltysynth.SoundFont) *TgProcessor {
	bot, err := telego.NewBot(cfg.TgBotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatal("can not start bot")
	}
	return &TgProcessor{
		tg: bot,
		sf: sf,
	}
}

func (p *TgProcessor) Start() {
	updates, _ := p.tg.UpdatesViaLongPolling(&telego.GetUpdatesParams{},
		telego.WithLongPollingUpdateInterval(2*time.Second))

	for upd := range updates {
		if err := p.Process(upd); err != nil {
			log.Print("err")
		}
	}
}

func (p *TgProcessor) Process(upd telego.Update) error {
	if upd.Message == nil {
		return fmt.Errorf("can't process message: empty message")
	}

	return p.doCmd(upd.Message.Text, upd.Message.Chat.ID, upd.Message.Chat.Username)
}

func (p *TgProcessor) doCmd(text string, chatID int64, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	case ExampleCmd:
		return p.sendExample(chatID)
	case InstructionsCmd:
		return p.sendInstructions(chatID)
	default:
		return p.sendResult(chatID, text)
	}
}

func (p *TgProcessor) sendHelp(id int64) error {
	_, err := p.tg.SendMessage(telegoutil.Message(telegoutil.ID(id), msgHelp))
	return err
}

func (p *TgProcessor) sendHello(id int64) error {
	_, err := p.tg.SendMessage(telegoutil.Message(telegoutil.ID(id), msgHello))
	return err
}

func (p *TgProcessor) sendExample(id int64) error {
	_, err := p.tg.SendMessage(telegoutil.Message(telegoutil.ID(id), msgExample))
	return err
}

func (p *TgProcessor) sendInstructions(id int64) error {
	_, err := p.tg.SendMessage(telegoutil.Message(telegoutil.ID(id), msgInstructions))
	return err
}

func (p *TgProcessor) sendResult(id int64, text string) (err error) {
	defer func() {
		err = src.WrapIfErr("sendResult error:", err)
		if r := recover(); r != nil {
			_, err = p.tg.SendMessage(telegoutil.Message(telegoutil.ID(id), msgError))
		}
	}()
	mp3, err := tabs.TextToMp3(text, p.sf)
	if err != nil {
		_, err := p.tg.SendMessage(telegoutil.Message(telegoutil.ID(id), msgError))
		return err
	}
	params := AudioParams(&mp3, telegoutil.ID(id))
	_, err = p.tg.SendAudio(&params)
	return err
}

type Config struct {
	TgBotToken string
}

func MustLoad() Config {
	tgBotTokenToken := os.Getenv("TOKEN")

	flag.Parse()

	if tgBotTokenToken == "" {
		log.Fatal("token is not specified")
	}

	return Config{
		TgBotToken: tgBotTokenToken,
	}
}
