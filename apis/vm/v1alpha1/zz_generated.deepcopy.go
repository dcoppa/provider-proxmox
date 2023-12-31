//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskObservation) DeepCopyInto(out *DiskObservation) {
	*out = *in
	if in.Aio != nil {
		in, out := &in.Aio, &out.Aio
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(string)
		**out = **in
	}
	if in.Discard != nil {
		in, out := &in.Discard, &out.Discard
		*out = new(string)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.IopsMax != nil {
		in, out := &in.IopsMax, &out.IopsMax
		*out = new(float64)
		**out = **in
	}
	if in.IopsMaxLength != nil {
		in, out := &in.IopsMaxLength, &out.IopsMaxLength
		*out = new(float64)
		**out = **in
	}
	if in.IopsRd != nil {
		in, out := &in.IopsRd, &out.IopsRd
		*out = new(float64)
		**out = **in
	}
	if in.IopsRdMax != nil {
		in, out := &in.IopsRdMax, &out.IopsRdMax
		*out = new(float64)
		**out = **in
	}
	if in.IopsRdMaxLength != nil {
		in, out := &in.IopsRdMaxLength, &out.IopsRdMaxLength
		*out = new(float64)
		**out = **in
	}
	if in.IopsWr != nil {
		in, out := &in.IopsWr, &out.IopsWr
		*out = new(float64)
		**out = **in
	}
	if in.IopsWrMax != nil {
		in, out := &in.IopsWrMax, &out.IopsWrMax
		*out = new(float64)
		**out = **in
	}
	if in.IopsWrMaxLength != nil {
		in, out := &in.IopsWrMaxLength, &out.IopsWrMaxLength
		*out = new(float64)
		**out = **in
	}
	if in.Iothread != nil {
		in, out := &in.Iothread, &out.Iothread
		*out = new(float64)
		**out = **in
	}
	if in.Mbps != nil {
		in, out := &in.Mbps, &out.Mbps
		*out = new(float64)
		**out = **in
	}
	if in.MbpsRd != nil {
		in, out := &in.MbpsRd, &out.MbpsRd
		*out = new(float64)
		**out = **in
	}
	if in.MbpsRdMax != nil {
		in, out := &in.MbpsRdMax, &out.MbpsRdMax
		*out = new(float64)
		**out = **in
	}
	if in.MbpsWr != nil {
		in, out := &in.MbpsWr, &out.MbpsWr
		*out = new(float64)
		**out = **in
	}
	if in.MbpsWrMax != nil {
		in, out := &in.MbpsWrMax, &out.MbpsWrMax
		*out = new(float64)
		**out = **in
	}
	if in.Media != nil {
		in, out := &in.Media, &out.Media
		*out = new(string)
		**out = **in
	}
	if in.Replicate != nil {
		in, out := &in.Replicate, &out.Replicate
		*out = new(float64)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Slot != nil {
		in, out := &in.Slot, &out.Slot
		*out = new(float64)
		**out = **in
	}
	if in.Ssd != nil {
		in, out := &in.Ssd, &out.Ssd
		*out = new(float64)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskObservation.
func (in *DiskObservation) DeepCopy() *DiskObservation {
	if in == nil {
		return nil
	}
	out := new(DiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskParameters) DeepCopyInto(out *DiskParameters) {
	*out = *in
	if in.Aio != nil {
		in, out := &in.Aio, &out.Aio
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(string)
		**out = **in
	}
	if in.Discard != nil {
		in, out := &in.Discard, &out.Discard
		*out = new(string)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.IopsMax != nil {
		in, out := &in.IopsMax, &out.IopsMax
		*out = new(float64)
		**out = **in
	}
	if in.IopsMaxLength != nil {
		in, out := &in.IopsMaxLength, &out.IopsMaxLength
		*out = new(float64)
		**out = **in
	}
	if in.IopsRd != nil {
		in, out := &in.IopsRd, &out.IopsRd
		*out = new(float64)
		**out = **in
	}
	if in.IopsRdMax != nil {
		in, out := &in.IopsRdMax, &out.IopsRdMax
		*out = new(float64)
		**out = **in
	}
	if in.IopsRdMaxLength != nil {
		in, out := &in.IopsRdMaxLength, &out.IopsRdMaxLength
		*out = new(float64)
		**out = **in
	}
	if in.IopsWr != nil {
		in, out := &in.IopsWr, &out.IopsWr
		*out = new(float64)
		**out = **in
	}
	if in.IopsWrMax != nil {
		in, out := &in.IopsWrMax, &out.IopsWrMax
		*out = new(float64)
		**out = **in
	}
	if in.IopsWrMaxLength != nil {
		in, out := &in.IopsWrMaxLength, &out.IopsWrMaxLength
		*out = new(float64)
		**out = **in
	}
	if in.Iothread != nil {
		in, out := &in.Iothread, &out.Iothread
		*out = new(float64)
		**out = **in
	}
	if in.Mbps != nil {
		in, out := &in.Mbps, &out.Mbps
		*out = new(float64)
		**out = **in
	}
	if in.MbpsRd != nil {
		in, out := &in.MbpsRd, &out.MbpsRd
		*out = new(float64)
		**out = **in
	}
	if in.MbpsRdMax != nil {
		in, out := &in.MbpsRdMax, &out.MbpsRdMax
		*out = new(float64)
		**out = **in
	}
	if in.MbpsWr != nil {
		in, out := &in.MbpsWr, &out.MbpsWr
		*out = new(float64)
		**out = **in
	}
	if in.MbpsWrMax != nil {
		in, out := &in.MbpsWrMax, &out.MbpsWrMax
		*out = new(float64)
		**out = **in
	}
	if in.Media != nil {
		in, out := &in.Media, &out.Media
		*out = new(string)
		**out = **in
	}
	if in.Replicate != nil {
		in, out := &in.Replicate, &out.Replicate
		*out = new(float64)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Slot != nil {
		in, out := &in.Slot, &out.Slot
		*out = new(float64)
		**out = **in
	}
	if in.Ssd != nil {
		in, out := &in.Ssd, &out.Ssd
		*out = new(float64)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskParameters.
func (in *DiskParameters) DeepCopy() *DiskParameters {
	if in == nil {
		return nil
	}
	out := new(DiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostpciObservation) DeepCopyInto(out *HostpciObservation) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Pcie != nil {
		in, out := &in.Pcie, &out.Pcie
		*out = new(float64)
		**out = **in
	}
	if in.Rombar != nil {
		in, out := &in.Rombar, &out.Rombar
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostpciObservation.
func (in *HostpciObservation) DeepCopy() *HostpciObservation {
	if in == nil {
		return nil
	}
	out := new(HostpciObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostpciParameters) DeepCopyInto(out *HostpciParameters) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Pcie != nil {
		in, out := &in.Pcie, &out.Pcie
		*out = new(float64)
		**out = **in
	}
	if in.Rombar != nil {
		in, out := &in.Rombar, &out.Rombar
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostpciParameters.
func (in *HostpciParameters) DeepCopy() *HostpciParameters {
	if in == nil {
		return nil
	}
	out := new(HostpciParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.Bridge != nil {
		in, out := &in.Bridge, &out.Bridge
		*out = new(string)
		**out = **in
	}
	if in.Firewall != nil {
		in, out := &in.Firewall, &out.Firewall
		*out = new(bool)
		**out = **in
	}
	if in.LinkDown != nil {
		in, out := &in.LinkDown, &out.LinkDown
		*out = new(bool)
		**out = **in
	}
	if in.Macaddr != nil {
		in, out := &in.Macaddr, &out.Macaddr
		*out = new(string)
		**out = **in
	}
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = new(float64)
		**out = **in
	}
	if in.Rate != nil {
		in, out := &in.Rate, &out.Rate
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.Bridge != nil {
		in, out := &in.Bridge, &out.Bridge
		*out = new(string)
		**out = **in
	}
	if in.Firewall != nil {
		in, out := &in.Firewall, &out.Firewall
		*out = new(bool)
		**out = **in
	}
	if in.LinkDown != nil {
		in, out := &in.LinkDown, &out.LinkDown
		*out = new(bool)
		**out = **in
	}
	if in.Macaddr != nil {
		in, out := &in.Macaddr, &out.Macaddr
		*out = new(string)
		**out = **in
	}
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = new(float64)
		**out = **in
	}
	if in.Rate != nil {
		in, out := &in.Rate, &out.Rate
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Qemu) DeepCopyInto(out *Qemu) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Qemu.
func (in *Qemu) DeepCopy() *Qemu {
	if in == nil {
		return nil
	}
	out := new(Qemu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Qemu) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuList) DeepCopyInto(out *QemuList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Qemu, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuList.
func (in *QemuList) DeepCopy() *QemuList {
	if in == nil {
		return nil
	}
	out := new(QemuList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QemuList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuObservation) DeepCopyInto(out *QemuObservation) {
	*out = *in
	if in.AdditionalWait != nil {
		in, out := &in.AdditionalWait, &out.AdditionalWait
		*out = new(float64)
		**out = **in
	}
	if in.Agent != nil {
		in, out := &in.Agent, &out.Agent
		*out = new(float64)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.AutomaticReboot != nil {
		in, out := &in.AutomaticReboot, &out.AutomaticReboot
		*out = new(bool)
		**out = **in
	}
	if in.Balloon != nil {
		in, out := &in.Balloon, &out.Balloon
		*out = new(float64)
		**out = **in
	}
	if in.Bios != nil {
		in, out := &in.Bios, &out.Bios
		*out = new(string)
		**out = **in
	}
	if in.Boot != nil {
		in, out := &in.Boot, &out.Boot
		*out = new(string)
		**out = **in
	}
	if in.Bootdisk != nil {
		in, out := &in.Bootdisk, &out.Bootdisk
		*out = new(string)
		**out = **in
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.CiWait != nil {
		in, out := &in.CiWait, &out.CiWait
		*out = new(float64)
		**out = **in
	}
	if in.Cicustom != nil {
		in, out := &in.Cicustom, &out.Cicustom
		*out = new(string)
		**out = **in
	}
	if in.Ciuser != nil {
		in, out := &in.Ciuser, &out.Ciuser
		*out = new(string)
		**out = **in
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = new(string)
		**out = **in
	}
	if in.CloneWait != nil {
		in, out := &in.CloneWait, &out.CloneWait
		*out = new(float64)
		**out = **in
	}
	if in.CloudinitCdromStorage != nil {
		in, out := &in.CloudinitCdromStorage, &out.CloudinitCdromStorage
		*out = new(string)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.DefaultIPv4Address != nil {
		in, out := &in.DefaultIPv4Address, &out.DefaultIPv4Address
		*out = new(string)
		**out = **in
	}
	if in.DefineConnectionInfo != nil {
		in, out := &in.DefineConnectionInfo, &out.DefineConnectionInfo
		*out = new(bool)
		**out = **in
	}
	if in.Desc != nil {
		in, out := &in.Desc, &out.Desc
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]DiskObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceCreate != nil {
		in, out := &in.ForceCreate, &out.ForceCreate
		*out = new(bool)
		**out = **in
	}
	if in.ForceRecreateOnChangeOf != nil {
		in, out := &in.ForceRecreateOnChangeOf, &out.ForceRecreateOnChangeOf
		*out = new(string)
		**out = **in
	}
	if in.FullClone != nil {
		in, out := &in.FullClone, &out.FullClone
		*out = new(bool)
		**out = **in
	}
	if in.Hagroup != nil {
		in, out := &in.Hagroup, &out.Hagroup
		*out = new(string)
		**out = **in
	}
	if in.Hastate != nil {
		in, out := &in.Hastate, &out.Hastate
		*out = new(string)
		**out = **in
	}
	if in.Hostpci != nil {
		in, out := &in.Hostpci, &out.Hostpci
		*out = make([]HostpciObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hotplug != nil {
		in, out := &in.Hotplug, &out.Hotplug
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig0 != nil {
		in, out := &in.Ipconfig0, &out.Ipconfig0
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig1 != nil {
		in, out := &in.Ipconfig1, &out.Ipconfig1
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig10 != nil {
		in, out := &in.Ipconfig10, &out.Ipconfig10
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig11 != nil {
		in, out := &in.Ipconfig11, &out.Ipconfig11
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig12 != nil {
		in, out := &in.Ipconfig12, &out.Ipconfig12
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig13 != nil {
		in, out := &in.Ipconfig13, &out.Ipconfig13
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig14 != nil {
		in, out := &in.Ipconfig14, &out.Ipconfig14
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig15 != nil {
		in, out := &in.Ipconfig15, &out.Ipconfig15
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig2 != nil {
		in, out := &in.Ipconfig2, &out.Ipconfig2
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig3 != nil {
		in, out := &in.Ipconfig3, &out.Ipconfig3
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig4 != nil {
		in, out := &in.Ipconfig4, &out.Ipconfig4
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig5 != nil {
		in, out := &in.Ipconfig5, &out.Ipconfig5
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig6 != nil {
		in, out := &in.Ipconfig6, &out.Ipconfig6
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig7 != nil {
		in, out := &in.Ipconfig7, &out.Ipconfig7
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig8 != nil {
		in, out := &in.Ipconfig8, &out.Ipconfig8
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig9 != nil {
		in, out := &in.Ipconfig9, &out.Ipconfig9
		*out = new(string)
		**out = **in
	}
	if in.Iso != nil {
		in, out := &in.Iso, &out.Iso
		*out = new(string)
		**out = **in
	}
	if in.Kvm != nil {
		in, out := &in.Kvm, &out.Kvm
		*out = new(bool)
		**out = **in
	}
	if in.Machine != nil {
		in, out := &in.Machine, &out.Machine
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.Nameserver != nil {
		in, out := &in.Nameserver, &out.Nameserver
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Numa != nil {
		in, out := &in.Numa, &out.Numa
		*out = new(bool)
		**out = **in
	}
	if in.Onboot != nil {
		in, out := &in.Onboot, &out.Onboot
		*out = new(bool)
		**out = **in
	}
	if in.VMstate != nil {
		in, out := &in.VMstate, &out.VMstate
		*out = new(string)
		**out = **in
	}
	if in.OsNetworkConfig != nil {
		in, out := &in.OsNetworkConfig, &out.OsNetworkConfig
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.Pxe != nil {
		in, out := &in.Pxe, &out.Pxe
		*out = new(bool)
		**out = **in
	}
	if in.QemuOs != nil {
		in, out := &in.QemuOs, &out.QemuOs
		*out = new(string)
		**out = **in
	}
	if in.RebootRequired != nil {
		in, out := &in.RebootRequired, &out.RebootRequired
		*out = new(bool)
		**out = **in
	}
	if in.SSHForwardIP != nil {
		in, out := &in.SSHForwardIP, &out.SSHForwardIP
		*out = new(string)
		**out = **in
	}
	if in.SSHHost != nil {
		in, out := &in.SSHHost, &out.SSHHost
		*out = new(string)
		**out = **in
	}
	if in.SSHPort != nil {
		in, out := &in.SSHPort, &out.SSHPort
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.Scsihw != nil {
		in, out := &in.Scsihw, &out.Scsihw
		*out = new(string)
		**out = **in
	}
	if in.Searchdomain != nil {
		in, out := &in.Searchdomain, &out.Searchdomain
		*out = new(string)
		**out = **in
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = make([]SerialObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sockets != nil {
		in, out := &in.Sockets, &out.Sockets
		*out = new(float64)
		**out = **in
	}
	if in.Sshkeys != nil {
		in, out := &in.Sshkeys, &out.Sshkeys
		*out = new(string)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = new(string)
		**out = **in
	}
	if in.Tablet != nil {
		in, out := &in.Tablet, &out.Tablet
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.TargetNode != nil {
		in, out := &in.TargetNode, &out.TargetNode
		*out = new(string)
		**out = **in
	}
	if in.UnusedDisk != nil {
		in, out := &in.UnusedDisk, &out.UnusedDisk
		*out = make([]UnusedDiskObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Usb != nil {
		in, out := &in.Usb, &out.Usb
		*out = make([]UsbObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vcpus != nil {
		in, out := &in.Vcpus, &out.Vcpus
		*out = new(float64)
		**out = **in
	}
	if in.Vga != nil {
		in, out := &in.Vga, &out.Vga
		*out = make([]VgaObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vmid != nil {
		in, out := &in.Vmid, &out.Vmid
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuObservation.
func (in *QemuObservation) DeepCopy() *QemuObservation {
	if in == nil {
		return nil
	}
	out := new(QemuObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuParameters) DeepCopyInto(out *QemuParameters) {
	*out = *in
	if in.AdditionalWait != nil {
		in, out := &in.AdditionalWait, &out.AdditionalWait
		*out = new(float64)
		**out = **in
	}
	if in.Agent != nil {
		in, out := &in.Agent, &out.Agent
		*out = new(float64)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.AutomaticReboot != nil {
		in, out := &in.AutomaticReboot, &out.AutomaticReboot
		*out = new(bool)
		**out = **in
	}
	if in.Balloon != nil {
		in, out := &in.Balloon, &out.Balloon
		*out = new(float64)
		**out = **in
	}
	if in.Bios != nil {
		in, out := &in.Bios, &out.Bios
		*out = new(string)
		**out = **in
	}
	if in.Boot != nil {
		in, out := &in.Boot, &out.Boot
		*out = new(string)
		**out = **in
	}
	if in.Bootdisk != nil {
		in, out := &in.Bootdisk, &out.Bootdisk
		*out = new(string)
		**out = **in
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(string)
		**out = **in
	}
	if in.CiWait != nil {
		in, out := &in.CiWait, &out.CiWait
		*out = new(float64)
		**out = **in
	}
	if in.Cicustom != nil {
		in, out := &in.Cicustom, &out.Cicustom
		*out = new(string)
		**out = **in
	}
	if in.CipasswordSecretRef != nil {
		in, out := &in.CipasswordSecretRef, &out.CipasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Ciuser != nil {
		in, out := &in.Ciuser, &out.Ciuser
		*out = new(string)
		**out = **in
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = new(string)
		**out = **in
	}
	if in.CloneWait != nil {
		in, out := &in.CloneWait, &out.CloneWait
		*out = new(float64)
		**out = **in
	}
	if in.CloudinitCdromStorage != nil {
		in, out := &in.CloudinitCdromStorage, &out.CloudinitCdromStorage
		*out = new(string)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.DefineConnectionInfo != nil {
		in, out := &in.DefineConnectionInfo, &out.DefineConnectionInfo
		*out = new(bool)
		**out = **in
	}
	if in.Desc != nil {
		in, out := &in.Desc, &out.Desc
		*out = new(string)
		**out = **in
	}
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		*out = make([]DiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceCreate != nil {
		in, out := &in.ForceCreate, &out.ForceCreate
		*out = new(bool)
		**out = **in
	}
	if in.ForceRecreateOnChangeOf != nil {
		in, out := &in.ForceRecreateOnChangeOf, &out.ForceRecreateOnChangeOf
		*out = new(string)
		**out = **in
	}
	if in.FullClone != nil {
		in, out := &in.FullClone, &out.FullClone
		*out = new(bool)
		**out = **in
	}
	if in.Hagroup != nil {
		in, out := &in.Hagroup, &out.Hagroup
		*out = new(string)
		**out = **in
	}
	if in.Hastate != nil {
		in, out := &in.Hastate, &out.Hastate
		*out = new(string)
		**out = **in
	}
	if in.Hostpci != nil {
		in, out := &in.Hostpci, &out.Hostpci
		*out = make([]HostpciParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hotplug != nil {
		in, out := &in.Hotplug, &out.Hotplug
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig0 != nil {
		in, out := &in.Ipconfig0, &out.Ipconfig0
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig1 != nil {
		in, out := &in.Ipconfig1, &out.Ipconfig1
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig10 != nil {
		in, out := &in.Ipconfig10, &out.Ipconfig10
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig11 != nil {
		in, out := &in.Ipconfig11, &out.Ipconfig11
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig12 != nil {
		in, out := &in.Ipconfig12, &out.Ipconfig12
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig13 != nil {
		in, out := &in.Ipconfig13, &out.Ipconfig13
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig14 != nil {
		in, out := &in.Ipconfig14, &out.Ipconfig14
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig15 != nil {
		in, out := &in.Ipconfig15, &out.Ipconfig15
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig2 != nil {
		in, out := &in.Ipconfig2, &out.Ipconfig2
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig3 != nil {
		in, out := &in.Ipconfig3, &out.Ipconfig3
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig4 != nil {
		in, out := &in.Ipconfig4, &out.Ipconfig4
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig5 != nil {
		in, out := &in.Ipconfig5, &out.Ipconfig5
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig6 != nil {
		in, out := &in.Ipconfig6, &out.Ipconfig6
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig7 != nil {
		in, out := &in.Ipconfig7, &out.Ipconfig7
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig8 != nil {
		in, out := &in.Ipconfig8, &out.Ipconfig8
		*out = new(string)
		**out = **in
	}
	if in.Ipconfig9 != nil {
		in, out := &in.Ipconfig9, &out.Ipconfig9
		*out = new(string)
		**out = **in
	}
	if in.Iso != nil {
		in, out := &in.Iso, &out.Iso
		*out = new(string)
		**out = **in
	}
	if in.Kvm != nil {
		in, out := &in.Kvm, &out.Kvm
		*out = new(bool)
		**out = **in
	}
	if in.Machine != nil {
		in, out := &in.Machine, &out.Machine
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.Nameserver != nil {
		in, out := &in.Nameserver, &out.Nameserver
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Numa != nil {
		in, out := &in.Numa, &out.Numa
		*out = new(bool)
		**out = **in
	}
	if in.Onboot != nil {
		in, out := &in.Onboot, &out.Onboot
		*out = new(bool)
		**out = **in
	}
	if in.VMstate != nil {
		in, out := &in.VMstate, &out.VMstate
		*out = new(string)
		**out = **in
	}
	if in.OsNetworkConfig != nil {
		in, out := &in.OsNetworkConfig, &out.OsNetworkConfig
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.Pxe != nil {
		in, out := &in.Pxe, &out.Pxe
		*out = new(bool)
		**out = **in
	}
	if in.QemuOs != nil {
		in, out := &in.QemuOs, &out.QemuOs
		*out = new(string)
		**out = **in
	}
	if in.SSHForwardIP != nil {
		in, out := &in.SSHForwardIP, &out.SSHForwardIP
		*out = new(string)
		**out = **in
	}
	if in.SSHPrivateKeySecretRef != nil {
		in, out := &in.SSHPrivateKeySecretRef, &out.SSHPrivateKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.Scsihw != nil {
		in, out := &in.Scsihw, &out.Scsihw
		*out = new(string)
		**out = **in
	}
	if in.Searchdomain != nil {
		in, out := &in.Searchdomain, &out.Searchdomain
		*out = new(string)
		**out = **in
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = make([]SerialParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sockets != nil {
		in, out := &in.Sockets, &out.Sockets
		*out = new(float64)
		**out = **in
	}
	if in.Sshkeys != nil {
		in, out := &in.Sshkeys, &out.Sshkeys
		*out = new(string)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = new(string)
		**out = **in
	}
	if in.Tablet != nil {
		in, out := &in.Tablet, &out.Tablet
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.TargetNode != nil {
		in, out := &in.TargetNode, &out.TargetNode
		*out = new(string)
		**out = **in
	}
	if in.Usb != nil {
		in, out := &in.Usb, &out.Usb
		*out = make([]UsbParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vcpus != nil {
		in, out := &in.Vcpus, &out.Vcpus
		*out = new(float64)
		**out = **in
	}
	if in.Vga != nil {
		in, out := &in.Vga, &out.Vga
		*out = make([]VgaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vmid != nil {
		in, out := &in.Vmid, &out.Vmid
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuParameters.
func (in *QemuParameters) DeepCopy() *QemuParameters {
	if in == nil {
		return nil
	}
	out := new(QemuParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuSpec) DeepCopyInto(out *QemuSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuSpec.
func (in *QemuSpec) DeepCopy() *QemuSpec {
	if in == nil {
		return nil
	}
	out := new(QemuSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuStatus) DeepCopyInto(out *QemuStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuStatus.
func (in *QemuStatus) DeepCopy() *QemuStatus {
	if in == nil {
		return nil
	}
	out := new(QemuStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SerialObservation) DeepCopyInto(out *SerialObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SerialObservation.
func (in *SerialObservation) DeepCopy() *SerialObservation {
	if in == nil {
		return nil
	}
	out := new(SerialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SerialParameters) DeepCopyInto(out *SerialParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SerialParameters.
func (in *SerialParameters) DeepCopy() *SerialParameters {
	if in == nil {
		return nil
	}
	out := new(SerialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnusedDiskObservation) DeepCopyInto(out *UnusedDiskObservation) {
	*out = *in
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.Slot != nil {
		in, out := &in.Slot, &out.Slot
		*out = new(float64)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnusedDiskObservation.
func (in *UnusedDiskObservation) DeepCopy() *UnusedDiskObservation {
	if in == nil {
		return nil
	}
	out := new(UnusedDiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnusedDiskParameters) DeepCopyInto(out *UnusedDiskParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnusedDiskParameters.
func (in *UnusedDiskParameters) DeepCopy() *UnusedDiskParameters {
	if in == nil {
		return nil
	}
	out := new(UnusedDiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsbObservation) DeepCopyInto(out *UsbObservation) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Usb3 != nil {
		in, out := &in.Usb3, &out.Usb3
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsbObservation.
func (in *UsbObservation) DeepCopy() *UsbObservation {
	if in == nil {
		return nil
	}
	out := new(UsbObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsbParameters) DeepCopyInto(out *UsbParameters) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Usb3 != nil {
		in, out := &in.Usb3, &out.Usb3
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsbParameters.
func (in *UsbParameters) DeepCopy() *UsbParameters {
	if in == nil {
		return nil
	}
	out := new(UsbParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VgaObservation) DeepCopyInto(out *VgaObservation) {
	*out = *in
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VgaObservation.
func (in *VgaObservation) DeepCopy() *VgaObservation {
	if in == nil {
		return nil
	}
	out := new(VgaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VgaParameters) DeepCopyInto(out *VgaParameters) {
	*out = *in
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VgaParameters.
func (in *VgaParameters) DeepCopy() *VgaParameters {
	if in == nil {
		return nil
	}
	out := new(VgaParameters)
	in.DeepCopyInto(out)
	return out
}
