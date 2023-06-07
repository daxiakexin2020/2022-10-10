package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strconv"
	"strings"
)

type es struct {
	c *elasticsearch.Client
}

type TestDoc struct {
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
	salary  int    `json:"salary"`
	Address string `json:"address"`
}

var GES *es

func main() {
	c, err := client()
	log.Println("client es err:", err)
	if err != nil {
		return
	}
	GES = &es{c: c}
	handle()
}

func handle() {
	indexName := "test01"
	//GES.deleteIndex(indexName)
	//	GES.createIndex(indexName)
	//	GES.getIndex(indexName)
	//GES.insertDocs(indexName)
	//GES.getIndex(indexName)
	GES.getDocs(indexName)
	GES.deleteDocs(indexName)
	GES.getDocs(indexName)
}

func (e *es) createIndex(name string) {
	create, err := e.c.API.Indices.Create(name)
	log.Printf("create index result:%v,err:%v\n", create, err)
}

func (e *es) getIndex(name string) {
	get, err := e.c.API.Indices.Get([]string{name})
	log.Printf("get index result:%v,err:%v\n", get, err)
}

func (e *es) deleteIndex(name string) {
	response, err := e.c.API.Indices.Delete([]string{name})
	log.Printf("deleete index result:%v,err:%v", response, err)
}

func (e *es) insertDocs(index string) {
	t1 := &TestDoc{Name: "zz", Age: 32, salary: 1, Address: "baotou"}
	t2 := &TestDoc{Name: "kx", Age: 31, salary: 2, Address: "baotou"}
	marshal1, _ := json.Marshal(t1)
	marshal2, _ := json.Marshal(t2)
	data := []string{string(marshal1), string(marshal2)}
	for i, doc := range data {
		req := esapi.IndexRequest{
			Index:      index,
			DocumentID: strconv.Itoa(i),
			Body:       strings.NewReader(doc),
			Refresh:    "true",
		}
		do, err := req.Do(context.Background(), e.c)
		log.Printf("req do result:%v,err:%v\n", do, err)
		defer do.Body.Close()
		if err == nil {
			var resMap map[string]interface{}
			if err := json.NewDecoder(do.Body).Decode(&resMap); err != nil {
				log.Printf("json decode err:%v\n", err)
			} else {
				log.Printf("\nIndexRequest() RESPONSE:")
				fmt.Println("Status:", do.Status())
				fmt.Println("Result:", resMap["result"])
				fmt.Println("Version:", int(resMap["_version"].(float64)))
				fmt.Println("resMap:", resMap)
				fmt.Println("\n")
			}
		}
	}
}

func (e *es) getDocs(index string) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			//"match": map[string]interface{}{
			//	"name": "zz",
			//},
			"match_all": map[string]interface{}{},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("json encode err:%v\n", err)
	}

	res, err := e.c.Search(
		e.c.Search.WithContext(context.Background()),
		e.c.Search.WithIndex(index),
		e.c.Search.WithBody(&buf),
		e.c.Search.WithTrackTotalHits(true),
		e.c.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return
	}

	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %v\n", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		data := hit.(map[string]interface{})["_source"]
		marshal, _ := json.Marshal(data)
		td := &TestDoc{}
		json.Unmarshal(marshal, td)
		log.Printf("Unmarshal data:%+v\n", td)
	}
}

func (e *es) deleteDocs(index string) {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: "1",
	}
	_, err := req.Do(context.Background(), e.c)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
}

func client() (*elasticsearch.Client, error) {
	c, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return c, nil
}
