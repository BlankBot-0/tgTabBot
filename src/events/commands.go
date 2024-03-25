package events

import (
	"context"
	"log"
	"strings"
)

const (
	HelpCmd         = "/help"
	StartCmd        = "/start"
	ExampleCmd      = "/example"
	InstructionsCmd = "/instructions"
)

func (p *TgProcessor) doCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(ctx, chatID)
	case StartCmd:
		return p.sendHello(ctx, chatID)
	case ExampleCmd:
		return p.sendExample(ctx, chatID)
	case InstructionsCmd:
		return p.sendInstructions(ctx, chatID)
	default:
		// TODO: implement actual action
		return p.tg.SendMessage(ctx, chatID, msgUnknownCommand)
	}
}

func (p *TgProcessor) sendHelp(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgHelp)
}

func (p *TgProcessor) sendHello(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgHello)
}

func (p *TgProcessor) sendExample(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgExample)
}

func (p *TgProcessor) sendInstructions(ctx context.Context, chatID int) error {
	return p.tg.SendMessage(ctx, chatID, msgInstructions)
}
