package config

import (
	"context"
	"fmt"
	"strings"

	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/constant"
)

func checkPFB(pfb string) bool {
	if pfb == "" {
		defaultPFBName := fmt.Sprintf("%s-%s-%s", GetProjectModule(), strings.ToLower(GetEnv()), GetCountry())
		return GetServiceName() == defaultPFBName
	}
	destPFB := strings.ReplaceAll(pfb, "-", "")
	pfbRule, isPFB := GetPFBRuleFromEnv()
	if isPFB && pfbRule == destPFB {
		return true
	}
	return false
}

func checkMainAPI(ctx context.Context, mainAPI string) bool {
	if mainAPI == "" {
		return true
	}
	apiName, ok := ctx.Value(constant.APIName).(string)
	if ok && apiName == mainAPI {
		return true
	}
	return false
}

func checkSubAPI(command string, subAPIList []string) bool {
	if len(subAPIList) == 0 {
		return true
	}
	for _, cmd := range subAPIList {
		if cmd != command {
			continue
		}
		return true
	}
	return false
}

// CheckMockerToggle ...
func CheckMockerToggle(ctx context.Context, command string) bool {
	env := strings.ToLower(GetEnv())
	if env == "live" || env == "liveish" {
		return false
	}

	mockConfig := GetSpexMockConfig()
	for _, tc := range mockConfig.Toggle {
		if tc.Method == constant.MockMethod &&
			checkPFB(tc.PFB) &&
			checkMainAPI(ctx, tc.MainAPI) &&
			checkSubAPI(command, tc.SubAPIList) {
			return true
		}
	}
	return false
}

// CheckRecorderToggle ...
func CheckRecorderToggle(ctx context.Context, command string) bool {
	env := strings.ToLower(GetEnv())
	if env == "live" || env == "liveish" {
		return false
	}

	mockConfig := GetSpexMockConfig()
	for _, tc := range mockConfig.Toggle {
		if tc.Method == constant.RecordMethod &&
			checkPFB(tc.PFB) &&
			checkMainAPI(ctx, tc.MainAPI) &&
			checkSubAPI(command, tc.SubAPIList) {
			return true
		}
	}
	return false
}
