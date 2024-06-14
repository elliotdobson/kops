// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the IAM roles that have the specified path prefix. If there are none, the
// operation returns an empty list. For more information about roles, see [IAM roles]in the
// IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. This operation does not return the following attributes, even
// though they are an attribute of the returned object:
//
//   - PermissionsBoundary
//
//   - RoleLastUsed
//
//   - Tags
//
// To view all of the information for a role, see GetRole.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
func (c *Client) ListRoles(ctx context.Context, params *ListRolesInput, optFns ...func(*Options)) (*ListRolesOutput, error) {
	if params == nil {
		params = &ListRolesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoles", params, optFns, c.addOperationListRolesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRolesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRolesInput struct {

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	MaxItems *int32

	//  The path prefix for filtering the results. For example, the prefix
	// /application_abc/component_xyz/ gets all roles whose path starts with
	// /application_abc/component_xyz/ .
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/),
	// listing all roles. This parameter allows (through its [regex pattern]) a string of characters
	// consisting of either a forward slash (/) by itself or a string that must begin
	// and end with forward slashes. In addition, it can contain any ASCII character
	// from the ! ( \u0021 ) through the DEL character ( \u007F ), including most
	// punctuation characters, digits, and upper and lowercased letters.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	PathPrefix *string

	noSmithyDocumentSerde
}

// Contains the response to a successful ListRoles request.
type ListRolesOutput struct {

	// A list of roles.
	//
	// This member is required.
	Roles []types.Role

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRolesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListRoles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListRoles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRoles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoles(options.Region), middleware.Before); err != nil {
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

// ListRolesAPIClient is a client that implements the ListRoles operation.
type ListRolesAPIClient interface {
	ListRoles(context.Context, *ListRolesInput, ...func(*Options)) (*ListRolesOutput, error)
}

var _ ListRolesAPIClient = (*Client)(nil)

// ListRolesPaginatorOptions is the paginator options for ListRoles
type ListRolesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRolesPaginator is a paginator for ListRoles
type ListRolesPaginator struct {
	options   ListRolesPaginatorOptions
	client    ListRolesAPIClient
	params    *ListRolesInput
	nextToken *string
	firstPage bool
}

// NewListRolesPaginator returns a new ListRolesPaginator
func NewListRolesPaginator(client ListRolesAPIClient, params *ListRolesInput, optFns ...func(*ListRolesPaginatorOptions)) *ListRolesPaginator {
	if params == nil {
		params = &ListRolesInput{}
	}

	options := ListRolesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRolesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRolesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoles page.
func (p *ListRolesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRolesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListRoles(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRoles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRoles",
	}
}
