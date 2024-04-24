package webhook

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HandleIncomingWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}

	log.Println("Received webhook payload: ", string(body))
}

func StartServer() {
	http.HandleFunc("/webhook", HandleIncomingWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
