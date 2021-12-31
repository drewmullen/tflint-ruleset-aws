// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsPlacementGroupInvalidStrategyRule checks the pattern is valid
type AwsPlacementGroupInvalidStrategyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsPlacementGroupInvalidStrategyRule returns new rule with default attributes
func NewAwsPlacementGroupInvalidStrategyRule() *AwsPlacementGroupInvalidStrategyRule {
	return &AwsPlacementGroupInvalidStrategyRule{
		resourceType:  "aws_placement_group",
		attributeName: "strategy",
		enum: []string{
			"cluster",
			"spread",
			"partition",
		},
	}
}

// Name returns the rule name
func (r *AwsPlacementGroupInvalidStrategyRule) Name() string {
	return "aws_placement_group_invalid_strategy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsPlacementGroupInvalidStrategyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsPlacementGroupInvalidStrategyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsPlacementGroupInvalidStrategyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsPlacementGroupInvalidStrategyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as strategy`, truncateLongMessage(val)),
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
