package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	// renamed to aws2 to avoid name collision
	aws2 "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformVpcTemplate(t *testing.T) {
	t.Parallel()

	terraformDir := test_structure.CopyTerraformFolderToTemp(t, "../../", "terraform")

	defer test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
		terraform.Destroy(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := createTerraformOptions(t, terraformDir)
		test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)
	})

	test_structure.RunTestStage(t, "test VPC cidr is acceptable", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, terraformDir)
		testVpcCidrIsCorrect(t, terraformOptions)
	})

}

func createTerraformOptions(t *testing.T, terraformDir string) *terraform.Options {

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	// Issue found with this is you can come across dodgy regions without as much support and fine code breaks, like
	// regions not having enough AZ's to support 3 separate subnets in TF code
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	terraformOptions := &terraform.Options{
		TerraformDir: "../../terraform",

		Vars: map[string]interface{}{
			"aws_region": awsRegion,
		},

		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	}

	return terraformOptions

}

// test: testVpcCidrIsCorrect
// assert that the user configured vpcCidr within the terraform template is
// an acceptable value
func testVpcCidrIsCorrect(t *testing.T, terraformOptions *terraform.Options) { // {{{
	vpcCidr := terraform.Output(t, terraformOptions, "vpc_cidr")

	assert.Equal(t, vpcCidr, "10.0.0.0/16")
} // }}}

// test: aws_orgs_org - organisations feature set is set to ALL

// test: aws_orgs_acc - email is valid according to our personal domains

// test: aws_orgs_acc - iam user access to billing is enabled

// potential test: ensure that role_name fits acceptance standards
// basically, we will have policy around how to name our roles

// test: aws_org_pol - ensure that service is limited to ec2 for certain accs

// test: aws_org_pol - must be other tests there
