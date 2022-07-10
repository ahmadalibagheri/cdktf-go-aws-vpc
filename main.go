package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws"
	"cdk.tf/go/stack/generated/hashicorp/aws/ec2"
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	team                 = "DevOps"
	company              = "Your Comapny"
	destinationCidrBlock = "0.0.0.0/0"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("us-east-1"),
	})

	newVpc := vpc.NewVpc(stack, jsii.String("vpc"), &vpc.VpcConfig{
		CidrBlock: jsii.String("10.0.0.0/16"),
		Tags:      newTag("CDKtf-Golang-Demo-Private-Subnet-A", team, company),
	})

	privateSubnetA := vpc.NewSubnet(stack, jsii.String("private-subnet-a"),
		newPrivateSubnetConfig("CDKtf-Golang-Demo-Private-Subnet-A", getAvailabilityZones()[0],
			newVpc.Id(), jsii.String("10.0.1.0/24")))

	privateSubnetB := vpc.NewSubnet(stack, jsii.String("private-subnet-b"),
		newPrivateSubnetConfig("CDKtf-Golang-Demo-Private-Subnet-B", getAvailabilityZones()[1],
			newVpc.Id(), jsii.String("10.0.2.0/24")))

	publicSubnetA := vpc.NewSubnet(stack, jsii.String("public-subnet-a"),
		newPublicSubnetConfig("CDKtf-Golang-Demo-Public-Subnet-A", getAvailabilityZones()[0],
			newVpc.Id(), jsii.String("10.0.6.0/24")))

	publicSubnetB := vpc.NewSubnet(stack, jsii.String("public-subnet-b"),
		newPublicSubnetConfig("CDKtf-Golang-Demo-Public-Subnet-B", getAvailabilityZones()[1],
			newVpc.Id(), jsii.String("10.0.7.0/24")))

	internetGateway := vpc.NewInternetGateway(stack, jsii.String("internet-gateway"),
		&vpc.InternetGatewayConfig{
			VpcId: newVpc.Id(),
			Tags:  newTag("CDKtf-Golang-Demo-IG", team, company),
		})

	publicIPAddressA := ec2.NewEip(stack, jsii.String("eip-a"), &ec2.EipConfig{
		Vpc:  true,
		Tags: newTag("CDKtf-Golang-Demo-Public-eip-A", team, company),
	})

	publicIPAddressB := ec2.NewEip(stack, jsii.String("eip-b"), &ec2.EipConfig{
		Vpc:  true,
		Tags: newTag("CDKtf-Golang-Demo-Public-eip-B", team, company),
	})

	natGatewayA := vpc.NewNatGateway(stack, jsii.String("net-gateway-a"), &vpc.NatGatewayConfig{
		AllocationId: publicIPAddressA.Id(),
		SubnetId:     publicSubnetA.Id(),
		Tags:         newTag("CDKtf-Golang-Demo-Public-NG-A", team, company),
	})

	natGatewayB := vpc.NewNatGateway(stack, jsii.String("net-gateway-b"), &vpc.NatGatewayConfig{
		AllocationId: publicIPAddressB.Id(),
		SubnetId:     publicSubnetB.Id(),
		Tags:         newTag("CDKtf-Golang-Demo-Public-NG-B", team, company),
	})

	publicRouteTable := vpc.NewRouteTable(stack, jsii.String("public-route-table"), &vpc.RouteTableConfig{
		VpcId: newVpc.Id(),
		Tags:  newTag("CDKtf-Golang-Demo-Public-RT", team, company),
	})

	vpc.NewRoute(stack, jsii.String("route"), &vpc.RouteConfig{
		DestinationCidrBlock: jsii.String(destinationCidrBlock),
		RouteTableId:         publicRouteTable.Id(),
		GatewayId:            internetGateway.Id(),
	})

	vpc.NewRouteTableAssociation(stack, jsii.String("route-table-association-pub-sub-a"),
		&vpc.RouteTableAssociationConfig{
			RouteTableId: publicRouteTable.Id(),
			SubnetId:     publicSubnetA.Id(),
		})

	vpc.NewRouteTableAssociation(stack, jsii.String("route-table-association-pub-sub-b"),
		&vpc.RouteTableAssociationConfig{
			RouteTableId: publicRouteTable.Id(),
			SubnetId:     publicSubnetB.Id(),
		})

	privateRouteTableA := vpc.NewRouteTable(stack, jsii.String("private-route-table-a"), &vpc.RouteTableConfig{
		VpcId: newVpc.Id(),
		Tags:  newTag("CDKtf-Golang-Demo-Private-RT-A", team, company),
	})

	vpc.NewRoute(stack, jsii.String("private-route-a"), &vpc.RouteConfig{
		DestinationCidrBlock: jsii.String(destinationCidrBlock),
		RouteTableId:         privateRouteTableA.Id(),
		GatewayId:            natGatewayA.Id(),
	})

	vpc.NewRouteTableAssociation(stack, jsii.String("route-table-association-private-sub-a"),
		&vpc.RouteTableAssociationConfig{
			RouteTableId: privateRouteTableA.Id(),
			SubnetId:     privateSubnetA.Id(),
		})

	privateRouteTableB := vpc.NewRouteTable(stack, jsii.String("private-route-table-b"), &vpc.RouteTableConfig{
		VpcId: newVpc.Id(),
		Tags:  newTag("CDKtf-Golang-Demo-Private-RT-B", team, company),
	})

	vpc.NewRoute(stack, jsii.String("private-route-b"), &vpc.RouteConfig{
		DestinationCidrBlock: jsii.String(destinationCidrBlock),
		RouteTableId:         privateRouteTableB.Id(),
		GatewayId:            natGatewayB.Id(),
	})

	vpc.NewRouteTableAssociation(stack, jsii.String("route-table-association-private-sub-b"),
		&vpc.RouteTableAssociationConfig{
			RouteTableId: privateRouteTableB.Id(),
			SubnetId:     privateSubnetB.Id(),
		})

	cdktf.NewTerraformOutput(stack, jsii.String("vpc-id"), &cdktf.TerraformOutputConfig{
		Value: newVpc.Id(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "cdktf-go-aws-vpc")
	cdktf.NewRemoteBackend(stack, &cdktf.RemoteBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("jigsaw373"),
		Workspaces:   cdktf.NewNamedRemoteWorkspace(jsii.String("cdktf-go-aws-vpc")),
	})

	app.Synth()
}
