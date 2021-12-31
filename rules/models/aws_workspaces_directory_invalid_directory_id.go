// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWorkspacesDirectoryInvalidDirectoryIDRule checks the pattern is valid
type AwsWorkspacesDirectoryInvalidDirectoryIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWorkspacesDirectoryInvalidDirectoryIDRule returns new rule with default attributes
func NewAwsWorkspacesDirectoryInvalidDirectoryIDRule() *AwsWorkspacesDirectoryInvalidDirectoryIDRule {
	return &AwsWorkspacesDirectoryInvalidDirectoryIDRule{
		resourceType:  "aws_workspaces_directory",
		attributeName: "directory_id",
		max:           65,
		min:           10,
		pattern:       regexp.MustCompile(`^d-[0-9a-f]{8,63}$`),
	}
}

// Name returns the rule name
func (r *AwsWorkspacesDirectoryInvalidDirectoryIDRule) Name() string {
	return "aws_workspaces_directory_invalid_directory_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorkspacesDirectoryInvalidDirectoryIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWorkspacesDirectoryInvalidDirectoryIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWorkspacesDirectoryInvalidDirectoryIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorkspacesDirectoryInvalidDirectoryIDRule) Check(runner tflint.Runner) error {
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
					"directory_id must be 65 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"directory_id must be 10 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^d-[0-9a-f]{8,63}$`),
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
