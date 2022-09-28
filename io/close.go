package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("cannot open the file")
	}
	// ファイルは開いて使わなくなったら必ず閉じるものなので、Close()はdeferでの呼び出しと相性が良い
	// エラー処理ないバージョン(readのとき)
	// defer f.Close()

	// エラー処理ありバージョン(writeのとき、closeできない可能性がある)
	// deferを使いつつエラーを扱うためには無名関数を使う
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
