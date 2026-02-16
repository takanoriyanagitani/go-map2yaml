package gmap2yaml

import (
	"context"
	"io"

	gy "github.com/goccy/go-yaml"
	my "github.com/takanoriyanagitani/go-map2yaml"
)

type Config struct {
	Options []gy.EncodeOption
}

//nolint:gochecknoglobals
var ConfigDefault Config

func (c Config) EnableAutoInt() Config {
	return Config{
		Options: append(c.Options, gy.AutoInt()),
	}
}

func (c Config) WithIndent(spaces int) Config {
	return Config{
		Options: append(c.Options, gy.Indent(spaces)),
	}
}

func (c Config) WithFlow(isFlowStyle bool) Config {
	return Config{
		Options: append(c.Options, gy.Flow(isFlowStyle)),
	}
}

func (c Config) WithIndentSequence(indent bool) Config {
	return Config{
		Options: append(c.Options, gy.IndentSequence(indent)),
	}
}

func (c Config) UseLiteralStyleIfMultiline(useLiteralStyleIfMultiline bool) Config {
	return Config{
		Options: append(c.Options, gy.UseLiteralStyleIfMultiline(useLiteralStyleIfMultiline)),
	}
}

func (c Config) UseSingleQuote(sq bool) Config {
	return Config{
		Options: append(c.Options, gy.UseSingleQuote(sq)),
	}
}

func (c Config) ToEncoder(wtr io.Writer) Encoder {
	var enc *gy.Encoder = gy.NewEncoder(wtr, c.Options...)
	return Encoder{Encoder: enc}
}

type Encoder struct{ *gy.Encoder }

func (g Encoder) Close() error { return g.Encoder.Close() }

func (g Encoder) Write(ctx context.Context, imap my.InputMap) error {
	return g.Encoder.EncodeContext(ctx, imap)
}

func (g Encoder) AsWriteCloser() my.WriteCloser {
	return g
}
