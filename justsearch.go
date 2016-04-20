package main
import (
	"flag"
	"github.com/golang/glog"
	"github.com/openedinc/opened-go"
)

func init() {
  flag.Parse()
}  

func main() {
	glog.V(1).Infof("My OpenEd Resource Find Script")
token,err := opened.GetToken ("","","","")
  query_params:=make(map[string]string)
  query_params["contribution_name"]="LearningRegistry"

  results,err:=opened.SearchResources(query_params,token)
  if err!=nil {
    glog.Errorf("Error from SearchResources,%+v",err)    
  } 
  glog.V(1).Infof("%d results returned",len(results.Resources))
  for _,resource:= range results.Resources {
    glog.V(2).Infof("Resource Title: %s (%d)",resource.Title,resource.Id)  
    
  }

}