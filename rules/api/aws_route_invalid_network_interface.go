// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsRouteInvalidNetworkInterfaceRule checks whether attribute value actually exists
type AwsRouteInvalidNetworkInterfaceRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsRouteInvalidNetworkInterfaceRule returns new rule with default attributes
func NewAwsRouteInvalidNetworkInterfaceRule() *AwsRouteInvalidNetworkInterfaceRule {
	return &AwsRouteInvalidNetworkInterfaceRule{
		resourceType:  "aws_route",
		attributeName: "network_interface_id",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsRouteInvalidNetworkInterfaceRule) Name() string {
	return "aws_route_invalid_network_interface"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRouteInvalidNetworkInterfaceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRouteInvalidNetworkInterfaceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRouteInvalidNetworkInterfaceRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsRouteInvalidNetworkInterfaceRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeNetworkInterfaces
func (r *AwsRouteInvalidNetworkInterfaceRule) Check(rr tflint.Runner) error {
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
			log.Print("[DEBUG] invoking DescribeNetworkInterfaces")
			var err error
			r.data, err = runner.AwsClient.DescribeNetworkInterfaces()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeNetworkInterfaces",
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
					fmt.Sprintf(`"%s" is invalid network interface ID.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	}

	return nil
}
