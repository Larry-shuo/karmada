//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Karmada Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	search "github.com/karmada-io/karmada/pkg/apis/search"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BackendStoreConfig)(nil), (*search.BackendStoreConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_BackendStoreConfig_To_search_BackendStoreConfig(a.(*BackendStoreConfig), b.(*search.BackendStoreConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.BackendStoreConfig)(nil), (*BackendStoreConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_BackendStoreConfig_To_v1alpha1_BackendStoreConfig(a.(*search.BackendStoreConfig), b.(*BackendStoreConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*OpenSearchConfig)(nil), (*search.OpenSearchConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_OpenSearchConfig_To_search_OpenSearchConfig(a.(*OpenSearchConfig), b.(*search.OpenSearchConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.OpenSearchConfig)(nil), (*OpenSearchConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_OpenSearchConfig_To_v1alpha1_OpenSearchConfig(a.(*search.OpenSearchConfig), b.(*OpenSearchConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Proxying)(nil), (*search.Proxying)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Proxying_To_search_Proxying(a.(*Proxying), b.(*search.Proxying), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.Proxying)(nil), (*Proxying)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_Proxying_To_v1alpha1_Proxying(a.(*search.Proxying), b.(*Proxying), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceRegistry)(nil), (*search.ResourceRegistry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceRegistry_To_search_ResourceRegistry(a.(*ResourceRegistry), b.(*search.ResourceRegistry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.ResourceRegistry)(nil), (*ResourceRegistry)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_ResourceRegistry_To_v1alpha1_ResourceRegistry(a.(*search.ResourceRegistry), b.(*ResourceRegistry), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceRegistryList)(nil), (*search.ResourceRegistryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceRegistryList_To_search_ResourceRegistryList(a.(*ResourceRegistryList), b.(*search.ResourceRegistryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.ResourceRegistryList)(nil), (*ResourceRegistryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_ResourceRegistryList_To_v1alpha1_ResourceRegistryList(a.(*search.ResourceRegistryList), b.(*ResourceRegistryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceRegistrySpec)(nil), (*search.ResourceRegistrySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceRegistrySpec_To_search_ResourceRegistrySpec(a.(*ResourceRegistrySpec), b.(*search.ResourceRegistrySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.ResourceRegistrySpec)(nil), (*ResourceRegistrySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_ResourceRegistrySpec_To_v1alpha1_ResourceRegistrySpec(a.(*search.ResourceRegistrySpec), b.(*ResourceRegistrySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceRegistryStatus)(nil), (*search.ResourceRegistryStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceRegistryStatus_To_search_ResourceRegistryStatus(a.(*ResourceRegistryStatus), b.(*search.ResourceRegistryStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.ResourceRegistryStatus)(nil), (*ResourceRegistryStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_ResourceRegistryStatus_To_v1alpha1_ResourceRegistryStatus(a.(*search.ResourceRegistryStatus), b.(*ResourceRegistryStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceSelector)(nil), (*search.ResourceSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceSelector_To_search_ResourceSelector(a.(*ResourceSelector), b.(*search.ResourceSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.ResourceSelector)(nil), (*ResourceSelector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_ResourceSelector_To_v1alpha1_ResourceSelector(a.(*search.ResourceSelector), b.(*ResourceSelector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Search)(nil), (*search.Search)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Search_To_search_Search(a.(*Search), b.(*search.Search), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*search.Search)(nil), (*Search)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_search_Search_To_v1alpha1_Search(a.(*search.Search), b.(*Search), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_BackendStoreConfig_To_search_BackendStoreConfig(in *BackendStoreConfig, out *search.BackendStoreConfig, s conversion.Scope) error {
	out.OpenSearch = (*search.OpenSearchConfig)(unsafe.Pointer(in.OpenSearch))
	return nil
}

// Convert_v1alpha1_BackendStoreConfig_To_search_BackendStoreConfig is an autogenerated conversion function.
func Convert_v1alpha1_BackendStoreConfig_To_search_BackendStoreConfig(in *BackendStoreConfig, out *search.BackendStoreConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_BackendStoreConfig_To_search_BackendStoreConfig(in, out, s)
}

func autoConvert_search_BackendStoreConfig_To_v1alpha1_BackendStoreConfig(in *search.BackendStoreConfig, out *BackendStoreConfig, s conversion.Scope) error {
	out.OpenSearch = (*OpenSearchConfig)(unsafe.Pointer(in.OpenSearch))
	return nil
}

// Convert_search_BackendStoreConfig_To_v1alpha1_BackendStoreConfig is an autogenerated conversion function.
func Convert_search_BackendStoreConfig_To_v1alpha1_BackendStoreConfig(in *search.BackendStoreConfig, out *BackendStoreConfig, s conversion.Scope) error {
	return autoConvert_search_BackendStoreConfig_To_v1alpha1_BackendStoreConfig(in, out, s)
}

func autoConvert_v1alpha1_OpenSearchConfig_To_search_OpenSearchConfig(in *OpenSearchConfig, out *search.OpenSearchConfig, s conversion.Scope) error {
	out.Addresses = *(*[]string)(unsafe.Pointer(&in.Addresses))
	out.SecretRef = in.SecretRef
	return nil
}

// Convert_v1alpha1_OpenSearchConfig_To_search_OpenSearchConfig is an autogenerated conversion function.
func Convert_v1alpha1_OpenSearchConfig_To_search_OpenSearchConfig(in *OpenSearchConfig, out *search.OpenSearchConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_OpenSearchConfig_To_search_OpenSearchConfig(in, out, s)
}

func autoConvert_search_OpenSearchConfig_To_v1alpha1_OpenSearchConfig(in *search.OpenSearchConfig, out *OpenSearchConfig, s conversion.Scope) error {
	out.Addresses = *(*[]string)(unsafe.Pointer(&in.Addresses))
	out.SecretRef = in.SecretRef
	return nil
}

// Convert_search_OpenSearchConfig_To_v1alpha1_OpenSearchConfig is an autogenerated conversion function.
func Convert_search_OpenSearchConfig_To_v1alpha1_OpenSearchConfig(in *search.OpenSearchConfig, out *OpenSearchConfig, s conversion.Scope) error {
	return autoConvert_search_OpenSearchConfig_To_v1alpha1_OpenSearchConfig(in, out, s)
}

func autoConvert_v1alpha1_Proxying_To_search_Proxying(in *Proxying, out *search.Proxying, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_Proxying_To_search_Proxying is an autogenerated conversion function.
func Convert_v1alpha1_Proxying_To_search_Proxying(in *Proxying, out *search.Proxying, s conversion.Scope) error {
	return autoConvert_v1alpha1_Proxying_To_search_Proxying(in, out, s)
}

func autoConvert_search_Proxying_To_v1alpha1_Proxying(in *search.Proxying, out *Proxying, s conversion.Scope) error {
	return nil
}

// Convert_search_Proxying_To_v1alpha1_Proxying is an autogenerated conversion function.
func Convert_search_Proxying_To_v1alpha1_Proxying(in *search.Proxying, out *Proxying, s conversion.Scope) error {
	return autoConvert_search_Proxying_To_v1alpha1_Proxying(in, out, s)
}

func autoConvert_v1alpha1_ResourceRegistry_To_search_ResourceRegistry(in *ResourceRegistry, out *search.ResourceRegistry, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ResourceRegistrySpec_To_search_ResourceRegistrySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ResourceRegistryStatus_To_search_ResourceRegistryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ResourceRegistry_To_search_ResourceRegistry is an autogenerated conversion function.
func Convert_v1alpha1_ResourceRegistry_To_search_ResourceRegistry(in *ResourceRegistry, out *search.ResourceRegistry, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceRegistry_To_search_ResourceRegistry(in, out, s)
}

func autoConvert_search_ResourceRegistry_To_v1alpha1_ResourceRegistry(in *search.ResourceRegistry, out *ResourceRegistry, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_search_ResourceRegistrySpec_To_v1alpha1_ResourceRegistrySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_search_ResourceRegistryStatus_To_v1alpha1_ResourceRegistryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_search_ResourceRegistry_To_v1alpha1_ResourceRegistry is an autogenerated conversion function.
func Convert_search_ResourceRegistry_To_v1alpha1_ResourceRegistry(in *search.ResourceRegistry, out *ResourceRegistry, s conversion.Scope) error {
	return autoConvert_search_ResourceRegistry_To_v1alpha1_ResourceRegistry(in, out, s)
}

func autoConvert_v1alpha1_ResourceRegistryList_To_search_ResourceRegistryList(in *ResourceRegistryList, out *search.ResourceRegistryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]search.ResourceRegistry)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ResourceRegistryList_To_search_ResourceRegistryList is an autogenerated conversion function.
func Convert_v1alpha1_ResourceRegistryList_To_search_ResourceRegistryList(in *ResourceRegistryList, out *search.ResourceRegistryList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceRegistryList_To_search_ResourceRegistryList(in, out, s)
}

func autoConvert_search_ResourceRegistryList_To_v1alpha1_ResourceRegistryList(in *search.ResourceRegistryList, out *ResourceRegistryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ResourceRegistry)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_search_ResourceRegistryList_To_v1alpha1_ResourceRegistryList is an autogenerated conversion function.
func Convert_search_ResourceRegistryList_To_v1alpha1_ResourceRegistryList(in *search.ResourceRegistryList, out *ResourceRegistryList, s conversion.Scope) error {
	return autoConvert_search_ResourceRegistryList_To_v1alpha1_ResourceRegistryList(in, out, s)
}

func autoConvert_v1alpha1_ResourceRegistrySpec_To_search_ResourceRegistrySpec(in *ResourceRegistrySpec, out *search.ResourceRegistrySpec, s conversion.Scope) error {
	out.TargetCluster = in.TargetCluster
	out.ResourceSelectors = *(*[]search.ResourceSelector)(unsafe.Pointer(&in.ResourceSelectors))
	out.BackendStore = (*search.BackendStoreConfig)(unsafe.Pointer(in.BackendStore))
	return nil
}

// Convert_v1alpha1_ResourceRegistrySpec_To_search_ResourceRegistrySpec is an autogenerated conversion function.
func Convert_v1alpha1_ResourceRegistrySpec_To_search_ResourceRegistrySpec(in *ResourceRegistrySpec, out *search.ResourceRegistrySpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceRegistrySpec_To_search_ResourceRegistrySpec(in, out, s)
}

func autoConvert_search_ResourceRegistrySpec_To_v1alpha1_ResourceRegistrySpec(in *search.ResourceRegistrySpec, out *ResourceRegistrySpec, s conversion.Scope) error {
	out.TargetCluster = in.TargetCluster
	out.ResourceSelectors = *(*[]ResourceSelector)(unsafe.Pointer(&in.ResourceSelectors))
	out.BackendStore = (*BackendStoreConfig)(unsafe.Pointer(in.BackendStore))
	return nil
}

// Convert_search_ResourceRegistrySpec_To_v1alpha1_ResourceRegistrySpec is an autogenerated conversion function.
func Convert_search_ResourceRegistrySpec_To_v1alpha1_ResourceRegistrySpec(in *search.ResourceRegistrySpec, out *ResourceRegistrySpec, s conversion.Scope) error {
	return autoConvert_search_ResourceRegistrySpec_To_v1alpha1_ResourceRegistrySpec(in, out, s)
}

func autoConvert_v1alpha1_ResourceRegistryStatus_To_search_ResourceRegistryStatus(in *ResourceRegistryStatus, out *search.ResourceRegistryStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_ResourceRegistryStatus_To_search_ResourceRegistryStatus is an autogenerated conversion function.
func Convert_v1alpha1_ResourceRegistryStatus_To_search_ResourceRegistryStatus(in *ResourceRegistryStatus, out *search.ResourceRegistryStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceRegistryStatus_To_search_ResourceRegistryStatus(in, out, s)
}

func autoConvert_search_ResourceRegistryStatus_To_v1alpha1_ResourceRegistryStatus(in *search.ResourceRegistryStatus, out *ResourceRegistryStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_search_ResourceRegistryStatus_To_v1alpha1_ResourceRegistryStatus is an autogenerated conversion function.
func Convert_search_ResourceRegistryStatus_To_v1alpha1_ResourceRegistryStatus(in *search.ResourceRegistryStatus, out *ResourceRegistryStatus, s conversion.Scope) error {
	return autoConvert_search_ResourceRegistryStatus_To_v1alpha1_ResourceRegistryStatus(in, out, s)
}

func autoConvert_v1alpha1_ResourceSelector_To_search_ResourceSelector(in *ResourceSelector, out *search.ResourceSelector, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	out.Namespace = in.Namespace
	return nil
}

// Convert_v1alpha1_ResourceSelector_To_search_ResourceSelector is an autogenerated conversion function.
func Convert_v1alpha1_ResourceSelector_To_search_ResourceSelector(in *ResourceSelector, out *search.ResourceSelector, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceSelector_To_search_ResourceSelector(in, out, s)
}

func autoConvert_search_ResourceSelector_To_v1alpha1_ResourceSelector(in *search.ResourceSelector, out *ResourceSelector, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	out.Namespace = in.Namespace
	return nil
}

// Convert_search_ResourceSelector_To_v1alpha1_ResourceSelector is an autogenerated conversion function.
func Convert_search_ResourceSelector_To_v1alpha1_ResourceSelector(in *search.ResourceSelector, out *ResourceSelector, s conversion.Scope) error {
	return autoConvert_search_ResourceSelector_To_v1alpha1_ResourceSelector(in, out, s)
}

func autoConvert_v1alpha1_Search_To_search_Search(in *Search, out *search.Search, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_Search_To_search_Search is an autogenerated conversion function.
func Convert_v1alpha1_Search_To_search_Search(in *Search, out *search.Search, s conversion.Scope) error {
	return autoConvert_v1alpha1_Search_To_search_Search(in, out, s)
}

func autoConvert_search_Search_To_v1alpha1_Search(in *search.Search, out *Search, s conversion.Scope) error {
	return nil
}

// Convert_search_Search_To_v1alpha1_Search is an autogenerated conversion function.
func Convert_search_Search_To_v1alpha1_Search(in *search.Search, out *Search, s conversion.Scope) error {
	return autoConvert_search_Search_To_v1alpha1_Search(in, out, s)
}
