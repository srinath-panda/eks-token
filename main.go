package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func main() {
	str := "{\"kind\": \"ExecCredential\", \"apiVersion\": \"client.authentication.k8s.io/v1beta1\", \"spec\": {}, \"status\": {\"expirationTimestamp\": \"%v\", \"token\": \"%v\"}}"

	namePtr := flag.String("name", "", "Name of the EKS cluster")
	profilePtr := flag.String("profile", "", "Name of the aws profile")
	flag.Parse()
	var err error

	if *namePtr == "" {
		panic("Please pass the correct value for --name")
	}
	if *profilePtr == "" {
		panic("Please pass the correct value for --profile")
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{},
		Profile:           aws.StringValue(profilePtr),
	})
	if err != nil {
		panic(err)
	}

	opts := &token.GetTokenOptions{
		ClusterID: aws.StringValue(namePtr),
		Session:   sess,
	}
	gen, err := token.NewGenerator(true, false)
	if err != nil {
		panic(err)
	}
	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf(str, tok.Expiration.UTC().Format(time.RFC3339), tok.Token)
}
