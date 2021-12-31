// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule checks the pattern is valid
type AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsSagemakerUserProfileInvalidSingleSignOnUserValueRule returns new rule with default attributes
func NewAwsSagemakerUserProfileInvalidSingleSignOnUserValueRule() *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule {
	return &AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule{
		resourceType:  "aws_sagemaker_user_profile",
		attributeName: "single_sign_on_user_value",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Name() string {
	return "aws_sagemaker_user_profile_invalid_single_sign_on_user_value"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerUserProfileInvalidSingleSignOnUserValueRule) Check(runner tflint.Runner) error {
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
					"single_sign_on_user_value must be 256 characters or less",
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
