resource "aws_iam_policy" "orgs_policy" {
  name = "organisations_policy"
  path = "/"
  description = "Policy to allow management of organisations"

  policy = <<EOF

{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "organizations:CreateAccount",
        "organizations:CreateOrganization",
        "organizations:CreateOrganizationalUnit"
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
