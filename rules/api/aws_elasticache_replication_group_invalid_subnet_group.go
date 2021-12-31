// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsElastiCacheReplicationGroupInvalidSubnetGroupRule checks whether attribute value actually exists
type AwsElastiCacheReplicationGroupInvalidSubnetGroupRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsElastiCacheReplicationGroupInvalidSubnetGroupRule returns new rule with default attributes
func NewAwsElastiCacheReplicationGroupInvalidSubnetGroupRule() *AwsElastiCacheReplicationGroupInvalidSubnetGroupRule {
	return &AwsElastiCacheReplicationGroupInvalidSubnetGroupRule{
		resourceType:  "aws_elasticache_replication_group",
		attributeName: "subnet_group_name",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsElastiCacheReplicationGroupInvalidSubnetGroupRule) Name() string {
	return "aws_elasticache_replication_group_invalid_subnet_group"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastiCacheReplicationGroupInvalidSubnetGroupRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastiCacheReplicationGroupInvalidSubnetGroupRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastiCacheReplicationGroupInvalidSubnetGroupRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsElastiCacheReplicationGroupInvalidSubnetGroupRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeCacheSubnetGroups
func (r *AwsElastiCacheReplicationGroupInvalidSubnetGroupRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

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

		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeCacheSubnetGroups")
			var err error
			r.data, err = runner.AwsClient.DescribeCacheSubnetGroups()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeCacheSubnetGroups",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid subnet group name.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	}

	return nil
}
