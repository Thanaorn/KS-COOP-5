package service

import (
	"context"
	"fmt"
)

func (wr WebhookService) WebhookService() error {
	data, err := wr.storage.FindAll(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("data = ", data)

	return nil
}

func (wr WebhookService) test() error {

	return nil
}
