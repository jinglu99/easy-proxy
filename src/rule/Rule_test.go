package rule

import (
	"testing"
)

func TestPattern(t *testing.T) {
	var pairs = []struct {
		a string
		b string
		c string
		d string
	}{
		{"baidu.com->google.com", "baidu.com", "google.com", ""},
		{"baidu.com->google.com:80", "baidu.com", "google.com", "80"},
	}

	t.Parallel()
	for _, v := range pairs {
		t.Log("testing ", v)
		rs := CompiledPattern().FindStringSubmatch(v.a)
		t.Log("parse rs", rs)
		if rs[1] != v.b {
			t.Fatal("get", rs[1])
		}

		if rs[2] != v.c {
			t.Fatal("get", rs[2])
		}

		if rs[4] != v.d {
			t.Fatal("get", rs[3])
		}

		t.Log("success!")
	}
}

func TestParseRule(t *testing.T) {
	var pairs = []struct {
		a string
		b string
		c string
		d int
	}{
		{"baidu.com->google.com", "baidu.com", "google.com", 0},
		{"baidu.com->google.com:80", "baidu.com", "google.com", 80},
	}

	t.Parallel()
	for _, v := range pairs {
		t.Log("testing ", v)
		a, b, c, _ := parseRule(v.a)
		t.Log("parsed rs", a, b, c)
		if a != v.b {
			t.Fatal("get", a)
		}

		if b != v.c {
			t.Fatal("get", b)
		}

		if c != v.d {
			t.Fatal("get", c)
		}

		t.Log("success!")
	}

}

func TestRule_IsMatch(t *testing.T) {
	var pairs = []struct {
		rule  Rule
		host  string
		match bool
	}{
		{MustNewRule(".*\\.baidu\\.com->www.google.com"), "www.baidu.com", true},
		{MustNewRule(".*\\.baidu\\.com->www.google.com"), "map.baidu.com", true},
		{MustNewRule(".*\\.baidu\\.com->www.google.com"), "www.google.com", false},
		{MustNewRule(".*\\.google.com->www.baidu.com"), "email.google.com", true},
	}

	t.Parallel()
	for _, v := range pairs {
		if r := v.rule.IsMatch(v.host); r != v.match {
			t.Fatalf("test:%s, except: %t, actual: %t", v.host, v.match, r)
		}
	}
}
