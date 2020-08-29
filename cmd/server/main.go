package main

import (
	"net/http"

	"github.com/johnhckuo/Crawler/pkg/crawler/warframe"
	RM "github.com/johnhckuo/Crawler/pkg/crawler/warframe/riven.market"
)

func main() {

	// starting server
	/*
		mux := http.NewServeMux()

		mux.HandleFunc("/", homeHandler)
		mux.HandleFunc("/contact", contactHandler)

		log.Println(fmt.Sprintf("Server running on http://localhost%s 🐹", ":4000"))
		err := http.ListenAndServe(":4000", mux)
		if err != nil {
			log.Fatalf("could not run the server %v", err)
			return
		}
	*/
	handler()
}

// listen to request
/*
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// validation

	// get riven prices

	targetWeapon := "Grakata"
	var crawler riven.Crawler = &RM.Crawler{}
	crawler.GetRivenByWeapon(&targetWeapon)

	w.Write([]byte("sss"))

}
*/

func handler() {

	// validation

	// get riven prices

	targetWeapon := "Grakata"
	var crawler warframe.Crawler = &RM.Crawler{}
	crawler.GetRivenByWeapon(&targetWeapon)

	//w.Write([]byte("sss"))

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from contact handler"))
}
