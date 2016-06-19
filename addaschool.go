package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/openedinc/opened-go"
	"encoding/json"
)

func main() {
    


params := map[string]string{
  "nces_id":"BB981479",
  "name":"Woodlands Christian Academy",
  "address":"5900 Academy Way",
  "city": "The Woodlands",
  "state": "TX",
  "zip": "77384",
  "low_grade": "K",
  "high_grade": "12",
}
jsonParams, _ := json.Marshal(params)
client := &http.Client{}
req, _ := http.NewRequest("POST", "https://partner.opened.com/1/schools", &jsonParams)
req.Header.Add("Content-Type", `W/"application/json"`)
token,err := opened.GetToken ("","","","") //needed
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
