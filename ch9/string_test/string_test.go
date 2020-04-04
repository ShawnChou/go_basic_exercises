package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))
	t.Logf("%x", c[0])
	t.Log(s)

}

func TestStringLoop(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		//都已c的第一个参数为值
		t.Logf("%[1]c %[1]d", c)
	}
}
