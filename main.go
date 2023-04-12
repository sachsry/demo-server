package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Env     string `envconfig:"ENV" required:"true"`
	DevURL  string `envconfig:"DEV_URL"`
	ProdURL string `envconfig:"PROD_URL"`
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/demo", DemoConfig)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am testing out something")
	fmt.Fprintf(w, "Hello, from %s!", r.Host)
}

func DemoConfig(w http.ResponseWriter, r *http.Request) {
	var c config
	err := envconfig.Process("", &c)
	resp := make(map[string]string)
	if err != nil {
		resp["error"] = err.Error()
	} else {
		resp["env"] = c.Env
		resp["devUrl"] = c.DevURL
		resp["prodUrl"] = c.ProdURL
	}

	json, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		fmt.Fprintf(w, "encountered unexpected error: %s", marshalErr.Error())
		return
	}

	w.Write(json)
}
