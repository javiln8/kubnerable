package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"k8s.io/api/core/v1"
)

type Findings struct {
	Findings []Finding
}

type Finding struct {
	PodName         string `json:"pod_name"`
	ContainerName   string `json:"container_name"`
	VulnerabilityId int    `json:"vulnerability_id"`
	Description     string `json:"description"`
	ResourceName    string `json:"resource_name"`
	ResourceValue   string `json:"resource_value"`
	ExploitCommand  string `json:"exploit_command"`
}

// AnalyzePods checks all the vulnerabilities for all running pods in the cluster
func AnalyzePods(vulsFilePath string) {
	findings := new(Findings)
	pods := GetPods(GetClient(GetKubeconfig()))
	vuls := GetVulnerabilities(vulsFilePath)

	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			for _, vul := range vuls.Vulnerabilities {
				for _, check := range vul.Checks {
					EvaluateCheck(check, vul, findings, container, pod)
				}
			}
		}
	}
	FindingsToJson(*findings)
}

// EvaluateCheck evaluates if a vulnerable resource/parameter exists, based on the vulnerabilities YAML
func EvaluateCheck(check Check, vul Vulnerability, findings *Findings, container v1.Container, pod v1.Pod) {
	switch check.Resource {
	case "container.SecurityContext":
		securityContextMap := ResourceToStringMap(container.SecurityContext)
		EvaluateTests(check, securityContextMap, vul, findings, container.Name, pod)
	case "pod.SecurityContext":
		securityContextMap := ResourceToStringMap(pod.Spec.SecurityContext)
		EvaluateTests(check, securityContextMap, vul, findings, "*", pod)
	case "pod.Spec":
		specMap := ResourceToStringMap(pod.Spec)
		EvaluateTests(check, specMap, vul, findings, "*", pod)
	}
}

// EvaluateTests runs all tests for each vulnerability check
func EvaluateTests(check Check, resourceMap map[string]interface{}, vul Vulnerability, findings *Findings, containerName string, pod v1.Pod) {
	for _, test := range check.Compares {
		if setting, ok := resourceMap[test.Item]; ok {
			if CheckTest(setting, test.Operation, test.Value) {
				if !CheckDuplicatedFinding(findings, pod.Name, containerName, vul.Id) {
					PrintVulnerability(vul, containerName, check, test, pod.Name, pod.Namespace)

					finding := Finding{
						PodName:         pod.Name,
						ContainerName:   containerName,
						VulnerabilityId: vul.Id,
						Description:     check.Description,
						ResourceName:    test.Item,
						ResourceValue:   test.Value,
					}
					findings.Findings = append(findings.Findings, finding)
				}
			}
		}
	}
}

// CheckDuplicatedFinding checks that the new vulnerability find is not already stored
func CheckDuplicatedFinding(findings *Findings, podName string, containerName string, vulId int) bool {
	temporalFinding := Finding{
		PodName:         podName,
		ContainerName:   containerName,
		VulnerabilityId: vulId,
	}

	for _, finding := range findings.Findings {
		if temporalFinding.PodName == finding.PodName && temporalFinding.ContainerName == finding.ContainerName &&
			temporalFinding.VulnerabilityId == finding.VulnerabilityId {
			return true
		}
	}
	return false
}

// FindingsToJson generates a JSON report of the run findings into disk
func FindingsToJson(findings Findings) {
	boldText := "\033[1m%s\033[0m"

	findingsReport, err := json.MarshalIndent(findings, "", "  ")
	CheckError(err, "Could not marshal findings into JSON")
	err = ioutil.WriteFile("findings_report.json", findingsReport, 0644)
	CheckError(err, "Could not generate the JSON report")

	fmt.Printf(boldText, "Vulnerabilities found stored at findings_report.json\n")
}
