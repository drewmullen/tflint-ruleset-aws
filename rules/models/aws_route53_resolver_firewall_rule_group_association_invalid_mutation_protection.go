// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule checks the pattern is valid
type AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule returns new rule with default attributes
func NewAwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule() *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule {
	return &AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule{
		resourceType:  "aws_route53_resolver_firewall_rule_group_association",
		attributeName: "mutation_protection",
		enum: []string{
			"ENABLED",
			"DISABLED",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule) Name() string {
	return "aws_route53_resolver_firewall_rule_group_association_invalid_mutation_protection"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidMutationProtectionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as mutation_protection`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
