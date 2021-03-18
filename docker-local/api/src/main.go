package main

import (
	"fmt"
	"log"
	"net/http"

	// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
	// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
	mecab "github.com/bluele/mecab-golang"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ichi-infra-challenge/docker-local/api/src/model"
)

const (
	ipadic = "/usr/local/lib/mecab/dic/mecab-ipadic-neologd"
)

func main() {
	// Start HTTP server
	r := gin.Default()
	r.GET("/search", searchExample)
	r.GET("/create", createData)
	r.POST("/analyze", morphologicalAnalyze)

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

	text := c.PostForm("text")
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
