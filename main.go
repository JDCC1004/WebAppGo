package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type PageData struct {
	HostName 	string
	RandomImgs 	[]ImageData
	Name 		string
}

type ImageData struct {
	Base64 string
	Name string
}

func main() {
	puerto := flag.String("puerto", "8000", "Puerto del Servidor")
	flag.Parse()
	
	http.HandleFunc("/", handler)
	fmt.Println("Servidor corriendo en  http://localhost:" + *puerto)
	if err := http.ListenAndServe(":"+*puerto, nil); err != nil {
		fmt.Println("Error al iniciar el servidor: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmp1 := template.Must(template.ParseFiles("Templates/index.html"))

	hostName, err := os.Hostname()
	if err != nil {
		hostName = "Desconocido"
	}

	randomImages := getRandomImg(4)

	data := PageData{
		HostName: hostName,
		RandomImgs: randomImages,
		Name: "Julian Cruz - Eric Bedoya",
	}

	tmp1.Execute(w, data)
}