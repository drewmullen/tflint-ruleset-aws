// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigRemediationConfigurationInvalidTargetIDRule checks the pattern is valid
type AwsConfigRemediationConfigurationInvalidTargetIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigRemediationConfigurationInvalidTargetIDRule returns new rule with default attributes
func NewAwsConfigRemediationConfigurationInvalidTargetIDRule() *AwsConfigRemediationConfigurationInvalidTargetIDRule {
	return &AwsConfigRemediationConfigurationInvalidTargetIDRule{
		resourceType:  "aws_config_remediation_configuration",
		attributeName: "target_id",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigRemediationConfigurationInvalidTargetIDRule) Name() string {
	return "aws_config_remediation_configuration_invalid_target_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigRemediationConfigurationInvalidTargetIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigRemediationConfigurationInvalidTargetIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigRemediationConfigurationInvalidTargetIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigRemediationConfigurationInvalidTargetIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"target_id must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"target_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
