package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	fail = syntax()

	account         = os.Args[1]
	resourceid      = os.Args[2]
	tagslice        = os.Args[3:]
	credentialsPath = os.Getenv("HOME") + "/.aws/credentials"
	svc             = authAWS(account)
)

func main() {
	var listoftags []*ec2.Tag

	for _, slc := range tagslice {
		asplit := strings.Split(slc, ":")
		ts := &ec2.Tag{
			Key:   aws.String(asplit[0]),
			Value: aws.String(asplit[1]),
		}
		listoftags = append(listoftags, ts)
	}

	t := &ec2.CreateTagsInput{
		Resources: []*string{aws.String(resourceid)},
		Tags:      listoftags}

	_, e := svc.CreateTags(t)
	if e != nil {
		log.Fatalln("Die: ", e)
	}
	fmt.Printf("Setting: %v\n", listoftags)
	fmt.Println("Done!")
}

func syntax() string {
	if len(os.Args) != 4 {
		fmt.Printf("- %v: ERR\n- %v account_role instanceid Key:Name\n- You can specifiy as many Key:Name combos as you wish.\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}
	return "Die"
}

func authAWS(account string) *ec2.EC2 {
	credentialObject := credentials.NewSharedCredentials(credentialsPath, account)
	svc := ec2.New(session.New(aws.NewConfig().WithRegion("us-east-1").
		WithMaxRetries(2).WithCredentials(credentialObject)))
	return svc
}
