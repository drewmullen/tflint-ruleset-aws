// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule checks the pattern is valid
type AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule returns new rule with default attributes
func NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule() *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule {
	return &AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule{
		resourceType:  "aws_worklink_website_certificate_authority_association",
		attributeName: "fleet_arn",
		max:           2048,
		min:           20,
	}
}

// Name returns the rule name
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule) Name() string {
	return "aws_worklink_website_certificate_authority_association_invalid_fleet_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule) Check(runner tflint.Runner) error {
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
					"fleet_arn must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"fleet_arn must be 20 characters or higher",
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
