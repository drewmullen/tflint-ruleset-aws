// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyDomainAssociationInvalidDomainNameRule checks the pattern is valid
type AwsAmplifyDomainAssociationInvalidDomainNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsAmplifyDomainAssociationInvalidDomainNameRule returns new rule with default attributes
func NewAwsAmplifyDomainAssociationInvalidDomainNameRule() *AwsAmplifyDomainAssociationInvalidDomainNameRule {
	return &AwsAmplifyDomainAssociationInvalidDomainNameRule{
		resourceType:  "aws_amplify_domain_association",
		attributeName: "domain_name",
		max:           255,
	}
}

// Name returns the rule name
func (r *AwsAmplifyDomainAssociationInvalidDomainNameRule) Name() string {
	return "aws_amplify_domain_association_invalid_domain_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyDomainAssociationInvalidDomainNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyDomainAssociationInvalidDomainNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyDomainAssociationInvalidDomainNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyDomainAssociationInvalidDomainNameRule) Check(runner tflint.Runner) error {
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
					"domain_name must be 255 characters or less",
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
