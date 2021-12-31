// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchMetricAlarmInvalidAlarmNameRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidAlarmNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchMetricAlarmInvalidAlarmNameRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidAlarmNameRule() *AwsCloudwatchMetricAlarmInvalidAlarmNameRule {
	return &AwsCloudwatchMetricAlarmInvalidAlarmNameRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "alarm_name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidAlarmNameRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_alarm_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidAlarmNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidAlarmNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidAlarmNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidAlarmNameRule) Check(runner tflint.Runner) error {
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
					"alarm_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"alarm_name must be 1 characters or higher",
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
