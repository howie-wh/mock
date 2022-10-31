package config

import (
	"os"
	"strings"
)

// GetEnv is a method to get current environment
func GetEnv() string {
	environment := os.Getenv("ENV")

	if len(environment) == 0 {
		environment = "test"
	}
	return environment
}

// GetServiceName is a method to get Service Name
func GetServiceName() string {
	serviceName := os.Getenv("SERVICE_NAME")
	return strings.ToLower(serviceName)
}

// GetCountry is a method to get current country
func GetCountry() string {
	country := os.Getenv("CID")

	if len(country) == 0 {
		country = "ID"
	}
	return strings.ToLower(country)
}

// GetPFBRuleFromEnv get PFB related according to the env variable SERVICE_NAME
func GetPFBRuleFromEnv() (pfbRule string, isPFB bool) {
	// ServiceName, according to shopee_deploy/run.py:
	// For default env, '%s-%s-%s-%s' % (project_name, module_name, env, cid)
	// For pfb, '%s-%s-%s-%s-%s' % (project_name, module_name, get_pfb_dir(branch_name), env, cid)
	serviceName := os.Getenv("SERVICE_NAME")
	if strings.Count(serviceName, "-") > 3 {
		items := strings.Split(serviceName, "-")

		serviceName = strings.Join(items[2:len(items)-2], "")
		if len(serviceName) == 0 {
			return "", false
		}
		return strings.ReplaceAll(serviceName, "-", ""), true
	}
	return "", false
}

// GetPFB ...
func GetPFB() string {
	serviceName := GetServiceName()
	if strings.Count(serviceName, "-") > 3 {
		items := strings.Split(serviceName, "-")
		return strings.Join(items[2:len(items)-2], "-")
	}
	return ""
}

// GetProjectModule ...
func GetProjectModule() string {
	projectName := os.Getenv("PROJECT_NAME")
	if len(projectName) == 0 {
		projectName = "discoverlanding"
	}

	moduleName := os.Getenv("MODULE_NAME")
	if len(moduleName) == 0 {
		moduleName = "shopbff"
	}

	return projectName + "-" + moduleName
}
