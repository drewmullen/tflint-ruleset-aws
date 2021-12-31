// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ZoneAssociationInvalidVpcIDRule checks the pattern is valid
type AwsRoute53ZoneAssociationInvalidVpcIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRoute53ZoneAssociationInvalidVpcIDRule returns new rule with default attributes
func NewAwsRoute53ZoneAssociationInvalidVpcIDRule() *AwsRoute53ZoneAssociationInvalidVpcIDRule {
	return &AwsRoute53ZoneAssociationInvalidVpcIDRule{
		resourceType:  "aws_route53_zone_association",
		attributeName: "vpc_id",
		max:           1024,
	}
}

// Name returns the rule name
func (r *AwsRoute53ZoneAssociationInvalidVpcIDRule) Name() string {
	return "aws_route53_zone_association_invalid_vpc_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ZoneAssociationInvalidVpcIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ZoneAssociationInvalidVpcIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ZoneAssociationInvalidVpcIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ZoneAssociationInvalidVpcIDRule) Check(runner tflint.Runner) error {
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
					"vpc_id must be 1024 characters or less",
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
