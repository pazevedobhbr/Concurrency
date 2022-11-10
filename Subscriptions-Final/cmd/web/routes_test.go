package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

func Test_Routes_Exists(t *testing.T) {
	testRoutes := testApp.routes()

	chiRoutes := testRoutes.(chi.Router)

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false
	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if foundRoute == route {
			found = true
		}
		return nil
	})
	if !found {
		t.Errorf("Route %s not found", route)
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	auth := testApp.IsAuthenticated(req)

	if auth {
		t.Error("returns true for authenticated, when it should be false")
	}

	testApp.Session.Put(ctx, "userID", 1)

	auth = testApp.IsAuthenticated(req)

	if !auth {
		t.Error("returns false for authenticated, when it should be true")
	}
}

func TestConfig_render(t *testing.T) {
	pathToTemplates = "./templates"

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != 200 {
		t.Error("failed to render page")
	}
}
