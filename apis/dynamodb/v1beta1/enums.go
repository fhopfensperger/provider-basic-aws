/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

type SSEStatus string

const (
	SSEStatus_ENABLING  SSEStatus = "ENABLING"
	SSEStatus_ENABLED   SSEStatus = "ENABLED"
	SSEStatus_DISABLING SSEStatus = "DISABLING"
	SSEStatus_DISABLED  SSEStatus = "DISABLED"
	SSEStatus_UPDATING  SSEStatus = "UPDATING"
)

type TableStatus_SDK string

const (
	TableStatus_SDK_CREATING                            TableStatus_SDK = "CREATING"
	TableStatus_SDK_UPDATING                            TableStatus_SDK = "UPDATING"
	TableStatus_SDK_DELETING                            TableStatus_SDK = "DELETING"
	TableStatus_SDK_ACTIVE                              TableStatus_SDK = "ACTIVE"
	TableStatus_SDK_INACCESSIBLE_ENCRYPTION_CREDENTIALS TableStatus_SDK = "INACCESSIBLE_ENCRYPTION_CREDENTIALS"
	TableStatus_SDK_ARCHIVING                           TableStatus_SDK = "ARCHIVING"
	TableStatus_SDK_ARCHIVED                            TableStatus_SDK = "ARCHIVED"
)
