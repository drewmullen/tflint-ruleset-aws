// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule checks the pattern is valid
type AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule returns new rule with default attributes
func NewAwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule() *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule {
	return &AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule{
		resourceType:  "aws_elastic_beanstalk_environment",
		attributeName: "cname_prefix",
		max:           63,
		min:           4,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Name() string {
	return "aws_elastic_beanstalk_environment_invalid_cname_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Check(runner tflint.Runner) error {
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
					"cname_prefix must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"cname_prefix must be 4 characters or higher",
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
