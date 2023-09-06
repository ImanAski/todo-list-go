package handlers

import (
	"net/http"

	"github.com/ImanAski/todo-list-go/utils"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	rnd = renderer.New()
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	utils.CheckError(err)

}
