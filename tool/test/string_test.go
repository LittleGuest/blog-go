package tool

import (
	"blog/tool"
	"testing"
)

func TestToCamelCase(t *testing.T) {
	t.Log(tool.ToCamelCase("dd_dd", "_"))
}

func TestToPascal(t *testing.T) {
	t.Log(tool.ToPascal("dd_dd", "_"))
}

func TestFirstLetterToUpper(t *testing.T) {
	t.Log(tool.FirstLetterToUpper("gopher"))
	t.Log(tool.FirstLetterToUpper("Gopher"))
}

func TestToString(t *testing.T) {
	t.Log(tool.ToString(333))
	t.Log(tool.ToString("333"))

	stru := struct {
		Id   string
		Name string
	}{
		Id:   "999",
		Name: "gopher",
	}
	t.Log(tool.ToString(stru))
}

func TestIsBlank(t *testing.T) {
	t.Log(tool.IsBlank("gopher"))
	t.Log(tool.IsBlank(""))
	t.Log(tool.IsBlank(" "))
}

func TestIsNotBlank(t *testing.T) {
	t.Log(tool.IsNotBlank("gopher"))
	t.Log(tool.IsNotBlank(""))
	t.Log(tool.IsNotBlank(" "))
}

func TestXml2JsonString(t *testing.T) {
	x := `
	<xml>
	<data></data>
	</xml>
	`
	t.Log(tool.Xml2JsonString(x))
}
