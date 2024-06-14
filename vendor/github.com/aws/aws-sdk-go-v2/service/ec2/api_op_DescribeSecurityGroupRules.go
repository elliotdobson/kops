// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more of your security group rules.
func (c *Client) DescribeSecurityGroupRules(ctx context.Context, params *DescribeSecurityGroupRulesInput, optFns ...func(*Options)) (*DescribeSecurityGroupRulesOutput, error) {
	if params == nil {
		params = &DescribeSecurityGroupRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSecurityGroupRules", params, optFns, c.addOperationDescribeSecurityGroupRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSecurityGroupRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityGroupRulesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//
	//   - group-id - The ID of the security group.
	//
	//   - security-group-rule-id - The ID of the security group rule.
	//
	//   - tag : - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	Filters []types.Filter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. This value
	// can be between 5 and 1000. If this parameter is not specified, then all items
	// are returned. For more information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// The IDs of the security group rules.
	SecurityGroupRuleIds []string

	noSmithyDocumentSerde
}

type DescribeSecurityGroupRulesOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about security group rules.
	SecurityGroupRules []types.SecurityGroupRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSecurityGroupRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSecurityGroupRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSecurityGroupRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeSecurityGroupRules"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityGroupRules(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeSecurityGroupRulesAPIClient is a client that implements the
// DescribeSecurityGroupRules operation.
type DescribeSecurityGroupRulesAPIClient interface {
	DescribeSecurityGroupRules(context.Context, *DescribeSecurityGroupRulesInput, ...func(*Options)) (*DescribeSecurityGroupRulesOutput, error)
}

var _ DescribeSecurityGroupRulesAPIClient = (*Client)(nil)

// DescribeSecurityGroupRulesPaginatorOptions is the paginator options for
// DescribeSecurityGroupRules
type DescribeSecurityGroupRulesPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. This value
	// can be between 5 and 1000. If this parameter is not specified, then all items
	// are returned. For more information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSecurityGroupRulesPaginator is a paginator for
// DescribeSecurityGroupRules
type DescribeSecurityGroupRulesPaginator struct {
	options   DescribeSecurityGroupRulesPaginatorOptions
	client    DescribeSecurityGroupRulesAPIClient
	params    *DescribeSecurityGroupRulesInput
	nextToken *string
	firstPage bool
}

// NewDescribeSecurityGroupRulesPaginator returns a new
// DescribeSecurityGroupRulesPaginator
func NewDescribeSecurityGroupRulesPaginator(client DescribeSecurityGroupRulesAPIClient, params *DescribeSecurityGroupRulesInput, optFns ...func(*DescribeSecurityGroupRulesPaginatorOptions)) *DescribeSecurityGroupRulesPaginator {
	if params == nil {
		params = &DescribeSecurityGroupRulesInput{}
	}

	options := DescribeSecurityGroupRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSecurityGroupRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSecurityGroupRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSecurityGroupRules page.
func (p *DescribeSecurityGroupRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSecurityGroupRulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeSecurityGroupRules(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeSecurityGroupRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeSecurityGroupRules",
	}
}
