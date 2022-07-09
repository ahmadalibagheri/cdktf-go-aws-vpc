package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/jsii-runtime-go"
)

func newTag(name, team, company string) *map[string]*string {
	return &map[string]*string{
		"Name":    jsii.String(name),
		"Team":    jsii.String(team),
		"Company": jsii.String(company),
	}
}
func newPrivateSubnetConfig(name string, az, vpcId, cidrBlock *string) *vpc.SubnetConfig {
	return &vpc.SubnetConfig{
		AvailabilityZone:    az,
		VpcId:               vpcId,
		MapPublicIpOnLaunch: false,
		CidrBlock:           cidrBlock,
		Tags:                newTag(name, team, company),
	}
}

func newPublicSubnetConfig(name string, az, vpcId, cidrBlock *string) *vpc.SubnetConfig {
	return &vpc.SubnetConfig{
		AvailabilityZone:    az,
		VpcId:               vpcId,
		MapPublicIpOnLaunch: true,
		CidrBlock:           cidrBlock,
		Tags:                newTag(name, team, company),
	}
}

func getAvailabilityZones() []*string {
	return []*string{
		jsii.String("us-east-1a"),
		jsii.String("us-east-1b"),
		jsii.String("us-east-1c"),
		jsii.String("us-east-1d"),
		jsii.String("us-east-1e"),
		jsii.String("us-east-1f"),
	}
}
