# Gnomock Elasticsearch

Gnomock Elasticsearch is a [Gnomock](https://github.com/johanhugg/gnomock)
preset for running tests against a real Elasticsearch container, without mocks.

```go
package elastic_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/johanhugg/gnomock"
	"github.com/johanhugg/gnomock/preset/elastic"
	"github.com/stretchr/testify/require"
)

func TestPreset(t *testing.T) {
	t.Parallel()

	p := elastic.Preset(
		elastic.WithInputFile("./testdata/titles"),
		elastic.WithInputFile("./testdata/names"),
	)

	c, err := gnomock.Start(p)
	require.NoError(t, err)

	defer func() { require.NoError(t, gnomock.Stop(c)) }()

	cfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s", c.DefaultAddress())},
	}

	client, err := elasticsearch.NewClient(cfg)
	require.NoError(t, err)

	res, err := client.Search(
		client.Search.WithIndex("titles"),
		client.Search.WithQuery("gnomock"),
	)
	require.NoError(t, err)
	require.False(t, res.IsError(), res.String())

	var out struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}
	require.NoError(t, json.NewDecoder(res.Body).Decode(&out))
	require.NoError(t, res.Body.Close())
	require.Equal(t, 1, out.Hits.Total.Value)

	res, err = client.Search(
		client.Search.WithIndex("titles"),
		client.Search.WithQuery("unknown"),
	)
	require.NoError(t, err)
	require.False(t, res.IsError(), res.String())

	require.NoError(t, json.NewDecoder(res.Body).Decode(&out))
	require.NoError(t, res.Body.Close())
	require.Equal(t, 0, out.Hits.Total.Value)
}
```
