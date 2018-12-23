package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tebeka/selenium"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type TouTiao struct {
	Id      int
	Words   string
	Company string
	Rate    int
}

func GrabToutiaoProcess() []TouTiao {
	var qs []TouTiao

	r := NewCrawler("PC")
	defer r.ReleaseCrawler()

	v, _ := r.crawlWeiboReSou()
	qs = append(qs, v...)

	v, _ = r.crawlBaiduReSou()
	qs = append(qs, v...)

	v, _ = r.crawlToutiaoReSou()
	qs = append(qs, v...)

	return qs
}

////////////

func (r *GoCrawler) crawlWeiboReSou() ([]TouTiao, bool) {
	url := "https://s.weibo.com/top/summary?Refer=top_hot&topnav=1&wvr=6"
	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "a",
		Value: "www",
	})
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		return nil, false //"0 Failed to load page"
	}

	//       source,_ := r.webDriver.PageSource()
	melem, err := r.webDriver.FindElement(selenium.ByID, "pl_top_realtimehot")
	if err != nil {
		logs.Error(err)
		return nil, false
	} else {
		var retv = []TouTiao{} //   make(map[int]TouTiao, 30)
		src, _ := melem.Text()
		sv := strings.Split(src, "\n")
		if len(sv) > 10 {
			idx := 0
			inum := 1
			for _, s := range sv {
				idx++
				if idx < 3 {
					continue
				}

				ssv := strings.Split(s, " ")
				{
					itm := len(ssv)
					if (itm > 3) && ((ssv[itm-1] == "热") || (ssv[itm-1] == "沸") || (ssv[itm-1] == "新")) {
						sname := ""
						for j := 1; j < itm-2; j++ {
							sname += ssv[j]
						}

						//	fmt.Println(ssv[0], sname, " Hot=", ssv[itm-2])
						value := 0
						value, _ = strconv.Atoi(ssv[itm-2])
						ti := TouTiao{inum, sname, "WEIBO", value}
						inum++
						retv = append(retv, ti)
						if len(retv) == 30 {
							break
						}

					} else if (itm > 3) && (ssv[itm-1] == "荐") {
						continue
					} else {
						sname := ""
						for j := 1; j < itm-1; j++ {
							sname += ssv[j]
						}
						//	fmt.Println(ssv[0], sname, " Hot=", ssv[itm-1])
						value := 0
						value, _ = strconv.Atoi(ssv[itm-1])
						ti := TouTiao{inum, sname, "WEIBO", value}
						inum++
						retv = append(retv, ti)
						if len(retv) == 30 {
							break
						}

					}
				}
				//	fmt.Println(s)
			}
		}
		return retv, true
	}
	return nil, false
}

func (r *GoCrawler) crawlBaiduReSou() ([]TouTiao, bool) {
	url := "http://top.baidu.com/buzz?b=1&fr=20811"
	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "a",
		Value: "www",
	})
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		return nil, false //"0 Failed to load page"
	}
	/*
		source,_ := r.webDriver.PageSource()
				fmt.Println(source)
						return ""
	*/
	melem, err := r.webDriver.FindElement(selenium.ByClassName, "list-table")
	if err != nil {
		logs.Error(err)
		return nil, false // ""
	} else {
		retv := []TouTiao{}
		// inum := 1
		s, _ := melem.Text()
		sv := strings.Split(s, "\n")
		if len(sv) > 200 {

			idnow := 0
			for idx := 1; idx <= 30; idx++ {
				ids := fmt.Sprintf("%d", idx)
				for id := idnow; id < len(sv); id++ {
					if ids == sv[id] {
						value, _ := strconv.Atoi(sv[id+3])
						ti := TouTiao{idx, sv[id+1], "BAIDU", value}
						retv = append(retv, ti)
						idnow = id + 3
						break
					}
				}

			}
			/*
				value := 0
				value, _ = strconv.Atoi(sv[4])
				ti := TouTiao{inum, sv[2], "BAIDU", value}
				retv = append(retv, ti)
				inum++
				value, _ = strconv.Atoi(sv[10])
				ti = TouTiao{inum, sv[8], "BAIDU", value}
				retv = append(retv, ti)
				inum++
				value, _ = strconv.Atoi(sv[16])
				ti = TouTiao{inum, sv[14], "BAIDU", value}
				retv = append(retv, ti)
				inum++
				for i := 0; i < 27; i++ {
					//			fmt.Println(sv[19+i*4], ":", sv[20+i*4], ",", sv[22+i*4])
					value, _ = strconv.Atoi(sv[22+i*4])
					ti = TouTiao{inum, sv[20+i*4], "BAIDU", value}
					retv = append(retv, ti)
					inum++
				}
			*/

		}
		return retv, true
	}
	return nil, false
}

func (r *GoCrawler) crawlToutiaoReSou() ([]TouTiao, bool) {
	url := "https://is.snssdk.com/2/wap/search/extra/hot_word_list/"
	r.webDriver.AddCookie(&selenium.Cookie{
		Name:  "a",
		Value: "www",
	})
	err := r.webDriver.Get(url)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to load page: %s\n", err))
		return nil, false //"0 Failed to load page"
	}
	/*
		source,_ := r.webDriver.PageSource()
				fmt.Println(source)
						return ""
	*/
	melem, err := r.webDriver.FindElement(selenium.ByClassName, "list")
	if err != nil {
		logs.Error(err)
		return nil, false //  ""
	} else {
		retv := []TouTiao{}
		inum := 1
		s, _ := melem.Text()
		sv := strings.Split(s, "\n")
		if len(sv) == 60 {
			for i := 0; i < 30; i++ {
				// fmt.Println(sv[i*2], ":", sv[i*2+1])
				value := 0
				ti := TouTiao{inum, sv[i*2+1], "TOUTIAO", value}
				retv = append(retv, ti)
				inum++
			}
		}
		return retv, true
	}
	return nil, false
}

func FetchKeyWords() string {
	key := time.Now().Format("2006-01-02") + ":key"
	if v, b := GetKVStore("RESOU", key); b {
		//	fmt.Println(v)
		return v
	} else {
		return ""
	}
}

func PickKeyWords(tts []TouTiao) {
	fLEVEL := 0.88
	s1, s2, s3 := "", "", ""
	for _, v := range tts {
		if v.Company == "BAIDU" {
			s1 = s1 + ";" + v.Words
		}
		if v.Company == "TOUTIAO" {
			s2 = s2 + ";" + v.Words
		}
		if v.Company == "WEIBO" {
			s3 = s3 + ";" + v.Words
		}
	}
	/*
	  1. get keyword by weibo
	  2. get keyword by toutiao
	  3. get keyword by baidu
	  4. merge and pick  [hot & big]
	*/
	retv := map[string]float64{}
	//	retv, retv1 := map[string]float64{}
	if retv1, b := PullwordGet(s1, 0, 1, fLEVEL); b {
		// merge to retv
		for k1, v1 := range retv1 {
			//	fmt.Println(k1, ":", v1)
			retv[k1] = v1
		}
	}

	if retv1, b := PullwordGet(s2, 0, 1, fLEVEL); b {
		// merge to retv
		for k1, v1 := range retv1 {
			//	fmt.Println(k1, ":", v1)
			if _, b := retv[k1]; !b {
				retv[k1] = v1
			} else {
				retv[k1] += v1
			}
		}
	}

	if retv1, b := PullwordGet(s3, 0, 1, fLEVEL); b {
		// mergo to retv
		for k1, v1 := range retv1 {
			//	fmt.Println(k1, ":", v1)
			if _, b := retv[k1]; !b {
				retv[k1] = v1
			} else {
				retv[k1] += v1
			}
		}
	}
	///      print
	words := ""
	rand.Seed(time.Now().UnixNano())

	for k1, v1 := range retv {
		//	if len(k1) > 6 || v1 >= 2 {
		if len(k1) > 6 || v1 >= 3 {
			x := rand.Intn(6)
			spanclass := ""
			switch x {
			case 0:
				spanclass = "tipsA"
			case 1:
				spanclass = "tipsB"
			case 2:
				spanclass = "tipsC"
			case 3:
				spanclass = "tipsD"
			case 4:
				spanclass = "tipsE"
			case 5:
				spanclass = "tipsF"
			}
			if len(words) == 0 {
				words = "<span class=\"" + spanclass + "\">" + k1 + "</span>"
			} else {
				words = words + "<span class=\"" + spanclass + "\">" + k1 + "</span>"
			}
		}
	}
	// fmt.Println(words)
	// save to KVStore
	key := time.Now().Format("2006-01-02") + ":key"
	SetKVStore("RESOU", key, words)

}
