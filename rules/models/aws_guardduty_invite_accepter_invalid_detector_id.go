// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyInviteAccepterInvalidDetectorIDRule checks the pattern is valid
type AwsGuarddutyInviteAccepterInvalidDetectorIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyInviteAccepterInvalidDetectorIDRule returns new rule with default attributes
func NewAwsGuarddutyInviteAccepterInvalidDetectorIDRule() *AwsGuarddutyInviteAccepterInvalidDetectorIDRule {
	return &AwsGuarddutyInviteAccepterInvalidDetectorIDRule{
		resourceType:  "aws_guardduty_invite_accepter",
		attributeName: "detector_id",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyInviteAccepterInvalidDetectorIDRule) Name() string {
	return "aws_guardduty_invite_accepter_invalid_detector_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyInviteAccepterInvalidDetectorIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyInviteAccepterInvalidDetectorIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyInviteAccepterInvalidDetectorIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyInviteAccepterInvalidDetectorIDRule) Check(runner tflint.Runner) error {
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
					"detector_id must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"detector_id must be 1 characters or higher",
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
