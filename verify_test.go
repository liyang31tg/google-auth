package googleAuth

import (
	"testing"
)

func TestVerify(t *testing.T) {
	secretId := "rtbhtWNgMaFuJFuxIAQzSrVKdtWzfFCP"
	result := ReturnCode(secretId)
	t.Log(result)
	t.Error("err")
}
