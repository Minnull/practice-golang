package db

import (
	"errors"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	// 期望返回100，但是存在error，所以方法内部不会返回db查询的结果，最终会返回-1
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exists"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
