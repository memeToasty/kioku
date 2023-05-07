package handler_test

import (
	"testing"

	"github.com/kioku-project/kioku/services/register/handler"

	pblogin "github.com/kioku-project/kioku/services/login/proto"
	"github.com/kioku-project/kioku/store"
)

func TestRegister(t *testing.T) {
	var mockStore store.Store = nil
	var mockLoginService pblogin.LoginService = nil

	registerService := handler.New(mockStore, mockLoginService)

	if registerService == nil {
		t.Errorf("Received invalid Register Service\n")
	}
}
