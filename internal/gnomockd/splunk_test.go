package gnomockd_test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/johanhugg/gnomock"
	"github.com/johanhugg/gnomock/internal/gnomockd"
	"github.com/johanhugg/gnomock/preset/splunk"
	"github.com/stretchr/testify/require"
)

func TestSplunk(t *testing.T) {
	t.Parallel()

	h := gnomockd.Handler()
	bs, err := os.ReadFile("./testdata/splunk.json")
	require.NoError(t, err)

	buf := bytes.NewBuffer(bs)
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/start/splunk", buf)
	h.ServeHTTP(w, r)

	res := w.Result()

	defer func() { require.NoError(t, res.Body.Close()) }()

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	require.Equalf(t, http.StatusOK, res.StatusCode, string(body))

	var c *gnomock.Container

	err = json.Unmarshal(body, &c)
	require.NoError(t, err)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	data := url.Values{}
	data.Add("search", "search index=sales | stats count")
	data.Add("earliest", "1546300800")
	data.Add("latest", "1609372800")
	data.Add("output_mode", "json")
	buf = bytes.NewBufferString(data.Encode())

	addr := fmt.Sprintf("https://%s/services/search/jobs/export", c.Address(splunk.APIPort))
	req, err := http.NewRequest(http.MethodPost, addr, buf)
	require.NoError(t, err)
	req.SetBasicAuth("admin", "12345678")
	res, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	defer func() { require.NoError(t, res.Body.Close()) }()

	out := struct {
		Result struct {
			Count string `json:"count"`
		} `json:"result"`
	}{}

	bs, err = io.ReadAll(res.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bs, &out))
	require.Equal(t, "525", out.Result.Count)

	bs, err = json.Marshal(c)
	require.NoError(t, err)

	buf = bytes.NewBuffer(bs)
	w, r = httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/stop", buf)
	h.ServeHTTP(w, r)

	res = w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
}
