package misc

import (
	"net/http"

	"github.com/aws/aws-xray-sdk-go/xray"
)

// TransportWrapper X-Ray セグメント情報を保持したトランスポータ
type TransportWrapper struct {
	Transport     http.RoundTripper
	Authorization string
	XraySeg       *xray.Segment
}

func (t *TransportWrapper) transport() http.RoundTripper {
	if t.Transport == nil {
		return http.DefaultTransport
	}
	return t.Transport
}

// CancelRequest リクエストのキャンセル
func (t *TransportWrapper) CancelRequest(req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	if c, ok := t.transport().(canceler); ok {
		c.CancelRequest(req)
	}
}

// RoundTrip カスタムヘッダを付与
func (t *TransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Authorization != "" {
		req.Header.Set("Authorization", t.Authorization)
	}
	if t.XraySeg != nil {
		req.Header.Set("x-amzn-trace-id", t.XraySeg.DownstreamHeader().String())
	}
	return t.transport().RoundTrip(req)
}
