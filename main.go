package main

import (
	"api/src/api/router"
	"api/src/configs"
	"fmt"
	"log"
	"net/http"
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
