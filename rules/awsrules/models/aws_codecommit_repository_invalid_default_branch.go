// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCodecommitRepositoryInvalidDefaultBranchRule checks the pattern is valid
type AwsCodecommitRepositoryInvalidDefaultBranchRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodecommitRepositoryInvalidDefaultBranchRule returns new rule with default attributes
func NewAwsCodecommitRepositoryInvalidDefaultBranchRule() *AwsCodecommitRepositoryInvalidDefaultBranchRule {
	return &AwsCodecommitRepositoryInvalidDefaultBranchRule{
		resourceType:  "aws_codecommit_repository",
		attributeName: "default_branch",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Name() string {
	return "aws_codecommit_repository_invalid_default_branch"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodecommitRepositoryInvalidDefaultBranchRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"default_branch must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"default_branch must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
