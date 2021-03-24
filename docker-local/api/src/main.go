package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	mecab "github.com/bluele/mecab-golang"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ichi-infra-challenge/docker-local/api/src/model"
	"github.com/saintfish/chardet"
)

const (
	ipadic = "/usr/local/lib/mecab/dic/mecab-ipadic-neologd"
)

func main() {
	// Start HTTP server
	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost",
			//"*",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Origin",
			"Authorization",
			"Accept",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("/search", searchExample)
	r.GET("/create", createData)
	r.GET("/analyze", morphologicalAnalyze)
	r.GET("/scrape", scrapeText)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

//あいまい文字列検索
func searchExample(c *gin.Context) {
	searchTerm := c.Query("search")
	if searchTerm == "" {
		c.JSON(http.StatusInternalServerError, "search not specified")
		return
	}

	resp, err := model.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func createData(c *gin.Context) {

	name := c.Query("search")
	resp, err := model.Create(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func morphologicalAnalyze(c *gin.Context) {

	text := c.Query("text")
	// mecab here
	mecab, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer mecab.Destroy()
	data := parseToNode(mecab, text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	// cData := make([]map[string]interface{}, 0)

	// for ; !node.IsZero(); node = node.Next() {
	// 	entry2 := make(map[string]interface{})
	// 	entry2[node.Surface()] = node.Feature()
	// 	cData.append(cData, entry2)
	// }
	c.IndentedJSON(http.StatusOK, data)
	//c.HTML(200, "index.html", gin.H{
	//	"data": data,
	//})
}

func parseToNode(m *mecab.MeCab, text string) []map[string]interface{} {
	// var key string
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(text)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)

	data := make([]map[string]interface{}, 0)

	for {
		entry := make(map[string]interface{})
		fmt.Printf("%s\t%s\n", node.Surface(), node.Feature())
		entry[node.Surface()] = node.Feature()
		data = append(data, entry)
		// features := strings.Split(node.Feature(), ",")
		// if features[0] != "助詞" && features[0] != "助動詞" {

		//     key = key + " " + features[7]
		// }
		if node.Next() != nil {
			break
		}
	}
	fmt.Println(data)
	return data
}

type Title struct {
	Title string
	Time string
}

func scrapeText(c *gin.Context) {
	var urls []string
	urls = strings.Split(c.Query("urls"), ",")

	fmt.Println(urls)
	//syncでURLの数だけ待ち合わせ
	var wg sync.WaitGroup
	wg.Add(len(urls))

	//titlesを返すresult
	var results []Title

	//titleをfetchする関数定義
	fetchTitle := func(url string) {
		start := time.Now()
		defer wg.Done()
		// Getリクエスト
		res, _ := http.Get(url)
		defer res.Body.Close()

		// 読み取り
		buffer, _ := ioutil.ReadAll(res.Body)

		// 文字コード判定
		detector := chardet.NewTextDetector()
		detectResult, _ := detector.DetectBest(buffer)
		fmt.Println(detectResult.Charset)

		// 文字コード変換
		bufferReader := bytes.NewReader(buffer)
		reader, _ := charset.NewReaderLabel(detectResult.Charset, bufferReader)

		// HTMLパース
		document, _ := goquery.NewDocumentFromReader(reader)

		// titleを抜き出し
		title := document.Find("title").Text()
		end := time.Now();
		time := (end.Sub(start)).Seconds()
		result := Title{title, strconv.FormatFloat(time, 'f', -1, 64) }
		results= append(results, result)
		fmt.Println(results)

	}

	//urlsの回数分スレッド実行
	for _, url := range urls {
		go fetchTitle(url)
	}

	//fetchTitle goroutineが終わるまで、wg.Wait()で待つ
	wg.Wait()

	c.JSON(http.StatusOK, results)

}