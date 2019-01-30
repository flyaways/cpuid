package main

import (
	"fmt"

	prettyjson "github.com/hokaccha/go-prettyjson"
	cpuid "github.com/intel-go/cpuid"
	klauspost "github.com/klauspost/cpuid"
)

func main() {
	cpu := map[string]interface{}{}
	cpu["Vendor"] = cpuid.VendorIdentificatorString
	cpu["ProcessorBrand"] = cpuid.ProcessorBrandString
	cpu["SteppingId"] = cpuid.SteppingId
	cpu["ProcessorType"] = cpuid.ProcessorType
	cpu["DisplayFamily"] = cpuid.DisplayFamily
	cpu["DisplayModel"] = cpuid.DisplayModel
	cpu["CacheLineSize"] = cpuid.CacheLineSize
	cpu["MaxLogocalCPUId"] = cpuid.MaxLogocalCPUId
	cpu["InitialAPICId"] = cpuid.InitialAPICId
	cpu["MonLineSizeMin"] = cpuid.MonLineSizeMin
	cpu["MonLineSizeMax"] = cpuid.MonLineSizeMax
	cpu["PhysicalCores"] = klauspost.CPU.PhysicalCores
	cpu["ThreadsPerCore"] = klauspost.CPU.ThreadsPerCore
	cpu["LogicalCores"] = klauspost.CPU.LogicalCores
	cpu["Cache.L1"] = klauspost.CPU.Cache.L1D
	cpu["Cache.L2"] = klauspost.CPU.Cache.L2
	cpu["Cache.L3"] = klauspost.CPU.Cache.L3
	cpu["SSE"] = klauspost.CPU.SSE()
	cpu["MonitorIBE"] = cpuid.MonitorIBE
	cpu["MonitorEMX"] = cpuid.MonitorEMX
	cpu["EnabledAVX"] = cpuid.EnabledAVX
	cpu["EnabledAVX512"] = cpuid.EnabledAVX512
	cpu["ThermalSensorInterruptThresholds"] = cpuid.ThermalSensorInterruptThresholds

	featureNames := []string{}
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasFeature(1 << i) {
			featureNames = append(featureNames, cpuid.FeatureNames[1<<i])
		}
	}
	cpu["Features"] = featureNames

	extendedFeatures := []string{}
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasExtendedFeature(1 << i) {
			extendedFeatures = append(extendedFeatures, cpuid.ExtendedFeatureNames[1<<i])
		}
	}
	cpu["extendedFeatures"] = extendedFeatures

	extraFeatures := []string{}
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasExtraFeature(1 << i) {
			extraFeatures = append(extraFeatures, cpuid.ExtraFeatureNames[1<<i])
		}
	}
	cpu["extraFeatures"] = extraFeatures

	thermalAndPowerFeatures := []string{}
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasThermalAndPowerFeature(1 << i) {
			if name, found := cpuid.ThermalAndPowerFeatureNames[1<<i]; found {
				thermalAndPowerFeatures = append(thermalAndPowerFeatures, name)
			}
		}
	}
	cpu["thermalAndPowerFeatures"] = thermalAndPowerFeatures

	cpu["cacheDescriptor"] = cpuid.CacheDescriptors

	data, _ := prettyjson.Marshal(&cpu)
	fmt.Println(string(data))
}
