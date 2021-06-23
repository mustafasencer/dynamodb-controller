// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package global_table

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/dynamodb-controller/apis/v1alpha1"
	svcsdkapi "github.com/aws/aws-sdk-go/service/dynamodb"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.DynamoDB{}
	_ = &svcapitypes.GlobalTable{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = svcsdkapi.New
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdkapi.DescribeGlobalTableOutput
	resp, err = rm.sdkapi.DescribeGlobalTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeGlobalTable", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "GlobalTableNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.GlobalTableDescription.CreationDateTime != nil {
		ko.Status.CreationDateTime = &metav1.Time{*resp.GlobalTableDescription.CreationDateTime}
	} else {
		ko.Status.CreationDateTime = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.GlobalTableDescription.GlobalTableArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.GlobalTableDescription.GlobalTableArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.GlobalTableDescription.GlobalTableName != nil {
		ko.Spec.GlobalTableName = resp.GlobalTableDescription.GlobalTableName
	} else {
		ko.Spec.GlobalTableName = nil
	}
	if resp.GlobalTableDescription.GlobalTableStatus != nil {
		ko.Status.GlobalTableStatus = resp.GlobalTableDescription.GlobalTableStatus
	} else {
		ko.Status.GlobalTableStatus = nil
	}
	if resp.GlobalTableDescription.ReplicationGroup != nil {
		f4 := []*svcapitypes.Replica{}
		for _, f4iter := range resp.GlobalTableDescription.ReplicationGroup {
			f4elem := &svcapitypes.Replica{}
			if f4iter.RegionName != nil {
				f4elem.RegionName = f4iter.RegionName
			}
			f4 = append(f4, f4elem)
		}
		ko.Spec.ReplicationGroup = f4
	} else {
		ko.Spec.ReplicationGroup = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.GlobalTableName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeGlobalTableInput, error) {
	res := &svcsdk.DescribeGlobalTableInput{}

	if r.ko.Spec.GlobalTableName != nil {
		res.SetGlobalTableName(*r.ko.Spec.GlobalTableName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdkapi.CreateGlobalTableOutput
	resp, err = rm.sdkapi.CreateGlobalTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateGlobalTable", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.GlobalTableDescription.CreationDateTime != nil {
		ko.Status.CreationDateTime = &metav1.Time{*resp.GlobalTableDescription.CreationDateTime}
	} else {
		ko.Status.CreationDateTime = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.GlobalTableDescription.GlobalTableArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.GlobalTableDescription.GlobalTableArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.GlobalTableDescription.GlobalTableStatus != nil {
		ko.Status.GlobalTableStatus = resp.GlobalTableDescription.GlobalTableStatus
	} else {
		ko.Status.GlobalTableStatus = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateGlobalTableInput, error) {
	res := &svcsdk.CreateGlobalTableInput{}

	if r.ko.Spec.GlobalTableName != nil {
		res.SetGlobalTableName(*r.ko.Spec.GlobalTableName)
	}
	if r.ko.Spec.ReplicationGroup != nil {
		f1 := []*svcsdk.Replica{}
		for _, f1iter := range r.ko.Spec.ReplicationGroup {
			f1elem := &svcsdk.Replica{}
			if f1iter.RegionName != nil {
				f1elem.SetRegionName(*f1iter.RegionName)
			}
			f1 = append(f1, f1elem)
		}
		res.SetReplicationGroup(f1)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	customSetDeleteInput(r, input)
	_, err = rm.sdkapi.UpdateGlobalTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "UpdateGlobalTable", err)
	return err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.UpdateGlobalTableInput, error) {
	res := &svcsdk.UpdateGlobalTableInput{}

	if r.ko.Spec.GlobalTableName != nil {
		res.SetGlobalTableName(*r.ko.Spec.GlobalTableName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.GlobalTable,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	if terminalCondition != nil || recoverableCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
