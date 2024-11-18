package smi

import "github.com/hoony9x-furiosa-ai/68b4a517-730a-4911-b068-18f7cbcd58a7/pkg/smi/binding"

// DeviceErrorInfo represents a device error information.
type DeviceErrorInfo interface {
	// AxiPostErrorCount returns an axi post error count.
	AxiPostErrorCount() uint32
	// AxiFetchErrorCount returns an axi fetch error count.
	AxiFetchErrorCount() uint32
	// AxiDiscardErrorCount returns an axi discard error count.
	AxiDiscardErrorCount() uint32
	// AxiDoorbellErrorCount returns an axi doorbell error count.
	AxiDoorbellErrorCount() uint32
	// PciePostErrorCount returns a pcie post error count.
	PciePostErrorCount() uint32
	// PcieFetchErrorCount returns a pcie fetch error count.
	PcieFetchErrorCount() uint32
	// PcieDiscardErrorCount returns a pcie discard error count.
	PcieDiscardErrorCount() uint32
	// PcieDoorbellErrorCount returns a pcie doorbell error count.
	PcieDoorbellErrorCount() uint32
	// DeviceErrorCount returns a device error count.
	DeviceErrorCount() uint32
}

var _ DeviceErrorInfo = new(deviceErrorInfo)

type deviceErrorInfo struct {
	raw binding.FuriosaSmiDeviceErrorInfo
}

func newDeviceErrorInfo(raw binding.FuriosaSmiDeviceErrorInfo) DeviceErrorInfo {
	return &deviceErrorInfo{
		raw: raw,
	}
}

func (d *deviceErrorInfo) AxiPostErrorCount() uint32 {
	return d.raw.AxiPostErrorCount
}

func (d *deviceErrorInfo) AxiFetchErrorCount() uint32 {
	return d.raw.AxiFetchErrorCount
}

func (d *deviceErrorInfo) AxiDiscardErrorCount() uint32 {
	return d.raw.AxiDiscardErrorCount
}

func (d *deviceErrorInfo) AxiDoorbellErrorCount() uint32 {
	return d.raw.AxiDoorbellErrorCount
}

func (d *deviceErrorInfo) PciePostErrorCount() uint32 {
	return d.raw.PciePostErrorCount
}

func (d *deviceErrorInfo) PcieFetchErrorCount() uint32 {
	return d.raw.PcieFetchErrorCount
}

func (d *deviceErrorInfo) PcieDiscardErrorCount() uint32 {
	return d.raw.PcieDiscardErrorCount
}

func (d *deviceErrorInfo) PcieDoorbellErrorCount() uint32 {
	return d.raw.PcieDoorbellErrorCount
}

func (d *deviceErrorInfo) DeviceErrorCount() uint32 {
	return d.raw.DeviceErrorCount
}
