package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")

	link, errRes := h.service.RedirectTitleLink(title)
	if errRes != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Link não encontrado"))
		return
	}
	fmt.Println(link)

	http.Redirect(w, r, link.Link, http.StatusFound)
}
