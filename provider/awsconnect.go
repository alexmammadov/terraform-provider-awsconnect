package provider

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connect"
)

// func connectService() *connect.Connect {
// 	sess := session.Must(session.NewSession(&aws.Config{
// 		Region: aws.String(os.Getenv("AWS_REGION")),
// 		Credentials: credentials.NewStaticCredentials(
// 			os.Getenv("AWS_ACCESS_KEY_ID"),
// 			os.Getenv("AWS_SECRET_ACCESS_KEY"),
// 			os.Getenv("AWS_SESSION_TOKEN")),
// 	}))

// 	return connect.New(sess)
// }

func connectService() *connect.Connect {
	cfg := defaults.Config()
	handlers := defaults.Handlers()

	p := defaults.RemoteCredProvider(*cfg, handlers)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewCredentials(p),
	}))

	return connect.New(sess)
}

// func connectService() *connect.Connect {
// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		Config:            aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))},
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))

// 	return connect.New(sess)
// }
