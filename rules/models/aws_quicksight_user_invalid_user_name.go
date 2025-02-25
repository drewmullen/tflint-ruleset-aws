// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightUserInvalidUserNameRule checks the pattern is valid
type AwsQuicksightUserInvalidUserNameRule struct {
	resourceType  string
	attributeName string
	min           int
	pattern       *regexp.Regexp
}

// NewAwsQuicksightUserInvalidUserNameRule returns new rule with default attributes
func NewAwsQuicksightUserInvalidUserNameRule() *AwsQuicksightUserInvalidUserNameRule {
	return &AwsQuicksightUserInvalidUserNameRule{
		resourceType:  "aws_quicksight_user",
		attributeName: "user_name",
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsQuicksightUserInvalidUserNameRule) Name() string {
	return "aws_quicksight_user_invalid_user_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightUserInvalidUserNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightUserInvalidUserNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightUserInvalidUserNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightUserInvalidUserNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"user_name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\x{0020}-\x{00FF}]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
