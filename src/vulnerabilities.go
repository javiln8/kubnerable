package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Vulnerabilities struct {
	Version         string          `yaml:"version"`
	Description     string          `yaml:"description"`
	Vulnerabilities []Vulnerability `yaml:"vulnerabilities"`
}

type Vulnerability struct {
	Risk        string  `yaml:"risk"`
	CvssScore   float64 `yaml:"cvss_score"`
	Id          int     `yaml:"id"`
	Title       string  `yaml:"title"`
	Description string  `yaml:"description"`
	Checks      []Check `yaml:"checks"`
	Remediation string  `yaml:"remediation"`
}

type Check struct {
	Description string    `yaml:"description"`
	Resource    string    `yaml:"resource"`
	Compares    []Compare `yaml:"compare"`
}

type Compare struct {
	Item      string `yaml:"item"`
	Operation string `yaml:"operation"`
	Value     string `yaml:"value"`
}

// GetVulnerabilities stores the vulnerabilities database in its pertinent data structures
func GetVulnerabilities(vulsFilePath string) *Vulnerabilities {
	vulnerabilitiesYaml, err := ioutil.ReadFile(vulsFilePath)
	CheckError(err, "Could not find vulnerabilities.yaml")

	vulnerabilities := new(Vulnerabilities)
	err = yaml.Unmarshal(vulnerabilitiesYaml, vulnerabilities)
	CheckError(err, "Could not unmarshal vulnerabilities.yaml")

	return vulnerabilities
}

// CheckTest checks the current test conditional, comparing the Kubernetes resource with the vulnerability database
func CheckTest(a interface{}, operator, b string) bool {
	switch operator {
	case "==":
		a = fmt.Sprintf("%v", a) // interface to string
		if a == b {
			return true
		}
	// Container capabilities check
	case "add":
		aMap := ResourceToStringMap(a)
		for _, capability := range aMap[operator].([]interface{}) {
			if capability == b {
				return true
			}
		}

	// Pod with mounted volumes in hostPath
	case "path":
		for _, volume := range a.([]interface{}) {
			volumeMap := ResourceToStringMap(volume)
			if volumeMap["hostPath"] != nil {
				hostPath := ResourceToStringMap(volumeMap["hostPath"])
				if hostPath["path"] == b {
					return true
				}
			}
		}

	default:
		log.Fatalln("Could not perform the vulnerability check")
	}
	return false
}

// PrintVulnerability generates a colored output for each vulnerability finding
func PrintVulnerability(vul Vulnerability, containerName string, check Check, test Compare, podName string, namespace string) {
	greenColor := "\033[1;32m%s\033[0m"
	yellowColor := "\033[1;33m%s\033[0m"
	redColor := "\033[1;31m%s\033[0m"

	fmt.Printf(greenColor, vul.Title)
	fmt.Print("\n", vul.Description, "\nRisk:        ")

	if vul.Risk == "High" {
		fmt.Printf(redColor, vul.Risk)
		fmt.Print("\nCVSS Score:  ")
		fmt.Printf(redColor, FloatToString(vul.CvssScore))
	}

	if vul.Risk == "Medium" {
		fmt.Printf(yellowColor, vul.Risk)
		fmt.Print("\nCVSS Score:  ")
		fmt.Printf(yellowColor, FloatToString(vul.CvssScore))
	}

	// Kubernetes context related prints
	fmt.Printf("\nCheck:       %s: %s %s %s", check.Description, test.Item, test.Operation, test.Value)
	fmt.Printf("\nContainer:   %s", containerName)
	fmt.Printf("\nPod:         %s", podName)
	fmt.Printf("\nNamespace:   %s", namespace)

	fmt.Printf("\nRemediation: %s\n", vul.Remediation)
	fmt.Print("\n\n")
}
