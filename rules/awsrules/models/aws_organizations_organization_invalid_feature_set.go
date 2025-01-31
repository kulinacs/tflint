// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsOrganizationsOrganizationInvalidFeatureSetRule checks the pattern is valid
type AwsOrganizationsOrganizationInvalidFeatureSetRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOrganizationsOrganizationInvalidFeatureSetRule returns new rule with default attributes
func NewAwsOrganizationsOrganizationInvalidFeatureSetRule() *AwsOrganizationsOrganizationInvalidFeatureSetRule {
	return &AwsOrganizationsOrganizationInvalidFeatureSetRule{
		resourceType:  "aws_organizations_organization",
		attributeName: "feature_set",
		enum: []string{
			"ALL",
			"CONSOLIDATED_BILLING",
		},
	}
}

// Name returns the rule name
func (r *AwsOrganizationsOrganizationInvalidFeatureSetRule) Name() string {
	return "aws_organizations_organization_invalid_feature_set"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsOrganizationInvalidFeatureSetRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsOrganizationsOrganizationInvalidFeatureSetRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsOrganizationInvalidFeatureSetRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsOrganizationInvalidFeatureSetRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`feature_set is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
