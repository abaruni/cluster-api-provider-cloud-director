// +build !ignore_autogenerated_conversions

/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha4

import (
	unsafe "unsafe"

	v1beta1 "github.com/vmware/cluster-api-provider-cloud-director/api/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*APIEndpoint)(nil), (*v1beta1.APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_APIEndpoint_To_v1beta1_APIEndpoint(a.(*APIEndpoint), b.(*v1beta1.APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.APIEndpoint)(nil), (*APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_APIEndpoint_To_v1alpha4_APIEndpoint(a.(*v1beta1.APIEndpoint), b.(*APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UserCredentialsContext)(nil), (*v1beta1.UserCredentialsContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_UserCredentialsContext_To_v1beta1_UserCredentialsContext(a.(*UserCredentialsContext), b.(*v1beta1.UserCredentialsContext), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.UserCredentialsContext)(nil), (*UserCredentialsContext)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_UserCredentialsContext_To_v1alpha4_UserCredentialsContext(a.(*v1beta1.UserCredentialsContext), b.(*UserCredentialsContext), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDCluster)(nil), (*v1beta1.VCDCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDCluster_To_v1beta1_VCDCluster(a.(*VCDCluster), b.(*v1beta1.VCDCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDCluster)(nil), (*VCDCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDCluster_To_v1alpha4_VCDCluster(a.(*v1beta1.VCDCluster), b.(*VCDCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDClusterList)(nil), (*v1beta1.VCDClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDClusterList_To_v1beta1_VCDClusterList(a.(*VCDClusterList), b.(*v1beta1.VCDClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDClusterList)(nil), (*VCDClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDClusterList_To_v1alpha4_VCDClusterList(a.(*v1beta1.VCDClusterList), b.(*VCDClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDClusterSpec)(nil), (*v1beta1.VCDClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDClusterSpec_To_v1beta1_VCDClusterSpec(a.(*VCDClusterSpec), b.(*v1beta1.VCDClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDClusterStatus)(nil), (*v1beta1.VCDClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDClusterStatus_To_v1beta1_VCDClusterStatus(a.(*VCDClusterStatus), b.(*v1beta1.VCDClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachine)(nil), (*v1beta1.VCDMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachine_To_v1beta1_VCDMachine(a.(*VCDMachine), b.(*v1beta1.VCDMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachine)(nil), (*VCDMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachine_To_v1alpha4_VCDMachine(a.(*v1beta1.VCDMachine), b.(*VCDMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineList)(nil), (*v1beta1.VCDMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineList_To_v1beta1_VCDMachineList(a.(*VCDMachineList), b.(*v1beta1.VCDMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachineList)(nil), (*VCDMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineList_To_v1alpha4_VCDMachineList(a.(*v1beta1.VCDMachineList), b.(*VCDMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineStatus)(nil), (*v1beta1.VCDMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineStatus_To_v1beta1_VCDMachineStatus(a.(*VCDMachineStatus), b.(*v1beta1.VCDMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachineStatus)(nil), (*VCDMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineStatus_To_v1alpha4_VCDMachineStatus(a.(*v1beta1.VCDMachineStatus), b.(*VCDMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplate)(nil), (*v1beta1.VCDMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(a.(*VCDMachineTemplate), b.(*v1beta1.VCDMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachineTemplate)(nil), (*VCDMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplate_To_v1alpha4_VCDMachineTemplate(a.(*v1beta1.VCDMachineTemplate), b.(*VCDMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateList)(nil), (*v1beta1.VCDMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(a.(*VCDMachineTemplateList), b.(*v1beta1.VCDMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachineTemplateList)(nil), (*VCDMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateList_To_v1alpha4_VCDMachineTemplateList(a.(*v1beta1.VCDMachineTemplateList), b.(*VCDMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateResource)(nil), (*v1beta1.VCDMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(a.(*VCDMachineTemplateResource), b.(*v1beta1.VCDMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachineTemplateResource)(nil), (*VCDMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateResource_To_v1alpha4_VCDMachineTemplateResource(a.(*v1beta1.VCDMachineTemplateResource), b.(*VCDMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateSpec)(nil), (*v1beta1.VCDMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(a.(*VCDMachineTemplateSpec), b.(*v1beta1.VCDMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachineTemplateSpec)(nil), (*VCDMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateSpec_To_v1alpha4_VCDMachineTemplateSpec(a.(*v1beta1.VCDMachineTemplateSpec), b.(*VCDMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VCDMachineTemplateStatus)(nil), (*v1beta1.VCDMachineTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(a.(*VCDMachineTemplateStatus), b.(*v1beta1.VCDMachineTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VCDMachineTemplateStatus)(nil), (*VCDMachineTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineTemplateStatus_To_v1alpha4_VCDMachineTemplateStatus(a.(*v1beta1.VCDMachineTemplateStatus), b.(*VCDMachineTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*VCDMachineSpec)(nil), (*v1beta1.VCDMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_VCDMachineSpec_To_v1beta1_VCDMachineSpec(a.(*VCDMachineSpec), b.(*v1beta1.VCDMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.VCDClusterSpec)(nil), (*VCDClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDClusterSpec_To_v1alpha4_VCDClusterSpec(a.(*v1beta1.VCDClusterSpec), b.(*VCDClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.VCDClusterStatus)(nil), (*VCDClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDClusterStatus_To_v1alpha4_VCDClusterStatus(a.(*v1beta1.VCDClusterStatus), b.(*VCDClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.VCDMachineSpec)(nil), (*VCDMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VCDMachineSpec_To_v1alpha4_VCDMachineSpec(a.(*v1beta1.VCDMachineSpec), b.(*VCDMachineSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha4_APIEndpoint_To_v1beta1_APIEndpoint(in *APIEndpoint, out *v1beta1.APIEndpoint, s conversion.Scope) error {
	out.Host = in.Host
	out.Port = in.Port
	return nil
}

// Convert_v1alpha4_APIEndpoint_To_v1beta1_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha4_APIEndpoint_To_v1beta1_APIEndpoint(in *APIEndpoint, out *v1beta1.APIEndpoint, s conversion.Scope) error {
	return autoConvert_v1alpha4_APIEndpoint_To_v1beta1_APIEndpoint(in, out, s)
}

func autoConvert_v1beta1_APIEndpoint_To_v1alpha4_APIEndpoint(in *v1beta1.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	out.Host = in.Host
	out.Port = in.Port
	return nil
}

// Convert_v1beta1_APIEndpoint_To_v1alpha4_APIEndpoint is an autogenerated conversion function.
func Convert_v1beta1_APIEndpoint_To_v1alpha4_APIEndpoint(in *v1beta1.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	return autoConvert_v1beta1_APIEndpoint_To_v1alpha4_APIEndpoint(in, out, s)
}

func autoConvert_v1alpha4_UserCredentialsContext_To_v1beta1_UserCredentialsContext(in *UserCredentialsContext, out *v1beta1.UserCredentialsContext, s conversion.Scope) error {
	out.Username = in.Username
	out.Password = in.Password
	out.RefreshToken = in.RefreshToken
	return nil
}

// Convert_v1alpha4_UserCredentialsContext_To_v1beta1_UserCredentialsContext is an autogenerated conversion function.
func Convert_v1alpha4_UserCredentialsContext_To_v1beta1_UserCredentialsContext(in *UserCredentialsContext, out *v1beta1.UserCredentialsContext, s conversion.Scope) error {
	return autoConvert_v1alpha4_UserCredentialsContext_To_v1beta1_UserCredentialsContext(in, out, s)
}

func autoConvert_v1beta1_UserCredentialsContext_To_v1alpha4_UserCredentialsContext(in *v1beta1.UserCredentialsContext, out *UserCredentialsContext, s conversion.Scope) error {
	out.Username = in.Username
	out.Password = in.Password
	out.RefreshToken = in.RefreshToken
	return nil
}

// Convert_v1beta1_UserCredentialsContext_To_v1alpha4_UserCredentialsContext is an autogenerated conversion function.
func Convert_v1beta1_UserCredentialsContext_To_v1alpha4_UserCredentialsContext(in *v1beta1.UserCredentialsContext, out *UserCredentialsContext, s conversion.Scope) error {
	return autoConvert_v1beta1_UserCredentialsContext_To_v1alpha4_UserCredentialsContext(in, out, s)
}

func autoConvert_v1alpha4_VCDCluster_To_v1beta1_VCDCluster(in *VCDCluster, out *v1beta1.VCDCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_VCDClusterSpec_To_v1beta1_VCDClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_VCDClusterStatus_To_v1beta1_VCDClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_VCDCluster_To_v1beta1_VCDCluster is an autogenerated conversion function.
func Convert_v1alpha4_VCDCluster_To_v1beta1_VCDCluster(in *VCDCluster, out *v1beta1.VCDCluster, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDCluster_To_v1beta1_VCDCluster(in, out, s)
}

func autoConvert_v1beta1_VCDCluster_To_v1alpha4_VCDCluster(in *v1beta1.VCDCluster, out *VCDCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VCDClusterSpec_To_v1alpha4_VCDClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VCDClusterStatus_To_v1alpha4_VCDClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDCluster_To_v1alpha4_VCDCluster is an autogenerated conversion function.
func Convert_v1beta1_VCDCluster_To_v1alpha4_VCDCluster(in *v1beta1.VCDCluster, out *VCDCluster, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDCluster_To_v1alpha4_VCDCluster(in, out, s)
}

func autoConvert_v1alpha4_VCDClusterList_To_v1beta1_VCDClusterList(in *VCDClusterList, out *v1beta1.VCDClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.VCDCluster, len(*in))
		for i := range *in {
			if err := Convert_v1alpha4_VCDCluster_To_v1beta1_VCDCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha4_VCDClusterList_To_v1beta1_VCDClusterList is an autogenerated conversion function.
func Convert_v1alpha4_VCDClusterList_To_v1beta1_VCDClusterList(in *VCDClusterList, out *v1beta1.VCDClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDClusterList_To_v1beta1_VCDClusterList(in, out, s)
}

func autoConvert_v1beta1_VCDClusterList_To_v1alpha4_VCDClusterList(in *v1beta1.VCDClusterList, out *VCDClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCDCluster, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_VCDCluster_To_v1alpha4_VCDCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_VCDClusterList_To_v1alpha4_VCDClusterList is an autogenerated conversion function.
func Convert_v1beta1_VCDClusterList_To_v1alpha4_VCDClusterList(in *v1beta1.VCDClusterList, out *VCDClusterList, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDClusterList_To_v1alpha4_VCDClusterList(in, out, s)
}

func autoConvert_v1alpha4_VCDClusterSpec_To_v1beta1_VCDClusterSpec(in *VCDClusterSpec, out *v1beta1.VCDClusterSpec, s conversion.Scope) error {
	if err := Convert_v1alpha4_APIEndpoint_To_v1beta1_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	out.Site = in.Site
	out.Org = in.Org
	out.Ovdc = in.Ovdc
	out.OvdcNetwork = in.OvdcNetwork
	if err := Convert_v1alpha4_UserCredentialsContext_To_v1beta1_UserCredentialsContext(&in.UserCredentialsContext, &out.UserCredentialsContext, s); err != nil {
		return err
	}
	out.DefaultComputePolicy = in.DefaultComputePolicy
	return nil
}

// Convert_v1alpha4_VCDClusterSpec_To_v1beta1_VCDClusterSpec is an autogenerated conversion function.
func Convert_v1alpha4_VCDClusterSpec_To_v1beta1_VCDClusterSpec(in *VCDClusterSpec, out *v1beta1.VCDClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDClusterSpec_To_v1beta1_VCDClusterSpec(in, out, s)
}

func autoConvert_v1beta1_VCDClusterSpec_To_v1alpha4_VCDClusterSpec(in *v1beta1.VCDClusterSpec, out *VCDClusterSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_APIEndpoint_To_v1alpha4_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	out.Site = in.Site
	out.Org = in.Org
	out.Ovdc = in.Ovdc
	out.OvdcNetwork = in.OvdcNetwork
	if err := Convert_v1beta1_UserCredentialsContext_To_v1alpha4_UserCredentialsContext(&in.UserCredentialsContext, &out.UserCredentialsContext, s); err != nil {
		return err
	}
	// WARNING: in.DefaultStorageClassOptions requires manual conversion: does not exist in peer-type
	out.DefaultComputePolicy = in.DefaultComputePolicy
	// WARNING: in.RDEId requires manual conversion: does not exist in peer-type
	// WARNING: in.ParentUID requires manual conversion: does not exist in peer-type
	// WARNING: in.UseAsManagementCluster requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha4_VCDClusterStatus_To_v1beta1_VCDClusterStatus(in *VCDClusterStatus, out *v1beta1.VCDClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.Conditions = *(*apiv1alpha4.Conditions)(unsafe.Pointer(&in.Conditions))
	out.InfraId = in.InfraId
	return nil
}

// Convert_v1alpha4_VCDClusterStatus_To_v1beta1_VCDClusterStatus is an autogenerated conversion function.
func Convert_v1alpha4_VCDClusterStatus_To_v1beta1_VCDClusterStatus(in *VCDClusterStatus, out *v1beta1.VCDClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDClusterStatus_To_v1beta1_VCDClusterStatus(in, out, s)
}

func autoConvert_v1beta1_VCDClusterStatus_To_v1alpha4_VCDClusterStatus(in *v1beta1.VCDClusterStatus, out *VCDClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.Conditions = *(*apiv1alpha4.Conditions)(unsafe.Pointer(&in.Conditions))
	out.InfraId = in.InfraId
	// WARNING: in.ParentUID requires manual conversion: does not exist in peer-type
	// WARNING: in.UseAsManagementCluster requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha4_VCDMachine_To_v1beta1_VCDMachine(in *VCDMachine, out *v1beta1.VCDMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_VCDMachineSpec_To_v1beta1_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_VCDMachineStatus_To_v1beta1_VCDMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_VCDMachine_To_v1beta1_VCDMachine is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachine_To_v1beta1_VCDMachine(in *VCDMachine, out *v1beta1.VCDMachine, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachine_To_v1beta1_VCDMachine(in, out, s)
}

func autoConvert_v1beta1_VCDMachine_To_v1alpha4_VCDMachine(in *v1beta1.VCDMachine, out *VCDMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VCDMachineSpec_To_v1alpha4_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VCDMachineStatus_To_v1alpha4_VCDMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachine_To_v1alpha4_VCDMachine is an autogenerated conversion function.
func Convert_v1beta1_VCDMachine_To_v1alpha4_VCDMachine(in *v1beta1.VCDMachine, out *VCDMachine, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachine_To_v1alpha4_VCDMachine(in, out, s)
}

func autoConvert_v1alpha4_VCDMachineList_To_v1beta1_VCDMachineList(in *VCDMachineList, out *v1beta1.VCDMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.VCDMachine, len(*in))
		for i := range *in {
			if err := Convert_v1alpha4_VCDMachine_To_v1beta1_VCDMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha4_VCDMachineList_To_v1beta1_VCDMachineList is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachineList_To_v1beta1_VCDMachineList(in *VCDMachineList, out *v1beta1.VCDMachineList, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachineList_To_v1beta1_VCDMachineList(in, out, s)
}

func autoConvert_v1beta1_VCDMachineList_To_v1alpha4_VCDMachineList(in *v1beta1.VCDMachineList, out *VCDMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCDMachine, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_VCDMachine_To_v1alpha4_VCDMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_VCDMachineList_To_v1alpha4_VCDMachineList is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineList_To_v1alpha4_VCDMachineList(in *v1beta1.VCDMachineList, out *VCDMachineList, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineList_To_v1alpha4_VCDMachineList(in, out, s)
}

func autoConvert_v1alpha4_VCDMachineSpec_To_v1beta1_VCDMachineSpec(in *VCDMachineSpec, out *v1beta1.VCDMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.Catalog = in.Catalog
	out.Template = in.Template
	// WARNING: in.ComputePolicy requires manual conversion: does not exist in peer-type
	out.Bootstrapped = in.Bootstrapped
	return nil
}

func autoConvert_v1beta1_VCDMachineSpec_To_v1alpha4_VCDMachineSpec(in *v1beta1.VCDMachineSpec, out *VCDMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.Catalog = in.Catalog
	out.Template = in.Template
	// WARNING: in.SizingPolicy requires manual conversion: does not exist in peer-type
	// WARNING: in.StorageProfile requires manual conversion: does not exist in peer-type
	out.Bootstrapped = in.Bootstrapped
	return nil
}

func autoConvert_v1alpha4_VCDMachineStatus_To_v1beta1_VCDMachineStatus(in *VCDMachineStatus, out *v1beta1.VCDMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.Addresses = *(*[]apiv1alpha4.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.Conditions = *(*apiv1alpha4.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha4_VCDMachineStatus_To_v1beta1_VCDMachineStatus is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachineStatus_To_v1beta1_VCDMachineStatus(in *VCDMachineStatus, out *v1beta1.VCDMachineStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachineStatus_To_v1beta1_VCDMachineStatus(in, out, s)
}

func autoConvert_v1beta1_VCDMachineStatus_To_v1alpha4_VCDMachineStatus(in *v1beta1.VCDMachineStatus, out *VCDMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.Addresses = *(*[]apiv1alpha4.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.Conditions = *(*apiv1alpha4.Conditions)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1beta1_VCDMachineStatus_To_v1alpha4_VCDMachineStatus is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineStatus_To_v1alpha4_VCDMachineStatus(in *v1beta1.VCDMachineStatus, out *VCDMachineStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineStatus_To_v1alpha4_VCDMachineStatus(in, out, s)
}

func autoConvert_v1alpha4_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(in *VCDMachineTemplate, out *v1beta1.VCDMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(in *VCDMachineTemplate, out *v1beta1.VCDMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplate_To_v1alpha4_VCDMachineTemplate(in *v1beta1.VCDMachineTemplate, out *VCDMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_VCDMachineTemplateSpec_To_v1alpha4_VCDMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_VCDMachineTemplateStatus_To_v1alpha4_VCDMachineTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplate_To_v1alpha4_VCDMachineTemplate is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplate_To_v1alpha4_VCDMachineTemplate(in *v1beta1.VCDMachineTemplate, out *VCDMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplate_To_v1alpha4_VCDMachineTemplate(in, out, s)
}

func autoConvert_v1alpha4_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(in *VCDMachineTemplateList, out *v1beta1.VCDMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.VCDMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha4_VCDMachineTemplate_To_v1beta1_VCDMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha4_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(in *VCDMachineTemplateList, out *v1beta1.VCDMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachineTemplateList_To_v1beta1_VCDMachineTemplateList(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplateList_To_v1alpha4_VCDMachineTemplateList(in *v1beta1.VCDMachineTemplateList, out *VCDMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VCDMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_VCDMachineTemplate_To_v1alpha4_VCDMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplateList_To_v1alpha4_VCDMachineTemplateList is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateList_To_v1alpha4_VCDMachineTemplateList(in *v1beta1.VCDMachineTemplateList, out *VCDMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateList_To_v1alpha4_VCDMachineTemplateList(in, out, s)
}

func autoConvert_v1alpha4_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(in *VCDMachineTemplateResource, out *v1beta1.VCDMachineTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha4_VCDMachineSpec_To_v1beta1_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(in *VCDMachineTemplateResource, out *v1beta1.VCDMachineTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplateResource_To_v1alpha4_VCDMachineTemplateResource(in *v1beta1.VCDMachineTemplateResource, out *VCDMachineTemplateResource, s conversion.Scope) error {
	if err := Convert_v1beta1_VCDMachineSpec_To_v1alpha4_VCDMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplateResource_To_v1alpha4_VCDMachineTemplateResource is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateResource_To_v1alpha4_VCDMachineTemplateResource(in *v1beta1.VCDMachineTemplateResource, out *VCDMachineTemplateResource, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateResource_To_v1alpha4_VCDMachineTemplateResource(in, out, s)
}

func autoConvert_v1alpha4_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(in *VCDMachineTemplateSpec, out *v1beta1.VCDMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha4_VCDMachineTemplateResource_To_v1beta1_VCDMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(in *VCDMachineTemplateSpec, out *v1beta1.VCDMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachineTemplateSpec_To_v1beta1_VCDMachineTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplateSpec_To_v1alpha4_VCDMachineTemplateSpec(in *v1beta1.VCDMachineTemplateSpec, out *VCDMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_VCDMachineTemplateResource_To_v1alpha4_VCDMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_VCDMachineTemplateSpec_To_v1alpha4_VCDMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateSpec_To_v1alpha4_VCDMachineTemplateSpec(in *v1beta1.VCDMachineTemplateSpec, out *VCDMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateSpec_To_v1alpha4_VCDMachineTemplateSpec(in, out, s)
}

func autoConvert_v1alpha4_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(in *VCDMachineTemplateStatus, out *v1beta1.VCDMachineTemplateStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha4_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus is an autogenerated conversion function.
func Convert_v1alpha4_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(in *VCDMachineTemplateStatus, out *v1beta1.VCDMachineTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_VCDMachineTemplateStatus_To_v1beta1_VCDMachineTemplateStatus(in, out, s)
}

func autoConvert_v1beta1_VCDMachineTemplateStatus_To_v1alpha4_VCDMachineTemplateStatus(in *v1beta1.VCDMachineTemplateStatus, out *VCDMachineTemplateStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_VCDMachineTemplateStatus_To_v1alpha4_VCDMachineTemplateStatus is an autogenerated conversion function.
func Convert_v1beta1_VCDMachineTemplateStatus_To_v1alpha4_VCDMachineTemplateStatus(in *v1beta1.VCDMachineTemplateStatus, out *VCDMachineTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_VCDMachineTemplateStatus_To_v1alpha4_VCDMachineTemplateStatus(in, out, s)
}
