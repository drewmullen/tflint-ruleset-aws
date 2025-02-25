// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsQuicksightUserInvalidIdentityTypeRule checks the pattern is valid
type AwsQuicksightUserInvalidIdentityTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsQuicksightUserInvalidIdentityTypeRule returns new rule with default attributes
func NewAwsQuicksightUserInvalidIdentityTypeRule() *AwsQuicksightUserInvalidIdentityTypeRule {
	return &AwsQuicksightUserInvalidIdentityTypeRule{
		resourceType:  "aws_quicksight_user",
		attributeName: "identity_type",
		enum: []string{
			"IAM",
			"QUICKSIGHT",
		},
	}
}

// Name returns the rule name
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Name() string {
	return "aws_quicksight_user_invalid_identity_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsQuicksightUserInvalidIdentityTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as identity_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
