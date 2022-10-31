package mock

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/config"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/constant"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/logger"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/model"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/utils"

	"github.com/sirupsen/logrus"
)

var mocker = &mockerImpl{}

type mockerImpl struct {
}

func (m *mockerImpl) buildRequest(ctx context.Context, command string, request interface{}) (interface{}, error) {
	mainAPI, ok := ctx.Value(constant.APIName).(string)
	if !ok {
		return nil, fmt.Errorf("%s", "loss mainAPI")
	}

	var srcRequest = []byte("{}")
	var err error

	obj := reflect.ValueOf(request)
	switch obj.Kind() {
	case reflect.Ptr:
		if obj.Elem().Kind() == reflect.Struct {
			srcRequest, err = json.Marshal(request)
			if err != nil {
				return nil, err
			}
		}
	default:
	}

	return &model.MKAgentRequest{
		BasicParam: model.BasicParam{
			Project: config.GetProjectModule(),
			Env:     config.GetEnv(),
			Region:  config.GetCountry(),
			PFB:     config.GetPFB(),
			MainAPI: mainAPI,
			SubAPI:  command,
		},
		SrcRequest: string(srcRequest),
	}, nil
}

func (m *mockerImpl) forwardRequest(ctx context.Context, request, response interface{}) (uint32, error) {
	url := config.GetSpexMockConfig().MKCenter.MockURL
	var resp model.MKAgentResponse
	if err := utils.RequestPostWithContext(ctx, url, request, &resp); err != nil {
		return constant.ErrMockerForwardRequest, err
	}

	obj := reflect.ValueOf(response)
	switch obj.Kind() {
	case reflect.Ptr:
		if obj.Elem().Kind() == reflect.Struct {
			if err := json.Unmarshal([]byte(resp.Data), response); err != nil {
				return constant.ErrMockerResponseFormat, err
			}
		}
	default:
	}
	return resp.Error, nil
}

// Agent ...
func (m *mockerImpl) Agent(ctx context.Context, command string, request, response interface{}) uint32 {
	req, err := m.buildRequest(ctx, command, request)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"cmd":  command,
			"req":  req,
			"resp": response,
		}).Error(err)
		return constant.ErrMockerBuildRequest
	}

	code, err := m.forwardRequest(ctx, req, response)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"cmd":  command,
			"code": code,
			"req":  req,
			"resp": response,
		}).Error(err)
		return code
	}

	logger.Log.WithFields(logrus.Fields{
		"cmd":  command,
		"code": code,
		"req":  req,
		"resp": response,
	}).Debug("mocker_agent")

	return code
}
