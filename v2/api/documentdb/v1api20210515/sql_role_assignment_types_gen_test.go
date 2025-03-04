// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import (
	"encoding/json"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515storage"
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

func Test_SqlRoleAssignment_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlRoleAssignment to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlRoleAssignment, SqlRoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlRoleAssignment tests if a specific instance of SqlRoleAssignment round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlRoleAssignment(subject SqlRoleAssignment) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210515s.SqlRoleAssignment
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlRoleAssignment
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlRoleAssignment_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlRoleAssignment to SqlRoleAssignment via AssignProperties_To_SqlRoleAssignment & AssignProperties_From_SqlRoleAssignment returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlRoleAssignment, SqlRoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlRoleAssignment tests if a specific instance of SqlRoleAssignment can be assigned to v1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlRoleAssignment(subject SqlRoleAssignment) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlRoleAssignment
	err := copied.AssignProperties_To_SqlRoleAssignment(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlRoleAssignment
	err = actual.AssignProperties_From_SqlRoleAssignment(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlRoleAssignment_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlRoleAssignment via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlRoleAssignment, SqlRoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlRoleAssignment runs a test to see if a specific instance of SqlRoleAssignment round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlRoleAssignment(subject SqlRoleAssignment) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlRoleAssignment
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

// Generator of SqlRoleAssignment instances for property testing - lazily instantiated by SqlRoleAssignmentGenerator()
var sqlRoleAssignmentGenerator gopter.Gen

// SqlRoleAssignmentGenerator returns a generator of SqlRoleAssignment instances for property testing.
func SqlRoleAssignmentGenerator() gopter.Gen {
	if sqlRoleAssignmentGenerator != nil {
		return sqlRoleAssignmentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlRoleAssignment(generators)
	sqlRoleAssignmentGenerator = gen.Struct(reflect.TypeOf(SqlRoleAssignment{}), generators)

	return sqlRoleAssignmentGenerator
}

// AddRelatedPropertyGeneratorsForSqlRoleAssignment is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlRoleAssignment(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlRoleAssignment_SpecGenerator()
	gens["Status"] = DatabaseAccounts_SqlRoleAssignment_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlRoleAssignment_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_SqlRoleAssignment_Spec to DatabaseAccounts_SqlRoleAssignment_Spec via AssignProperties_To_DatabaseAccounts_SqlRoleAssignment_Spec & AssignProperties_From_DatabaseAccounts_SqlRoleAssignment_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_SqlRoleAssignment_Spec, DatabaseAccounts_SqlRoleAssignment_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_SqlRoleAssignment_Spec tests if a specific instance of DatabaseAccounts_SqlRoleAssignment_Spec can be assigned to v1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_SqlRoleAssignment_Spec(subject DatabaseAccounts_SqlRoleAssignment_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccounts_SqlRoleAssignment_Spec
	err := copied.AssignProperties_To_DatabaseAccounts_SqlRoleAssignment_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_SqlRoleAssignment_Spec
	err = actual.AssignProperties_From_DatabaseAccounts_SqlRoleAssignment_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccounts_SqlRoleAssignment_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlRoleAssignment_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlRoleAssignment_Spec, DatabaseAccounts_SqlRoleAssignment_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlRoleAssignment_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlRoleAssignment_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlRoleAssignment_Spec(subject DatabaseAccounts_SqlRoleAssignment_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlRoleAssignment_Spec
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

// Generator of DatabaseAccounts_SqlRoleAssignment_Spec instances for property testing - lazily instantiated by
// DatabaseAccounts_SqlRoleAssignment_SpecGenerator()
var databaseAccounts_SqlRoleAssignment_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlRoleAssignment_SpecGenerator returns a generator of DatabaseAccounts_SqlRoleAssignment_Spec instances for property testing.
func DatabaseAccounts_SqlRoleAssignment_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlRoleAssignment_SpecGenerator != nil {
		return databaseAccounts_SqlRoleAssignment_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlRoleAssignment_Spec(generators)
	databaseAccounts_SqlRoleAssignment_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlRoleAssignment_Spec{}), generators)

	return databaseAccounts_SqlRoleAssignment_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlRoleAssignment_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlRoleAssignment_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
}

func Test_DatabaseAccounts_SqlRoleAssignment_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccounts_SqlRoleAssignment_STATUS to DatabaseAccounts_SqlRoleAssignment_STATUS via AssignProperties_To_DatabaseAccounts_SqlRoleAssignment_STATUS & AssignProperties_From_DatabaseAccounts_SqlRoleAssignment_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccounts_SqlRoleAssignment_STATUS, DatabaseAccounts_SqlRoleAssignment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccounts_SqlRoleAssignment_STATUS tests if a specific instance of DatabaseAccounts_SqlRoleAssignment_STATUS can be assigned to v1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccounts_SqlRoleAssignment_STATUS(subject DatabaseAccounts_SqlRoleAssignment_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.DatabaseAccounts_SqlRoleAssignment_STATUS
	err := copied.AssignProperties_To_DatabaseAccounts_SqlRoleAssignment_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccounts_SqlRoleAssignment_STATUS
	err = actual.AssignProperties_From_DatabaseAccounts_SqlRoleAssignment_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccounts_SqlRoleAssignment_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlRoleAssignment_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlRoleAssignment_STATUS, DatabaseAccounts_SqlRoleAssignment_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlRoleAssignment_STATUS runs a test to see if a specific instance of DatabaseAccounts_SqlRoleAssignment_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlRoleAssignment_STATUS(subject DatabaseAccounts_SqlRoleAssignment_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlRoleAssignment_STATUS
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

// Generator of DatabaseAccounts_SqlRoleAssignment_STATUS instances for property testing - lazily instantiated by
// DatabaseAccounts_SqlRoleAssignment_STATUSGenerator()
var databaseAccounts_SqlRoleAssignment_STATUSGenerator gopter.Gen

// DatabaseAccounts_SqlRoleAssignment_STATUSGenerator returns a generator of DatabaseAccounts_SqlRoleAssignment_STATUS instances for property testing.
func DatabaseAccounts_SqlRoleAssignment_STATUSGenerator() gopter.Gen {
	if databaseAccounts_SqlRoleAssignment_STATUSGenerator != nil {
		return databaseAccounts_SqlRoleAssignment_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlRoleAssignment_STATUS(generators)
	databaseAccounts_SqlRoleAssignment_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlRoleAssignment_STATUS{}), generators)

	return databaseAccounts_SqlRoleAssignment_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlRoleAssignment_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlRoleAssignment_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
