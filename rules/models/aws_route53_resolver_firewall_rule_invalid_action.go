// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverFirewallRuleInvalidActionRule checks the pattern is valid
type AwsRoute53ResolverFirewallRuleInvalidActionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53ResolverFirewallRuleInvalidActionRule returns new rule with default attributes
func NewAwsRoute53ResolverFirewallRuleInvalidActionRule() *AwsRoute53ResolverFirewallRuleInvalidActionRule {
	return &AwsRoute53ResolverFirewallRuleInvalidActionRule{
		resourceType:  "aws_route53_resolver_firewall_rule",
		attributeName: "action",
		enum: []string{
			"ALLOW",
			"BLOCK",
			"ALERT",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Name() string {
	return "aws_route53_resolver_firewall_rule_invalid_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverFirewallRuleInvalidActionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as action`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
