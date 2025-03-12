package DBreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type DBreader interface {
	ReadData(data []byte) Recipes
}

type XMLReader struct {
	Recipe Recipes
}

type JSONReader struct {
	Recipe Recipes
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []Cake   `xml:"cake" json:"cake"`
}

type Cake struct {
	Name        string       `xml:"name" json:"name"`
	CookTime    string       `xml:"stovetime" json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

type Ingredient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
}

func (reader *XMLReader) ReadData(data []byte) Recipes {
	xml.Unmarshal(data, &reader.Recipe)
	return reader.Recipe
}

func (reader *JSONReader) ReadData(data []byte) Recipes {
	json.Unmarshal(data, &reader.Recipe)
	return reader.Recipe
}

func GetReader(filename string) (DBreader, string) {
	if strings.HasSuffix(filename, ".xml") {
		return &XMLReader{}, "xml"
	}
	if strings.HasSuffix(filename, ".json") {
		return &JSONReader{}, "json"
	}
	return nil, "unknown file format"
}

func Indent(recipe Recipes, ext string) (output []byte, err error) {
	if ext == "xml" {
		output, err = json.MarshalIndent(recipe, "", "    ")
	} else {
		output, err = xml.MarshalIndent(recipe, "", "    ")
	}

	return output, err
}

func ReadDB(filename string) ([]byte, error) {
	reader, ext := GetReader(filename)
	if ext != "xml" && ext != "json" {
		return nil, fmt.Errorf("invalid input")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	recipe := reader.ReadData(data)
	var output []byte

	output, err = Indent(recipe, ext)
	if err != nil {
		return nil, err
	}

	return output, nil
}
