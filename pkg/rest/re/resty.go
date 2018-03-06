package re

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	"git.containerum.net/ch/kube-client/pkg/rest"
	resty "github.com/go-resty/resty"
)

var (
	_ rest.REST = &Resty{}
)

// Resty -- resty client,
// implements REST interface
type Resty struct {
	request *resty.Request
}

// NewResty -- Resty constuctor
func NewResty() *Resty {
	return &Resty{
		request: resty.R(),
	}
}

// Get -- http get method
func (re *Resty) Get(reqconfig rest.Rq) error {
	resp, err := ToResty(reqconfig, re.request).
		Get(reqconfig.Path.Build())
	if err = rest.MapErrors(resp, err, http.StatusOK); err != nil {
		return err
	}
	if reqconfig.Result != nil {
		rest.CopyInterface(reqconfig.Result, resp.Result())
	}
	return nil
}

// Put -- http put method
func (re *Resty) Put(reqconfig rest.Rq) error {
	resp, err := ToResty(reqconfig, re.request).
		Put(reqconfig.Path.Build())
	if err = rest.MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusCreated); err != nil {
		return err
	}
	if reqconfig.Result != nil {
		rest.CopyInterface(reqconfig.Result, resp.Result())
	}
	return nil
}

// Post -- http post method
func (re *Resty) Post(reqconfig rest.Rq) error {
	resp, err := ToResty(reqconfig, re.request).
		Post(reqconfig.Path.Build())
	if err = rest.MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusCreated); err != nil {
		return err
	}
	if reqconfig.Result != nil {
		rest.CopyInterface(reqconfig.Result, resp.Result())
	}
	return nil
}

// Delete -- http delete method
func (re *Resty) Delete(reqconfig rest.Rq) error {
	resp, err := ToResty(reqconfig, re.request).
		Post(reqconfig.Path.Build())
	return rest.MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusNoContent)
}

func (re *Resty) SetToken(token string) {
	re.request.SetHeader(rest.HeaderUserToken, token)
}

// ToResty -- maps Rq data to resty request
func ToResty(rq rest.Rq, req *resty.Request) *resty.Request {
	if rq.Result != nil {
		req = req.SetResult(rq.Result)
	}
	if rq.Body != nil {
		req = req.SetBody(rq.Body)
	}
	if len(rq.Query) > 0 {
		req = req.SetQueryParams(rq.Query)
	}
	return req.SetError(cherry.Err{})
}
