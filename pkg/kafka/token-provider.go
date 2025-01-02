package kafka

import (
	"context"
	"time"

	"github.com/aws/aws-msk-iam-sasl-signer-go/signer"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type IAMTokenProvider struct {
	region              string
	credentialsProvider credentials.StaticCredentialsProvider
}

func NewIAMTokenProvider(region, accessKey, secretKey string) (*IAMTokenProvider, error) {
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	return &IAMTokenProvider{
		region:              region,
		credentialsProvider: creds,
	}, nil
}

func (p *IAMTokenProvider) Token() (kafka.OAuthBearerToken, error) {
	token, tokenExpirationTime, err := signer.GenerateAuthTokenFromCredentialsProvider(
		context.TODO(),
		p.region,
		&p.credentialsProvider,
	)
	if err != nil {
		return kafka.OAuthBearerToken{}, err
	}
	seconds := tokenExpirationTime / 1000
	nanoseconds := (tokenExpirationTime % 1000) * 1000000

	bearerToken := kafka.OAuthBearerToken{
		TokenValue: token,
		Expiration: time.Unix(seconds, nanoseconds),
	}

	return bearerToken, nil
}
