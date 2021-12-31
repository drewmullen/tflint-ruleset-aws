// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsALBTargetGroupInvalidTargetTypeRule checks the pattern is valid
type AwsALBTargetGroupInvalidTargetTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsALBTargetGroupInvalidTargetTypeRule returns new rule with default attributes
func NewAwsALBTargetGroupInvalidTargetTypeRule() *AwsALBTargetGroupInvalidTargetTypeRule {
	return &AwsALBTargetGroupInvalidTargetTypeRule{
		resourceType:  "aws_alb_target_group",
		attributeName: "target_type",
		enum: []string{
			"instance",
			"ip",
			"lambda",
			"alb",
		},
	}
}

// Name returns the rule name
func (r *AwsALBTargetGroupInvalidTargetTypeRule) Name() string {
	return "aws_alb_target_group_invalid_target_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsALBTargetGroupInvalidTargetTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsALBTargetGroupInvalidTargetTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsALBTargetGroupInvalidTargetTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsALBTargetGroupInvalidTargetTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as target_type`, truncateLongMessage(val)),
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
