package reader

import (
	"encoding/json"
	"os"

	"github.com/michael-valdron/schemer/pkg/data"
)

func ReadJsonBytes(jsonBytes []byte) data.DataBlob {
	var blob data.DataBlob
	var m map[string]interface{}

	json.Unmarshal(jsonBytes, &m)

	blob.InitDataBlob(m)

	return blob
}

func ReadJson(jsonStr string) data.DataBlob {
	return ReadJsonBytes([]byte(jsonStr))
}

func ReadJsonFile(filepath string) data.DataBlob {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	return ReadJsonBytes(content)
}
