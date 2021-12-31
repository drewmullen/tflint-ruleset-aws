// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppsyncResolverInvalidFieldRule checks the pattern is valid
type AwsAppsyncResolverInvalidFieldRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAppsyncResolverInvalidFieldRule returns new rule with default attributes
func NewAwsAppsyncResolverInvalidFieldRule() *AwsAppsyncResolverInvalidFieldRule {
	return &AwsAppsyncResolverInvalidFieldRule{
		resourceType:  "aws_appsync_resolver",
		attributeName: "field",
		max:           65536,
		min:           1,
		pattern:       regexp.MustCompile(`^[_A-Za-z][_0-9A-Za-z]*$`),
	}
}

// Name returns the rule name
func (r *AwsAppsyncResolverInvalidFieldRule) Name() string {
	return "aws_appsync_resolver_invalid_field"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncResolverInvalidFieldRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncResolverInvalidFieldRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncResolverInvalidFieldRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncResolverInvalidFieldRule) Check(runner tflint.Runner) error {
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
					"field must be 65536 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"field must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[_A-Za-z][_0-9A-Za-z]*$`),
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
