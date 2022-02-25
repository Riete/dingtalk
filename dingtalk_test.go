package dingtalk

import (
	"testing"
)

func TestSender(t *testing.T) {
	dt := NewDingTalk("xxx", "xxx")
	dt.SendText("xxx", "xx", false, "xxx")
	dt.SendMarkdown("xxx", "xx", false, "xxx")

}
