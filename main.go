package main

import (
	"StarWarsAPI_GO/models"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"net/http"
)


func main(){
	request:="planets/1"

	safeRequest:=url.QueryEscape(request)

	url:=fmt.Sprintf("http://swapi.co/api/%s",safeRequest)

	//Build Request

	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil{
		log.Fatal("NewRequest: ", err)
		return
	}

	client:=http.Client{}

  // Sen the Request using the client.
  resp,err:=client.Do(req)
  if err!=nil{
  	log.Fatal("Do: ", err)
  	return
  }

defer resp.Body.Close()

var record models.Planets

if err:=json.NewDecoder(resp.Body).Decode(&record); err!=nil{
	log.Println(err)
}

fmt.Println("Hi I'm D4R5 from",record.Name)
fmt.Println("Our Planet's Gravity is",record.Gravity)
fmt.Println("Our Planet's population is",record.Population)
}



