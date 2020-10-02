package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeGetHandler(t *testing.T) {
	// 拿到我们程序的路由对象
	router := SetupRoute()
	w := httptest.NewRecorder()
	// 利用net/http发请求给我的gin server
	req, _ := http.NewRequest(http.MethodGet, "/someGet", nil)
	router.ServeHTTP(w, req) // 模拟我们写的程序处理上面的测试请求，把得到的响应结果写入w
	// 判断得到的响应是否与与预期的结果一致
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin", w.Body.String())
}

func TestMax(t *testing.T) {
	type test struct {
		inputA int
		inputB int
		want   int
	}
	tests := []test{
		{inputA: 10, inputB: 20, want: 20},
		{inputA: 1, inputB: 1, want: 1},
		{inputA: -10, inputB: 20, want: 20},
		{inputA: -10, inputB: -20, want: -10},
	}
	for _, tc := range tests {
		got := max(tc.inputA, tc.inputB)
		assert.Equal(t, got, tc.want)
	}
}
