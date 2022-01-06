package reader_test

import (
	"testing"

	"github.com/michael-valdron/schemer/pkg/data"
	"github.com/michael-valdron/schemer/pkg/reader"
	"github.com/stretchr/testify/assert"
)

func TestReadJsonBytes(t *testing.T) {
	expected := data.DataBlob{}
	expected.InitDataBlob(map[string]interface{}{"name": "Michael Valdron"})

	assert.Equal(t, expected, reader.ReadJsonBytes([]byte(`{"name": "Michael Valdron"}`)))
}

func TestReadJson(t *testing.T) {
	expected := data.DataBlob{}
	expected.InitDataBlob(map[string]interface{}{"name": "Michael Valdron"})

	assert.Equal(t, expected, reader.ReadJson(`{"name": "Michael Valdron"}`))
}

func TestReadJsonFile(t *testing.T) {
	expected := data.DataBlob{}
	expected.InitDataBlob(map[string]interface{}{"name": "John Smith"})

	assert.Equal(t, expected, reader.ReadJsonFile("../../test.json"))
}
