package alert_test

import (
	"fmt"
	"testing"

	"github.com/fcjr/alert"
)

func TestExampleInfo(t *testing.T) {
	_ = alert.Info("Example Message", "Example Message Text")
	_ = alert.Info("Example 中文标题", "Example 中文消息")
	_ = alert.Info("Example 테스트 메시지", "Example 테스트 메시지")
}
func TestExampleError(t *testing.T) {
	_ = alert.Error("Example Error", "Example Error Text")
}

func TestExampleQuestion(t *testing.T) {
	resp, _ := alert.Question("Example Question", "Yes or No?", "Yes", "No")
	fmt.Printf("Response: %t\n", resp)

	// Returns true iff 'Yes' is pressed
}
