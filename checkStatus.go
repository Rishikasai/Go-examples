package main
import (
	"fmt"
	 "net/http"
)

func main(){

	Sites:=[]string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http:/hi.com", // to see error
	}

	for _,link:=range Sites{
		checkStatus(link)

	}
}

func checkStatus(link string){
	_,err:=http.Get(link)
	if err!=nil{
		fmt.Println(link, err)
		return
	}
	fmt.Println(link, "is working")

}
