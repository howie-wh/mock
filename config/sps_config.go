package config

import (
	"git.garena.com/shopee/golang_splib/sps"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/logger"
	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/model"

	"github.com/sirupsen/logrus"
)

// SpexConfig is an interface to get the config by key
type SpexConfig interface {
	Get(key string) (interface{}, error)
}

var spexConfig SpexConfig

type defaultSpexConfig struct{}

// Get method is a function to get the config by key
func (sc defaultSpexConfig) Get(key string) (interface{}, error) {
	c, err := sps.GetConfigRegistry().Get(key)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// InitSpexConfig ...
func InitSpexConfig() bool {
	if spexConfig != nil {
		return true
	}
	cfg := sps.GetConfigRegistry()
	err := cfg.BindProto(model.MKConfigKey, &model.SpexMockConfig{})
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"config_key": model.MKConfigKey,
		}).Error(err)
		return false
	}
	spexConfig = &defaultSpexConfig{}
	return true
}

// GetSpexMockConfig is get config of SpexMockConfig
func GetSpexMockConfig() *model.SpexMockConfig {
	if spexConfig != nil || InitSpexConfig() {
		c, _ := spexConfig.Get(model.MKConfigKey)
		if c != nil {
			if value, ok := c.(*model.SpexMockConfig); ok {
				return value
			}
		}
	}
	return &model.SpexMockConfig{}
}
