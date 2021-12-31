// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloud9EnvironmentEc2InvalidDescriptionRule checks the pattern is valid
type AwsCloud9EnvironmentEc2InvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloud9EnvironmentEc2InvalidDescriptionRule returns new rule with default attributes
func NewAwsCloud9EnvironmentEc2InvalidDescriptionRule() *AwsCloud9EnvironmentEc2InvalidDescriptionRule {
	return &AwsCloud9EnvironmentEc2InvalidDescriptionRule{
		resourceType:  "aws_cloud9_environment_ec2",
		attributeName: "description",
		max:           200,
	}
}

// Name returns the rule name
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Name() string {
	return "aws_cloud9_environment_ec2_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloud9EnvironmentEc2InvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 200 characters or less",
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
