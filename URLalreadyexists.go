package main
import (
	"flag"
	//"github.com/golang/glog"
	"github.com/openedinc/opened-go"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
  flag.Parse()
}  

func main() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "url=hhttps://www.youtube.com/watch?v=wL4hICyMLKU", nil)
	token,err := opened.GetToken ("","","","")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
