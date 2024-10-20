package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	accountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	catalogURL string `envconfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {

	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	s, err := NewGraphQLServer(cfg.accountURL, cfg.catalogURL, cfg.OrderURL)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("Vikas", "/graphql"))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
