package mock

import (
	"context"

	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/config"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/constant"
)

func ignoreMock(code uint32) bool {
	if value, ok := constant.ActionMap[code]; ok {
		return value
	}
	return false
}

// ErrorTranslate ...
func ErrorTranslate(code uint32) error {
	if value, ok := constant.ErrMocker[code]; ok {
		return value
	}
	return nil
}

// MockerForward ...
func MockerForward(ctx context.Context, command string, request, response interface{}) (uint32, bool) {
	var code uint32 = constant.ErrMockerToggleOff
	if config.CheckMockerToggle(ctx, command) {
		code = mocker.Agent(ctx, command, request, response)
	}
	return code, ignoreMock(code)
}

// RecorderForward ...
func RecorderForward(ctx context.Context, command string, request, response interface{}) uint32 {
	var code uint32 = constant.ErrRecorderToggleOff
	if config.CheckRecorderToggle(ctx, command) {
		code = recorder.Agent(ctx, command, request, response)
	}
	return code
}
