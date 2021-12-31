// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule checks the pattern is valid
type AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule returns new rule with default attributes
func NewAwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule() *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule {
	return &AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule{
		resourceType:  "aws_codeartifact_repository_permissions_policy",
		attributeName: "domain",
		max:           50,
		min:           2,
		pattern:       regexp.MustCompile(`^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
	}
}

// Name returns the rule name
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule) Name() string {
	return "aws_codeartifact_repository_permissions_policy_invalid_domain"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodeartifactRepositoryPermissionsPolicyInvalidDomainRule) Check(runner tflint.Runner) error {
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
					"domain must be 50 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"domain must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`),
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
