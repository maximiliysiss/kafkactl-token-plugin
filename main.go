package main

import (
	"errors"
	"github.com/deviceinsight/kafkactl/v5/pkg/plugins/auth"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"os"
)

type tokenProvider struct {
	token  string
	logger hclog.Logger
}

func (t *tokenProvider) Token() (string, error) {
	if len(t.token) == 0 {
		return "", errors.New("empty token")
	}
	return t.token, nil
}

func (t *tokenProvider) Init(options map[string]any, brokers []string) error {

	v, ok := options["token"].(string)

	if !ok || len(v) == 0 {
		return errors.New("missing token option")
	}

	t.logger.Info("token option set")

	t.token = v
	return nil
}

func main() {

	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.DefaultLevel,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	tokenProvider := &tokenProvider{logger: logger}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: auth.TokenProviderPluginSpec.Handshake,
		Plugins: map[string]plugin.Plugin{
			auth.TokenProviderPluginSpec.InterfaceIdentifier: &auth.TokenProviderPlugin{Impl: tokenProvider},
		},
	})
}
