// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcsTaskSetInvalidLaunchTypeRule checks the pattern is valid
type AwsEcsTaskSetInvalidLaunchTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEcsTaskSetInvalidLaunchTypeRule returns new rule with default attributes
func NewAwsEcsTaskSetInvalidLaunchTypeRule() *AwsEcsTaskSetInvalidLaunchTypeRule {
	return &AwsEcsTaskSetInvalidLaunchTypeRule{
		resourceType:  "aws_ecs_task_set",
		attributeName: "launch_type",
		enum: []string{
			"EC2",
			"FARGATE",
			"EXTERNAL",
		},
	}
}

// Name returns the rule name
func (r *AwsEcsTaskSetInvalidLaunchTypeRule) Name() string {
	return "aws_ecs_task_set_invalid_launch_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcsTaskSetInvalidLaunchTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcsTaskSetInvalidLaunchTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcsTaskSetInvalidLaunchTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcsTaskSetInvalidLaunchTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as launch_type`, truncateLongMessage(val)),
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
