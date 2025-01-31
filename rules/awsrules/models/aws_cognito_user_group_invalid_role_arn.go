// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCognitoUserGroupInvalidRoleArnRule checks the pattern is valid
type AwsCognitoUserGroupInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserGroupInvalidRoleArnRule returns new rule with default attributes
func NewAwsCognitoUserGroupInvalidRoleArnRule() *AwsCognitoUserGroupInvalidRoleArnRule {
	return &AwsCognitoUserGroupInvalidRoleArnRule{
		resourceType:  "aws_cognito_user_group",
		attributeName: "role_arn",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Name() string {
	return "aws_cognito_user_group_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserGroupInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`role_arn does not match valid pattern ^arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
