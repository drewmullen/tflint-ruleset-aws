// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule checks the pattern is valid
type AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule returns new rule with default attributes
func NewAwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule() *AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule {
	return &AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule{
		resourceType:  "aws_sagemaker_studio_lifecycle_config",
		attributeName: "studio_lifecycle_config_app_type",
		enum: []string{
			"JupyterServer",
			"KernelGateway",
		},
	}
}

// Name returns the rule name
func (r *AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule) Name() string {
	return "aws_sagemaker_studio_lifecycle_config_invalid_studio_lifecycle_config_app_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerStudioLifecycleConfigInvalidStudioLifecycleConfigAppTypeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as studio_lifecycle_config_app_type`, truncateLongMessage(val)),
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
