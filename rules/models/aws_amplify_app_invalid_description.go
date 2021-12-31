// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyAppInvalidDescriptionRule checks the pattern is valid
type AwsAmplifyAppInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsAmplifyAppInvalidDescriptionRule returns new rule with default attributes
func NewAwsAmplifyAppInvalidDescriptionRule() *AwsAmplifyAppInvalidDescriptionRule {
	return &AwsAmplifyAppInvalidDescriptionRule{
		resourceType:  "aws_amplify_app",
		attributeName: "description",
		max:           1000,
	}
}

// Name returns the rule name
func (r *AwsAmplifyAppInvalidDescriptionRule) Name() string {
	return "aws_amplify_app_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyAppInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyAppInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyAppInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyAppInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 1000 characters or less",
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
