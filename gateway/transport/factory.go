package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Zhan9Yunhua/blog-svr/common"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Zhan9Yunhua/logger"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	kithttp "github.com/go-kit/kit/transport/http"
)

func svcFactory(method string, path string) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		tgt, err := url.Parse(instance)
		logger.Infoln("listening svc: ", tgt)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path

		var (
			enc kithttp.EncodeRequestFunc
			dec kithttp.DecodeResponseFunc
		)

		method = strings.ToUpper(method)

		if method == "GET" {
			enc, dec = encodeGetRequest, decodeResponse
		} else {
			enc, dec = encodeJsonRequest, decodeResponse
		}

		return kithttp.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}

// 客户端到内部服务：转换Get请求
func encodeGetRequest(_ context.Context, req *http.Request, request interface{}) error {
	data, ok := request.(common.RequestUrlParams)
	if ok {
		req.URL.Path = strings.Replace(req.URL.Path, "{param}", data.Param, -1)
	}

	return nil
}

// 客户端到内部服务：转换Json请求
func encodeJsonRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		logger.Debugln(err.Error())
		return err
	}
	req.Body = ioutil.NopCloser(&buf)

	return nil
}

// 内部 -> 外部响应
func decodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var innerResponse common.InnerResponse
	var outputResponse common.OutputResponse

	if err := json.NewDecoder(r.Body).Decode(&innerResponse); err != nil {
		return nil, err
	}
	if innerResponse.Err == "" {
		outputResponse.Msg = innerResponse.Msg
		outputResponse.Code = innerResponse.Code
		outputResponse.Data = innerResponse.Data
	} else {
		outputResponse.Msg = innerResponse.Err
		outputResponse.Code = common.Error
		outputResponse.Data = nil
	}
	return outputResponse, nil
}