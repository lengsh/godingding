package libs

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/lengsh/godingding/libs"
	"testing"
	"time"
)

func TestFindKeyWords(t *testing.T) {
	var qy []libs.TouTiao
	key := time.Now().Format("2006-01-02")

	if value, ret := libs.GetKVStore("RESOU", key); ret {
		err := json.Unmarshal([]byte(value), &qy)
		if err != nil {
			logs.Error("Unmarshal failed, ", err)
			qy = nil
		}

	}
	libs.PickKeyWords(qy)
	fmt.Println(libs.FetchKeyWords())
}
