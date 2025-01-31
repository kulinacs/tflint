package terraformrules

import (
	"fmt"
	"log"

	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/project"
	"github.com/wata727/tflint/tflint"
)

// TerraformDocumentedOutputsRule checks whether outputs have descriptions
type TerraformDocumentedOutputsRule struct{}

// NewTerraformDocumentedOutputsRule returns a new rule
func NewTerraformDocumentedOutputsRule() *TerraformDocumentedOutputsRule {
	return &TerraformDocumentedOutputsRule{}
}

// Name returns the rule name
func (r *TerraformDocumentedOutputsRule) Name() string {
	return "terraform_documented_outputs"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformDocumentedOutputsRule) Enabled() bool {
	return false
}

// Type returns the rule severity
func (r *TerraformDocumentedOutputsRule) Type() string {
	return issue.NOTICE
}

// Link returns the rule reference link
func (r *TerraformDocumentedOutputsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether outputs have descriptions
func (r *TerraformDocumentedOutputsRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	for _, output := range runner.TFConfig.Module.Outputs {
		if output.Description == "" {
			runner.EmitIssue(
				r,
				fmt.Sprintf("`%s` output has no description", output.Name),
				output.DeclRange,
			)
		}
	}

	return nil
}
