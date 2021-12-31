// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSignerSigningProfilePermissionInvalidProfileVersionRule checks the pattern is valid
type AwsSignerSigningProfilePermissionInvalidProfileVersionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSignerSigningProfilePermissionInvalidProfileVersionRule returns new rule with default attributes
func NewAwsSignerSigningProfilePermissionInvalidProfileVersionRule() *AwsSignerSigningProfilePermissionInvalidProfileVersionRule {
	return &AwsSignerSigningProfilePermissionInvalidProfileVersionRule{
		resourceType:  "aws_signer_signing_profile_permission",
		attributeName: "profile_version",
		max:           10,
		min:           10,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9]{10}$`),
	}
}

// Name returns the rule name
func (r *AwsSignerSigningProfilePermissionInvalidProfileVersionRule) Name() string {
	return "aws_signer_signing_profile_permission_invalid_profile_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSignerSigningProfilePermissionInvalidProfileVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSignerSigningProfilePermissionInvalidProfileVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSignerSigningProfilePermissionInvalidProfileVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSignerSigningProfilePermissionInvalidProfileVersionRule) Check(runner tflint.Runner) error {
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
					"profile_version must be 10 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"profile_version must be 10 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9]{10}$`),
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
