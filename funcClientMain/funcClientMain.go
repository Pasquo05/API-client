package funcClientMain

import (
	prjclient "Pasquo05/API-client/prjClient"
	"fmt"
	"net/http"
)

type client struct {
	url string
	cl  *http.Client
}

func NewClient(urlInput string) client {
	c := client{
		url: urlInput,
		cl:  &http.Client{},
	}
	return c
}

func (c client) DeleteBook() {

	Id := ""
	fmt.Scanln(&Id)
	prjclient.GetRequest(c.url)
}

func NewBook() prjclient.Book {

	/*
		fmt.Println("Scrivere id , titolo , desc")

		idInput := ""
		titleInput := ""
		descInput := ""

		fmt.Scanln(&idInput)
		fmt.Scanln(&titleInput)
		fmt.Scanln(&idInput)
	*/
	value := prjclient.Book{Id: "1", Title: "ciao", Desc: "ciao"}

	return value
}

func (c client) CreateBook() {

	book := NewBook()

	url := fmt.Sprintf("%s/book/post", c.url)

	prjclient.PostRequest(url, book)

}

func (a client) GetBooks() {
	url := fmt.Sprintf("%s/books", a.url)
	prjclient.GetRequest(url)
}

func (a client) GetBook() {

	Id := ""
	fmt.Scanln(&Id)
	url := fmt.Sprintf("%s/book/%s", a.url, Id)
	prjclient.GetRequest(url)

}
