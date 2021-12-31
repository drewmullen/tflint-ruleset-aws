// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3BucketObjectInvalidStorageClassRule checks the pattern is valid
type AwsS3BucketObjectInvalidStorageClassRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketObjectInvalidStorageClassRule returns new rule with default attributes
func NewAwsS3BucketObjectInvalidStorageClassRule() *AwsS3BucketObjectInvalidStorageClassRule {
	return &AwsS3BucketObjectInvalidStorageClassRule{
		resourceType:  "aws_s3_bucket_object",
		attributeName: "storage_class",
		enum: []string{
			"STANDARD",
			"REDUCED_REDUNDANCY",
			"STANDARD_IA",
			"ONEZONE_IA",
			"INTELLIGENT_TIERING",
			"GLACIER",
			"DEEP_ARCHIVE",
			"OUTPOSTS",
			"GLACIER_IR",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketObjectInvalidStorageClassRule) Name() string {
	return "aws_s3_bucket_object_invalid_storage_class"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketObjectInvalidStorageClassRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3BucketObjectInvalidStorageClassRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketObjectInvalidStorageClassRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketObjectInvalidStorageClassRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as storage_class`, truncateLongMessage(val)),
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
