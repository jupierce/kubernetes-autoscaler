// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package apps

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/client-go/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_ControllerRevision, InType: reflect.TypeOf(&ControllerRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_ControllerRevisionList, InType: reflect.TypeOf(&ControllerRevisionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_PartitionStatefulSetStrategy, InType: reflect.TypeOf(&PartitionStatefulSetStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_StatefulSet, InType: reflect.TypeOf(&StatefulSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_StatefulSetList, InType: reflect.TypeOf(&StatefulSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_StatefulSetSpec, InType: reflect.TypeOf(&StatefulSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_StatefulSetStatus, InType: reflect.TypeOf(&StatefulSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apps_StatefulSetUpdateStrategy, InType: reflect.TypeOf(&StatefulSetUpdateStrategy{})},
	)
}

// DeepCopy_apps_ControllerRevision is an autogenerated deepcopy function.
func DeepCopy_apps_ControllerRevision(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ControllerRevision)
		out := out.(*ControllerRevision)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		// in.Data is kind 'Interface'
		if in.Data != nil {
			if newVal, err := c.DeepCopy(&in.Data); err != nil {
				return err
			} else {
				out.Data = *newVal.(*runtime.Object)
			}
		}
		return nil
	}
}

// DeepCopy_apps_ControllerRevisionList is an autogenerated deepcopy function.
func DeepCopy_apps_ControllerRevisionList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ControllerRevisionList)
		out := out.(*ControllerRevisionList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ControllerRevision, len(*in))
			for i := range *in {
				if err := DeepCopy_apps_ControllerRevision(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_apps_PartitionStatefulSetStrategy is an autogenerated deepcopy function.
func DeepCopy_apps_PartitionStatefulSetStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PartitionStatefulSetStrategy)
		out := out.(*PartitionStatefulSetStrategy)
		*out = *in
		return nil
	}
}

// DeepCopy_apps_StatefulSet is an autogenerated deepcopy function.
func DeepCopy_apps_StatefulSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSet)
		out := out.(*StatefulSet)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_apps_StatefulSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_apps_StatefulSetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_apps_StatefulSetList is an autogenerated deepcopy function.
func DeepCopy_apps_StatefulSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetList)
		out := out.(*StatefulSetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]StatefulSet, len(*in))
			for i := range *in {
				if err := DeepCopy_apps_StatefulSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_apps_StatefulSetSpec is an autogenerated deepcopy function.
func DeepCopy_apps_StatefulSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetSpec)
		out := out.(*StatefulSetSpec)
		*out = *in
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if err := api.DeepCopy_api_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.VolumeClaimTemplates != nil {
			in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
			*out = make([]api.PersistentVolumeClaim, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_PersistentVolumeClaim(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if err := DeepCopy_apps_StatefulSetUpdateStrategy(&in.UpdateStrategy, &out.UpdateStrategy, c); err != nil {
			return err
		}
		if in.RevisionHistoryLimit != nil {
			in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_apps_StatefulSetStatus is an autogenerated deepcopy function.
func DeepCopy_apps_StatefulSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetStatus)
		out := out.(*StatefulSetStatus)
		*out = *in
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_apps_StatefulSetUpdateStrategy is an autogenerated deepcopy function.
func DeepCopy_apps_StatefulSetUpdateStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetUpdateStrategy)
		out := out.(*StatefulSetUpdateStrategy)
		*out = *in
		if in.Partition != nil {
			in, out := &in.Partition, &out.Partition
			*out = new(PartitionStatefulSetStrategy)
			**out = **in
		}
		return nil
	}
}