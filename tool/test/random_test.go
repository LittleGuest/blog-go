package tool

import (
	"blog/tool"
	"testing"
)

func TestRandomNumbersToString(t *testing.T) {
	t.Logf("随机：%s", tool.RandomNumbersToString(-1))
	t.Logf("随机：%s", tool.RandomNumbersToString(0))
	t.Logf("随机：%s", tool.RandomNumbersToString(6))
}
