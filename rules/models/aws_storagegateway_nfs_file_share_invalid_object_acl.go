// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayNfsFileShareInvalidObjectACLRule checks the pattern is valid
type AwsStoragegatewayNfsFileShareInvalidObjectACLRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsStoragegatewayNfsFileShareInvalidObjectACLRule returns new rule with default attributes
func NewAwsStoragegatewayNfsFileShareInvalidObjectACLRule() *AwsStoragegatewayNfsFileShareInvalidObjectACLRule {
	return &AwsStoragegatewayNfsFileShareInvalidObjectACLRule{
		resourceType:  "aws_storagegateway_nfs_file_share",
		attributeName: "object_acl",
		enum: []string{
			"private",
			"public-read",
			"public-read-write",
			"authenticated-read",
			"bucket-owner-read",
			"bucket-owner-full-control",
			"aws-exec-read",
		},
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Name() string {
	return "aws_storagegateway_nfs_file_share_invalid_object_acl"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayNfsFileShareInvalidObjectACLRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as object_acl`, truncateLongMessage(val)),
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
