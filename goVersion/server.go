package main

/*
type NumberImg struct {
	File string `json: "file"`
	Text string `json: "text"`
}

func runServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/img", ImgToNumber).Methods("POST")
	log.Fatal(http.ListenAndServe(":3020", router))
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func ImgToNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	decoder := json.NewDecoder(r.Body)

	var t NumberImg
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Text)
	fmt.Fprintf(w, "data")
}*/
