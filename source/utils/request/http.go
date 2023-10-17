package request

import (
	"bytes"
	"fmt"
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	client  *http.Client
	timeout time.Duration
	logger  logger.ILogger
}

type QueryMap map[string]string

func NewClient(logger logger.ILogger) *HttpClient {
	return &HttpClient{
		timeout: 30 * time.Second,
		logger:  logger,
	}
}

func (c *HttpClient) HttpGet(url string, query map[string]string, headers map[string]string) ([]byte, error) {
	_url := fmt.Sprintf("%s?%s", url, queryMapEncode(query))
	return c.DoRequest(http.MethodGet, _url, nil, headers)
}

func (c *HttpClient) DoRequest(method string, url string, body []byte, headers map[string]string) ([]byte, error) {
	c.client = &http.Client{
		Timeout: c.timeout,
	}

	var bodyReader io.Reader

	if len(body) > 0 {
		bodyReader = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		c.logger.Error(xerrors.Errorf("http DoRequest make request error: %w", err))
		return nil, err
	}

	//req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	//req.Header.Set("Accept", "application/json; charset=UTF-8")
	//req.Header.Set("Accept-Encoding", "utf-8")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := c.client.Do(req)
	if err != nil {
		c.logger.Error(xerrors.Errorf("http DoRequest send request error: %w", err))
		return nil, err
	}

	readRes, err := io.ReadAll(res.Body)

	if err != nil {
		c.logger.Error(xerrors.Errorf("http DoRequest read response error: %w", err))
		return nil, err
	}

	return readRes, nil
}

func queryMapEncode(qm QueryMap) string {
	uq := url.Values{}
	for k, v := range qm {
		if v != "" {
			uq.Add(k, v)
		}
	}
	return uq.Encode()
}
