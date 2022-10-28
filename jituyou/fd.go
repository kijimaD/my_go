// 同じようなエラーチェックの繰り返し

_, err = fd.Write(p0[a:b])
if err =! nil {
	return err
}
_, err = fd.Write(p1[c:d])
if err =! nil {
	return err
}
_, err = fd.Write(p2[c:d])
if err =! nil {
	return err
}

// すべての書き込みが成功しないと意味をなさない場合、書き込むときに発生したエラーを内部で保持するようなラッパーを定義すると、呼び出し元のコードはシンプルになる。1つでもエラーが発生した場合はそれ以降の処理はスキップし、最終的なエラーのチェックは1回で済むようにする

type errWriter struct {
	w io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	// 本処理。エラーがあった場合は構造体の方に保持する
	_, ew.err = ew.w.Write(buf)
}

// 使いかた
ew := &errWriter{w: fd}
ew.write(p0[a:b])
ew.write(p0[c:d])
ew.write(p0[e:f])

if ew.err != nil {
	return ew.err
}
