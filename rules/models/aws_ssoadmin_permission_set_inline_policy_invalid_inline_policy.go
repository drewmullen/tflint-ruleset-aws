// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule checks the pattern is valid
type AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule returns new rule with default attributes
func NewAwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule() *AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule {
	return &AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule{
		resourceType:  "aws_ssoadmin_permission_set_inline_policy",
		attributeName: "inline_policy",
		max:           10240,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
	}
}

// Name returns the rule name
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule) Name() string {
	return "aws_ssoadmin_permission_set_inline_policy_invalid_inline_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminPermissionSetInlinePolicyInvalidInlinePolicyRule) Check(runner tflint.Runner) error {
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
					"inline_policy must be 10240 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"inline_policy must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\x{0009}\x{000A}\x{000D}\x{0020}-\x{00FF}]+$`),
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
