# cdktf-go-aws-vpc

The Cloud Development Kit for Terraform (CDKTF) allows you to define your infrastructure in a familiar programming language such as TypeScript, Python, Go, C#, or Java.

In this tutorial, you will provision an EC2 instance on AWS using your preferred programming language.

## Prerequisites

* [Terraform](https://www.terraform.io/downloads) >= v1.0
* [CDK for Terraform](https://learn.hashicorp.com/tutorials/terraform/cdktf-install) >= v0.8
* A [Terraform Cloud](https://app.terraform.io/) account, with [CLI authentication](https://learn.hashicorp.com/tutorials/terraform/cloud-login) configured
* [an AWS account](https://portal.aws.amazon.com/billing/signup?nc2=h_ct&src=default&redirect_url=https%3A%2F%2Faws.amazon.com%2Fregistration-confirmation#/start)
* AWS Credentials [configured for use with Terraform](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication)


Credentials can be provided by using the AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, and optionally AWS_SESSION_TOKEN environment variables. The region can be set using the AWS_REGION or AWS_DEFAULT_REGION environment variables.

```shell
$ export AWS_ACCESS_KEY_ID="anaccesskey"
$ export AWS_SECRET_ACCESS_KEY="asecretkey"
$ export AWS_REGION="us-west-2"
```

## Install project dependencies

```shell
mkdir cdktf-go-aws-vpc
cd cdktf-go-aws-vpc
cdktf init --template="go"
```

## Install AWS provider
Open `cdktf.json` in your text editor, and add `aws` as one of the Terraform providers that you will use in the application.
```JSON
{
  "language": "go",
  "app": "go run *.go",
  "codeMakerOutput": "generated",
  "projectId": "02f2d864-a2f2-49e8-ab52-b472e233755e",
  "sendCrashReports": "false",
  "terraformProviders": [
	 "hashicorp/aws@~> 3.67.0"
  ],
  "terraformModules": [],
  "context": {
    "excludeStackIdFromLogicalIds": "true",
    "allowSepCharsInLogicalIds": "true"
  }
}
```
Run `cdktf get` to install the AWS provider you added to `cdktf.json`.
```SHELL
cdktf get
```

CDKTF uses a library called `jsii` to allow Go code to interact with CDK, 
which is written in TypeScript. 
Ensure that the jsii runtime is installed by running `go mod tidy`.

```SHELL
go mod tidy
```
## Provision infrastructure
```shell
cdktf deploy
```
After the instance is created, visit the AWS EC2 Dashboard.

## Clean up your infrastructure
```shell
cdktf destroy
```
