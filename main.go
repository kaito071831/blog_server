package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/kaito071831/blog_server/models"
	"github.com/kaito071831/blog_server/schema"
)


func main(){
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		var b models.BlogData

		if r.Method == "OPTIONS" {
		} else if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			w.WriteHeader(400)
			return
		}

		schema, err := graphql.NewSchema(schema.Schema)
		if err != nil {
			log.Fatalf("スキーマの取得に失敗しました: %v", err)
		}

		result := graphql.Do(graphql.Params{
			Context: r.Context(),
			Schema: schema,
			RequestString: b.Query,
			OperationName: b.Operation,
		})

		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("結果の書き込みに失敗しました: %v", err)
		}

	})

	fmt.Println("Listening on :8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalln(err)
	}
}