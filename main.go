package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	check(err)

	client := sts.NewFromConfig(cfg)

	resp, err := client.GetCallerIdentity(ctx, nil)
	check(err)

	fmt.Println(*resp.Account)
	fmt.Println(*resp.Arn)
	fmt.Println(*resp.UserId)
}
