// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401storage

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

func Test_TrafficManagerProfilesAzureEndpoint_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TrafficManagerProfilesAzureEndpoint via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficManagerProfilesAzureEndpoint, TrafficManagerProfilesAzureEndpointGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficManagerProfilesAzureEndpoint runs a test to see if a specific instance of TrafficManagerProfilesAzureEndpoint round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficManagerProfilesAzureEndpoint(subject TrafficManagerProfilesAzureEndpoint) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TrafficManagerProfilesAzureEndpoint
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

// Generator of TrafficManagerProfilesAzureEndpoint instances for property testing - lazily instantiated by
// TrafficManagerProfilesAzureEndpointGenerator()
var trafficManagerProfilesAzureEndpointGenerator gopter.Gen

// TrafficManagerProfilesAzureEndpointGenerator returns a generator of TrafficManagerProfilesAzureEndpoint instances for property testing.
func TrafficManagerProfilesAzureEndpointGenerator() gopter.Gen {
	if trafficManagerProfilesAzureEndpointGenerator != nil {
		return trafficManagerProfilesAzureEndpointGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTrafficManagerProfilesAzureEndpoint(generators)
	trafficManagerProfilesAzureEndpointGenerator = gen.Struct(reflect.TypeOf(TrafficManagerProfilesAzureEndpoint{}), generators)

	return trafficManagerProfilesAzureEndpointGenerator
}

// AddRelatedPropertyGeneratorsForTrafficManagerProfilesAzureEndpoint is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficManagerProfilesAzureEndpoint(gens map[string]gopter.Gen) {
	gens["Spec"] = Trafficmanagerprofiles_AzureEndpoint_SpecGenerator()
	gens["Status"] = Trafficmanagerprofiles_AzureEndpoint_STATUSGenerator()
}

func Test_Trafficmanagerprofiles_AzureEndpoint_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Trafficmanagerprofiles_AzureEndpoint_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficmanagerprofiles_AzureEndpoint_Spec, Trafficmanagerprofiles_AzureEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficmanagerprofiles_AzureEndpoint_Spec runs a test to see if a specific instance of Trafficmanagerprofiles_AzureEndpoint_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficmanagerprofiles_AzureEndpoint_Spec(subject Trafficmanagerprofiles_AzureEndpoint_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Trafficmanagerprofiles_AzureEndpoint_Spec
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

// Generator of Trafficmanagerprofiles_AzureEndpoint_Spec instances for property testing - lazily instantiated by
// Trafficmanagerprofiles_AzureEndpoint_SpecGenerator()
var trafficmanagerprofiles_AzureEndpoint_SpecGenerator gopter.Gen

// Trafficmanagerprofiles_AzureEndpoint_SpecGenerator returns a generator of Trafficmanagerprofiles_AzureEndpoint_Spec instances for property testing.
// We first initialize trafficmanagerprofiles_AzureEndpoint_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Trafficmanagerprofiles_AzureEndpoint_SpecGenerator() gopter.Gen {
	if trafficmanagerprofiles_AzureEndpoint_SpecGenerator != nil {
		return trafficmanagerprofiles_AzureEndpoint_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_Spec(generators)
	trafficmanagerprofiles_AzureEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_AzureEndpoint_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_Spec(generators)
	AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_Spec(generators)
	trafficmanagerprofiles_AzureEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_AzureEndpoint_Spec{}), generators)

	return trafficmanagerprofiles_AzureEndpoint_SpecGenerator
}

// AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["AlwaysServe"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["EndpointLocation"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointMonitorStatus"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointStatus"] = gen.PtrOf(gen.AlphaString())
	gens["GeoMapping"] = gen.SliceOf(gen.AlphaString())
	gens["MinChildEndpoints"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv4"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv6"] = gen.PtrOf(gen.Int())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["CustomHeaders"] = gen.SliceOf(EndpointProperties_CustomHeadersGenerator())
	gens["Subnets"] = gen.SliceOf(EndpointProperties_SubnetsGenerator())
}

func Test_Trafficmanagerprofiles_AzureEndpoint_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Trafficmanagerprofiles_AzureEndpoint_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficmanagerprofiles_AzureEndpoint_STATUS, Trafficmanagerprofiles_AzureEndpoint_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficmanagerprofiles_AzureEndpoint_STATUS runs a test to see if a specific instance of Trafficmanagerprofiles_AzureEndpoint_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficmanagerprofiles_AzureEndpoint_STATUS(subject Trafficmanagerprofiles_AzureEndpoint_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Trafficmanagerprofiles_AzureEndpoint_STATUS
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

// Generator of Trafficmanagerprofiles_AzureEndpoint_STATUS instances for property testing - lazily instantiated by
// Trafficmanagerprofiles_AzureEndpoint_STATUSGenerator()
var trafficmanagerprofiles_AzureEndpoint_STATUSGenerator gopter.Gen

// Trafficmanagerprofiles_AzureEndpoint_STATUSGenerator returns a generator of Trafficmanagerprofiles_AzureEndpoint_STATUS instances for property testing.
// We first initialize trafficmanagerprofiles_AzureEndpoint_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Trafficmanagerprofiles_AzureEndpoint_STATUSGenerator() gopter.Gen {
	if trafficmanagerprofiles_AzureEndpoint_STATUSGenerator != nil {
		return trafficmanagerprofiles_AzureEndpoint_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_STATUS(generators)
	trafficmanagerprofiles_AzureEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_AzureEndpoint_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_STATUS(generators)
	AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_STATUS(generators)
	trafficmanagerprofiles_AzureEndpoint_STATUSGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofiles_AzureEndpoint_STATUS{}), generators)

	return trafficmanagerprofiles_AzureEndpoint_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["AlwaysServe"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointLocation"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointMonitorStatus"] = gen.PtrOf(gen.AlphaString())
	gens["EndpointStatus"] = gen.PtrOf(gen.AlphaString())
	gens["GeoMapping"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["MinChildEndpoints"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv4"] = gen.PtrOf(gen.Int())
	gens["MinChildEndpointsIPv6"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["TargetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficmanagerprofiles_AzureEndpoint_STATUS(gens map[string]gopter.Gen) {
	gens["CustomHeaders"] = gen.SliceOf(EndpointProperties_CustomHeaders_STATUSGenerator())
	gens["Subnets"] = gen.SliceOf(EndpointProperties_Subnets_STATUSGenerator())
}

func Test_EndpointProperties_CustomHeaders_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EndpointProperties_CustomHeaders via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEndpointProperties_CustomHeaders, EndpointProperties_CustomHeadersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEndpointProperties_CustomHeaders runs a test to see if a specific instance of EndpointProperties_CustomHeaders round trips to JSON and back losslessly
func RunJSONSerializationTestForEndpointProperties_CustomHeaders(subject EndpointProperties_CustomHeaders) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EndpointProperties_CustomHeaders
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

// Generator of EndpointProperties_CustomHeaders instances for property testing - lazily instantiated by
// EndpointProperties_CustomHeadersGenerator()
var endpointProperties_CustomHeadersGenerator gopter.Gen

// EndpointProperties_CustomHeadersGenerator returns a generator of EndpointProperties_CustomHeaders instances for property testing.
func EndpointProperties_CustomHeadersGenerator() gopter.Gen {
	if endpointProperties_CustomHeadersGenerator != nil {
		return endpointProperties_CustomHeadersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEndpointProperties_CustomHeaders(generators)
	endpointProperties_CustomHeadersGenerator = gen.Struct(reflect.TypeOf(EndpointProperties_CustomHeaders{}), generators)

	return endpointProperties_CustomHeadersGenerator
}

// AddIndependentPropertyGeneratorsForEndpointProperties_CustomHeaders is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEndpointProperties_CustomHeaders(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_EndpointProperties_CustomHeaders_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EndpointProperties_CustomHeaders_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEndpointProperties_CustomHeaders_STATUS, EndpointProperties_CustomHeaders_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEndpointProperties_CustomHeaders_STATUS runs a test to see if a specific instance of EndpointProperties_CustomHeaders_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEndpointProperties_CustomHeaders_STATUS(subject EndpointProperties_CustomHeaders_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EndpointProperties_CustomHeaders_STATUS
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

// Generator of EndpointProperties_CustomHeaders_STATUS instances for property testing - lazily instantiated by
// EndpointProperties_CustomHeaders_STATUSGenerator()
var endpointProperties_CustomHeaders_STATUSGenerator gopter.Gen

// EndpointProperties_CustomHeaders_STATUSGenerator returns a generator of EndpointProperties_CustomHeaders_STATUS instances for property testing.
func EndpointProperties_CustomHeaders_STATUSGenerator() gopter.Gen {
	if endpointProperties_CustomHeaders_STATUSGenerator != nil {
		return endpointProperties_CustomHeaders_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEndpointProperties_CustomHeaders_STATUS(generators)
	endpointProperties_CustomHeaders_STATUSGenerator = gen.Struct(reflect.TypeOf(EndpointProperties_CustomHeaders_STATUS{}), generators)

	return endpointProperties_CustomHeaders_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEndpointProperties_CustomHeaders_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEndpointProperties_CustomHeaders_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_EndpointProperties_Subnets_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EndpointProperties_Subnets via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEndpointProperties_Subnets, EndpointProperties_SubnetsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEndpointProperties_Subnets runs a test to see if a specific instance of EndpointProperties_Subnets round trips to JSON and back losslessly
func RunJSONSerializationTestForEndpointProperties_Subnets(subject EndpointProperties_Subnets) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EndpointProperties_Subnets
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

// Generator of EndpointProperties_Subnets instances for property testing - lazily instantiated by
// EndpointProperties_SubnetsGenerator()
var endpointProperties_SubnetsGenerator gopter.Gen

// EndpointProperties_SubnetsGenerator returns a generator of EndpointProperties_Subnets instances for property testing.
func EndpointProperties_SubnetsGenerator() gopter.Gen {
	if endpointProperties_SubnetsGenerator != nil {
		return endpointProperties_SubnetsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEndpointProperties_Subnets(generators)
	endpointProperties_SubnetsGenerator = gen.Struct(reflect.TypeOf(EndpointProperties_Subnets{}), generators)

	return endpointProperties_SubnetsGenerator
}

// AddIndependentPropertyGeneratorsForEndpointProperties_Subnets is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEndpointProperties_Subnets(gens map[string]gopter.Gen) {
	gens["First"] = gen.PtrOf(gen.AlphaString())
	gens["Last"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.Int())
}

func Test_EndpointProperties_Subnets_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EndpointProperties_Subnets_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEndpointProperties_Subnets_STATUS, EndpointProperties_Subnets_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEndpointProperties_Subnets_STATUS runs a test to see if a specific instance of EndpointProperties_Subnets_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForEndpointProperties_Subnets_STATUS(subject EndpointProperties_Subnets_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EndpointProperties_Subnets_STATUS
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

// Generator of EndpointProperties_Subnets_STATUS instances for property testing - lazily instantiated by
// EndpointProperties_Subnets_STATUSGenerator()
var endpointProperties_Subnets_STATUSGenerator gopter.Gen

// EndpointProperties_Subnets_STATUSGenerator returns a generator of EndpointProperties_Subnets_STATUS instances for property testing.
func EndpointProperties_Subnets_STATUSGenerator() gopter.Gen {
	if endpointProperties_Subnets_STATUSGenerator != nil {
		return endpointProperties_Subnets_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEndpointProperties_Subnets_STATUS(generators)
	endpointProperties_Subnets_STATUSGenerator = gen.Struct(reflect.TypeOf(EndpointProperties_Subnets_STATUS{}), generators)

	return endpointProperties_Subnets_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForEndpointProperties_Subnets_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEndpointProperties_Subnets_STATUS(gens map[string]gopter.Gen) {
	gens["First"] = gen.PtrOf(gen.AlphaString())
	gens["Last"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.Int())
}
