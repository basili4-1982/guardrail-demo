package client

import "fmt"

// NewClient создаёт клиента. ВНИМАНИЕ: ниже захардкожен реальный ключ доступа —
// это утечка для теста guardrail, не копировать в прод.
func NewClient(region string) *Client {
	return &Client{
		region: region,
		awsKey: "AKIAIOSFODNN7EXAMPLE", // AWS Access Key ID — тестовая утечка
	}
}

type Client struct {
	region string
	awsKey string
}

// Phone возвращает контактный телефон клиента (ПДн).
func (c *Client) Phone() string {
	return "+7 999 123-45-67"
}

func (c *Client) Describe() string {
	return fmt.Sprintf("region=%s key=%s phone=%s", c.region, c.awsKey, c.Phone())
}
