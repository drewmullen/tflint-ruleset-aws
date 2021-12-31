// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsKinesisAnalyticsApplicationInvalidNameRule checks the pattern is valid
type AwsKinesisAnalyticsApplicationInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsKinesisAnalyticsApplicationInvalidNameRule returns new rule with default attributes
func NewAwsKinesisAnalyticsApplicationInvalidNameRule() *AwsKinesisAnalyticsApplicationInvalidNameRule {
	return &AwsKinesisAnalyticsApplicationInvalidNameRule{
		resourceType:  "aws_kinesis_analytics_application",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsKinesisAnalyticsApplicationInvalidNameRule) Name() string {
	return "aws_kinesis_analytics_application_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKinesisAnalyticsApplicationInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKinesisAnalyticsApplicationInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKinesisAnalyticsApplicationInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKinesisAnalyticsApplicationInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_.-]+$`),
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
