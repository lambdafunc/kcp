/*
Copyright The KCP Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// CowboySpecApplyConfiguration represents a declarative configuration of the CowboySpec type for use
// with apply.
type CowboySpecApplyConfiguration struct {
	Intent *string `json:"intent,omitempty"`
}

// CowboySpecApplyConfiguration constructs a declarative configuration of the CowboySpec type for use with
// apply.
func CowboySpec() *CowboySpecApplyConfiguration {
	return &CowboySpecApplyConfiguration{}
}

// WithIntent sets the Intent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Intent field is set to the value of the last call.
func (b *CowboySpecApplyConfiguration) WithIntent(value string) *CowboySpecApplyConfiguration {
	b.Intent = &value
	return b
}
