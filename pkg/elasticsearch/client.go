package elasticsearch

import (
	"big_mall_api/configs"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

type Client struct {
	client *elasticsearch.Client
}

func NewClient(cfg *configs.ESConfig) (*Client, error) {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port),
		},
		Username: cfg.Username,
		Password: cfg.Password,
	}

	client, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) GetClient() *elasticsearch.Client {
	return c.client
}
