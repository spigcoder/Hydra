package errs

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	err := Wrap(Unknown, "I don't know what happened").WithMeta(Metadata{
		"foo": "bar",
	})
	assert.Equal(t, Unknown, err.Code)
	assert.Equal(t, "bar", err.Meta["foo"])
	assert.Equal(t, nil, err.Err)
}

// 测试 WithErr 方法
func TestError_WithErr(t *testing.T) {
	err := Error{Code: http.StatusOK, Message: "ok"}
	assert.Equal(t, nil, err.Err)
	// 使用 WithErr 方法设置 Err 属性
	assert.Equal(t, "mock error", err.WithErr(errors.New("mock error")).Err.Error())
}

// 测试 WithMeta 方法
func TestError_WithMeta(t *testing.T) {
	err := Error{Code: http.StatusOK, Message: "ok"}
	assert.Equal(t, nil, err.Err)
	// 使用 WithMeta 方法设置 Meta 属性
	assert.Equal(t, "bar", err.WithMeta(Metadata{"foo": "bar"}).Meta["foo"])
}
