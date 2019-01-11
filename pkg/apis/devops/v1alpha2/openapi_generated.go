// +build !ignore_autogenerated

/*
Copyright 2019 The KubeSphere Authors.

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
// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProject":       schema_pkg_apis_devops_v1alpha2_DevOpsProject(ref),
		"github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProjectList":   schema_pkg_apis_devops_v1alpha2_DevOpsProjectList(ref),
		"github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProjectSpec":   schema_pkg_apis_devops_v1alpha2_DevOpsProjectSpec(ref),
		"github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProjectStatus": schema_pkg_apis_devops_v1alpha2_DevOpsProjectStatus(ref),
	}
}

func schema_pkg_apis_devops_v1alpha2_DevOpsProject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevOpsProject is the Schema for the devopsprojects API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProjectSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProjectStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProjectSpec", "github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProjectStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_devops_v1alpha2_DevOpsProjectList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevOpsProjectList contains a list of DevOpsProject",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProject"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/runzexia/kubesphere-crd-sample/pkg/apis/devops/v1alpha2.DevOpsProject", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_devops_v1alpha2_DevOpsProjectSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevOpsProjectSpec defines the desired state of DevOpsProject",
				Properties: map[string]spec.Schema{
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Description: "DisplayName is DisplayName of DevOpsProject",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is Description of DevOpsProject",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_devops_v1alpha2_DevOpsProjectStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DevOpsProjectStatus defines the observed state of DevOpsProject",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is Status of DevOpsProject",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
