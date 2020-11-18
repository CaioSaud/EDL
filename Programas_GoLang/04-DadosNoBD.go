//sem usar gochannels e goroutines, apenas salvando em batch

package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type produto struct {
	ID   string
	Cnpj string
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("error connecting to mongo: %s", err.Error())
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("")
	}

	defer client.Disconnect(context.Background())

	database := client.Database("admin")

	f, err := os.Open("dump_compliance.csv")
	if err != nil {
		log.Fatalf("error ao abrir o arquivo: %s", err.Error())
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("error ao ler as linhas do csv: %s", err.Error())
	}

	now := time.Now()
	var produtos []interface{}
	for i, line := range lines {
		if i == 0 {
			continue
		}

		id := line[1]
		cnpj := line[2]

		p := produto{
			ID:   id,
			Cnpj: cnpj,
		}

		produtos = append(produtos, p)
	}

	err = save(database, produtos)
	if err != nil {
		log.Fatalf("error ao salvar os produtos: %s", err.Error())
	}

	log.Println(time.Now().Sub(now))
}

func save(db *mongo.Database, ps []interface{}) error {
	res, err := db.Collection("produtos").InsertMany(context.Background(), ps)
	if err != nil {
		return err
	}

	log.Printf("%d documentos inseridos com sucesso\n", len(res.InsertedIDs))
	return nil
}
