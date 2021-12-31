// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAccountAlternateContactInvalidEmailAddressRule checks the pattern is valid
type AwsAccountAlternateContactInvalidEmailAddressRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsAccountAlternateContactInvalidEmailAddressRule returns new rule with default attributes
func NewAwsAccountAlternateContactInvalidEmailAddressRule() *AwsAccountAlternateContactInvalidEmailAddressRule {
	return &AwsAccountAlternateContactInvalidEmailAddressRule{
		resourceType:  "aws_account_alternate_contact",
		attributeName: "email_address",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.-]+@[\w.-]+\.[\w]+$`),
	}
}

// Name returns the rule name
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Name() string {
	return "aws_account_alternate_contact_invalid_email_address"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAccountAlternateContactInvalidEmailAddressRule) Check(runner tflint.Runner) error {
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
					"email_address must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"email_address must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=,.-]+@[\w.-]+\.[\w]+$`),
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
