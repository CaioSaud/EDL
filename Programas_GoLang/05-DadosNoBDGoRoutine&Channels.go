//Com goroutines e gochannels salvando sequencialmente

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
	//criação do channel para comunicação entre goroutines
	parsed := make(chan produto)
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
				//envio do produto por meio do channel
				parsed <- p
			}(line)
		}

		wgParse.Wait()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for p := range parsed {
			err := save(database, p)
			if err != nil {
				log.Fatalf("error ao salvar os produtos: %s", err.Error())
			}
		}
	}()

	wg.Wait()
	log.Println(time.Now().Sub(now))
}

func save(db *mongo.Database, p produto) error {
	_, err := db.Collection("produtos").InsertOne(context.Background(), p)
	if err != nil {
		return err
	}

	log.Printf("ID %s inserido com sucesso\n", p.ID)
	return nil
}
