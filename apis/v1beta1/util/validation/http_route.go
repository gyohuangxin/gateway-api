/*
Copyright 2022 The Kubernetes Authors.

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

package validation

import gatewayv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"

// containsInSectionNameSlice checks whether the provided SectionName
// is in the target SectionName slice.
func ContainsInSectionNameSlice(items []gatewayv1b1.SectionName, item *gatewayv1b1.SectionName) bool {
	for _, eachItem := range items {
		if eachItem == *item {
			return true
		}
	}
	return false
}