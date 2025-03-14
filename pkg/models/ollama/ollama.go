package ollama

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/katallaxie/m/pkg/chats"
	"github.com/katallaxie/m/pkg/models"
)

var _ models.Chatter = (*Ollama)(nil)

// Opts ...
type Opts struct {
	// BaseURL is the base URL.
	BaseURL string `json:"base_url"`
	// Timeout is the timeout.
	Timeout time.Duration `json:"timeout"`
	// Model is the model.
	Model string `json:"model"`
	// Client is the HTTP client.
	Client *http.Client `json:"-"`
	// Format is the format.
	Format json.RawMessage `json:"format"`
	// KeepAlive is the keep alive.
	KeepAlive bool `json:"keep_alive"`
	// Options is the options.
	Opts *api.Options `json:"options"`
}

// Opt ...
type Opt func(*Opts)

// Defaults ...
func Defaults() *Opts {
	return &Opts{}
}

// Ollama ...
type Ollama struct {
	client *api.Client
	opts   *Opts
}

// New ...
func New(opts ...Opt) (*Ollama, error) {
	options := Defaults()

	client := &http.Client{Timeout: options.Timeout}
	options.Client = client

	for _, opt := range opts {
		opt(options)
	}

	baseURL, err := url.Parse(options.BaseURL)
	if err != nil {
		return nil, err
	}

	model := new(Ollama)
	model.client = api.NewClient(baseURL, options.Client)
	model.opts = options

	return model, nil
}

// WithBaseURL ...
func WithBaseURL(baseURL string) Opt {
	return func(o *Opts) {
		o.BaseURL = baseURL
	}
}

// Generate ...
func (o *Ollama) Generate(ctx context.Context, chat chats.Chat, opts ...models.Opt) (chats.Chat, error) {
	req := &api.ChatRequest{}
	req.Model = chat.Model
	req.Messages = make([]api.Message, 0)
	req.Messages = append(req.Messages, api.Message{
		Role:    "system",
		Content: "You are an fun and friendly AI assistant that adds emojies to the answers.",
	})

	for _, m := range chat.Messages {
		req.Messages = append(req.Messages, api.Message{
			Role:    "user",
			Content: m.Content,
		})
	}

	fn := func(res api.ChatResponse) error {
		fmt.Print(res.Message.Content)

		return nil
	}

	err := o.client.Chat(ctx, req, fn)
	if err != nil {
		return chat, err
	}

	return chat, nil
}
