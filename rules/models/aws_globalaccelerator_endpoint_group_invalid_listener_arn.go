// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule checks the pattern is valid
type AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsGlobalacceleratorEndpointGroupInvalidListenerArnRule returns new rule with default attributes
func NewAwsGlobalacceleratorEndpointGroupInvalidListenerArnRule() *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule {
	return &AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule{
		resourceType:  "aws_globalaccelerator_endpoint_group",
		attributeName: "listener_arn",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Name() string {
	return "aws_globalaccelerator_endpoint_group_invalid_listener_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGlobalacceleratorEndpointGroupInvalidListenerArnRule) Check(runner tflint.Runner) error {
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
					"listener_arn must be 255 characters or less",
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
