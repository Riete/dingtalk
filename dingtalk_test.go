package dingtalk

import (
	"testing"
)

func TestSender(t *testing.T) {
	url := "xxx"
	secret := "xxx"
	SendDingTalkText("test", "test", url, secret, false)
	SendDingTalkMarkdown("test", "test", url, secret, false)
}
