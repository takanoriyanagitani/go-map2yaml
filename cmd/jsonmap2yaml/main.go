package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"os"

	my "github.com/takanoriyanagitani/go-map2yaml"
	gm "github.com/takanoriyanagitani/go-map2yaml/ser/goccy"
)

func sub(ctx context.Context, cfg gm.Config) error {
	var rdr io.Reader = bufio.NewReader(os.Stdin)
	var dec *json.Decoder = json.NewDecoder(rdr)

	var buf map[string]any

	err := dec.Decode(&buf)
	if nil != err {
		return err
	}

	var wtr *bufio.Writer = bufio.NewWriter(os.Stdout)

	var enc gm.Encoder = cfg.ToEncoder(wtr)
	var wcls my.WriteCloser = enc.AsWriteCloser()

	err = wcls.Write(ctx, buf)

	return errors.Join(
		err,
		wcls.Close(),
		wtr.Flush(),
	)
}

func main() {
	var indent int
	flag.IntVar(&indent, "indent", 2, "indent spaces")

	var singleQuote bool
	flag.BoolVar(&singleQuote, "single-quote", false, "use single quotes")

	var flow bool
	flag.BoolVar(&flow, "flow", false, "use flow style")

	var autoInt bool
	flag.BoolVar(&autoInt, "auto-int", false, "enable auto int conversion")

	var literalMultiline bool
	flag.BoolVar(&literalMultiline, "literal-multiline", false, "use literal style for multiline strings")

	var indentSequence bool
	flag.BoolVar(&indentSequence, "indent-sequence", false, "indent sequence items")

	flag.Parse()

	cfg := gm.ConfigDefault
	if 0 < indent {
		cfg = cfg.WithIndent(indent)
	}
	if singleQuote {
		cfg = cfg.UseSingleQuote(singleQuote)
	}
	if flow {
		cfg = cfg.WithFlow(flow)
	}
	if autoInt {
		cfg = cfg.EnableAutoInt()
	}
	if literalMultiline {
		cfg = cfg.UseLiteralStyleIfMultiline(literalMultiline)
	}
	if indentSequence {
		cfg = cfg.WithIndentSequence(indentSequence)
	}

	err := sub(context.Background(), cfg)
	if nil != err {
		log.Printf("%v\n", err)
	}
}
