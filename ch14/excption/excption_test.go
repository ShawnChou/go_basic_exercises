package excption_test

import (
	"errors"
	"os"
	"testing"
)

func EchoInt(i interface{}) (string, error) {
	if _, ok := i.(int); !ok {
		return "haha", errors.New("error!!!")
	}
	return "hoho", nil
}

func TestError(t *testing.T) {
	if i, err := EchoInt("asdf"); err != nil {
		//t.Error(err)
		t.Log(err)
	} else {
		t.Log(i)
	}
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("lala ", err)
		}
	}()
	//panic("haha")
	os.Exit(-1)
}
