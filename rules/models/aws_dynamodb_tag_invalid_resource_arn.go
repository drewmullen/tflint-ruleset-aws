// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDynamoDBTagInvalidResourceArnRule checks the pattern is valid
type AwsDynamoDBTagInvalidResourceArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsDynamoDBTagInvalidResourceArnRule returns new rule with default attributes
func NewAwsDynamoDBTagInvalidResourceArnRule() *AwsDynamoDBTagInvalidResourceArnRule {
	return &AwsDynamoDBTagInvalidResourceArnRule{
		resourceType:  "aws_dynamodb_tag",
		attributeName: "resource_arn",
		max:           1283,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsDynamoDBTagInvalidResourceArnRule) Name() string {
	return "aws_dynamodb_tag_invalid_resource_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBTagInvalidResourceArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBTagInvalidResourceArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBTagInvalidResourceArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBTagInvalidResourceArnRule) Check(runner tflint.Runner) error {
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
					"resource_arn must be 1283 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"resource_arn must be 1 characters or higher",
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
