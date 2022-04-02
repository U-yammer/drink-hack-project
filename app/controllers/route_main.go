package controllers

import (
	"log"
	"net/http"
)

/*
	http.ResponseWriter: httpのレスポンスを作るために必要なもの
	*http.Request: httpのリクエスト．
	この二つを引数にすると，ハンドラ（リクエストを処理できる関数）として定義される仕様
*/
func top (w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		renderView(w, "Hello", "layout", "public_navbar","top")
	} else {
		renderView(w, nil, "layout", "private_navbar", "index")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		renderView(w, nil, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		renderView(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			println(err)
		}
		user, err := sess.
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
	}
}