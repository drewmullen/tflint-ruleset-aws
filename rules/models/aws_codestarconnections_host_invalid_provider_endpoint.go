// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodestarconnectionsHostInvalidProviderEndpointRule checks the pattern is valid
type AwsCodestarconnectionsHostInvalidProviderEndpointRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodestarconnectionsHostInvalidProviderEndpointRule returns new rule with default attributes
func NewAwsCodestarconnectionsHostInvalidProviderEndpointRule() *AwsCodestarconnectionsHostInvalidProviderEndpointRule {
	return &AwsCodestarconnectionsHostInvalidProviderEndpointRule{
		resourceType:  "aws_codestarconnections_host",
		attributeName: "provider_endpoint",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AwsCodestarconnectionsHostInvalidProviderEndpointRule) Name() string {
	return "aws_codestarconnections_host_invalid_provider_endpoint"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodestarconnectionsHostInvalidProviderEndpointRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodestarconnectionsHostInvalidProviderEndpointRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodestarconnectionsHostInvalidProviderEndpointRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodestarconnectionsHostInvalidProviderEndpointRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"provider_endpoint must be 512 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"provider_endpoint must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
