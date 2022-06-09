package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"os"
)

type CdkBazelTestStackProps struct {
	awscdk.StackProps
}

func NewCdkBazelTestStack(scope constructs.Construct, id string, props *CdkBazelTestStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("CdkBazelTestQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })
	ami := awsec2.MachineImage_Lookup(&awsec2.LookupMachineImageProps{
		Name:   jsii.String("whslabs-*"),
		Owners: jsii.Strings("102933037533"),
	})

	vpc := awsec2.NewVpc(stack, jsii.String("Vpc"), &awsec2.VpcProps{
		Cidr: jsii.String("10.0.0.0/16"),
	})

	awsec2.NewInstance(stack, jsii.String("Instance"), &awsec2.InstanceProps{
		InstanceType: awsec2.NewInstanceType(jsii.String("t3.large")),
		MachineImage: ami,
		Vpc:          vpc,
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkBazelTestStack(app, "CdkBazelTestStack", &CdkBazelTestStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
