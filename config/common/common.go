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

package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"time"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/crossplane-contrib/provider-jet-aws/config/common"

	// PathARNExtractor is the golang path to ARNExtractor function
	// in this package.
	PathARNExtractor = SelfPackagePath + ".ARNExtractor()"

	// VersionV1Alpha2 is used as minimum version for all manually configured
	// resources.
	VersionV1Alpha2 = "v1alpha2"

	// Default Terraform read timeout
	TerraformReadTimeout = 10 * time.Minute

	// Default Terraform create timeout
	TerraformCreateTimeout = 30 * time.Minute

	// Default Terraform update timeout
	TerraformUpdateTimeout = 30 * time.Minute

	// Default Terraform delete timeout
	TerraformDeleteTimeout = 30 * time.Minute
)

// ARNExtractor extracts ARN of the resources from "status.atProvider.arn" which
// is quite common among all AWS resources.
func ARNExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.atProvider.arn")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}
