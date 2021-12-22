package controllers_test

import (
	"ariefsn/game-currency/global/helper"
	"ariefsn/game-currency/global/models"
	"ariefsn/game-currency/services/currency-service/controllers"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, *models.ResponseModel) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	// http client that doesn't redirect
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	defer resp.Body.Close()

	res := models.NewResponseModel()

	err = json.Unmarshal(respBody, res)

	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	return resp, res
}

func buildPayload(payload map[string]interface{}) *bytes.Buffer {
	jsonValue, _ := json.Marshal(payload)

	buff := bytes.NewBuffer(jsonValue)

	return buff
}

func TestPing(t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		r := chi.NewRouter()
		controllers.NewRootController().Register(r)
		ts := httptest.NewServer(r)
		helper.InitRender()
		defer ts.Close()

		_, body := testRequest(t, ts, http.MethodGet, "/", nil)

		assert.NotNil(t, body)
		assert.True(t, body.Success)
		assert.Contains(t, body.Data, "running")
	})
}
