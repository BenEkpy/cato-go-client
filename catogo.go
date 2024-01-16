package catogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type httpClient struct {
	httpclient *http.Client
	token      string
	baseurl    string
}

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type validResponse struct {
	Data   interface{}   `json:"data,omitempty"`
	Errors []interface{} `json:"errors,omitempty"`
}

func APIClient(token string) *httpClient {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	return &httpClient{
		httpclient: client,
		baseurl:    "https://api.catonetworks.com/api/v1/graphql2",
		token:      token,
	}
}

func (c *httpClient) post(reqBody graphQLRequest, respBody *validResponse) error {

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.baseurl, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return err
	}

	req.Header.Set("x-api-key", c.token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpclient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	byteRespBody, err := io.ReadAll(res.Body)


	if res.StatusCode == http.StatusOK {

		json.Unmarshal(byteRespBody, respBody)

		if respBody.Errors != nil {
			json_error, _ := json.Marshal(respBody.Errors)
			return fmt.Errorf(string(json_error))
		}

	} else {

		// test to remove & improve error handling
		json.Unmarshal(byteRespBody, respBody)
		test, _ := json.Marshal(respBody)
		fmt.Println(string(test))

		return fmt.Errorf("unknown error")

	}

	return nil

}
