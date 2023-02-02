package router

import (
	"fmt"
	"strings"
	"testing"
)

func init() {

}

func Test_match(t *testing.T) {
	match("/hello/test")
}

func TestMatch(t *testing.T) {
	url, r, err := Match("/hello/hahah")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
	t.Log(r)
}

func TestTirm(t *testing.T) {
	print(strings.TrimPrefix("/abc/x", "/abc"))
}

func TestMain(m *testing.M) {
	InitRouter()
	m.Run()
	fmt.Println("end")
}
