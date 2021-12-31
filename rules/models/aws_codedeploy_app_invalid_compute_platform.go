// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodedeployAppInvalidComputePlatformRule checks the pattern is valid
type AwsCodedeployAppInvalidComputePlatformRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodedeployAppInvalidComputePlatformRule returns new rule with default attributes
func NewAwsCodedeployAppInvalidComputePlatformRule() *AwsCodedeployAppInvalidComputePlatformRule {
	return &AwsCodedeployAppInvalidComputePlatformRule{
		resourceType:  "aws_codedeploy_app",
		attributeName: "compute_platform",
		enum: []string{
			"Server",
			"Lambda",
			"ECS",
		},
	}
}

// Name returns the rule name
func (r *AwsCodedeployAppInvalidComputePlatformRule) Name() string {
	return "aws_codedeploy_app_invalid_compute_platform"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodedeployAppInvalidComputePlatformRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodedeployAppInvalidComputePlatformRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodedeployAppInvalidComputePlatformRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodedeployAppInvalidComputePlatformRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as compute_platform`, truncateLongMessage(val)),
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
