package mock

import (
	"context"
	"encoding/json"
	"fmt"

	"git.garena.com/shopee/golang_splib/sps"
	hsps "git.garena.com/shopee/platform/golang_splib/sps"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/config"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/constant"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/logger"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/model"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/utils"

	"github.com/sirupsen/logrus"
)

func getTraceId(ctx context.Context) string {
	spexHeader := sps.FromIncomingContext(ctx)
	var requestID string
	if requestID = spexHeader.RequestID(); requestID == "" {
		spexHeader = hsps.FromIncomingContext(ctx)
		requestID = spexHeader.RequestID()
	}
	return requestID
}

var recorder = &recorderImpl{}

type recorderImpl struct {
}

func (r *recorderImpl) buildRequest(ctx context.Context, command string, request, response interface{}) (interface{}, error) {
	mainAPI, ok := ctx.Value(constant.APIName).(string)
	if !ok {
		return nil, fmt.Errorf("%s", "loss mainAPI")
	}

	var originReq, originResp []byte
	var err error
	if originReq, err = json.Marshal(request); err != nil {
		return nil, err
	}
	if originResp, err = json.Marshal(response); err != nil {
		return nil, err
	}

	return &model.RecordAgentRequest{
		BasicParam: model.BasicParam{
			Project: config.GetProjectModule(),
			Env:     config.GetEnv(),
			Region:  config.GetCountry(),
			PFB:     config.GetPFB(),
			MainAPI: mainAPI,
			SubAPI:  command,
		},
		RequestID:  getTraceId(ctx),
		OriginReq:  string(originReq),
		OriginResp: string(originResp),
	}, nil
}

func (r *recorderImpl) forwardRequest(ctx context.Context, request interface{}) (uint32, error) {
	url := config.GetSpexMockConfig().MKCenter.RecordURL
	var resp model.RecordAgentResponse
	if err := utils.RequestPostWithContext(ctx, url, request, &resp); err != nil {
		return constant.ErrRecorderForwardRequest, err
	}

	if resp.Error != 0 {
		return resp.Error, fmt.Errorf(resp.ErrorMsg)
	}
	return 0, nil
}

// Agent ...
func (r *recorderImpl) Agent(ctx context.Context, command string, request, response interface{}) uint32 {
	req, err := r.buildRequest(ctx, command, request, response)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"cmd":  command,
			"req":  req,
			"resp": response,
		}).Error(err)
		return constant.ErrRecorderBuildRequest
	}

	code, err := r.forwardRequest(ctx, req)
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
	}).Debug("record_agent")

	return code
}
