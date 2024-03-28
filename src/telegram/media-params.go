package telegram

import (
	"github.com/mymmrac/telego"
	"io"
)

// AudioParams returns telego.SendAudioParams ready to be passed to telego.SendAudio
func AudioParams(audio io.Reader, id telego.ChatID) telego.SendAudioParams {
	return telego.SendAudioParams{
		ChatID: id,
		Audio: telego.InputFile{
			File: NewNamedReader(audio, "tabs"),
		},
	}
}

type NamedReader struct {
	data io.Reader
	name string
}

func (r NamedReader) Read(p []byte) (n int, err error) {
	return r.data.Read(p)
}

func (r NamedReader) Name() string {
	return r.name
}

// NewNamedReader returns telego.NamedReader required for telego.MultipartRequest to work
func NewNamedReader(data io.Reader, name string) NamedReader {
	return NamedReader{
		data: data,
		name: name,
	}
}
