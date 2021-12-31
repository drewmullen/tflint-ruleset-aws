// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftClusterInvalidOwnerAccountRule checks the pattern is valid
type AwsRedshiftClusterInvalidOwnerAccountRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftClusterInvalidOwnerAccountRule returns new rule with default attributes
func NewAwsRedshiftClusterInvalidOwnerAccountRule() *AwsRedshiftClusterInvalidOwnerAccountRule {
	return &AwsRedshiftClusterInvalidOwnerAccountRule{
		resourceType:  "aws_redshift_cluster",
		attributeName: "owner_account",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftClusterInvalidOwnerAccountRule) Name() string {
	return "aws_redshift_cluster_invalid_owner_account"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftClusterInvalidOwnerAccountRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftClusterInvalidOwnerAccountRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftClusterInvalidOwnerAccountRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftClusterInvalidOwnerAccountRule) Check(runner tflint.Runner) error {
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
					"owner_account must be 2147483647 characters or less",
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
