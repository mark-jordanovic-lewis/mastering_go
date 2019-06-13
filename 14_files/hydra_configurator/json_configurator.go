package hydra_configurator

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func decodeJSONConfigFile(config interface{}, filename string) error {
	fmt.Printf("Decoding %v JSON\n", filename)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	return decodeJSONConfig(config, file)
}

func decodeJSONConfig(config interface{}, reader io.Reader) error {
	return json.NewDecoder(reader).Decode(config)
}