package tests

import (
	DBreader "Day-01/pkg/readDB"
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"
)

const XMLdata = `<recipes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>40 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>3</itemcount>
                <itemunit>cups</itemunit>
            </item>
            <item>
                <itemname>Vanilla extract</itemname>
                <itemcount>1.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>7</itemcount>
                <itemunit></itemunit> <!-- itemunit may be empty  -->
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
                <itemunit>pieces</itemunit>
            </item>
        </ingredients>
    </cake>
</recipes>`

const JSONdata = `{
    "cake": [
      {
        "name": "Red Velvet Strawberry Cake",
        "time": "45 min",
        "ingredients": [
          {
            "ingredient_name": "Flour",
            "ingredient_count": "2",
            "ingredient_unit": "mugs"
          },
          {
            "ingredient_name": "Strawberries",
            "ingredient_count": "8"  
          },
          {
            "ingredient_name": "Coffee Beans",
            "ingredient_count": "2.5",
            "ingredient_unit": "tablespoons"
          },
          {
            "ingredient_name": "Cinnamon",
            "ingredient_count": "1"
          }
        ]
      }
	]
}`

func Test_XMLReader_ReadData(t *testing.T) {
	reader := DBreader.XMLReader{}
	recipe := reader.ReadData([]byte(XMLdata))

	if len(recipe.Cakes) != 1 {
		t.Errorf("Expected 1 cake, got %d", len(recipe.Cakes))
	}

	if recipe.Cakes[0].Name != "Red Velvet Strawberry Cake" {
		t.Errorf("Expected cake's name: Red Velvet Strawberry Cake, got %s", recipe.Cakes[0].Name)
	}
}

func Test_JSONReader_ReadData(t *testing.T) {
	reader := DBreader.JSONReader{}
	recipe := reader.ReadData([]byte(JSONdata))

	if len(recipe.Cakes) != 1 {
		t.Errorf("Expected 1 cake, got %d", len(recipe.Cakes))
	}

	if recipe.Cakes[0].Name != "Red Velvet Strawberry Cake" {
		t.Errorf("Expected cake's name: Red Velvet Strawberry Cake, got %s", recipe.Cakes[0].Name)
	}
}

func Test_GetReader(t *testing.T) {
	reader, ext := DBreader.GetReader("test.xml")

	if _, ok := reader.(*DBreader.XMLReader); !ok || ext != "xml" {
		t.Errorf("Expected DBReader.XMLReader type and ext xml, got %T type and ext %s", reader, ext)
	}

	reader, ext = DBreader.GetReader("test.json")

	if _, ok := reader.(*DBreader.JSONReader); !ok || ext != "json" {
		t.Errorf("Expected DBReader.JSONReader type and ext json, got %T type and ext %s", reader, ext)
	}
}

func Test_Indent(t *testing.T) {
	var recipe DBreader.Recipes
	xml.Unmarshal([]byte(JSONdata), &recipe)

	output, err := DBreader.Indent(recipe, "json")
	if err != nil || !strings.Contains(string(output), "<recipes>") {
		t.Errorf("Expected XML output, got %s", output)
	}

	json.Unmarshal([]byte(XMLdata), &recipe)

	output, err = DBreader.Indent(recipe, "xml")
	if err != nil || !strings.Contains(string(output), "cake") {
		t.Errorf("Expected JSON output, got %s", output)
	}
}
