//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package trait

import (
	"k8s.io/api/networking/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AffinityTrait) DeepCopyInto(out *AffinityTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(bool)
		**out = **in
	}
	if in.PodAntiAffinity != nil {
		in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.NodeAffinityLabels != nil {
		in, out := &in.NodeAffinityLabels, &out.NodeAffinityLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodAffinityLabels != nil {
		in, out := &in.PodAffinityLabels, &out.PodAffinityLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodAntiAffinityLabels != nil {
		in, out := &in.PodAntiAffinityLabels, &out.PodAntiAffinityLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AffinityTrait.
func (in *AffinityTrait) DeepCopy() *AffinityTrait {
	if in == nil {
		return nil
	}
	out := new(AffinityTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderTrait) DeepCopyInto(out *BuilderTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.Verbose != nil {
		in, out := &in.Verbose, &out.Verbose
		*out = new(bool)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncrementalImageBuild != nil {
		in, out := &in.IncrementalImageBuild, &out.IncrementalImageBuild
		*out = new(bool)
		**out = **in
	}
	if in.MavenProfiles != nil {
		in, out := &in.MavenProfiles, &out.MavenProfiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TasksRequestCPU != nil {
		in, out := &in.TasksRequestCPU, &out.TasksRequestCPU
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TasksRequestMemory != nil {
		in, out := &in.TasksRequestMemory, &out.TasksRequestMemory
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TasksLimitCPU != nil {
		in, out := &in.TasksLimitCPU, &out.TasksLimitCPU
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TasksLimitMemory != nil {
		in, out := &in.TasksLimitMemory, &out.TasksLimitMemory
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderTrait.
func (in *BuilderTrait) DeepCopy() *BuilderTrait {
	if in == nil {
		return nil
	}
	out := new(BuilderTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CamelTrait) DeepCopyInto(out *CamelTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CamelTrait.
func (in *CamelTrait) DeepCopy() *CamelTrait {
	if in == nil {
		return nil
	}
	out := new(CamelTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerTrait) DeepCopyInto(out *ContainerTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerTrait.
func (in *ContainerTrait) DeepCopy() *ContainerTrait {
	if in == nil {
		return nil
	}
	out := new(ContainerTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronTrait) DeepCopyInto(out *CronTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Fallback != nil {
		in, out := &in.Fallback, &out.Fallback
		*out = new(bool)
		**out = **in
	}
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
	if in.StartingDeadlineSeconds != nil {
		in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronTrait.
func (in *CronTrait) DeepCopy() *CronTrait {
	if in == nil {
		return nil
	}
	out := new(CronTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependenciesTrait) DeepCopyInto(out *DependenciesTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependenciesTrait.
func (in *DependenciesTrait) DeepCopy() *DependenciesTrait {
	if in == nil {
		return nil
	}
	out := new(DependenciesTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployerTrait) DeepCopyInto(out *DeployerTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.UseSSA != nil {
		in, out := &in.UseSSA, &out.UseSSA
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployerTrait.
func (in *DeployerTrait) DeepCopy() *DeployerTrait {
	if in == nil {
		return nil
	}
	out := new(DeployerTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTrait) DeepCopyInto(out *DeploymentTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
	if in.RollingUpdateMaxUnavailable != nil {
		in, out := &in.RollingUpdateMaxUnavailable, &out.RollingUpdateMaxUnavailable
		*out = new(int)
		**out = **in
	}
	if in.RollingUpdateMaxSurge != nil {
		in, out := &in.RollingUpdateMaxSurge, &out.RollingUpdateMaxSurge
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTrait.
func (in *DeploymentTrait) DeepCopy() *DeploymentTrait {
	if in == nil {
		return nil
	}
	out := new(DeploymentTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentTrait) DeepCopyInto(out *EnvironmentTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.ContainerMeta != nil {
		in, out := &in.ContainerMeta, &out.ContainerMeta
		*out = new(bool)
		**out = **in
	}
	if in.HTTPProxy != nil {
		in, out := &in.HTTPProxy, &out.HTTPProxy
		*out = new(bool)
		**out = **in
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentTrait.
func (in *EnvironmentTrait) DeepCopy() *EnvironmentTrait {
	if in == nil {
		return nil
	}
	out := new(EnvironmentTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerTrait) DeepCopyInto(out *ErrorHandlerTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerTrait.
func (in *ErrorHandlerTrait) DeepCopy() *ErrorHandlerTrait {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCTrait) DeepCopyInto(out *GCTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.DiscoveryCache != nil {
		in, out := &in.DiscoveryCache, &out.DiscoveryCache
		*out = new(DiscoveryCacheType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCTrait.
func (in *GCTrait) DeepCopy() *GCTrait {
	if in == nil {
		return nil
	}
	out := new(GCTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthTrait) DeepCopyInto(out *HealthTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.LivenessProbeEnabled != nil {
		in, out := &in.LivenessProbeEnabled, &out.LivenessProbeEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ReadinessProbeEnabled != nil {
		in, out := &in.ReadinessProbeEnabled, &out.ReadinessProbeEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StartupProbeEnabled != nil {
		in, out := &in.StartupProbeEnabled, &out.StartupProbeEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthTrait.
func (in *HealthTrait) DeepCopy() *HealthTrait {
	if in == nil {
		return nil
	}
	out := new(HealthTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTrait) DeepCopyInto(out *IngressTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PathType != nil {
		in, out := &in.PathType, &out.PathType
		*out = new(v1.PathType)
		**out = **in
	}
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTrait.
func (in *IngressTrait) DeepCopy() *IngressTrait {
	if in == nil {
		return nil
	}
	out := new(IngressTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioTrait) DeepCopyInto(out *IstioTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Inject != nil {
		in, out := &in.Inject, &out.Inject
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioTrait.
func (in *IstioTrait) DeepCopy() *IstioTrait {
	if in == nil {
		return nil
	}
	out := new(IstioTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JVMTrait) DeepCopyInto(out *JVMTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(bool)
		**out = **in
	}
	if in.DebugSuspend != nil {
		in, out := &in.DebugSuspend, &out.DebugSuspend
		*out = new(bool)
		**out = **in
	}
	if in.PrintCommand != nil {
		in, out := &in.PrintCommand, &out.PrintCommand
		*out = new(bool)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JVMTrait.
func (in *JVMTrait) DeepCopy() *JVMTrait {
	if in == nil {
		return nil
	}
	out := new(JVMTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JolokiaTrait) DeepCopyInto(out *JolokiaTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.CaCert != nil {
		in, out := &in.CaCert, &out.CaCert
		*out = new(string)
		**out = **in
	}
	if in.ClientPrincipal != nil {
		in, out := &in.ClientPrincipal, &out.ClientPrincipal
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DiscoveryEnabled != nil {
		in, out := &in.DiscoveryEnabled, &out.DiscoveryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ExtendedClientCheck != nil {
		in, out := &in.ExtendedClientCheck, &out.ExtendedClientCheck
		*out = new(bool)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
	if in.UseSslClientAuthentication != nil {
		in, out := &in.UseSslClientAuthentication, &out.UseSslClientAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JolokiaTrait.
func (in *JolokiaTrait) DeepCopy() *JolokiaTrait {
	if in == nil {
		return nil
	}
	out := new(JolokiaTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletsTrait) DeepCopyInto(out *KameletsTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletsTrait.
func (in *KameletsTrait) DeepCopy() *KameletsTrait {
	if in == nil {
		return nil
	}
	out := new(KameletsTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServiceTrait) DeepCopyInto(out *KnativeServiceTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(int)
		**out = **in
	}
	if in.MinScale != nil {
		in, out := &in.MinScale, &out.MinScale
		*out = new(int)
		**out = **in
	}
	if in.MaxScale != nil {
		in, out := &in.MaxScale, &out.MaxScale
		*out = new(int)
		**out = **in
	}
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServiceTrait.
func (in *KnativeServiceTrait) DeepCopy() *KnativeServiceTrait {
	if in == nil {
		return nil
	}
	out := new(KnativeServiceTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeTrait) DeepCopyInto(out *KnativeTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.ChannelSources != nil {
		in, out := &in.ChannelSources, &out.ChannelSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ChannelSinks != nil {
		in, out := &in.ChannelSinks, &out.ChannelSinks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EndpointSources != nil {
		in, out := &in.EndpointSources, &out.EndpointSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EndpointSinks != nil {
		in, out := &in.EndpointSinks, &out.EndpointSinks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EventSources != nil {
		in, out := &in.EventSources, &out.EventSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EventSinks != nil {
		in, out := &in.EventSinks, &out.EventSinks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FilterSourceChannels != nil {
		in, out := &in.FilterSourceChannels, &out.FilterSourceChannels
		*out = new(bool)
		**out = **in
	}
	if in.SinkBinding != nil {
		in, out := &in.SinkBinding, &out.SinkBinding
		*out = new(bool)
		**out = **in
	}
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
	if in.NamespaceLabel != nil {
		in, out := &in.NamespaceLabel, &out.NamespaceLabel
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeTrait.
func (in *KnativeTrait) DeepCopy() *KnativeTrait {
	if in == nil {
		return nil
	}
	out := new(KnativeTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingTrait) DeepCopyInto(out *LoggingTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = new(bool)
		**out = **in
	}
	if in.JSON != nil {
		in, out := &in.JSON, &out.JSON
		*out = new(bool)
		**out = **in
	}
	if in.JSONPrettyPrint != nil {
		in, out := &in.JSONPrettyPrint, &out.JSONPrettyPrint
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingTrait.
func (in *LoggingTrait) DeepCopy() *LoggingTrait {
	if in == nil {
		return nil
	}
	out := new(LoggingTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTrait) DeepCopyInto(out *MountTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.Configs != nil {
		in, out := &in.Configs, &out.Configs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HotReload != nil {
		in, out := &in.HotReload, &out.HotReload
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTrait.
func (in *MountTrait) DeepCopy() *MountTrait {
	if in == nil {
		return nil
	}
	out := new(MountTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenAPITrait) DeepCopyInto(out *OpenAPITrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.Configmaps != nil {
		in, out := &in.Configmaps, &out.Configmaps
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenAPITrait.
func (in *OpenAPITrait) DeepCopy() *OpenAPITrait {
	if in == nil {
		return nil
	}
	out := new(OpenAPITrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerTrait) DeepCopyInto(out *OwnerTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.TargetAnnotations != nil {
		in, out := &in.TargetAnnotations, &out.TargetAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TargetLabels != nil {
		in, out := &in.TargetLabels, &out.TargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerTrait.
func (in *OwnerTrait) DeepCopy() *OwnerTrait {
	if in == nil {
		return nil
	}
	out := new(OwnerTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PDBTrait) DeepCopyInto(out *PDBTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PDBTrait.
func (in *PDBTrait) DeepCopy() *PDBTrait {
	if in == nil {
		return nil
	}
	out := new(PDBTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformBaseTrait) DeepCopyInto(out *PlatformBaseTrait) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(Configuration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformBaseTrait.
func (in *PlatformBaseTrait) DeepCopy() *PlatformBaseTrait {
	if in == nil {
		return nil
	}
	out := new(PlatformBaseTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformTrait) DeepCopyInto(out *PlatformTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.CreateDefault != nil {
		in, out := &in.CreateDefault, &out.CreateDefault
		*out = new(bool)
		**out = **in
	}
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(bool)
		**out = **in
	}
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformTrait.
func (in *PlatformTrait) DeepCopy() *PlatformTrait {
	if in == nil {
		return nil
	}
	out := new(PlatformTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTrait) DeepCopyInto(out *PodTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTrait.
func (in *PodTrait) DeepCopy() *PodTrait {
	if in == nil {
		return nil
	}
	out := new(PodTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusTrait) DeepCopyInto(out *PrometheusTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.PodMonitor != nil {
		in, out := &in.PodMonitor, &out.PodMonitor
		*out = new(bool)
		**out = **in
	}
	if in.PodMonitorLabels != nil {
		in, out := &in.PodMonitorLabels, &out.PodMonitorLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusTrait.
func (in *PrometheusTrait) DeepCopy() *PrometheusTrait {
	if in == nil {
		return nil
	}
	out := new(PrometheusTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullSecretTrait) DeepCopyInto(out *PullSecretTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.ImagePullerDelegation != nil {
		in, out := &in.ImagePullerDelegation, &out.ImagePullerDelegation
		*out = new(bool)
		**out = **in
	}
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullSecretTrait.
func (in *PullSecretTrait) DeepCopy() *PullSecretTrait {
	if in == nil {
		return nil
	}
	out := new(PullSecretTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuarkusTrait) DeepCopyInto(out *QuarkusTrait) {
	*out = *in
	in.PlatformBaseTrait.DeepCopyInto(&out.PlatformBaseTrait)
	if in.PackageTypes != nil {
		in, out := &in.PackageTypes, &out.PackageTypes
		*out = make([]QuarkusPackageType, len(*in))
		copy(*out, *in)
	}
	if in.Modes != nil {
		in, out := &in.Modes, &out.Modes
		*out = make([]QuarkusMode, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuarkusTrait.
func (in *QuarkusTrait) DeepCopy() *QuarkusTrait {
	if in == nil {
		return nil
	}
	out := new(QuarkusTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RawMessage) DeepCopyInto(out *RawMessage) {
	{
		in := &in
		*out = make(RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawMessage.
func (in RawMessage) DeepCopy() RawMessage {
	if in == nil {
		return nil
	}
	out := new(RawMessage)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryTrait) DeepCopyInto(out *RegistryTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryTrait.
func (in *RegistryTrait) DeepCopy() *RegistryTrait {
	if in == nil {
		return nil
	}
	out := new(RegistryTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTrait) DeepCopyInto(out *RouteTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTrait.
func (in *RouteTrait) DeepCopy() *RouteTrait {
	if in == nil {
		return nil
	}
	out := new(RouteTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingTrait) DeepCopyInto(out *ServiceBindingTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingTrait.
func (in *ServiceBindingTrait) DeepCopy() *ServiceBindingTrait {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTrait) DeepCopyInto(out *ServiceTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(bool)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(ServiceType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTrait.
func (in *ServiceTrait) DeepCopy() *ServiceTrait {
	if in == nil {
		return nil
	}
	out := new(ServiceTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TolerationTrait) DeepCopyInto(out *TolerationTrait) {
	*out = *in
	in.Trait.DeepCopyInto(&out.Trait)
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TolerationTrait.
func (in *TolerationTrait) DeepCopy() *TolerationTrait {
	if in == nil {
		return nil
	}
	out := new(TolerationTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trait) DeepCopyInto(out *Trait) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(Configuration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trait.
func (in *Trait) DeepCopy() *Trait {
	if in == nil {
		return nil
	}
	out := new(Trait)
	in.DeepCopyInto(out)
	return out
}
