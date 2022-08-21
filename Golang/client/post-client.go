package client

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetPost(title string) (string, error) {
	Url := "https://api.themoviedb.org/3/search/movie?api_key=72066e73ae85af67c2bc69cdfd9abca7&query=" + title
	request, err := http.Get(Url)

	if err != nil {
		log.Fatal(err)
	}

	defer request.Body.Close()
	//Create a variable of the same type as our model
	var pResp posts.post
	//Decode the data
	if err := json.NewDecoder(request.Body).Decode(&pResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	//Invoke the text output function & return it with nil as the error value
	return pResp.TextOutput(), nil

}
