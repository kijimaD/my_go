// ステータスコードが200以外の場合のエラーを扱う構造体
type HTTPError struct {
	StatusCode int
	URL string
}

// このメソッドを実装することでHTTPError 構造体のポインターは Error インターフェースを満たす
func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code =%d, url = %s", he.StatusCode, he.URL)
}

// 利用するコード例

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL: url,
		}
	}
	return io.ReadAll(resp.Body)
}
