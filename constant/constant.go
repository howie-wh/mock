package constant

import (
	"fmt"
	"time"
)

const (
	// ErrMockerToggleOff mocker toggle off
	ErrMockerToggleOff uint32 = 50
	// ErrMockerEmptyResponse mock center not data
	ErrMockerEmptyResponse uint32 = 100
	// ErrMockerBuildRequest mock request build error
	ErrMockerBuildRequest uint32 = 200
	// ErrMockerForwardRequest mock request forward error
	ErrMockerForwardRequest uint32 = 201
	// ErrMockerResponseFormat mock response data format error
	ErrMockerResponseFormat uint32 = 202

	// ErrRecorderToggleOff recorder toggle off
	ErrRecorderToggleOff uint32 = 51
	// ErrRecorderBuildRequest mock request build error
	ErrRecorderBuildRequest uint32 = 300
	// ErrRecorderForwardRequest mock request forward error
	ErrRecorderForwardRequest uint32 = 301

	// APIName ...
	APIName = "API_NAME"

	// LogFilePath ...
	LogFilePath = "./logmock"
	// LogFileName ...
	LogFileName = "mock"
	// LogLevel ...
	LogLevel = "debug"

	// MockMethod ...
	MockMethod int = 1
	// RecordMethod ...
	RecordMethod int = 2

	// HTTPExpire ...
	HTTPExpire = 2 * time.Minute
)

var (
	// ActionMap ...
	ActionMap = map[uint32]bool{
		ErrMockerEmptyResponse:  true,
		ErrMockerToggleOff:      true,
		ErrMockerBuildRequest:   false,
		ErrMockerForwardRequest: false,
		ErrMockerResponseFormat: false,
	}
	// ErrMocker ...
	ErrMocker = map[uint32]error{
		ErrMockerBuildRequest:   fmt.Errorf("build mock request struct error, code:[%d]", ErrMockerBuildRequest),
		ErrMockerForwardRequest: fmt.Errorf("forward mock request error, code:[%d]", ErrMockerForwardRequest),
		ErrMockerResponseFormat: fmt.Errorf("mock response data format error, error:[%d]", ErrMockerResponseFormat),
	}
)
