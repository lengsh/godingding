package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
	"time"
)

func TestGetKVStore(t *testing.T) {
	k := time.Now().Format("2006-01-02")
	v, e := libs.GetKVStore("TT", k)
	fmt.Println("key = ", k, ", v=", v, ", e=", e)

}
func TestSetKVStore(t *testing.T) {
	k := time.Now().Format("2006-01-02")
	s := time.Now().Format("2006-01-02 15:04:05")
	qs := libs.SetKVStore("TT", k, s)
	fmt.Println("value = ", qs)
}
