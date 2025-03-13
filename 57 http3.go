package main
import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r:=mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)){
		fmt.Fprintf(w, "Welcome to the homepage!")
	}
}
r.handleFunc("/about", func(w http.ResponseWriter, r *http.Request)){
	fmt.Fprintf(w, "Welcome to about page")
}

r.handleFunc("/contact", func(w http.ResponseWriter, r *http.Request)){
	fmt.Fprintf(w, "Welcome to contact page contact us on github")
}
fmt.Println("server started at :9090")
http.ListenAndServe(":9090",r)