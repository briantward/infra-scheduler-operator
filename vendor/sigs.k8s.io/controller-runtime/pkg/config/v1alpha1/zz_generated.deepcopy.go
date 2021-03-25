// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerHealth) DeepCopyInto(out *ControllerHealth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerHealth.
func (in *ControllerHealth) DeepCopy() *ControllerHealth {
	if in == nil {
		return nil
	}
	out := new(ControllerHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfiguration) DeepCopyInto(out *ControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfiguration.
func (in *ControllerManagerConfiguration) DeepCopy() *ControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfigurationSpec) DeepCopyInto(out *ControllerManagerConfigurationSpec) {
	*out = *in
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(configv1alpha1.LeaderElectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.GracefulShutdownTimeout != nil {
		in, out := &in.GracefulShutdownTimeout, &out.GracefulShutdownTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	out.Metrics = in.Metrics
	out.Health = in.Health
	in.Webhook.DeepCopyInto(&out.Webhook)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfigurationSpec.
func (in *ControllerManagerConfigurationSpec) DeepCopy() *ControllerManagerConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerMetrics) DeepCopyInto(out *ControllerMetrics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerMetrics.
func (in *ControllerMetrics) DeepCopy() *ControllerMetrics {
	if in == nil {
		return nil
	}
	out := new(ControllerMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerWebhook) DeepCopyInto(out *ControllerWebhook) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerWebhook.
func (in *ControllerWebhook) DeepCopy() *ControllerWebhook {
	if in == nil {
		return nil
	}
	out := new(ControllerWebhook)
	in.DeepCopyInto(out)
	return out
}