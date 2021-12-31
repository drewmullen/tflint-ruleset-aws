// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsMacie2FindingsFilterInvalidActionRule checks the pattern is valid
type AwsMacie2FindingsFilterInvalidActionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsMacie2FindingsFilterInvalidActionRule returns new rule with default attributes
func NewAwsMacie2FindingsFilterInvalidActionRule() *AwsMacie2FindingsFilterInvalidActionRule {
	return &AwsMacie2FindingsFilterInvalidActionRule{
		resourceType:  "aws_macie2_findings_filter",
		attributeName: "action",
		enum: []string{
			"ARCHIVE",
			"NOOP",
		},
	}
}

// Name returns the rule name
func (r *AwsMacie2FindingsFilterInvalidActionRule) Name() string {
	return "aws_macie2_findings_filter_invalid_action"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacie2FindingsFilterInvalidActionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacie2FindingsFilterInvalidActionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacie2FindingsFilterInvalidActionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacie2FindingsFilterInvalidActionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as action`, truncateLongMessage(val)),
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
