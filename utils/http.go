package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"git.garena.com/shopee/seller-server/seller-marketing/marketing-common/mock/constant"
)

// RequestPostWithContext ...
func RequestPostWithContext(_ context.Context, url string, req interface{}, resp interface{}) error {
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), constant.HTTPExpire)
	defer cancel()
	reqBytes := bytes.NewReader(reqJSON)
	reqData, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBytes)
	reqData.Header.Set("Content-Type", "application/json;charset=UTF-8")

	var respRaw *http.Response
	c := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	if respRaw, err = c.Do(reqData); err != nil {
		return err
	}
	defer respRaw.Body.Close()

	var respBytes []byte
	if respBytes, err = ioutil.ReadAll(respRaw.Body); err != nil {
		return err
	}
	if err = json.Unmarshal(respBytes, resp); err != nil {
		return err
	}

	return nil
}
