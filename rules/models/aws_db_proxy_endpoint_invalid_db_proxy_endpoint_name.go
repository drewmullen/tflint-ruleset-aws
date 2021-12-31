// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDBProxyEndpointInvalidDBProxyEndpointNameRule checks the pattern is valid
type AwsDBProxyEndpointInvalidDBProxyEndpointNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsDBProxyEndpointInvalidDBProxyEndpointNameRule returns new rule with default attributes
func NewAwsDBProxyEndpointInvalidDBProxyEndpointNameRule() *AwsDBProxyEndpointInvalidDBProxyEndpointNameRule {
	return &AwsDBProxyEndpointInvalidDBProxyEndpointNameRule{
		resourceType:  "aws_db_proxy_endpoint",
		attributeName: "db_proxy_endpoint_name",
		max:           63,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsDBProxyEndpointInvalidDBProxyEndpointNameRule) Name() string {
	return "aws_db_proxy_endpoint_invalid_db_proxy_endpoint_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDBProxyEndpointInvalidDBProxyEndpointNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDBProxyEndpointInvalidDBProxyEndpointNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDBProxyEndpointInvalidDBProxyEndpointNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDBProxyEndpointInvalidDBProxyEndpointNameRule) Check(runner tflint.Runner) error {
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
					"db_proxy_endpoint_name must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"db_proxy_endpoint_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$`),
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
