package controllers

import (
	"drink_hack_project/app/models"
	"drink_hack_project/config"
	"fmt"
	"html/template"
	"net/http"
)

// 動的にhtmlファイルをレンダリングしている
func renderView(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	/* 引数 filenames から，app/views/templates/xxx.html の配列を作成する */
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	/* files から，htmlファイルを解析し，解析データを配列として持つ */
	templates := template.Must(template.ParseFiles(files...))

	/* 実行．layout.htmlがエントリポイントになっている(layout.htmlがベースとなって，htmlファイルが展開される) */
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	/* http request の設定 */
	files := http.FileServer(http.Dir(config.Config.Static))

	/* プロジェクト内では，"/app/views/" = /static/ とするが，httpアクセスする際はstaticはなくなる */
	http.Handle("/static/", http.StripPrefix("/static/", files))


	http.HandleFunc("/", top)          //  "/" にアクセスすると，topハンドラにルーティングする
	http.HandleFunc("/signup", signup) //  "/signup" にアクセスすると，signupハンドラをcallする
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)

	return http.ListenAndServe(":" + config.Config.Port, nil) // handler: nil にするとデフォルトで page not found を返す
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error){
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{
			UUID: cookie.Value,
		}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}