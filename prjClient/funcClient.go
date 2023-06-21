package prjclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func PostRequest(url string, bookInput interface{}) {

	book, err := json.Marshal(bookInput)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(book))
	if err != nil {
		fmt.Println("Errore durante l'esecuzione della richiesta:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Errore durante la lettura della risposta:", err)
		return
	}

	fmt.Println(string(body))

}

func GetRequest(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Errore durante l'esecuzione della richiesta:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Errore durante la lettura della risposta:", err)
		return
	}

	fmt.Println(string(body))
}

func GetDelete(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Errore durante l'esecuzione della richiesta:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Errore durante la lettura della risposta:", err)
		return
	}

	fmt.Println(string(body))
}
