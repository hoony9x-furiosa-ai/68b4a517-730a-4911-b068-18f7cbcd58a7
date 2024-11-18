package smi

import (
	"testing"

	"github.com/hoony9x-furiosa-ai/68b4a517-730a-4911-b068-18f7cbcd58a7/pkg/smi/binding"
	"github.com/stretchr/testify/assert"
)

func testDeviceTemperature(t *testing.T, arch Arch, expected deviceTemperature) {
	mockDevice := GetStaticMockDevice(arch, 0)

	temperature, err := mockDevice.DeviceTemperature()
	assert.NoError(t, err)

	assert.Equal(t, expected.SocPeak(), temperature.SocPeak())
	assert.Equal(t, expected.Ambient(), temperature.Ambient())
}

func TestDeviceTemperature(t *testing.T) {
	tests := []struct {
		description string
		arch        Arch
		expected    deviceTemperature
	}{
		{
			description: "Test Warboy Device Temperature",
			arch:        ArchWarboy,
			expected:    deviceTemperature{binding.FuriosaSmiDeviceTemperature{SocPeak: 20.0, Ambient: 10.0}},
		},
		{
			description: "Test RNGD Device Temperature",
			arch:        ArchRngd,
			expected:    deviceTemperature{binding.FuriosaSmiDeviceTemperature{SocPeak: 20.0, Ambient: 10.0}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			testDeviceTemperature(t, tc.arch, tc.expected)
		})
	}
}

func testPowerConsumption(t *testing.T, arch Arch, expected float64) {
	mockDevice := GetStaticMockDevice(arch, 0)

	power, err := mockDevice.PowerConsumption()
	assert.NoError(t, err)

	assert.Equal(t, expected, power)
}

func TestPowerConsumption(t *testing.T) {
	tests := []struct {
		description string
		arch        Arch
		expected    float64
	}{
		{
			description: "Test Warboy Device Power Consumption",
			arch:        ArchWarboy,
			expected:    100.0,
		},
		{
			description: "Test RNGD Device Power Consumption",
			arch:        ArchRngd,
			expected:    100.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			testPowerConsumption(t, tc.arch, tc.expected)
		})
	}
}

func testCoreUtilization(t *testing.T, arch Arch) {
	mockDevice := GetStaticMockDevice(arch, 0)

	_, err := mockDevice.CoreUtilization() // Currently, not to check the value.
	assert.NoError(t, err)
}

func TestCoreUtilization(t *testing.T) {
	tests := []struct {
		description string
		arch        Arch
	}{
		{
			description: "Test Warboy Device CoreUtilization",
			arch:        ArchWarboy,
		},
		{
			description: "Test RNGD Device CoreUtilization",
			arch:        ArchRngd,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			testCoreUtilization(t, tc.arch)
		})
	}
}
