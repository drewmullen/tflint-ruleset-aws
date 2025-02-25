// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule checks the pattern is valid
type AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule returns new rule with default attributes
func NewAwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule() *AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule {
	return &AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule{
		resourceType:  "aws_fsx_openzfs_file_system",
		attributeName: "weekly_maintenance_start_time",
		max:           7,
		min:           7,
		pattern:       regexp.MustCompile(`^[1-7]:([01]\d|2[0-3]):?([0-5]\d)$`),
	}
}

// Name returns the rule name
func (r *AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Name() string {
	return "aws_fsx_openzfs_file_system_invalid_weekly_maintenance_start_time"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxOpenzfsFileSystemInvalidWeeklyMaintenanceStartTimeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"weekly_maintenance_start_time must be 7 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"weekly_maintenance_start_time must be 7 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[1-7]:([01]\d|2[0-3]):?([0-5]\d)$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
