// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventBusPolicyInvalidEventBusNameRule checks the pattern is valid
type AwsCloudwatchEventBusPolicyInvalidEventBusNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventBusPolicyInvalidEventBusNameRule returns new rule with default attributes
func NewAwsCloudwatchEventBusPolicyInvalidEventBusNameRule() *AwsCloudwatchEventBusPolicyInvalidEventBusNameRule {
	return &AwsCloudwatchEventBusPolicyInvalidEventBusNameRule{
		resourceType:  "aws_cloudwatch_event_bus_policy",
		attributeName: "event_bus_name",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^[/\.\-_A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventBusPolicyInvalidEventBusNameRule) Name() string {
	return "aws_cloudwatch_event_bus_policy_invalid_event_bus_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventBusPolicyInvalidEventBusNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventBusPolicyInvalidEventBusNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventBusPolicyInvalidEventBusNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventBusPolicyInvalidEventBusNameRule) Check(runner tflint.Runner) error {
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
					"event_bus_name must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"event_bus_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[/\.\-_A-Za-z0-9]+$`),
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
