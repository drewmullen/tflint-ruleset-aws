// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDevicefarmProjectInvalidNameRule checks the pattern is valid
type AwsDevicefarmProjectInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsDevicefarmProjectInvalidNameRule returns new rule with default attributes
func NewAwsDevicefarmProjectInvalidNameRule() *AwsDevicefarmProjectInvalidNameRule {
	return &AwsDevicefarmProjectInvalidNameRule{
		resourceType:  "aws_devicefarm_project",
		attributeName: "name",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsDevicefarmProjectInvalidNameRule) Name() string {
	return "aws_devicefarm_project_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDevicefarmProjectInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDevicefarmProjectInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDevicefarmProjectInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDevicefarmProjectInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 256 characters or less",
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
