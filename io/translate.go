package main

import (
	"strings"
	"testing"
)

func TranslateIntoGerman(r io.Reader) string {
	data := make([]byte, 300)
	len, _ := r.Read(data)
	str := string(data[:len])

	result := strings.ReplaceAll(str, "Hello", "Guten Tag")
	return result
}

// strings.Reader型を使うことで、テスト内容をコード内で用意できる
func Test_Translate(t *testing.T) {
	// テスト
	tests := []struct{
		name string
		arg string
		want string
	} {
		{
			name: "normal"
			arg: "Hello, World!",
				want: "Guten Tag, WOrld!",
			},
			{
			name: "repeat",
				arg: "Hello World, Hello Golang!",
				want: "Guten Tag World, Guten Tag Golang!",
			},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateIntoGerman(strings.NewReader(tt.arg)); got != tt.want { // TranslateIntoGerman関数には、strings.NewReader(tt.args)で用意したstrings.Reader型を返す
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
