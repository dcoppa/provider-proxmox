/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DiskObservation struct {
	Aio *string `json:"aio,omitempty" tf:"aio,omitempty"`

	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	Cache *string `json:"cache,omitempty" tf:"cache,omitempty"`

	Discard *string `json:"discard,omitempty" tf:"discard,omitempty"`

	File *string `json:"file,omitempty" tf:"file,omitempty"`

	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	IopsMax *float64 `json:"iopsMax,omitempty" tf:"iops_max,omitempty"`

	IopsMaxLength *float64 `json:"iopsMaxLength,omitempty" tf:"iops_max_length,omitempty"`

	IopsRd *float64 `json:"iopsRd,omitempty" tf:"iops_rd,omitempty"`

	IopsRdMax *float64 `json:"iopsRdMax,omitempty" tf:"iops_rd_max,omitempty"`

	IopsRdMaxLength *float64 `json:"iopsRdMaxLength,omitempty" tf:"iops_rd_max_length,omitempty"`

	IopsWr *float64 `json:"iopsWr,omitempty" tf:"iops_wr,omitempty"`

	IopsWrMax *float64 `json:"iopsWrMax,omitempty" tf:"iops_wr_max,omitempty"`

	IopsWrMaxLength *float64 `json:"iopsWrMaxLength,omitempty" tf:"iops_wr_max_length,omitempty"`

	Iothread *float64 `json:"iothread,omitempty" tf:"iothread,omitempty"`

	Mbps *float64 `json:"mbps,omitempty" tf:"mbps,omitempty"`

	MbpsRd *float64 `json:"mbpsRd,omitempty" tf:"mbps_rd,omitempty"`

	MbpsRdMax *float64 `json:"mbpsRdMax,omitempty" tf:"mbps_rd_max,omitempty"`

	MbpsWr *float64 `json:"mbpsWr,omitempty" tf:"mbps_wr,omitempty"`

	MbpsWrMax *float64 `json:"mbpsWrMax,omitempty" tf:"mbps_wr_max,omitempty"`

	Media *string `json:"media,omitempty" tf:"media,omitempty"`

	Replicate *float64 `json:"replicate,omitempty" tf:"replicate,omitempty"`

	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	Slot *float64 `json:"slot,omitempty" tf:"slot,omitempty"`

	Ssd *float64 `json:"ssd,omitempty" tf:"ssd,omitempty"`

	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

type DiskParameters struct {

	// +kubebuilder:validation:Optional
	Aio *string `json:"aio,omitempty" tf:"aio,omitempty"`

	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// +kubebuilder:validation:Optional
	Cache *string `json:"cache,omitempty" tf:"cache,omitempty"`

	// +kubebuilder:validation:Optional
	Discard *string `json:"discard,omitempty" tf:"discard,omitempty"`

	// +kubebuilder:validation:Optional
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	IopsMax *float64 `json:"iopsMax,omitempty" tf:"iops_max,omitempty"`

	// +kubebuilder:validation:Optional
	IopsMaxLength *float64 `json:"iopsMaxLength,omitempty" tf:"iops_max_length,omitempty"`

	// +kubebuilder:validation:Optional
	IopsRd *float64 `json:"iopsRd,omitempty" tf:"iops_rd,omitempty"`

	// +kubebuilder:validation:Optional
	IopsRdMax *float64 `json:"iopsRdMax,omitempty" tf:"iops_rd_max,omitempty"`

	// +kubebuilder:validation:Optional
	IopsRdMaxLength *float64 `json:"iopsRdMaxLength,omitempty" tf:"iops_rd_max_length,omitempty"`

	// +kubebuilder:validation:Optional
	IopsWr *float64 `json:"iopsWr,omitempty" tf:"iops_wr,omitempty"`

	// +kubebuilder:validation:Optional
	IopsWrMax *float64 `json:"iopsWrMax,omitempty" tf:"iops_wr_max,omitempty"`

	// +kubebuilder:validation:Optional
	IopsWrMaxLength *float64 `json:"iopsWrMaxLength,omitempty" tf:"iops_wr_max_length,omitempty"`

	// +kubebuilder:validation:Optional
	Iothread *float64 `json:"iothread,omitempty" tf:"iothread,omitempty"`

	// +kubebuilder:validation:Optional
	Mbps *float64 `json:"mbps,omitempty" tf:"mbps,omitempty"`

	// +kubebuilder:validation:Optional
	MbpsRd *float64 `json:"mbpsRd,omitempty" tf:"mbps_rd,omitempty"`

	// +kubebuilder:validation:Optional
	MbpsRdMax *float64 `json:"mbpsRdMax,omitempty" tf:"mbps_rd_max,omitempty"`

	// +kubebuilder:validation:Optional
	MbpsWr *float64 `json:"mbpsWr,omitempty" tf:"mbps_wr,omitempty"`

	// +kubebuilder:validation:Optional
	MbpsWrMax *float64 `json:"mbpsWrMax,omitempty" tf:"mbps_wr_max,omitempty"`

	// +kubebuilder:validation:Optional
	Media *string `json:"media,omitempty" tf:"media,omitempty"`

	// +kubebuilder:validation:Optional
	Replicate *float64 `json:"replicate,omitempty" tf:"replicate,omitempty"`

	// +kubebuilder:validation:Required
	Size *string `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Slot *float64 `json:"slot,omitempty" tf:"slot,omitempty"`

	// +kubebuilder:validation:Optional
	Ssd *float64 `json:"ssd,omitempty" tf:"ssd,omitempty"`

	// +kubebuilder:validation:Required
	Storage *string `json:"storage" tf:"storage,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

type HostpciObservation struct {
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	Pcie *float64 `json:"pcie,omitempty" tf:"pcie,omitempty"`

	Rombar *float64 `json:"rombar,omitempty" tf:"rombar,omitempty"`
}

type HostpciParameters struct {

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Pcie *float64 `json:"pcie,omitempty" tf:"pcie,omitempty"`

	// +kubebuilder:validation:Optional
	Rombar *float64 `json:"rombar,omitempty" tf:"rombar,omitempty"`
}

type NetworkObservation struct {
	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	Firewall *bool `json:"firewall,omitempty" tf:"firewall,omitempty"`

	LinkDown *bool `json:"linkDown,omitempty" tf:"link_down,omitempty"`

	Macaddr *string `json:"macaddr,omitempty" tf:"macaddr,omitempty"`

	Model *string `json:"model,omitempty" tf:"model,omitempty"`

	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	Queues *float64 `json:"queues,omitempty" tf:"queues,omitempty"`

	Rate *float64 `json:"rate,omitempty" tf:"rate,omitempty"`

	// VLAN tag.
	Tag *float64 `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NetworkParameters struct {

	// +kubebuilder:validation:Optional
	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	// +kubebuilder:validation:Optional
	Firewall *bool `json:"firewall,omitempty" tf:"firewall,omitempty"`

	// +kubebuilder:validation:Optional
	LinkDown *bool `json:"linkDown,omitempty" tf:"link_down,omitempty"`

	// +kubebuilder:validation:Optional
	Macaddr *string `json:"macaddr,omitempty" tf:"macaddr,omitempty"`

	// +kubebuilder:validation:Required
	Model *string `json:"model" tf:"model,omitempty"`

	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// +kubebuilder:validation:Optional
	Queues *float64 `json:"queues,omitempty" tf:"queues,omitempty"`

	// +kubebuilder:validation:Optional
	Rate *float64 `json:"rate,omitempty" tf:"rate,omitempty"`

	// VLAN tag.
	// +kubebuilder:validation:Optional
	Tag *float64 `json:"tag,omitempty" tf:"tag,omitempty"`
}

type QemuObservation struct {

	// (do not use, api should manage timeouts)
	// value in second to wait after some operations, useful if system is not fast or during i/o intensive parallel Upbound official provider tasks
	AdditionalWait *float64 `json:"additionalWait,omitempty" tf:"additional_wait,omitempty"`

	Agent *float64 `json:"agent,omitempty" tf:"agent,omitempty"`

	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// : a valid boot order must be specified with Network type included (eg order=net0;scsi0)
	// Automatically reboot the VM if any of the modified parameters requires a reboot to take effect.
	AutomaticReboot *bool `json:"automaticReboot,omitempty" tf:"automatic_reboot,omitempty"`

	Balloon *float64 `json:"balloon,omitempty" tf:"balloon,omitempty"`

	// The VM bios, it can be seabios or ovmf
	Bios *string `json:"bios,omitempty" tf:"bios,omitempty"`

	// : a valid boot order must be specified with Network type included (eg order=net0;scsi0)
	// Boot order of the VM
	Boot *string `json:"boot,omitempty" tf:"boot,omitempty"`

	Bootdisk *string `json:"bootdisk,omitempty" tf:"bootdisk,omitempty"`

	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	CiWait *float64 `json:"ciWait,omitempty" tf:"ci_wait,omitempty"`

	Cicustom *string `json:"cicustom,omitempty" tf:"cicustom,omitempty"`

	Ciuser *string `json:"ciuser,omitempty" tf:"ciuser,omitempty"`

	Clone *string `json:"clone,omitempty" tf:"clone,omitempty"`

	// (do not use, api should manage timeouts)
	// value in second to wait after a vm has been cloned, useful if system is not fast or during i/o intensive parallel Upbound official provider tasks
	CloneWait *float64 `json:"cloneWait,omitempty" tf:"clone_wait,omitempty"`

	CloudinitCdromStorage *string `json:"cloudinitCdromStorage,omitempty" tf:"cloudinit_cdrom_storage,omitempty"`

	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// Use to track vm ipv4 address
	DefaultIPv4Address *string `json:"defaultIpv4Address,omitempty" tf:"default_ipv4_address,omitempty"`

	// By default define SSH for provisioner info
	DefineConnectionInfo *bool `json:"defineConnectionInfo,omitempty" tf:"define_connection_info,omitempty"`

	// The VM description
	Desc *string `json:"desc,omitempty" tf:"desc,omitempty"`

	Disk []DiskObservation `json:"disk,omitempty" tf:"disk,omitempty"`

	DiskGb *float64 `json:"diskGb,omitempty" tf:"disk_gb,omitempty"`

	ForceCreate *bool `json:"forceCreate,omitempty" tf:"force_create,omitempty"`

	ForceRecreateOnChangeOf *string `json:"forceRecreateOnChangeOf,omitempty" tf:"force_recreate_on_change_of,omitempty"`

	FullClone *bool `json:"fullClone,omitempty" tf:"full_clone,omitempty"`

	GuestAgentReadyTimeout *float64 `json:"guestAgentReadyTimeout,omitempty" tf:"guest_agent_ready_timeout,omitempty"`

	Hagroup *string `json:"hagroup,omitempty" tf:"hagroup,omitempty"`

	Hastate *string `json:"hastate,omitempty" tf:"hastate,omitempty"`

	Hostpci []HostpciObservation `json:"hostpci,omitempty" tf:"hostpci,omitempty"`

	Hotplug *string `json:"hotplug,omitempty" tf:"hotplug,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Ipconfig0 *string `json:"ipconfig0,omitempty" tf:"ipconfig0,omitempty"`

	Ipconfig1 *string `json:"ipconfig1,omitempty" tf:"ipconfig1,omitempty"`

	Ipconfig10 *string `json:"ipconfig10,omitempty" tf:"ipconfig10,omitempty"`

	Ipconfig11 *string `json:"ipconfig11,omitempty" tf:"ipconfig11,omitempty"`

	Ipconfig12 *string `json:"ipconfig12,omitempty" tf:"ipconfig12,omitempty"`

	Ipconfig13 *string `json:"ipconfig13,omitempty" tf:"ipconfig13,omitempty"`

	Ipconfig14 *string `json:"ipconfig14,omitempty" tf:"ipconfig14,omitempty"`

	Ipconfig15 *string `json:"ipconfig15,omitempty" tf:"ipconfig15,omitempty"`

	Ipconfig2 *string `json:"ipconfig2,omitempty" tf:"ipconfig2,omitempty"`

	Ipconfig3 *string `json:"ipconfig3,omitempty" tf:"ipconfig3,omitempty"`

	Ipconfig4 *string `json:"ipconfig4,omitempty" tf:"ipconfig4,omitempty"`

	Ipconfig5 *string `json:"ipconfig5,omitempty" tf:"ipconfig5,omitempty"`

	Ipconfig6 *string `json:"ipconfig6,omitempty" tf:"ipconfig6,omitempty"`

	Ipconfig7 *string `json:"ipconfig7,omitempty" tf:"ipconfig7,omitempty"`

	Ipconfig8 *string `json:"ipconfig8,omitempty" tf:"ipconfig8,omitempty"`

	Ipconfig9 *string `json:"ipconfig9,omitempty" tf:"ipconfig9,omitempty"`

	Iso *string `json:"iso,omitempty" tf:"iso,omitempty"`

	Kvm *bool `json:"kvm,omitempty" tf:"kvm,omitempty"`

	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// Specifies the Qemu machine type.
	Machine *string `json:"machine,omitempty" tf:"machine,omitempty"`

	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	Nameserver *string `json:"nameserver,omitempty" tf:"nameserver,omitempty"`

	Network []NetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	Nic *string `json:"nic,omitempty" tf:"nic,omitempty"`

	Numa *bool `json:"numa,omitempty" tf:"numa,omitempty"`

	// : a valid boot order must be specified with Network type included (eg order=net0;scsi0)
	// VM autostart on boot
	Onboot *bool `json:"onboot,omitempty" tf:"onboot,omitempty"`

	// VM autostart on create
	Oncreate *bool `json:"oncreate,omitempty" tf:"oncreate,omitempty"`

	OsNetworkConfig *string `json:"osNetworkConfig,omitempty" tf:"os_network_config,omitempty"`

	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	Pool *string `json:"pool,omitempty" tf:"pool,omitempty"`

	// (do not use, provider do not fully support preprovisioning anymore)
	Preprovision *bool `json:"preprovision,omitempty" tf:"preprovision,omitempty"`

	Pxe *bool `json:"pxe,omitempty" tf:"pxe,omitempty"`

	QemuOs *string `json:"qemuOs,omitempty" tf:"qemu_os,omitempty"`

	// Internal variable, true if any of the modified parameters requires a reboot to take effect.
	RebootRequired *bool `json:"rebootRequired,omitempty" tf:"reboot_required,omitempty"`

	// Use to pass instance ip address, redundant
	SSHForwardIP *string `json:"sshForwardIp,omitempty" tf:"ssh_forward_ip,omitempty"`

	SSHHost *string `json:"sshHost,omitempty" tf:"ssh_host,omitempty"`

	SSHPort *string `json:"sshPort,omitempty" tf:"ssh_port,omitempty"`

	SSHUser *string `json:"sshUser,omitempty" tf:"ssh_user,omitempty"`

	Scsihw *string `json:"scsihw,omitempty" tf:"scsihw,omitempty"`

	Searchdomain *string `json:"searchdomain,omitempty" tf:"searchdomain,omitempty"`

	Serial []SerialObservation `json:"serial,omitempty" tf:"serial,omitempty"`

	Sockets *float64 `json:"sockets,omitempty" tf:"sockets,omitempty"`

	Sshkeys *string `json:"sshkeys,omitempty" tf:"sshkeys,omitempty"`

	// Startup order of the VM
	Startup *string `json:"startup,omitempty" tf:"startup,omitempty"`

	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Enable tablet mode in the VM
	Tablet *bool `json:"tablet,omitempty" tf:"tablet,omitempty"`

	Tags *string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The node where VM goes to
	TargetNode *string `json:"targetNode,omitempty" tf:"target_node,omitempty"`

	// Record unused disks in proxmox. This is intended to be read-only for now.
	UnusedDisk []UnusedDiskObservation `json:"unusedDisk,omitempty" tf:"unused_disk,omitempty"`

	Usb []UsbObservation `json:"usb,omitempty" tf:"usb,omitempty"`

	Vcpus *float64 `json:"vcpus,omitempty" tf:"vcpus,omitempty"`

	Vga []VgaObservation `json:"vga,omitempty" tf:"vga,omitempty"`

	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`

	// The VM identifier in proxmox (100-999999999)
	Vmid *float64 `json:"vmid,omitempty" tf:"vmid,omitempty"`
}

type QemuParameters struct {

	// (do not use, api should manage timeouts)
	// value in second to wait after some operations, useful if system is not fast or during i/o intensive parallel Upbound official provider tasks
	// +kubebuilder:validation:Optional
	AdditionalWait *float64 `json:"additionalWait,omitempty" tf:"additional_wait,omitempty"`

	// +kubebuilder:validation:Optional
	Agent *float64 `json:"agent,omitempty" tf:"agent,omitempty"`

	// +kubebuilder:validation:Optional
	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// : a valid boot order must be specified with Network type included (eg order=net0;scsi0)
	// Automatically reboot the VM if any of the modified parameters requires a reboot to take effect.
	// +kubebuilder:validation:Optional
	AutomaticReboot *bool `json:"automaticReboot,omitempty" tf:"automatic_reboot,omitempty"`

	// +kubebuilder:validation:Optional
	Balloon *float64 `json:"balloon,omitempty" tf:"balloon,omitempty"`

	// The VM bios, it can be seabios or ovmf
	// +kubebuilder:validation:Optional
	Bios *string `json:"bios,omitempty" tf:"bios,omitempty"`

	// : a valid boot order must be specified with Network type included (eg order=net0;scsi0)
	// Boot order of the VM
	// +kubebuilder:validation:Optional
	Boot *string `json:"boot,omitempty" tf:"boot,omitempty"`

	// +kubebuilder:validation:Optional
	Bootdisk *string `json:"bootdisk,omitempty" tf:"bootdisk,omitempty"`

	// +kubebuilder:validation:Optional
	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	CiWait *float64 `json:"ciWait,omitempty" tf:"ci_wait,omitempty"`

	// +kubebuilder:validation:Optional
	Cicustom *string `json:"cicustom,omitempty" tf:"cicustom,omitempty"`

	// +kubebuilder:validation:Optional
	CipasswordSecretRef *v1.SecretKeySelector `json:"cipasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Ciuser *string `json:"ciuser,omitempty" tf:"ciuser,omitempty"`

	// +kubebuilder:validation:Optional
	Clone *string `json:"clone,omitempty" tf:"clone,omitempty"`

	// (do not use, api should manage timeouts)
	// value in second to wait after a vm has been cloned, useful if system is not fast or during i/o intensive parallel Upbound official provider tasks
	// +kubebuilder:validation:Optional
	CloneWait *float64 `json:"cloneWait,omitempty" tf:"clone_wait,omitempty"`

	// +kubebuilder:validation:Optional
	CloudinitCdromStorage *string `json:"cloudinitCdromStorage,omitempty" tf:"cloudinit_cdrom_storage,omitempty"`

	// +kubebuilder:validation:Optional
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// By default define SSH for provisioner info
	// +kubebuilder:validation:Optional
	DefineConnectionInfo *bool `json:"defineConnectionInfo,omitempty" tf:"define_connection_info,omitempty"`

	// The VM description
	// +kubebuilder:validation:Optional
	Desc *string `json:"desc,omitempty" tf:"desc,omitempty"`

	// +kubebuilder:validation:Optional
	Disk []DiskParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// +kubebuilder:validation:Optional
	DiskGb *float64 `json:"diskGb,omitempty" tf:"disk_gb,omitempty"`

	// +kubebuilder:validation:Optional
	ForceCreate *bool `json:"forceCreate,omitempty" tf:"force_create,omitempty"`

	// +kubebuilder:validation:Optional
	ForceRecreateOnChangeOf *string `json:"forceRecreateOnChangeOf,omitempty" tf:"force_recreate_on_change_of,omitempty"`

	// +kubebuilder:validation:Optional
	FullClone *bool `json:"fullClone,omitempty" tf:"full_clone,omitempty"`

	// +kubebuilder:validation:Optional
	GuestAgentReadyTimeout *float64 `json:"guestAgentReadyTimeout,omitempty" tf:"guest_agent_ready_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	Hagroup *string `json:"hagroup,omitempty" tf:"hagroup,omitempty"`

	// +kubebuilder:validation:Optional
	Hastate *string `json:"hastate,omitempty" tf:"hastate,omitempty"`

	// +kubebuilder:validation:Optional
	Hostpci []HostpciParameters `json:"hostpci,omitempty" tf:"hostpci,omitempty"`

	// +kubebuilder:validation:Optional
	Hotplug *string `json:"hotplug,omitempty" tf:"hotplug,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig0 *string `json:"ipconfig0,omitempty" tf:"ipconfig0,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig1 *string `json:"ipconfig1,omitempty" tf:"ipconfig1,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig10 *string `json:"ipconfig10,omitempty" tf:"ipconfig10,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig11 *string `json:"ipconfig11,omitempty" tf:"ipconfig11,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig12 *string `json:"ipconfig12,omitempty" tf:"ipconfig12,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig13 *string `json:"ipconfig13,omitempty" tf:"ipconfig13,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig14 *string `json:"ipconfig14,omitempty" tf:"ipconfig14,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig15 *string `json:"ipconfig15,omitempty" tf:"ipconfig15,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig2 *string `json:"ipconfig2,omitempty" tf:"ipconfig2,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig3 *string `json:"ipconfig3,omitempty" tf:"ipconfig3,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig4 *string `json:"ipconfig4,omitempty" tf:"ipconfig4,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig5 *string `json:"ipconfig5,omitempty" tf:"ipconfig5,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig6 *string `json:"ipconfig6,omitempty" tf:"ipconfig6,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig7 *string `json:"ipconfig7,omitempty" tf:"ipconfig7,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig8 *string `json:"ipconfig8,omitempty" tf:"ipconfig8,omitempty"`

	// +kubebuilder:validation:Optional
	Ipconfig9 *string `json:"ipconfig9,omitempty" tf:"ipconfig9,omitempty"`

	// +kubebuilder:validation:Optional
	Iso *string `json:"iso,omitempty" tf:"iso,omitempty"`

	// +kubebuilder:validation:Optional
	Kvm *bool `json:"kvm,omitempty" tf:"kvm,omitempty"`

	// +kubebuilder:validation:Optional
	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// Specifies the Qemu machine type.
	// +kubebuilder:validation:Optional
	Machine *string `json:"machine,omitempty" tf:"machine,omitempty"`

	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// +kubebuilder:validation:Optional
	Nameserver *string `json:"nameserver,omitempty" tf:"nameserver,omitempty"`

	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	Nic *string `json:"nic,omitempty" tf:"nic,omitempty"`

	// +kubebuilder:validation:Optional
	Numa *bool `json:"numa,omitempty" tf:"numa,omitempty"`

	// : a valid boot order must be specified with Network type included (eg order=net0;scsi0)
	// VM autostart on boot
	// +kubebuilder:validation:Optional
	Onboot *bool `json:"onboot,omitempty" tf:"onboot,omitempty"`

	// VM autostart on create
	// +kubebuilder:validation:Optional
	Oncreate *bool `json:"oncreate,omitempty" tf:"oncreate,omitempty"`

	// +kubebuilder:validation:Optional
	OsNetworkConfig *string `json:"osNetworkConfig,omitempty" tf:"os_network_config,omitempty"`

	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// +kubebuilder:validation:Optional
	Pool *string `json:"pool,omitempty" tf:"pool,omitempty"`

	// (do not use, provider do not fully support preprovisioning anymore)
	// +kubebuilder:validation:Optional
	Preprovision *bool `json:"preprovision,omitempty" tf:"preprovision,omitempty"`

	// +kubebuilder:validation:Optional
	Pxe *bool `json:"pxe,omitempty" tf:"pxe,omitempty"`

	// +kubebuilder:validation:Optional
	QemuOs *string `json:"qemuOs,omitempty" tf:"qemu_os,omitempty"`

	// Use to pass instance ip address, redundant
	// +kubebuilder:validation:Optional
	SSHForwardIP *string `json:"sshForwardIp,omitempty" tf:"ssh_forward_ip,omitempty"`

	// +kubebuilder:validation:Optional
	SSHPrivateKeySecretRef *v1.SecretKeySelector `json:"sshPrivateKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSHUser *string `json:"sshUser,omitempty" tf:"ssh_user,omitempty"`

	// +kubebuilder:validation:Optional
	Scsihw *string `json:"scsihw,omitempty" tf:"scsihw,omitempty"`

	// +kubebuilder:validation:Optional
	Searchdomain *string `json:"searchdomain,omitempty" tf:"searchdomain,omitempty"`

	// +kubebuilder:validation:Optional
	Serial []SerialParameters `json:"serial,omitempty" tf:"serial,omitempty"`

	// +kubebuilder:validation:Optional
	Sockets *float64 `json:"sockets,omitempty" tf:"sockets,omitempty"`

	// +kubebuilder:validation:Optional
	Sshkeys *string `json:"sshkeys,omitempty" tf:"sshkeys,omitempty"`

	// Startup order of the VM
	// +kubebuilder:validation:Optional
	Startup *string `json:"startup,omitempty" tf:"startup,omitempty"`

	// +kubebuilder:validation:Optional
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// Enable tablet mode in the VM
	// +kubebuilder:validation:Optional
	Tablet *bool `json:"tablet,omitempty" tf:"tablet,omitempty"`

	// +kubebuilder:validation:Optional
	Tags *string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The node where VM goes to
	// +kubebuilder:validation:Optional
	TargetNode *string `json:"targetNode,omitempty" tf:"target_node,omitempty"`

	// +kubebuilder:validation:Optional
	Usb []UsbParameters `json:"usb,omitempty" tf:"usb,omitempty"`

	// +kubebuilder:validation:Optional
	Vcpus *float64 `json:"vcpus,omitempty" tf:"vcpus,omitempty"`

	// +kubebuilder:validation:Optional
	Vga []VgaParameters `json:"vga,omitempty" tf:"vga,omitempty"`

	// +kubebuilder:validation:Optional
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`

	// The VM identifier in proxmox (100-999999999)
	// +kubebuilder:validation:Optional
	Vmid *float64 `json:"vmid,omitempty" tf:"vmid,omitempty"`
}

type SerialObservation struct {
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SerialParameters struct {

	// +kubebuilder:validation:Required
	ID *float64 `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type UnusedDiskObservation struct {
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	Slot *float64 `json:"slot,omitempty" tf:"slot,omitempty"`

	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`
}

type UnusedDiskParameters struct {
}

type UsbObservation struct {
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	Usb3 *bool `json:"usb3,omitempty" tf:"usb3,omitempty"`
}

type UsbParameters struct {

	// +kubebuilder:validation:Required
	Host *string `json:"host" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Usb3 *bool `json:"usb3,omitempty" tf:"usb3,omitempty"`
}

type VgaObservation struct {
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VgaParameters struct {

	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// QemuSpec defines the desired state of Qemu
type QemuSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QemuParameters `json:"forProvider"`
}

// QemuStatus defines the observed state of Qemu.
type QemuStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QemuObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Qemu is the Schema for the Qemus API. Creates and manages a Proxmox VM Qemu container
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmox}
type Qemu struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.targetNode)",message="targetNode is a required parameter"
	Spec   QemuSpec   `json:"spec"`
	Status QemuStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QemuList contains a list of Qemus
type QemuList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Qemu `json:"items"`
}

// Repository type metadata.
var (
	Qemu_Kind             = "Qemu"
	Qemu_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Qemu_Kind}.String()
	Qemu_KindAPIVersion   = Qemu_Kind + "." + CRDGroupVersion.String()
	Qemu_GroupVersionKind = CRDGroupVersion.WithKind(Qemu_Kind)
)

func init() {
	SchemeBuilder.Register(&Qemu{}, &QemuList{})
}
