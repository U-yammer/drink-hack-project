package controllers

import (
	"bytes"
	"drink_hack_project/app/models"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
)

/*
	http.ResponseWriter: httpのレスポンスを作るために必要なもの
	*http.Request: httpのリクエスト．
	この二つを引数にすると，ハンドラ（リクエストを処理できる関数）として定義される仕様
*/

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		renderView(w, "Hello", "layout", "public_navbar", "top")
	} else {
		//renderView(w, nil, "layout", "private_navbar", "index")
		http.Redirect(w, r, "/todos", 302)
	}
}

func getEncodePngImage(w http.ResponseWriter, filename string) string {
	file, err := os.Open("app/views/image/" + filename)
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}
	decodeImage, _, err := image.Decode(file)
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, decodeImage); err != nil {
		log.Fatalln("Unable to encode image.")
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index_do")
	sess, err := session(w, r)

	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()

		if err != nil {
			fmt.Println(err)
		}
		todos, _ := user.GetTodosByUser()

		user.Todos = todos
		encodeImage := getEncodePngImage(w, "account_icon.png")
		m := map[string]interface{}{
			"Image": encodeImage,
			"Name":  user.Name,
			"Todos": user.Todos,
		}

		renderView(w, m, "layout", "private_navbar", "index")
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

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}

		encodeImage := getEncodePngImage(w, "account_icon.png")
		m := map[string]interface{}{
			"Image":   encodeImage,
			"ID":      t.ID,
			"Content": t.Content,
		}

		renderView(w, m, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {

	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		t := &models.Todo{
			ID:      id,
			Content: content,
			UserID:  user.ID,
		}

		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}

		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}
