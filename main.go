package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type AgentStateUpdate struct {
  Name string
  State string
  Url string
}

type AgentStateResult struct {
	Update AgentStateUpdate `json:"agent-state-update"`
}

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Static assets
	r.StaticFile("/", "./assets/index.html")
	r.Static("/assets", "./assets")

	// Status
	r.GET("/status", func(c *gin.Context) {
		var username string = "admin"
		var passwd string = "opencast"
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://develop.opencast.org/capture-admin/agents/test.json", nil)
		req.SetBasicAuth(username, passwd)
		resp, err := client.Do(req)
		if err != nil{
			log.Fatal(err)
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		s := string(bodyText)
		var result AgentStateResult
		json.Unmarshal([]byte(s), &result)

		capturing, _ := json.Marshal(result.Update.State == "capturing")
		c.String(http.StatusOK, string(capturing))
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
