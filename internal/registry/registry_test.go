package registry_test

import (
	"testing"

	"github.com/johanhugg/gnomock"
	"github.com/johanhugg/gnomock/internal/registry"
	"github.com/stretchr/testify/require"
)

var p gnomock.Preset

func TestRegistry(t *testing.T) {
	registry.Register("preset", func() gnomock.Preset { return p })
	require.Equal(t, p, registry.Find("preset"))
	require.Nil(t, registry.Find("invalid"))
}
