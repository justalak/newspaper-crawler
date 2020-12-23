package main

import(
    "gopkg.in/olivere/elastic.v5"
    "golang.org/x/net/context"
    "log"
    "os"
    "encoding/json"
)


var client *elastic.Client
var err error
func init(){
    client,err = elastic.NewClient()
    if err!=nil{
        log.Fatal(err)
    }
}

func main() {
    var data []string

    file,err := os.Open("vietnamenet/vietnamnet.json")
    if err!=nil{
        log.Fatal(err)
    }
    defer file.Close()

    jsonDeocder :=  json.NewDecoder(file)
    if err := jsonDeocder.Decode(&data); err!=nil{
        log.Fatal("Decode: ",err)
    }

    bulkIndex("library","article",data)
}

func bulkIndex(index string,typ string ,data []string){
    ctx := context.Background()
    for _,item := range data{
        _,err := client.Index().Index(index).Type(typ).BodyJson(item).Do(ctx)   
        if err !=nil{
            log.Fatal(err)
        }
    }   
}
