// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventRuleInvalidDescriptionRule checks the pattern is valid
type AwsCloudwatchEventRuleInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloudwatchEventRuleInvalidDescriptionRule returns new rule with default attributes
func NewAwsCloudwatchEventRuleInvalidDescriptionRule() *AwsCloudwatchEventRuleInvalidDescriptionRule {
	return &AwsCloudwatchEventRuleInvalidDescriptionRule{
		resourceType:  "aws_cloudwatch_event_rule",
		attributeName: "description",
		max:           512,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Name() string {
	return "aws_cloudwatch_event_rule_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventRuleInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 512 characters or less",
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
