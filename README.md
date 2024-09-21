## MD provisório com informações basicas do projeto e lista de pacotes para poesterior criação de container docker.
 
#### comando para iniciar o servidor
```bash
 go run ./cmd/main.go
```
### Lista de pacotes instalados

```
Nesse projeto estou utilizando o pacote viper para utilizar arquivos .toml para configuração de ambiente
```

* go get github.com/gorilla/mux
* go get -u github.com/spf13/viper
* go get -u github.com/lib/pq  
* go get -u github.com/golang-jwt/jwt/v5


```
//generate 64 bits key
func init() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	string64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(string64)
}
```