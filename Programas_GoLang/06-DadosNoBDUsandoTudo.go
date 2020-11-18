//usando tudo e salvando em batch

package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"sync"
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
		log.Fatal("error pinging mongo")
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

	parsed := make(chan produto, 100)
	wg := sync.WaitGroup{}
	wgParse := sync.WaitGroup{}

	now := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(parsed)

		for i, line := range lines {
			if i == 0 {
				continue
			}

			wgParse.Add(1)
			go func(l []string) {
				defer wgParse.Done()

				id := l[1]
				cnpj := l[2]

				p := produto{
					ID:   id,
					Cnpj: cnpj,
				}

				parsed <- p
			}(line)
		}

		wgParse.Wait()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var produtos []interface{}
		for p := range parsed {
			produtos = append(produtos, p)
		}

		err := save(database, produtos)
		if err != nil {
			log.Fatalf("error ao salvar na base: %s", err.Error())
		}
	}()

	wg.Wait()
	log.Println(time.Now().Sub(now))
}

//salvando em batch
func save(db *mongo.Database, ps []interface{}) error {
	res, err := db.Collection("produtos").InsertMany(context.Background(), ps)
	if err != nil {
		return err
	}

	log.Printf("%d documentos inseridos com sucesso", len(res.InsertedIDs))
	return nil
}
