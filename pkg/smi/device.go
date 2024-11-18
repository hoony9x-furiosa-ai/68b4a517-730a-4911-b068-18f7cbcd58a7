package smi

import (
	"runtime"

	"github.com/hoony9x-furiosa-ai/68b4a517-730a-4911-b068-18f7cbcd58a7/pkg/smi/binding"
)

type furiosaSmiObserverInstance = *binding.FuriosaSmiObserver

// ListDevices lists all Furiosa NPU devices in the system.
func ListDevices() ([]Device, error) {
	var outDeviceHandle binding.FuriosaSmiDeviceHandles
	if ret := binding.FuriosaSmiGetDeviceHandles(&outDeviceHandle); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	var outObserverInstance = new(furiosaSmiObserverInstance)
	if ret := binding.FuriosaSmiCreateObserver(outObserverInstance); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	defer runtime.SetFinalizer(outObserverInstance, func(observerInstance *furiosaSmiObserverInstance) {
		_ = binding.FuriosaSmiDestroyObserver(observerInstance)
	})

	var devices []Device
	for i := 0; i < int(outDeviceHandle.Count); i++ {
		devices = append(devices, newDevice(outDeviceHandle.DeviceHandles[i], outObserverInstance))
	}

	return devices, nil
}

// Device represents the abstraction for a single Furiosa NPU device.
type Device interface {
	DeviceInfo() (DeviceInfo, error)
	DeviceFiles() ([]DeviceFile, error)
	CoreStatus() (map[uint32]CoreStatus, error)
	DeviceErrorInfo() (DeviceErrorInfo, error)
	Liveness() (bool, error)
	CoreUtilization() (CoreUtilization, error)
	MemoryUtilization() (MemoryUtilization, error)
	PowerConsumption() (float64, error)
	DeviceTemperature() (DeviceTemperature, error)
	DeviceToDeviceLinkType(target Device) (LinkType, error)
	P2PAccessible(target Device) (bool, error)
	DevicePerformanceCounter() (DevicePerformanceCounter, error)
}

var _ Device = new(device)

type device struct {
	observerInstance *furiosaSmiObserverInstance
	handle           binding.FuriosaSmiDeviceHandle
}

func newDevice(handle binding.FuriosaSmiDeviceHandle, observerInstance *furiosaSmiObserverInstance) Device {
	return &device{
		observerInstance: observerInstance,
		handle:           handle,
	}
}

// DeviceInfo returns `DeviceInfo` which contains information about NPU device. (e.g. arch, serial, ...)
func (d *device) DeviceInfo() (DeviceInfo, error) {
	var out binding.FuriosaSmiDeviceInfo
	if ret := binding.FuriosaSmiGetDeviceInfo(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	return newDeviceInfo(out), nil
}

// DeviceFiles list device files under this device.
func (d *device) DeviceFiles() ([]DeviceFile, error) {
	var out binding.FuriosaSmiDeviceFiles

	if ret := binding.FuriosaSmiGetDeviceFiles(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	var deviceFiles []DeviceFile
	for i := 0; i < int(out.Count); i++ {
		deviceFiles = append(deviceFiles, newDeviceFile(out.DeviceFiles[i]))
	}

	return deviceFiles, nil
}

// CoreStatus examine each core of the device, whether it is occupied or available.
func (d *device) CoreStatus() (map[uint32]CoreStatus, error) {
	var out binding.FuriosaSmiCoreStatuses

	if ret := binding.FuriosaSmiGetDeviceCoreStatus(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	coreStatusMap := make(map[uint32]CoreStatus, out.Count)
	for i := 0; i < int(out.Count); i++ {
		coreStatusMap[uint32(i)] = CoreStatus(out.CoreStatus[i])
	}

	return coreStatusMap, nil
}

// DeviceErrorInfo returns error states of the device.
func (d *device) DeviceErrorInfo() (DeviceErrorInfo, error) {
	var out binding.FuriosaSmiDeviceErrorInfo

	if ret := binding.FuriosaSmiGetDeviceErrorInfo(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	return newDeviceErrorInfo(out), nil
}

// Liveness returns a liveness state of the device.
func (d *device) Liveness() (bool, error) {
	var out bool

	if ret := binding.FuriosaSmiGetDeviceLiveness(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return false, toError(ret)
	}

	return out, nil
}

// CoreUtilization returns a core utilization of the device.
func (d *device) CoreUtilization() (CoreUtilization, error) {
	var out binding.FuriosaSmiCoreUtilization

	if ret := binding.FuriosaSmiGetCoreUtilization(*d.observerInstance, d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	return newCoreUtilization(out), nil
}

// MemoryUtilization returns a memory utilization of the device.
func (d *device) MemoryUtilization() (MemoryUtilization, error) {
	var out binding.FuriosaSmiMemoryUtilization

	if ret := binding.FuriosaSmiGetMemoryUtilization(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	return newMemoryUtilization(out), nil
}

// PowerConsumption returns a power consumption of the device.
func (d *device) PowerConsumption() (float64, error) {
	var out binding.FuriosaSmiDevicePowerConsumption

	if ret := binding.FuriosaSmiGetDevicePowerConsumption(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return 0, toError(ret)
	}

	return out.RmsTotal, nil
}

// DeviceTemperature returns a temperature of the device.
func (d *device) DeviceTemperature() (DeviceTemperature, error) {
	var out binding.FuriosaSmiDeviceTemperature

	if ret := binding.FuriosaSmiGetDeviceTemperature(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	return newDeviceTemperature(out), nil
}

// DeviceToDeviceLinkType returns a device link type between two devices.
func (d *device) DeviceToDeviceLinkType(target Device) (LinkType, error) {
	var linkType binding.FuriosaSmiDeviceToDeviceLinkType

	if ret := binding.FuriosaSmiGetDeviceToDeviceLinkType(d.handle, target.(*device).handle, &linkType); ret != binding.FuriosaSmiReturnCodeOk {
		return LinkTypeUnknown, toError(ret)
	}

	return LinkType(linkType), nil
}

// P2PAccessible returns whether two devices are p2p accessible each other or not.
func (d *device) P2PAccessible(target Device) (bool, error) {
	var out bool

	if ret := binding.FuriosaSmiGetP2pAccessible(d.handle, target.(*device).handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return false, toError(ret)
	}

	return out, nil
}

// DevicePerformanceCounter returns a performance counter of the device.
func (d *device) DevicePerformanceCounter() (DevicePerformanceCounter, error) {
	var out binding.FuriosaSmiDevicePerformanceCounter

	if ret := binding.FuriosaSmiGetDevicePerformanceCounter(d.handle, &out); ret != binding.FuriosaSmiReturnCodeOk {
		return nil, toError(ret)
	}

	return newDevicePerformanceCounter(out), nil
}
