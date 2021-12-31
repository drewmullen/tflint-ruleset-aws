// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsApprunnerServiceInvalidServiceNameRule checks the pattern is valid
type AwsApprunnerServiceInvalidServiceNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsApprunnerServiceInvalidServiceNameRule returns new rule with default attributes
func NewAwsApprunnerServiceInvalidServiceNameRule() *AwsApprunnerServiceInvalidServiceNameRule {
	return &AwsApprunnerServiceInvalidServiceNameRule{
		resourceType:  "aws_apprunner_service",
		attributeName: "service_name",
		max:           40,
		min:           4,
		pattern:       regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9-_]{3,39}$`),
	}
}

// Name returns the rule name
func (r *AwsApprunnerServiceInvalidServiceNameRule) Name() string {
	return "aws_apprunner_service_invalid_service_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsApprunnerServiceInvalidServiceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsApprunnerServiceInvalidServiceNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsApprunnerServiceInvalidServiceNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsApprunnerServiceInvalidServiceNameRule) Check(runner tflint.Runner) error {
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
					"service_name must be 40 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"service_name must be 4 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9][A-Za-z0-9-_]{3,39}$`),
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
