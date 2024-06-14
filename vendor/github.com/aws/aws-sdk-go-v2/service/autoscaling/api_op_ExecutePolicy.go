// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Executes the specified policy. This can be useful for testing the design of
// your scaling policy.
func (c *Client) ExecutePolicy(ctx context.Context, params *ExecutePolicyInput, optFns ...func(*Options)) (*ExecutePolicyOutput, error) {
	if params == nil {
		params = &ExecutePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecutePolicy", params, optFns, c.addOperationExecutePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecutePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecutePolicyInput struct {

	// The name or ARN of the policy.
	//
	// This member is required.
	PolicyName *string

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string

	// The breach threshold for the alarm.
	//
	// Required if the policy type is StepScaling and not supported otherwise.
	BreachThreshold *float64

	// Indicates whether Amazon EC2 Auto Scaling waits for the cooldown period to
	// complete before executing the policy.
	//
	// Valid only if the policy type is SimpleScaling . For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling] in
	// the Amazon EC2 Auto Scaling User Guide.
	//
	// [Scaling cooldowns for Amazon EC2 Auto Scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html
	HonorCooldown *bool

	// The metric value to compare to BreachThreshold . This enables you to execute a
	// policy of type StepScaling and determine which step adjustment to use. For
	// example, if the breach threshold is 50 and you want to use a step adjustment
	// with a lower bound of 0 and an upper bound of 10, you can set the metric value
	// to 59.
	//
	// If you specify a metric value that doesn't correspond to a step adjustment for
	// the policy, the call returns an error.
	//
	// Required if the policy type is StepScaling and not supported otherwise.
	MetricValue *float64

	noSmithyDocumentSerde
}

type ExecutePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecutePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpExecutePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpExecutePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExecutePolicy"); err != nil {
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
	if err = addOpExecutePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecutePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExecutePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExecutePolicy",
	}
}
