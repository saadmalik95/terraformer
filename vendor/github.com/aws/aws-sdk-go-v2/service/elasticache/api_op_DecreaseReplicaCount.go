// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DecreaseReplicaCountInput struct {
	_ struct{} `type:"structure"`

	// If True, the number of replica nodes is decreased immediately. ApplyImmediately=False
	// is not currently supported.
	//
	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	// The number of read replica nodes you want at the completion of this operation.
	// For Redis (cluster mode disabled) replication groups, this is the number
	// of replica nodes in the replication group. For Redis (cluster mode enabled)
	// replication groups, this is the number of replica nodes in each of the replication
	// group's node groups.
	//
	// The minimum number of replicas in a shard or replication group is:
	//
	//    * Redis (cluster mode disabled) If Multi-AZ with Automatic Failover is
	//    enabled: 1 If Multi-AZ with Automatic Failover is not enabled: 0
	//
	//    * Redis (cluster mode enabled): 0 (though you will not be able to failover
	//    to a replica if your primary node fails)
	NewReplicaCount *int64 `type:"integer"`

	// A list of ConfigureShard objects that can be used to configure each shard
	// in a Redis (cluster mode enabled) replication group. The ConfigureShard has
	// three members: NewReplicaCount, NodeGroupId, and PreferredAvailabilityZones.
	ReplicaConfiguration []ConfigureShard `locationNameList:"ConfigureShard" type:"list"`

	// A list of the node ids to remove from the replication group or node group
	// (shard).
	ReplicasToRemove []string `type:"list"`

	// The id of the replication group from which you want to remove replica nodes.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DecreaseReplicaCountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecreaseReplicaCountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DecreaseReplicaCountInput"}

	if s.ApplyImmediately == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyImmediately"))
	}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}
	if s.ReplicaConfiguration != nil {
		for i, v := range s.ReplicaConfiguration {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ReplicaConfiguration", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DecreaseReplicaCountOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s DecreaseReplicaCountOutput) String() string {
	return awsutil.Prettify(s)
}

const opDecreaseReplicaCount = "DecreaseReplicaCount"

// DecreaseReplicaCountRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Dynamically decreases the number of replics in a Redis (cluster mode disabled)
// replication group or the number of replica nodes in one or more node groups
// (shards) of a Redis (cluster mode enabled) replication group. This operation
// is performed with no cluster down time.
//
//    // Example sending a request using DecreaseReplicaCountRequest.
//    req := client.DecreaseReplicaCountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DecreaseReplicaCount
func (c *Client) DecreaseReplicaCountRequest(input *DecreaseReplicaCountInput) DecreaseReplicaCountRequest {
	op := &aws.Operation{
		Name:       opDecreaseReplicaCount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecreaseReplicaCountInput{}
	}

	req := c.newRequest(op, input, &DecreaseReplicaCountOutput{})
	return DecreaseReplicaCountRequest{Request: req, Input: input, Copy: c.DecreaseReplicaCountRequest}
}

// DecreaseReplicaCountRequest is the request type for the
// DecreaseReplicaCount API operation.
type DecreaseReplicaCountRequest struct {
	*aws.Request
	Input *DecreaseReplicaCountInput
	Copy  func(*DecreaseReplicaCountInput) DecreaseReplicaCountRequest
}

// Send marshals and sends the DecreaseReplicaCount API request.
func (r DecreaseReplicaCountRequest) Send(ctx context.Context) (*DecreaseReplicaCountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DecreaseReplicaCountResponse{
		DecreaseReplicaCountOutput: r.Request.Data.(*DecreaseReplicaCountOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DecreaseReplicaCountResponse is the response type for the
// DecreaseReplicaCount API operation.
type DecreaseReplicaCountResponse struct {
	*DecreaseReplicaCountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DecreaseReplicaCount request.
func (r *DecreaseReplicaCountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}