package main

import (
	"example/chat/trace"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

type Env struct {
	GOOGLE_SECRETKEY string `envconfig:"GOOGLE_SECRETKEY" default:""`
	GOOGLE_CLIENTID  string `envconfig:"GOOGLE_CLIENTID" default:""`
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse environment variables: %s\n", err.Error())
		os.Exit(1)
	}

	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	gomniauth.SetSecurityKey(env.GOOGLE_SECRETKEY)
	gomniauth.WithProviders(
		google.New("クライアントID", env.GOOGLE_CLIENTID, "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Webサーバを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
