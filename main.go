package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	smarthome "github.com/orktes/go-alexa-smarthome"
)

// main entry point
func main() {

	sm := smarthome.New(smarthome.AuthorizationFunc(func(req smarthome.AcceptGrantRequest) error {
		return nil
	}))

	// Get configurations from file
	config := &Config{}
	err := config.Parse()
	if err != nil {
		return
	}

	sess := session.Must(session.NewSession())

	// Try first with Environment variables and secondly with IAM role
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{},
			&ec2rolecreds.EC2RoleProvider{
				Client: ec2metadata.New(sess),
			},
		})

	dataPlaneConfig := &aws.Config{
		Region:      &config.Region,
		Credentials: creds,
		Endpoint:    &config.IOTEndpoint,
	}
	dataPlaneClientIOT := iotdataplane.New(sess, dataPlaneConfig)

	iotConfig := &aws.Config{
		Region:      &config.Region,
		Credentials: creds,
	}
	clientIOT := iot.New(sess, iotConfig)

	// mock data for a Thingler smart plug
	//mockThinglerSmartPlug(sm, dataPlaneClientIOT, config)

	// Get registred smart plugs
	getRegisteredSmartPlugs(sm, clientIOT, dataPlaneClientIOT, config)

	lambda.Start(sm.Handle)
}
