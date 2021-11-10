package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello,what's up", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 取得したCookieに応じてtodosを取得する
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
