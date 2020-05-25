package alert

import "testing"

func Test_sendMsg(t *testing.T) {
	if err:=sendMsg("this is a test");err!=nil{
		t.Fatal(err)
	}
}
