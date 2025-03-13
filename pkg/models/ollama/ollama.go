package ollama

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/ollama/ollama/api"

	"github.com/katallaxie/m/pkg/messages"
	"github.com/katallaxie/m/pkg/models"
)

var _ models.Model[string] = (*Ollama[string])(nil)

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
type Ollama[V models.Value] struct {
	client *api.Client
	opts   *Opts
}

// New ...
func New[V models.Value](opts ...Opt) (*Ollama[V], error) {
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

	model := new(Ollama[V])
	model.client = api.NewClient(baseURL, options.Client)
	model.opts = options

	return model, nil
}

// Generate ...
func (o *Ollama[V]) Generate(ctx context.Context, input []messages.Message[V], opts ...models.Opt) (messages.Message[V], error) {
	var req *api.ChatRequest
	var msg messages.Message[V]

	fn := func(res api.ChatResponse) error {
		return nil
	}

	err := o.client.Chat(ctx, req, fn)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
