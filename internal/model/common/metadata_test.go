package commonModel

import (
	"testing"

	"github.com/paper-indonesia/pg-mcp-server/constant"
	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {
	metadataKeyDelete := "test2"
	info := TMapAny{
		"test": "metadata",
	}
	info2 := TMapAny{
		"test2": "metadata2",
	}
	metadata := NewMetadata()
	metadata.Set(constant.MetadataAdditionalInfoKey, info)
	metadata.Set(metadataKeyDelete, info2)

	infoAny := metadata.Get(constant.MetadataAdditionalInfoKey)
	additionalInfo, ok := infoAny.(TMapAny)

	assert.True(t, ok)
	assert.NotEmpty(t, additionalInfo)
	assert.Equal(t, info.Json(), additionalInfo.Json())

	metadata.Del(metadataKeyDelete)
	deletedValue := metadata.Get(metadataKeyDelete)
	assert.Nil(t, deletedValue)

	jsonByte := metadata.Json()
	assert.NotEmpty(t, jsonByte)

	newMetadata := NewMetadata()
	newMetadata.FromStr(string(jsonByte))
	infofromNew := newMetadata.Get(constant.MetadataAdditionalInfoKey).(map[string]any)

	assert.Equal(t, info["test"], infofromNew["test"])
}
