package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "query":
		query()
	case "head":
		head()
	case "encoded":
		encoded()
	case "post":
		post()
	case "multi":
		multi()
	case "cookie":
		cookie()
	case "proxy":
		proxy()
	case "local":
		local()
	case "del":
		del()
	}
}

func query() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // io.Readerの内容をまとめてバイト配列に読み込む
	if err != nil {
		panic(err)
	}
	log.Println("body:", string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}

func head() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // io.Readerの内容をまとめてバイト配列に読み込む
	if err != nil {
		panic(err)
	}
	log.Println("body:", string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
}

func encoded() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func post() {
	reader := strings.NewReader("テキスト")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func multi() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	fileWriter, err := writer.CreateFormFile("thumbnail", "flagment")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("flagment")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func cookie() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Jar: jar,
	}

	// クッキーは初回アクセスでクッキーを受信し、2回め以降のアクセスでクッキーをサーバーに対して送信する
	for i := 0; i < 2; i++ {
		resp, err := client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(dump))
	}
}

func proxy() {
	proxyUrl, err := url.Parse("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	resp, err := client.Get("http://github.com")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))

	// リクエストを飛ばしたときに直にgithub.comに飛ばすのではなく、プロキシ先のローカルのサーバに飛ばす。
}

// ファイルスキームはローカルファイルにアクセスするときのスキーム
func local() {
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	client := http.Client{
		Transport: transport,
	}
	resp, err := client.Get("file://./main.go")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}

func del() {
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpRequest(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
