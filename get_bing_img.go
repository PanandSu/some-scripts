package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
)

func getHtml(url string) ([]string, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	//defer close要与下文的goquery.NewDocumentFromReader在一个函数,否则就关闭了,啥也读取不到

	imgs := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return imgs, err
	}
	doc.Find("a[data-src]").Each(func(i int, s *goquery.Selection) {
		val, ok := s.Attr("data-src")
		if ok {
			imgs = append(imgs, val)
		}
	})
	return imgs, nil
}

func save(imgs []string) error {
	err := os.Chdir("/Users/pansu/Projects/GoLand/snippets/imgs")
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i, img := range imgs {
		file, err := os.Create(strconv.Itoa(i) + "." + "jpg")
		if err != nil {
			fmt.Println(err)
			continue
		}
		resp, err := http.Get(img)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		file.Close()
	}
	return nil
}

func init() {
	imgs, err := getHtml("https://dailybing.com/index1.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(imgs)
}

func main() {
	const PageNum = 10
	var allImgs []string
	for i := 0; i < PageNum; i++ {
		url := "https://dailybing.com/index" + strconv.Itoa(i) + ".html"
		pageImg, err := getHtml(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allImgs = append(allImgs, pageImg...)
	}
	fmt.Println(allImgs)

	err := save(allImgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("program run successfully")
}

/*
[https://dailybing.com/api/v1/20240701zh-cnFHD430 https://dailybing.com/api/v1/20240630zh-cnFHD430 https://dailybing.com/api/v1/20240629zh-cnFHD430 https://dailybing.com/api/v1/20240628zh-cnFHD430 https://dailybing.com/api/v1/20240627zh-cnFHD430 https://dailybing.com/api/v1/20240626zh-cnFHD430 https://dailybing.com/api/v1/20240625zh-cnFHD430 https://dailybing.com/api/v1/20240624zh-cnFHD430 https://dailybing.com/api/v1/20240623zh-cnFHD430 https://dailybing.com/api/v1/20240622zh-cnFHD430 https://dailybing.com/api/v1/20240621zh-cnFHD430 https://dailybing.com/api/v1/20240620zh-cnFHD430 https://dailybing.com/api/v1/20240619zh-cnFHD430 https://dailybing.com/api/v1/20240618zh-cnFHD430 https://dailybing.com/api/v1/20240617zh-cnFHD430 https://dailybing.com/api/v1/20240616zh-cnFHD430 https://dailybing.com/api/v1/20240615zh-cnFHD430 https://dailybing.com/api/v1/20240614zh-cnFHD430 https://dailybing.com/api/v1/20240613zh-cnFHD430 https://dailybing.com/api/v1/20240612zh-cnFHD430 https://dailybing.com/api/v1/20240611zh-cnFHD430 https://dailybing.com/api/v1/20240610zh-cnFHD430 https://dailybing.com/api/v1/20240609zh-cnFHD430 https://dailybing.com/api/v1/20240608zh-cnFHD430]
[https://dailybing.com/api/v1/20240701zh-cnFHD430 https://dailybing.com/api/v1/20240630zh-cnFHD430 https://dailybing.com/api/v1/20240629zh-cnFHD430 https://dailybing.com/api/v1/20240628zh-cnFHD430 https://dailybing.com/api/v1/20240627zh-cnFHD430 https://dailybing.com/api/v1/20240626zh-cnFHD430 https://dailybing.com/api/v1/20240625zh-cnFHD430 https://dailybing.com/api/v1/20240624zh-cnFHD430 https://dailybing.com/api/v1/20240623zh-cnFHD430 https://dailybing.com/api/v1/20240622zh-cnFHD430 https://dailybing.com/api/v1/20240621zh-cnFHD430 https://dailybing.com/api/v1/20240620zh-cnFHD430 https://dailybing.com/api/v1/20240619zh-cnFHD430 https://dailybing.com/api/v1/20240618zh-cnFHD430 https://dailybing.com/api/v1/20240617zh-cnFHD430 https://dailybing.com/api/v1/20240616zh-cnFHD430 https://dailybing.com/api/v1/20240615zh-cnFHD430 https://dailybing.com/api/v1/20240614zh-cnFHD430 https://dailybing.com/api/v1/20240613zh-cnFHD430 https://dailybing.com/api/v1/20240612zh-cnFHD430 https://dailybing.com/api/v1/20240611zh-cnFHD430 https://dailybing.com/api/v1/20240610zh-cnFHD430 https://dailybing.com/api/v1/20240609zh-cnFHD430 https://dailybing.com/api/v1/20240608zh-cnFHD430 https://dailybing.com/api/v1/20240701zh-cnFHD430 https://dailybing.com/api/v1/20240630zh-cnFHD430 https://dailybing.com/api/v1/20240629zh-cnFHD430 https://dailybing.com/api/v1/20240628zh-cnFHD430 https://dailybing.com/api/v1/20240627zh-cnFHD430 https://dailybing.com/api/v1/20240626zh-cnFHD430 https://dailybing.com/api/v1/20240625zh-cnFHD430 https://dailybing.com/api/v1/20240624zh-cnFHD430 https://dailybing.com/api/v1/20240623zh-cnFHD430 https://dailybing.com/api/v1/20240622zh-cnFHD430 https://dailybing.com/api/v1/20240621zh-cnFHD430 https://dailybing.com/api/v1/20240620zh-cnFHD430 https://dailybing.com/api/v1/20240619zh-cnFHD430 https://dailybing.com/api/v1/20240618zh-cnFHD430 https://dailybing.com/api/v1/20240617zh-cnFHD430 https://dailybing.com/api/v1/20240616zh-cnFHD430 https://dailybing.com/api/v1/20240615zh-cnFHD430 https://dailybing.com/api/v1/20240614zh-cnFHD430 https://dailybing.com/api/v1/20240613zh-cnFHD430 https://dailybing.com/api/v1/20240612zh-cnFHD430 https://dailybing.com/api/v1/20240611zh-cnFHD430 https://dailybing.com/api/v1/20240610zh-cnFHD430 https://dailybing.com/api/v1/20240609zh-cnFHD430 https://dailybing.com/api/v1/20240608zh-cnFHD430 https://dailybing.com/api/v1/20240606zh-cnFHD430 https://dailybing.com/api/v1/20240605zh-cnFHD430 https://dailybing.com/api/v1/20240604zh-cnFHD430 https://dailybing.com/api/v1/20240603zh-cnFHD430 https://dailybing.com/api/v1/20240602zh-cnFHD430 https://dailybing.com/api/v1/20240601zh-cnFHD430 https://dailybing.com/api/v1/20240531zh-cnFHD430 https://dailybing.com/api/v1/20240530zh-cnFHD430 https://dailybing.com/api/v1/20240529zh-cnFHD430 https://dailybing.com/api/v1/20240528zh-cnFHD430 https://dailybing.com/api/v1/20240527zh-cnFHD430 https://dailybing.com/api/v1/20240526zh-cnFHD430 https://dailybing.com/api/v1/20240525zh-cnFHD430 https://dailybing.com/api/v1/20240524zh-cnFHD430 https://dailybing.com/api/v1/20240523zh-cnFHD430 https://dailybing.com/api/v1/20240522zh-cnFHD430 https://dailybing.com/api/v1/20240521zh-cnFHD430 https://dailybing.com/api/v1/20240520zh-cnFHD430 https://dailybing.com/api/v1/20240519zh-cnFHD430 https://dailybing.com/api/v1/20240518zh-cnFHD430 https://dailybing.com/api/v1/20240517zh-cnFHD430 https://dailybing.com/api/v1/20240516zh-cnFHD430 https://dailybing.com/api/v1/20240515zh-cnFHD430 https://dailybing.com/api/v1/20240514zh-cnFHD430 https://dailybing.com/api/v1/20240512zh-cnFHD430 https://dailybing.com/api/v1/20240511zh-cnFHD430 https://dailybing.com/api/v1/20240510zh-cnFHD430 https://dailybing.com/api/v1/20240509zh-cnFHD430 https://dailybing.com/api/v1/20240508zh-cnFHD430 https://dailybing.com/api/v1/20240507zh-cnFHD430 https://dailybing.com/api/v1/20240506zh-cnFHD430 https://dailybing.com/api/v1/20240505zh-cnFHD430 https://dailybing.com/api/v1/20240504zh-cnFHD430 https://dailybing.com/api/v1/20240503zh-cnFHD430 https://dailybing.com/api/v1/20240502zh-cnFHD430 https://dailybing.com/api/v1/20240501zh-cnFHD430 https://dailybing.com/api/v1/20240430zh-cnFHD430 https://dailybing.com/api/v1/20240429zh-cnFHD430 https://dailybing.com/api/v1/20240428zh-cnFHD430 https://dailybing.com/api/v1/20240427zh-cnFHD430 https://dailybing.com/api/v1/20240426zh-cnFHD430 https://dailybing.com/api/v1/20240425zh-cnFHD430 https://dailybing.com/api/v1/20240424zh-cnFHD430 https://dailybing.com/api/v1/20240423zh-cnFHD430 https://dailybing.com/api/v1/20240422zh-cnFHD430 https://dailybing.com/api/v1/20240421zh-cnFHD430 https://dailybing.com/api/v1/20240420zh-cnFHD430 https://dailybing.com/api/v1/20240419zh-cnFHD430 https://dailybing.com/api/v1/20240417zh-cnFHD430 https://dailybing.com/api/v1/20240416zh-cnFHD430 https://dailybing.com/api/v1/20240415zh-cnFHD430 https://dailybing.com/api/v1/20240414zh-cnFHD430 https://dailybing.com/api/v1/20240413zh-cnFHD430 https://dailybing.com/api/v1/20240412zh-cnFHD430 https://dailybing.com/api/v1/20240411zh-cnFHD430 https://dailybing.com/api/v1/20240410zh-cnFHD430 https://dailybing.com/api/v1/20240409zh-cnFHD430 https://dailybing.com/api/v1/20240408zh-cnFHD430 https://dailybing.com/api/v1/20240407zh-cnFHD430 https://dailybing.com/api/v1/20240406zh-cnFHD430 https://dailybing.com/api/v1/20240405zh-cnFHD430 https://dailybing.com/api/v1/20240404zh-cnFHD430 https://dailybing.com/api/v1/20240403zh-cnFHD430 https://dailybing.com/api/v1/20240402zh-cnFHD430 https://dailybing.com/api/v1/20240401zh-cnFHD430 https://dailybing.com/api/v1/20240331zh-cnFHD430 https://dailybing.com/api/v1/20240330zh-cnFHD430 https://dailybing.com/api/v1/20240329zh-cnFHD430 https://dailybing.com/api/v1/20240328zh-cnFHD430 https://dailybing.com/api/v1/20240327zh-cnFHD430 https://dailybing.com/api/v1/20240326zh-cnFHD430 https://dailybing.com/api/v1/20240325zh-cnFHD430 https://dailybing.com/api/v1/20240323zh-cnFHD430 https://dailybing.com/api/v1/20240322zh-cnFHD430 https://dailybing.com/api/v1/20240321zh-cnFHD430 https://dailybing.com/api/v1/20240320zh-cnFHD430 https://dailybing.com/api/v1/20240319zh-cnFHD430 https://dailybing.com/api/v1/20240318zh-cnFHD430 https://dailybing.com/api/v1/20240317zh-cnFHD430 https://dailybing.com/api/v1/20240316zh-cnFHD430 https://dailybing.com/api/v1/20240315zh-cnFHD430 https://dailybing.com/api/v1/20240314zh-cnFHD430 https://dailybing.com/api/v1/20240313zh-cnFHD430 https://dailybing.com/api/v1/20240312zh-cnFHD430 https://dailybing.com/api/v1/20240311zh-cnFHD430 https://dailybing.com/api/v1/20240310zh-cnFHD430 https://dailybing.com/api/v1/20240309zh-cnFHD430 https://dailybing.com/api/v1/20240308zh-cnFHD430 https://dailybing.com/api/v1/20240307zh-cnFHD430 https://dailybing.com/api/v1/20240306zh-cnFHD430 https://dailybing.com/api/v1/20240305zh-cnFHD430 https://dailybing.com/api/v1/20240304zh-cnFHD430 https://dailybing.com/api/v1/20240303zh-cnFHD430 https://dailybing.com/api/v1/20240302zh-cnFHD430 https://dailybing.com/api/v1/20240301zh-cnFHD430 https://dailybing.com/api/v1/20240229zh-cnFHD430 https://dailybing.com/api/v1/20240227zh-cnFHD430 https://dailybing.com/api/v1/20240226zh-cnFHD430 https://dailybing.com/api/v1/20240225zh-cnFHD430 https://dailybing.com/api/v1/20240224zh-cnFHD430 https://dailybing.com/api/v1/20240223zh-cnFHD430 https://dailybing.com/api/v1/20240222zh-cnFHD430 https://dailybing.com/api/v1/20240221zh-cnFHD430 https://dailybing.com/api/v1/20240220zh-cnFHD430 https://dailybing.com/api/v1/20240219zh-cnFHD430 https://dailybing.com/api/v1/20240218zh-cnFHD430 https://dailybing.com/api/v1/20240217zh-cnFHD430 https://dailybing.com/api/v1/20240216zh-cnFHD430 https://dailybing.com/api/v1/20240215zh-cnFHD430 https://dailybing.com/api/v1/20240214zh-cnFHD430 https://dailybing.com/api/v1/20240213zh-cnFHD430 https://dailybing.com/api/v1/20240212zh-cnFHD430 https://dailybing.com/api/v1/20240211zh-cnFHD430 https://dailybing.com/api/v1/20240210zh-cnFHD430 https://dailybing.com/api/v1/20240209zh-cnFHD430 https://dailybing.com/api/v1/20240208zh-cnFHD430 https://dailybing.com/api/v1/20240207zh-cnFHD430 https://dailybing.com/api/v1/20240206zh-cnFHD430 https://dailybing.com/api/v1/20240205zh-cnFHD430 https://dailybing.com/api/v1/20240204zh-cnFHD430 https://dailybing.com/api/v1/20240202zh-cnFHD430 https://dailybing.com/api/v1/20240201zh-cnFHD430 https://dailybing.com/api/v1/20240131zh-cnFHD430 https://dailybing.com/api/v1/20240130zh-cnFHD430 https://dailybing.com/api/v1/20240129zh-cnFHD430 https://dailybing.com/api/v1/20240128zh-cnFHD430 https://dailybing.com/api/v1/20240127zh-cnFHD430 https://dailybing.com/api/v1/20240126zh-cnFHD430 https://dailybing.com/api/v1/20240125zh-cnFHD430 https://dailybing.com/api/v1/20240124zh-cnFHD430 https://dailybing.com/api/v1/20240123zh-cnFHD430 https://dailybing.com/api/v1/20240122zh-cnFHD430 https://dailybing.com/api/v1/20240121zh-cnFHD430 https://dailybing.com/api/v1/20240120zh-cnFHD430 https://dailybing.com/api/v1/20240119zh-cnFHD430 https://dailybing.com/api/v1/20240118zh-cnFHD430 https://dailybing.com/api/v1/20240117zh-cnFHD430 https://dailybing.com/api/v1/20240116zh-cnFHD430 https://dailybing.com/api/v1/20240115zh-cnFHD430 https://dailybing.com/api/v1/20240114zh-cnFHD430 https://dailybing.com/api/v1/20240113zh-cnFHD430 https://dailybing.com/api/v1/20240112zh-cnFHD430 https://dailybing.com/api/v1/20240111zh-cnFHD430 https://dailybing.com/api/v1/20240110zh-cnFHD430 https://dailybing.com/api/v1/20240108zh-cnFHD430 https://dailybing.com/api/v1/20240107zh-cnFHD430 https://dailybing.com/api/v1/20240106zh-cnFHD430 https://dailybing.com/api/v1/20240105zh-cnFHD430 https://dailybing.com/api/v1/20240104zh-cnFHD430 https://dailybing.com/api/v1/20240103zh-cnFHD430 https://dailybing.com/api/v1/20240102zh-cnFHD430 https://dailybing.com/api/v1/20240101zh-cnFHD430 https://dailybing.com/api/v1/20231231zh-cnFHD430 https://dailybing.com/api/v1/20231230zh-cnFHD430 https://dailybing.com/api/v1/20231229zh-cnFHD430 https://dailybing.com/api/v1/20231228zh-cnFHD430 https://dailybing.com/api/v1/20231227zh-cnFHD430 https://dailybing.com/api/v1/20231226zh-cnFHD430 https://dailybing.com/api/v1/20231225zh-cnFHD430 https://dailybing.com/api/v1/20231224zh-cnFHD430 https://dailybing.com/api/v1/20231223zh-cnFHD430 https://dailybing.com/api/v1/20231222zh-cnFHD430 https://dailybing.com/api/v1/20231221zh-cnFHD430 https://dailybing.com/api/v1/20231220zh-cnFHD430 https://dailybing.com/api/v1/20231219zh-cnFHD430 https://dailybing.com/api/v1/20231218zh-cnFHD430 https://dailybing.com/api/v1/20231217zh-cnFHD430 https://dailybing.com/api/v1/20231216zh-cnFHD430 https://dailybing.com/api/v1/20231214zh-cnFHD430 https://dailybing.com/api/v1/20231213zh-cnFHD430 https://dailybing.com/api/v1/20231212zh-cnFHD430 https://dailybing.com/api/v1/20231211zh-cnFHD430 https://dailybing.com/api/v1/20231210zh-cnFHD430 https://dailybing.com/api/v1/20231209zh-cnFHD430 https://dailybing.com/api/v1/20231208zh-cnFHD430 https://dailybing.com/api/v1/20231207zh-cnFHD430 https://dailybing.com/api/v1/20231206zh-cnFHD430 https://dailybing.com/api/v1/20231205zh-cnFHD430 https://dailybing.com/api/v1/20231204zh-cnFHD430 https://dailybing.com/api/v1/20231203zh-cnFHD430 https://dailybing.com/api/v1/20231202zh-cnFHD430 https://dailybing.com/api/v1/20231201zh-cnFHD430 https://dailybing.com/api/v1/20231130zh-cnFHD430 https://dailybing.com/api/v1/20231129zh-cnFHD430 https://dailybing.com/api/v1/20231128zh-cnFHD430 https://dailybing.com/api/v1/20231127zh-cnFHD430 https://dailybing.com/api/v1/20231126zh-cnFHD430 https://dailybing.com/api/v1/20231125zh-cnFHD430 https://dailybing.com/api/v1/20231124zh-cnFHD430 https://dailybing.com/api/v1/20231123zh-cnFHD430 https://dailybing.com/api/v1/20231122zh-cnFHD430 https://dailybing.com/api/v1/20231121zh-cnFHD430]
program run successfully
*/
