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

// Uploads an SSH public key and associates it with the specified IAM user.
//
// The SSH public key uploaded by this operation can be used only for
// authenticating the associated IAM user to an CodeCommit repository. For more
// information about using SSH keys to authenticate to an CodeCommit repository,
// see [Set up CodeCommit for SSH connections]in the CodeCommit User Guide.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
func (c *Client) UploadSSHPublicKey(ctx context.Context, params *UploadSSHPublicKeyInput, optFns ...func(*Options)) (*UploadSSHPublicKeyOutput, error) {
	if params == nil {
		params = &UploadSSHPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadSSHPublicKey", params, optFns, c.addOperationUploadSSHPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadSSHPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadSSHPublicKeyInput struct {

	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM
	// format. The minimum bit-length of the public key is 2048 bits. For example, you
	// can generate a 2048-bit key, and the resulting PEM file is 1679 bytes long.
	//
	// The [regex pattern] used to validate this parameter is a string of characters consisting of
	// the following:
	//
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	SSHPublicKeyBody *string

	// The name of the IAM user to associate the SSH public key with.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

// Contains the response to a successful UploadSSHPublicKey request.
type UploadSSHPublicKeyOutput struct {

	// Contains information about the SSH public key.
	SSHPublicKey *types.SSHPublicKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUploadSSHPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUploadSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUploadSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UploadSSHPublicKey"); err != nil {
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
	if err = addOpUploadSSHPublicKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadSSHPublicKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUploadSSHPublicKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UploadSSHPublicKey",
	}
}
