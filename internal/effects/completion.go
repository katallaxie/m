package effects

import (
	"bufio"
	"bytes"
	"context"

	"github.com/katallaxie/m/internal/api"
	"github.com/katallaxie/m/internal/state"
	"github.com/katallaxie/pkg/fsmx"
	"github.com/katallaxie/pkg/slices"
	"github.com/katallaxie/prompts"
	"github.com/katallaxie/streams"
	"github.com/katallaxie/streams/sinks"
	"github.com/katallaxie/streams/sources"
)

func mapCompletionMessages(msg prompts.Completion) string {
	f := slices.First(msg.Choices...)
	return f.Message.GetContent()
}

type NoopWriterCloser struct {
	*bufio.Writer
}

func (w *NoopWriterCloser) Close() error {
	return w.Flush()
}

func FetchChatCompletion(api *api.Api, prompt string) fsmx.EffectFunc[state.State] {
	return func(ctx context.Context) (fsmx.Action, error) {
		res, err := api.CreatePrompt(ctx, prompt)
		if err != nil {
			return nil, err
		}

		b := bytes.NewBuffer(nil)
		bw := bufio.NewWriter(b)
		mw := &NoopWriterCloser{bw}

		sink, err := sinks.NewWriter(mw)
		if err != nil {
			return nil, err
		}

		source := sources.NewChanSource(res)
		source.Pipe(streams.NewPassThrough()).Pipe(streams.NewMap(mapCompletionMessages)).To(sink)

		return state.NewAddMessage(b.String()), nil
	}
}
