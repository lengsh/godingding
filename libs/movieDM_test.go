package libs

import (
	"fmt"
	"github.com/lengsh/godingding/libs"
	"testing"
)

func init() {
	/*
		logs.SetLogger(logs.AdapterFile, `{"filename":"./test.log","maxlines":10000,"maxsize":102400,"daily":true,"maxdays":2}`)
		logs.EnableFuncCallDepth(true)
		logs.SetLogFuncCallDepth(3)
	*/

}

func TestNewMovie(t *testing.T) {
	/*
		st := "2018-09-25"
		mv := libs.Movie{Company: "iqiyi", Name: "ceshi", Rate: 1.1, Releasetime: st}
		n := mv.NewMovie()
		if n == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("error to insert :", n)
		}
	*/
}

func TestRNewMovie(t *testing.T) {
	/*
		st := "2018-09-25"
		mv := libs.Movie{Company: "iqiyi", Name: "西红柿首富", Rate: 3.1, Releasetime: st}
		n := mv.NewMovie()
		if n == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("error to insert :", n)
		}

		mv = libs.Movie{Company: "youku", Name: "西红柿首富", Rate: 210.1, Releasetime: st}
		n = mv.NewMovie()
		if n == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("error to insert :", n)
		}
	*/
}

func TestQueryMovie(t *testing.T) {
	/*
		qs1 := libs.QueryTopMovies("IQIYI", 10)
		qs2 := libs.QueryTopMovies("TX", 10)

		var qs []libs.Movie

		qs = append(qs, qs1...)
		qs = append(qs, qs2...)
	*/

	libs.NewMylog("sys", "Update")
	fmt.Println(libs.GetLastLog())
}

func TestQueryTopMovie(t *testing.T) {
	/*
		qs := libs.QueryTopMovies("IQIYI", 10)
		if qs != nil {
			for _, s := range qs {
				fmt.Println("Id = ", s.Id)
				fmt.Printf(s.String())
			}
		}
	*/
}
