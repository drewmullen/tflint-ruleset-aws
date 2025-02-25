// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSecurityhubActionTargetInvalidDescriptionRule checks the pattern is valid
type AwsSecurityhubActionTargetInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSecurityhubActionTargetInvalidDescriptionRule returns new rule with default attributes
func NewAwsSecurityhubActionTargetInvalidDescriptionRule() *AwsSecurityhubActionTargetInvalidDescriptionRule {
	return &AwsSecurityhubActionTargetInvalidDescriptionRule{
		resourceType:  "aws_securityhub_action_target",
		attributeName: "description",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsSecurityhubActionTargetInvalidDescriptionRule) Name() string {
	return "aws_securityhub_action_target_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecurityhubActionTargetInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecurityhubActionTargetInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecurityhubActionTargetInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecurityhubActionTargetInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
