package hand

import "net/http"

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	if r.pat
	

}