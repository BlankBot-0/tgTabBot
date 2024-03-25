package events

import (
	"context"
	"errors"
	"tgScoreBot/src/client/telegram"
	"tgScoreBot/src/e"
)

var _ Processor = (*TgProcessor)(nil)

type TgProcessor struct {
	tg     *telegram.Client
	offset int
}

type Meta struct {
	ChatID   int
	Username string
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *telegram.Client) *TgProcessor {
	return &TgProcessor{
		tg: client,
	}
}

// TODO: make concurrent
func (p *TgProcessor) Fetch(ctx context.Context, limit int) ([]Event, error) {
	updates, err := p.tg.Updates(ctx, p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *TgProcessor) Process(ctx context.Context, event Event) error {
	switch event.Type {
	case Message:
		return p.processMessage(ctx, event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p *TgProcessor) processMessage(ctx context.Context, event Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.doCmd(ctx, event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("can't process message", err)
	}

	return nil
}

func meta(event Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}

	return res, nil
}

func event(upd telegram.Update) Event {
	updType := fetchType(upd)

	res := Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd telegram.Update) Type {
	if upd.Message == nil {
		return Unknown
	}

	return Message
}
