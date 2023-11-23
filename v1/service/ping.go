// author gmfan
// date 2023/8/1
package service

import (
	"acsupport/v1/models"
	"context"
)

func Pong(ctx context.Context) (resp any, err error) {
	return &models.PongVO{
		Pong: "pong",
	}, nil
}
