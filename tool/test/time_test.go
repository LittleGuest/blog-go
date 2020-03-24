package tool

import (
	"blog/tool"
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	t.Log(tool.FormatDate(time.Now()))
}

func TestFormatDateTime(t *testing.T) {
	t.Log(tool.FormatDateTime(time.Now()))
}

func TestStringToTime(t *testing.T) {
	timeStr := "2006-01-02 15:04:05"
	stt, _ := tool.StringToTime(timeStr)
	t.Log(stt)
}
