package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Case(root string, f func(string) []string) []string {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := r.RequestURI
		if req == "/" {
			http.Redirect(w, r, "/page1.html", 301)
		} else {
			http.ServeFile(w, r, "testdata/"+root+"/"+req)
		}
	}))
	defer ts.Close()
	return f(ts.URL)
}

func TestCrawler(t *testing.T) {
	for ii, tt := range []struct {
		root     string
		expected []string
	}{
		{root: "hosta", expected: []string{"/page1.html", "/page2.html", "/page3.html", "/page4.html", "/page5.html"}},
		{root: "hostb", expected: []string{"/page1.html", "/page2.html"}},
		{root: "hostd", expected: []string{"/page1.html", "/subdir/page2.html", "/page3.html"}},
	} {
		res := Case(tt.root, Crawl)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("Case [%d]: failed, expected %v, got %v", ii, tt.expected, res)
		}
	}
}
