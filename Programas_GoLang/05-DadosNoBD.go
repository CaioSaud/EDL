//exemplo sem goroutines e sem gochannels

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
//declaracao objeto produto
type produto struct {
	ID   string
	Cnpj string
}

func main() {
	//conexao com o banco de dados
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
	//acesso ao arquivos com os dados
	f, err := os.Open("dump_compliance.csv")
	if err != nil {
		log.Fatalf("error ao abrir o arquivo: %s", err.Error())
	}
	
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalf("error ao ler as linhas do csv: %s", err.Error())
	}

	// criando um array de produtos
	//var produtos []produto
	now := time.Now()
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

		err = save(database, p)
		if err != nil {
			log.Fatalf("error ao salvar os produtos: %s", err.Error())
		}
	}

	log.Println(time.Now().Sub(now))
}
//função para salvar no banco
func save(db *mongo.Database, p produto) error {
	_, err := db.Collection("produtos").InsertOne(context.Background(), p)
	if err != nil {
		return err
	}

	log.Printf("ID %s inserido com sucesso\n", p.ID)
	return nil
}
