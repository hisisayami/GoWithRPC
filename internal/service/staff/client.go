package staff

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// type for thr most common header keys
type HeaderKey string

const (
	contentType   HeaderKey = "Content-Type"
	authorization HeaderKey = "Authorization"
)

type HeaderValue string

const (
	applicationJSON HeaderValue = "application/json"
)

type API interface {
	Validate(ctx context.Context, input StaffInfoInput) (*StaffInfoResponse, error)
}

type newRequestFunc func(method string, url string, body io.Reader) (*http.Request, error)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	client         Doer
	host           string
	newRequestFunc newRequestFunc
}

func New(httpClient Doer, host string) API {
	return &client{
		host:           host,
		newRequestFunc: http.NewRequest,
		client:         httpClient,
	}
}

// custom error
type APIErr struct {
	CorrelationID string `json:"correlationID"`
	Err           struct {
		Code     int    `json:"code"`
		Message  string `json:"message"`
		Cause    string `json:"cause"`
		GRPCCode int    `json:"GRPCCode"`
	} `json:"error"`
}

func (e *APIErr) Error() string {
	return fmt.Sprintf("failed with status '%d' with correlation id '%s': %s", e.Err.Code, e.CorrelationID, e.Err.Message)
}

func handleAPIError(resp *http.Response, v []byte) error {
	if resp.StatusCode != http.StatusOK {
		var apiErr APIErr
		if err := json.Unmarshal(v, &apiErr); err != nil {
			return errors.Wrapf(err, "unexpected status code '%d': failed to unmarshall api error", resp.StatusCode)
		}

		return &apiErr
	}

	return nil
}

func (c *client) Validate(ctx context.Context, input StaffInfoInput) (*StaffInfoResponse, error) {

	p, err := json.Marshal(&input)
	if err != nil {
		return nil, errors.New("failed to marshall login payload")
	}

	req, err := c.newRequestFunc(http.MethodPost, fmt.Sprintf("%s/validateStaff", c.host), bytes.NewBuffer(p))
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct req")
	}

	req.Header.Add(string(contentType), string(applicationJSON))
	req.Header.Add(string(authorization), fmt.Sprintf("Bearer %s", "jwt"))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute request")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	if err := handleAPIError(resp, respBody); err != nil {
		return nil, err
	}

	var staffInfoResponse StaffInfoResponse
	err = json.Unmarshal(respBody, &staffInfoResponse)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarsall response body ")
	}

	return &staffInfoResponse, nil

}
