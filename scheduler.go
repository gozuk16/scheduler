package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

//var log string = "logだよ"

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 10s", MyJob{})
	jobrunner.Schedule("@every 5s", DailyJob{})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/jobrunner/status", JobJSON)

	// Load template file location relative to the current working directory
	// r.LoadHTMLGlob("../../bamzi/jobrunner/views/*.html")
	r.LoadHTMLGlob("views/*.html")

	// Returns html page at given endpoint based on the loaded
	// template from above
	r.GET("/jobrunner/html", JobHtml)

	r.Run(":8080")
}

// JobJSON ...
func JobJSON(c *gin.Context) {
	c.JSON(http.StatusOK, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {
	// Returns the template data pre-parsed
	//c.HTML(200, "", jobrunner.StatusPage())
	c.HTML(200, "Status-new.html", jobrunner.StatusPage())
	//c.HTML(200, "Status-new.html", gin.H{
	//    "a": "aaaa",
	//})
}

// MyJob ...
type MyJob struct {
}

// Run ...
func (e MyJob) Run() {
	fmt.Println("Run MyJob!")
	cmd := exec.Command("ls", "-al")
	out, _ := cmd.Output()
	fmt.Printf("%s", out)
	//    log = out
}

type DailyJob struct {
}

// Run ...
func (e DailyJob) Run() {
	fmt.Println("Run DailyJob!")
	// todo: validation path
	path := `jobs/daily/.`
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			//fmt.Println(p)
			procStart(p)
		}
		return nil
	})
	if err != nil {
		fmt.Println("job dir search err.")
		log.Fatal(err)
	}
}

// Process Start
func procStart(path string) {
	fmt.Println(path)

	// Todo: validate path
	openJobConf(path)

	// proc start
	out, err := exec.Command("ls", "bbb").CombinedOutput()
	if err != nil {
		log.Println("exec error.")
	}
	log.Printf("%s", string(out))

}

// open job file
func openJobConf(path string) {
	fmt.Println(path)

	// Todo: validate path



}


