// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElastiCacheUserInvalidEngineRule checks the pattern is valid
type AwsElastiCacheUserInvalidEngineRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsElastiCacheUserInvalidEngineRule returns new rule with default attributes
func NewAwsElastiCacheUserInvalidEngineRule() *AwsElastiCacheUserInvalidEngineRule {
	return &AwsElastiCacheUserInvalidEngineRule{
		resourceType:  "aws_elasticache_user",
		attributeName: "engine",
		pattern:       regexp.MustCompile(`^[a-zA-Z]*$`),
	}
}

// Name returns the rule name
func (r *AwsElastiCacheUserInvalidEngineRule) Name() string {
	return "aws_elasticache_user_invalid_engine"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheUserInvalidEngineRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheUserInvalidEngineRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheUserInvalidEngineRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastiCacheUserInvalidEngineRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z]*$`),
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
