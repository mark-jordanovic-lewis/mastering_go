package hydra_configurator

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func decodeXMLConfigFile(config interface{}, filename string) error {
	fmt.Println("Decoding", filename, "XML")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	return decodeXMLConfig(config, file)
}

func decodeXMLConfig(config interface{}, reader io.Reader) error {
	return xml.NewDecoder(reader).Decode(config)
}