// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule checks the pattern is valid
type AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule returns new rule with default attributes
func NewAwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule() *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule {
	return &AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule{
		resourceType:  "aws_elastic_beanstalk_configuration_template",
		attributeName: "description",
		max:           200,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Name() string {
	return "aws_elastic_beanstalk_configuration_template_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 200 characters or less",
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
