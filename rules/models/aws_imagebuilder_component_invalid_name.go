// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderComponentInvalidNameRule checks the pattern is valid
type AwsImagebuilderComponentInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsImagebuilderComponentInvalidNameRule returns new rule with default attributes
func NewAwsImagebuilderComponentInvalidNameRule() *AwsImagebuilderComponentInvalidNameRule {
	return &AwsImagebuilderComponentInvalidNameRule{
		resourceType:  "aws_imagebuilder_component",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`),
	}
}

// Name returns the rule name
func (r *AwsImagebuilderComponentInvalidNameRule) Name() string {
	return "aws_imagebuilder_component_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderComponentInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderComponentInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderComponentInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderComponentInvalidNameRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`),
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
