// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule checks the pattern is valid
type AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule returns new rule with default attributes
func NewAwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule() *AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule {
	return &AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule{
		resourceType:  "aws_fsx_windows_file_system",
		attributeName: "weekly_maintenance_start_time",
		max:           7,
		min:           7,
		pattern:       regexp.MustCompile(`^[1-7]:([01]\d|2[0-3]):?([0-5]\d)$`),
	}
}

// Name returns the rule name
func (r *AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Name() string {
	return "aws_fsx_windows_file_system_invalid_weekly_maintenance_start_time"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Check(runner tflint.Runner) error {
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
					"weekly_maintenance_start_time must be 7 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"weekly_maintenance_start_time must be 7 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[1-7]:([01]\d|2[0-3]):?([0-5]\d)$`),
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
