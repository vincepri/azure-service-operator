// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_VirtualNetwork_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork_STATUS_ARM, VirtualNetwork_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork_STATUS_ARM runs a test to see if a specific instance of VirtualNetwork_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork_STATUS_ARM(subject VirtualNetwork_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetwork_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualNetwork_STATUS_ARMGenerator()
var virtualNetwork_STATUS_ARMGenerator gopter.Gen

// VirtualNetwork_STATUS_ARMGenerator returns a generator of VirtualNetwork_STATUS_ARM instances for property testing.
// We first initialize virtualNetwork_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetwork_STATUS_ARMGenerator() gopter.Gen {
	if virtualNetwork_STATUS_ARMGenerator != nil {
		return virtualNetwork_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS_ARM(generators)
	virtualNetwork_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetwork_STATUS_ARM(generators)
	virtualNetwork_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_STATUS_ARM{}), generators)

	return virtualNetwork_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetwork_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetwork_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(VirtualNetworkPropertiesFormat_STATUS_ARMGenerator())
}

func Test_VirtualNetworkPropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPropertiesFormat_STATUS_ARM, VirtualNetworkPropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPropertiesFormat_STATUS_ARM runs a test to see if a specific instance of VirtualNetworkPropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPropertiesFormat_STATUS_ARM(subject VirtualNetworkPropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPropertiesFormat_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworkPropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualNetworkPropertiesFormat_STATUS_ARMGenerator()
var virtualNetworkPropertiesFormat_STATUS_ARMGenerator gopter.Gen

// VirtualNetworkPropertiesFormat_STATUS_ARMGenerator returns a generator of VirtualNetworkPropertiesFormat_STATUS_ARM instances for property testing.
// We first initialize virtualNetworkPropertiesFormat_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if virtualNetworkPropertiesFormat_STATUS_ARMGenerator != nil {
		return virtualNetworkPropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUS_ARM(generators)
	virtualNetworkPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUS_ARM(generators)
	virtualNetworkPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPropertiesFormat_STATUS_ARM{}), generators)

	return virtualNetworkPropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpace_STATUS_ARMGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUS_ARMGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResource_STATUS_ARMGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptions_STATUS_ARMGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResource_STATUS_ARMGenerator())
}

func Test_DhcpOptions_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptions_STATUS_ARM, DhcpOptions_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptions_STATUS_ARM runs a test to see if a specific instance of DhcpOptions_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptions_STATUS_ARM(subject DhcpOptions_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DhcpOptions_STATUS_ARM instances for property testing - lazily instantiated by
// DhcpOptions_STATUS_ARMGenerator()
var dhcpOptions_STATUS_ARMGenerator gopter.Gen

// DhcpOptions_STATUS_ARMGenerator returns a generator of DhcpOptions_STATUS_ARM instances for property testing.
func DhcpOptions_STATUS_ARMGenerator() gopter.Gen {
	if dhcpOptions_STATUS_ARMGenerator != nil {
		return dhcpOptions_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptions_STATUS_ARM(generators)
	dhcpOptions_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DhcpOptions_STATUS_ARM{}), generators)

	return dhcpOptions_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptions_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptions_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_VirtualNetworkBgpCommunities_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS_ARM, VirtualNetworkBgpCommunities_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS_ARM runs a test to see if a specific instance of VirtualNetworkBgpCommunities_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities_STATUS_ARM(subject VirtualNetworkBgpCommunities_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworkBgpCommunities_STATUS_ARM instances for property testing - lazily instantiated by
// VirtualNetworkBgpCommunities_STATUS_ARMGenerator()
var virtualNetworkBgpCommunities_STATUS_ARMGenerator gopter.Gen

// VirtualNetworkBgpCommunities_STATUS_ARMGenerator returns a generator of VirtualNetworkBgpCommunities_STATUS_ARM instances for property testing.
func VirtualNetworkBgpCommunities_STATUS_ARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunities_STATUS_ARMGenerator != nil {
		return virtualNetworkBgpCommunities_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS_ARM(generators)
	virtualNetworkBgpCommunities_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_STATUS_ARM{}), generators)

	return virtualNetworkBgpCommunities_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}
