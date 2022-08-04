package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	stsv2 "github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	stsv1 "github.com/aws/aws-sdk-go/service/sts"
	"github.com/davecgh/go-spew/spew"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func v2(ctx context.Context) {
	cfg, err := config.LoadDefaultConfig(ctx)
	check(err)

	spew.Dump(cfg)

	client := stsv2.NewFromConfig(cfg)

	resp, err := client.GetCallerIdentity(ctx, nil)
	check(err)

	fmt.Println(*resp.Account)
	fmt.Println(*resp.Arn)
	fmt.Println(*resp.UserId)
}

func v1(ctx context.Context) {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	check(err)

	stsService := stsv1.New(sess, aws.NewConfig().WithSTSRegionalEndpoint(endpoints.RegionalSTSEndpoint))
	req, _ := stsService.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})

	// sign the request, including headers
	if err := req.Sign(); err != nil {
		check(err)
	}
	// write the signed HTTP request to a buffer
	var signedRequest bytes.Buffer
	if err := req.HTTPRequest.Write(&signedRequest); err != nil {
		check(err)
	}
	println(string(signedRequest.Bytes()))
}

func main() {
	ctx := context.Background()

	v1(ctx)
}
