package main

import (
	"html/template"  //permite cargar plantillas HTML
	"net/http"       //implementaciones cliente servidor
	"fmt"
	"io"             //librerías E/S
	"io/ioutil"
	
)

    var tem *template.Template    

func init() {
	tem = template.Must(template.ParseGlob("templates/*"))   //retorna los html ubicados en la carpeta template
}

type persona struct {
	Nombre  string
	Apellido   string               //colección de datos
	Telefono string
	Correo string
	Subscrito bool
}



func ingresarPer(w http.ResponseWriter, r *http.Request) {

	nom := r.FormValue("nombre")
	ape := r.FormValue("apellido")                              
	tel := r.FormValue("telefono")        //retorna los valores ingresados en el html 
	co  := r.FormValue("correo")
	sub := r.FormValue("subscrito") == "on"
	
    tem.ExecuteTemplate(w, "index.html", persona{nom, ape, tel, co, sub})  


	
	}
	
	
func visualizar(w http.ResponseWriter, r *http.Request) {

    nom := r.FormValue("nombre")
	ape := r.FormValue("apellido")
	tel := r.FormValue("telefono")
	co  := r.FormValue("correo")
	sub := r.FormValue("subscrito") == "on"
	
	tem.ExecuteTemplate(w, "index2.html", persona{nom, ape, tel, co, sub})      //la información capturada de los FormValue los guarda en los atributos de la estructura persona
	}                                                                           // y los muestra en el index2.html


func cargarArchivo(w http.ResponseWriter, req *http.Request) {

	var inf string
	fmt.Println(req.Method)    //muestra la acción a ejecutar desde consola ya sea GET o POST
	if req.Method == http.MethodPost {

		// Abrir
		archivo, h, err := req.FormFile("arc")    // := operador de declaración corta no tiene un tipo especifico de daro
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return  
		}
		defer archivo.Close() //cierra el archivo pero al tener el comando defer lo hace hasta que finalice el if

		// Información desde consola
		fmt.Println("\nfile:", archivo, "\nheader:", h, "\nerr", err) 

		// Lee el archivo
		bs, err := ioutil.ReadAll(archivo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		inf = string(bs) //variable donde se muestra la información del txt
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
    <body bgcolor="637EC3" align="center">
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="arc">
	<br><br>
	<input type="submit">
	</form>
	<br>
	</body>
	`+inf)
}

func main() {
	http.HandleFunc("/", ingresarPer)
	http.HandleFunc("/datos", visualizar)
	http.HandleFunc("/files", cargarArchivo)
	http.ListenAndServe(":8080", nil)
}
