// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSpotFleetRequestInvalidAllocationStrategyRule checks the pattern is valid
type AwsSpotFleetRequestInvalidAllocationStrategyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSpotFleetRequestInvalidAllocationStrategyRule returns new rule with default attributes
func NewAwsSpotFleetRequestInvalidAllocationStrategyRule() *AwsSpotFleetRequestInvalidAllocationStrategyRule {
	return &AwsSpotFleetRequestInvalidAllocationStrategyRule{
		resourceType:  "aws_spot_fleet_request",
		attributeName: "allocation_strategy",
		enum: []string{
			"lowestPrice",
			"diversified",
			"capacityOptimized",
			"capacityOptimizedPrioritized",
		},
	}
}

// Name returns the rule name
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Name() string {
	return "aws_spot_fleet_request_invalid_allocation_strategy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSpotFleetRequestInvalidAllocationStrategyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as allocation_strategy`, truncateLongMessage(val)),
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
