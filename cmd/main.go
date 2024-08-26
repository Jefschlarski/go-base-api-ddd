package main

import (
	"fmt"
	"log"
	"net/http"
	"taskmanager/internal/api/router"
	"taskmanager/internal/configs"
)

func main() {

	err := configs.Load()
	if err != nil {
		log.Fatal("Erro ao tentar carregar as configurações", err)
	}

	apiConfigs := configs.GetApiConfig()

	r := router.GenRouter()

	fmt.Println("Servidor escutando a porta ", apiConfigs.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", apiConfigs.Url, apiConfigs.Port), r))
}
