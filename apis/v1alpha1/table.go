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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TableSpec defines the desired state of Table.
type TableSpec struct {
	// An array of attributes that describe the key schema for the table and indexes.
	// +kubebuilder:validation:Required
	AttributeDefinitions []*AttributeDefinition `json:"attributeDefinitions"`
	// Controls how you are charged for read and write throughput and how you manage
	// capacity. This setting can be changed later.
	//
	//    * PROVISIONED - We recommend using PROVISIONED for predictable workloads.
	//    PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
	//
	//    * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable
	//    workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
	BillingMode *string `json:"billingMode,omitempty"`
	// One or more global secondary indexes (the maximum is 20) to be created on
	// the table. Each global secondary index in the array includes the following:
	//
	//    * IndexName - The name of the global secondary index. Must be unique only
	//    for this table.
	//
	//    * KeySchema - Specifies the key schema for the global secondary index.
	//
	//    * Projection - Specifies attributes that are copied (projected) from the
	//    table into the index. These are in addition to the primary key attributes
	//    and index key attributes, which are automatically projected. Each attribute
	//    specification is composed of: ProjectionType - One of the following: KEYS_ONLY
	//    - Only the index and primary keys are projected into the index. INCLUDE
	//    - Only the specified table attributes are projected into the index. The
	//    list of projected attributes is in NonKeyAttributes. ALL - All of the
	//    table attributes are projected into the index. NonKeyAttributes - A list
	//    of one or more non-key attribute names that are projected into the secondary
	//    index. The total count of attributes provided in NonKeyAttributes, summed
	//    across all of the secondary indexes, must not exceed 100. If you project
	//    the same attribute into two different indexes, this counts as two distinct
	//    attributes when determining the total.
	//
	//    * ProvisionedThroughput - The provisioned throughput settings for the
	//    global secondary index, consisting of read and write capacity units.
	GlobalSecondaryIndexes []*GlobalSecondaryIndex `json:"globalSecondaryIndexes,omitempty"`
	// Specifies the attributes that make up the primary key for a table or an index.
	// The attributes in KeySchema must also be defined in the AttributeDefinitions
	// array. For more information, see Data Model (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	// Each KeySchemaElement in the array is composed of:
	//
	//    * AttributeName - The name of this key attribute.
	//
	//    * KeyType - The role that the key attribute will assume: HASH - partition
	//    key RANGE - sort key
	//
	// The partition key of an item is also known as its hash attribute. The term
	// "hash attribute" derives from the DynamoDB usage of an internal hash function
	// to evenly distribute data items across partitions, based on their partition
	// key values.
	//
	// The sort key of an item is also known as its range attribute. The term "range
	// attribute" derives from the way DynamoDB stores items with the same partition
	// key physically close together, in sorted order by the sort key value.
	//
	// For a simple primary key (partition key), you must provide exactly one element
	// with a KeyType of HASH.
	//
	// For a composite primary key (partition key and sort key), you must provide
	// exactly two elements, in this order: The first element must have a KeyType
	// of HASH, and the second element must have a KeyType of RANGE.
	//
	// For more information, see Working with Tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key)
	// in the Amazon DynamoDB Developer Guide.
	// +kubebuilder:validation:Required
	KeySchema []*KeySchemaElement `json:"keySchema"`
	// One or more local secondary indexes (the maximum is 5) to be created on the
	// table. Each index is scoped to a given partition key value. There is a 10
	// GB size limit per partition key value; otherwise, the size of a local secondary
	// index is unconstrained.
	//
	// Each local secondary index in the array includes the following:
	//
	//    * IndexName - The name of the local secondary index. Must be unique only
	//    for this table.
	//
	//    * KeySchema - Specifies the key schema for the local secondary index.
	//    The key schema must begin with the same partition key as the table.
	//
	//    * Projection - Specifies attributes that are copied (projected) from the
	//    table into the index. These are in addition to the primary key attributes
	//    and index key attributes, which are automatically projected. Each attribute
	//    specification is composed of: ProjectionType - One of the following: KEYS_ONLY
	//    - Only the index and primary keys are projected into the index. INCLUDE
	//    - Only the specified table attributes are projected into the index. The
	//    list of projected attributes is in NonKeyAttributes. ALL - All of the
	//    table attributes are projected into the index. NonKeyAttributes - A list
	//    of one or more non-key attribute names that are projected into the secondary
	//    index. The total count of attributes provided in NonKeyAttributes, summed
	//    across all of the secondary indexes, must not exceed 100. If you project
	//    the same attribute into two different indexes, this counts as two distinct
	//    attributes when determining the total.
	LocalSecondaryIndexes []*LocalSecondaryIndex `json:"localSecondaryIndexes,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// If you set BillingMode as PROVISIONED, you must specify this property. If
	// you set BillingMode as PAY_PER_REQUEST, you cannot specify this property.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
	// Represents the settings used to enable server-side encryption.
	SSESpecification *SSESpecification `json:"sseSpecification,omitempty"`
	// The settings for DynamoDB Streams on the table. These settings consist of:
	//
	//    * StreamEnabled - Indicates whether DynamoDB Streams is to be enabled
	//    (true) or disabled (false).
	//
	//    * StreamViewType - When an item in the table is modified, StreamViewType
	//    determines what information is written to the table's stream. Valid values
	//    for StreamViewType are: KEYS_ONLY - Only the key attributes of the modified
	//    item are written to the stream. NEW_IMAGE - The entire item, as it appears
	//    after it was modified, is written to the stream. OLD_IMAGE - The entire
	//    item, as it appeared before it was modified, is written to the stream.
	//    NEW_AND_OLD_IMAGES - Both the new and the old item images of the item
	//    are written to the stream.
	StreamSpecification *StreamSpecification `json:"streamSpecification,omitempty"`
	// The name of the table to create.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName"`
	// A list of key-value pairs to label the table. For more information, see Tagging
	// for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html).
	Tags []*Tag `json:"tags,omitempty"`
}

// TableStatus defines the observed state of Table
type TableStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Contains information about the table archive.
	ArchivalSummary *ArchivalSummary `json:"archivalSummary,omitempty"`
	// Contains the details for the read/write capacity mode.
	BillingModeSummary *BillingModeSummary `json:"billingModeSummary,omitempty"`
	// The date and time when the table was created, in UNIX epoch time (http://www.epochconverter.com/)
	// format.
	CreationDateTime *metav1.Time `json:"creationDateTime,omitempty"`
	// Represents the version of global tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html)
	// in use, if the table is replicated across AWS Regions.
	GlobalTableVersion *string `json:"globalTableVersion,omitempty"`
	// The number of items in the specified table. DynamoDB updates this value approximately
	// every six hours. Recent changes might not be reflected in this value.
	ItemCount *int64 `json:"itemCount,omitempty"`
	// The Amazon Resource Name (ARN) that uniquely identifies the latest stream
	// for this table.
	LatestStreamARN *string `json:"latestStreamARN,omitempty"`
	// A timestamp, in ISO 8601 format, for this stream.
	//
	// Note that LatestStreamLabel is not a unique identifier for the stream, because
	// it is possible that a stream from another table might have the same timestamp.
	// However, the combination of the following three elements is guaranteed to
	// be unique:
	//
	//    * AWS customer ID
	//
	//    * Table name
	//
	//    * StreamLabel
	LatestStreamLabel *string `json:"latestStreamLabel,omitempty"`
	// Represents replicas of the table.
	Replicas []*ReplicaDescription `json:"replicas,omitempty"`
	// Contains details for the restore.
	RestoreSummary *RestoreSummary `json:"restoreSummary,omitempty"`
	// The description of the server-side encryption status on the specified table.
	SSEDescription *SSEDescription `json:"sseDescription,omitempty"`
	// Unique identifier for the table for which the backup was created.
	TableID *string `json:"tableID,omitempty"`
	// The total size of the specified table, in bytes. DynamoDB updates this value
	// approximately every six hours. Recent changes might not be reflected in this
	// value.
	TableSizeBytes *int64 `json:"tableSizeBytes,omitempty"`
	// The current state of the table:
	//
	//    * CREATING - The table is being created.
	//
	//    * UPDATING - The table is being updated.
	//
	//    * DELETING - The table is being deleted.
	//
	//    * ACTIVE - The table is ready for use.
	//
	//    * INACCESSIBLE_ENCRYPTION_CREDENTIALS - The AWS KMS key used to encrypt
	//    the table in inaccessible. Table operations may fail due to failure to
	//    use the AWS KMS key. DynamoDB will initiate the table archival process
	//    when a table's AWS KMS key remains inaccessible for more than seven days.
	//
	//    * ARCHIVING - The table is being archived. Operations are not allowed
	//    until archival is complete.
	//
	//    * ARCHIVED - The table has been archived. See the ArchivalReason for more
	//    information.
	TableStatus *string `json:"tableStatus,omitempty"`
}

// Table is the Schema for the Tables API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec,omitempty"`
	Status            TableStatus `json:"status,omitempty"`
}

// TableList contains a list of Table
// +kubebuilder:object:root=true
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}
