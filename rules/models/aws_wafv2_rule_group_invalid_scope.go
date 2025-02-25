// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWafv2RuleGroupInvalidScopeRule checks the pattern is valid
type AwsWafv2RuleGroupInvalidScopeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsWafv2RuleGroupInvalidScopeRule returns new rule with default attributes
func NewAwsWafv2RuleGroupInvalidScopeRule() *AwsWafv2RuleGroupInvalidScopeRule {
	return &AwsWafv2RuleGroupInvalidScopeRule{
		resourceType:  "aws_wafv2_rule_group",
		attributeName: "scope",
		enum: []string{
			"CLOUDFRONT",
			"REGIONAL",
		},
	}
}

// Name returns the rule name
func (r *AwsWafv2RuleGroupInvalidScopeRule) Name() string {
	return "aws_wafv2_rule_group_invalid_scope"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafv2RuleGroupInvalidScopeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafv2RuleGroupInvalidScopeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafv2RuleGroupInvalidScopeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafv2RuleGroupInvalidScopeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as scope`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
