package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	root, err := os.OpenRoot(".")
	assert.NoError(t, err)
	defer root.Close()

	t.Run("シンボリックリンク先が下層にあると成功する", func(t *testing.T) {
		err = root.Mkdir("./down_sym/success", 0o777)
		assert.NoError(t, err)

		defer func() {
			err = os.RemoveAll("./down_sym/success")
			assert.NoError(t, err)
		}()
	})

	t.Run("シンボリックリンク先が上層にあると成功する", func(t *testing.T) {
		err = root.Mkdir("./up_sym/success", 0o777)
		assert.NoError(t, err)

		defer func() {
			err = os.RemoveAll("./up_sym/success")
			assert.NoError(t, err)
		}()
	})
}

// 今までの方法
// 直接作るしかない
func TestNormalMkdir(t *testing.T) {
	err := os.Mkdir("example_dir", 0o777)
	assert.NoError(t, err)
	err = os.Mkdir("example_dir/aaa", 0o777)
	assert.NoError(t, err)

	defer func() {
		err = os.RemoveAll("example_dir")
		assert.NoError(t, err)
	}()
}
