package tool

import (
	"blog/tool"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	//t.Log(CreateDirectory("test"))
	t.Log(tool.CreateDirectory("test/test"))
	//t.Log(CreateDirectory("test/test.log"))
}

func TestCreateFile(t *testing.T) {
	t.Log(tool.CreateFile("test1"))
	t.Log(tool.CreateFile("test.log"))
}

func TestGetFileSuffix(t *testing.T) {
	t.Log(tool.GetFileSuffix("test.log"))
	t.Log(tool.GetFileSuffix("test.log.png"))
}

func TestIsExist(t *testing.T) {
	t.Log(tool.IsExist("test"))
	t.Log(tool.IsExist("test-file"))
	t.Log(tool.IsExist("test/test.log"))
	t.Log(tool.IsExist("test/test.txt"))
}
