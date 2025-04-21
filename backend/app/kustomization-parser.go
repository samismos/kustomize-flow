package app

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

// ReadKustomization reads and parses the YAML file into the appropriate struct
func ReadKustomization(filePath string) (any, error) {
	var baseKustomization BaseKustomization
	var kustomizationConfig KustomizationConfig
	var kustomizationToolkit KustomizationToolkit

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &baseKustomization); err != nil {
		log.Fatalf("Error unmarshalling BaseKustomization: %v", err)
	}

	if strings.Contains(baseKustomization.APIVersion, "kustomize.config") {
		if err = yaml.Unmarshal(data, &kustomizationConfig); err == nil {
			return kustomizationConfig, nil
		}

	} else if strings.Contains(baseKustomization.APIVersion, "kustomize.toolkit") {
		if err = yaml.Unmarshal(data, &kustomizationToolkit); err == nil {

			if IsFluxVariableValid(kustomizationToolkit.Spec.PostBuild.Substitute.ApplicationName) {
				application_name = kustomizationToolkit.Spec.PostBuild.Substitute.ApplicationName
			}

			if IsFluxVariableValid(kustomizationToolkit.Spec.PostBuild.Substitute.KustomizationOverlay) {
				kustomization_overlay = kustomizationToolkit.Spec.PostBuild.Substitute.KustomizationOverlay
			}

			return kustomizationToolkit, nil
		}
	} else if strings.Contains(baseKustomization.APIVersion, "v1") {
		return baseKustomization, nil
	}

	return nil, fmt.Errorf("unknown kustomization format")
}

// PrintKustomization recursively prints the contents of a Kustomization
func PrintKustomization(v reflect.Value) {
	typeOfS := v.Type()

	for i := range v.NumField() {
		field := v.Field(i)
		fieldName := typeOfS.Field(i).Name
		fmt.Printf("%s:\n", fieldName)

		switch field.Kind() {
		case reflect.Slice:
			for j := range field.Len() {
				PrintValue(field.Index(j))
			}
		case reflect.Struct:
			PrintKustomization(field)
		default:
			PrintValue(field)
		}
	}
}

// PrintValue prints a single value
func PrintValue(v reflect.Value) {
	switch v.Kind() {
	case reflect.String:
		fmt.Printf(" - %v\n", v.Interface().(string))
	case reflect.Map:
		for _, key := range v.MapKeys() {
			fmt.Printf(" - %v: %v\n", key, v.MapIndex(key))
		}
	default:
		fmt.Printf(" - %v\n", v.Interface())
	}
}

func ReadAndPrintKustomization(filePath string) interface{} {
	kustomization, err := ReadKustomization(filePath)
	if err != nil {
		log.Fatalf("Error reading Kustomization: %v", err)
	}
	fmt.Println(Green + filePath + Reset)
	PrintKustomization(reflect.ValueOf(kustomization))
	return kustomization
}
