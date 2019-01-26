# The only services that are valid within aws_service_access_principals
#AWS Artifact
#AWS CloudTrail
#AWS Config
#AWS Directory Service
#AWS Firewall Manager
#AWS Resource Access Manager
#AWS Single Sign-On
#AWS License Manager
#AWS Service Catalog
resource "aws_organizations_organization" "organization" {

//  aws_service_access_principals = []
  feature_set = "ALL"

}

resource "aws_organizations_account" "account" {
  name = "${var.account_name}"
  email = "${var.account_email}"
  iam_user_access_to_billing = "ALLOW"
  role_name = "${var.account_role_name}"
}

resource "aws_organizations_policy" "ec2_policy" {
  name = ec2_policy

  content = <<CONTENT
{
  "Version": "2012-10-17",
  "Statement":  {
    "Action": "ec2:*",
    "Effect": "Allow",
    "Resource": "*"
  }
}
CONTENT
}

resource "aws_organizations_policy_attachment" "attachment" {
  policy_id = "${aws_organizations_policy.ec2_policy.id}"
  target_id = "${aws_organizations_account.account.id}"
}
