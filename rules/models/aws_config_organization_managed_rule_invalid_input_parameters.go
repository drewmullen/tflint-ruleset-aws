// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationManagedRuleInvalidInputParametersRule checks the pattern is valid
type AwsConfigOrganizationManagedRuleInvalidInputParametersRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationManagedRuleInvalidInputParametersRule returns new rule with default attributes
func NewAwsConfigOrganizationManagedRuleInvalidInputParametersRule() *AwsConfigOrganizationManagedRuleInvalidInputParametersRule {
	return &AwsConfigOrganizationManagedRuleInvalidInputParametersRule{
		resourceType:  "aws_config_organization_managed_rule",
		attributeName: "input_parameters",
		max:           2048,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationManagedRuleInvalidInputParametersRule) Name() string {
	return "aws_config_organization_managed_rule_invalid_input_parameters"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationManagedRuleInvalidInputParametersRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationManagedRuleInvalidInputParametersRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationManagedRuleInvalidInputParametersRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationManagedRuleInvalidInputParametersRule) Check(runner tflint.Runner) error {
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
					"input_parameters must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"input_parameters must be 1 characters or higher",
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
