// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEksAddonInvalidResolveConflictsRule checks the pattern is valid
type AwsEksAddonInvalidResolveConflictsRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEksAddonInvalidResolveConflictsRule returns new rule with default attributes
func NewAwsEksAddonInvalidResolveConflictsRule() *AwsEksAddonInvalidResolveConflictsRule {
	return &AwsEksAddonInvalidResolveConflictsRule{
		resourceType:  "aws_eks_addon",
		attributeName: "resolve_conflicts",
		enum: []string{
			"OVERWRITE",
			"NONE",
		},
	}
}

// Name returns the rule name
func (r *AwsEksAddonInvalidResolveConflictsRule) Name() string {
	return "aws_eks_addon_invalid_resolve_conflicts"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEksAddonInvalidResolveConflictsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEksAddonInvalidResolveConflictsRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEksAddonInvalidResolveConflictsRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEksAddonInvalidResolveConflictsRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as resolve_conflicts`, truncateLongMessage(val)),
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
