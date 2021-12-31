// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaFunctionInvalidS3ObjectVersionRule checks the pattern is valid
type AwsLambdaFunctionInvalidS3ObjectVersionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsLambdaFunctionInvalidS3ObjectVersionRule returns new rule with default attributes
func NewAwsLambdaFunctionInvalidS3ObjectVersionRule() *AwsLambdaFunctionInvalidS3ObjectVersionRule {
	return &AwsLambdaFunctionInvalidS3ObjectVersionRule{
		resourceType:  "aws_lambda_function",
		attributeName: "s3_object_version",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsLambdaFunctionInvalidS3ObjectVersionRule) Name() string {
	return "aws_lambda_function_invalid_s3_object_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaFunctionInvalidS3ObjectVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaFunctionInvalidS3ObjectVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaFunctionInvalidS3ObjectVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaFunctionInvalidS3ObjectVersionRule) Check(runner tflint.Runner) error {
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
					"s3_object_version must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"s3_object_version must be 1 characters or higher",
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
