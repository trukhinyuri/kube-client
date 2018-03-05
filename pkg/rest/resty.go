package rest

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	resty "github.com/go-resty/resty"
)

var (
	_ REST = &Resty{}
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
func (re *Resty) Get(reqconfig Rq) error {
	resp, err := reqconfig.ToResty(re.request).
		Get(reqconfig.Path.compile())
	if err = MapErrors(resp, err, http.StatusOK); err != nil {
		return err
	}
	if reqconfig.Result != nil {
		copyInterface(reqconfig.Result, resp.Result())
	}
	return nil
}

// Put -- http put method
func (re *Resty) Put(reqconfig Rq) error {
	resp, err := reqconfig.ToResty(re.request).
		Put(reqconfig.Path.compile())
	if err = MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusCreated); err != nil {
		return err
	}
	if reqconfig.Result != nil {
		copyInterface(reqconfig.Result, resp.Result())
	}
	return nil
}

// Post -- http post method
func (re *Resty) Post(reqconfig Rq) error {
	resp, err := reqconfig.ToResty(re.request).
		Post(reqconfig.Path.compile())
	if err = MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusCreated); err != nil {
		return err
	}
	if reqconfig.Result != nil {
		copyInterface(reqconfig.Result, resp.Result())
	}
	return nil
}

// Delete -- http delete method
func (re *Resty) Delete(reqconfig Rq) error {
	resp, err := reqconfig.ToResty(re.request).
		Post(reqconfig.Path.compile())
	return MapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted,
		http.StatusNoContent)
}

func (re *Resty) SetToken(token string) {
	re.request.SetHeader(HeaderUserToken, token)
}

// ToResty -- maps Rq data to resty request
func (rq *Rq) ToResty(req *resty.Request) *resty.Request {
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
