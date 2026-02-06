package delivery

import "net/http"

func (d *Delivery) GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/chats", d.chats)
	mux.HandleFunc("/chats/{id}/messages", d.message)
	mux.HandleFunc("/chats/{id}", d.doChats)
	return mux
}

func (d *Delivery) chats(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		d.newChat(w, r)
	}
}
func (d *Delivery) message(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		d.newMewssage(w, r)
	}
}
func (d *Delivery) doChats(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		d.deleteChat(w, r)
	}
	if r.Method == http.MethodGet {
		d.getChat(w, r)
	}
}
