package config

import (
	"github.com/jingleWang/easy-proxy/src/rule"
	"log"
	"testing"
)

func TestGetRulesStr(t *testing.T) {
	print(GetRulesStrs())
}

func TestWriteReadRules(t *testing.T) {
	//rule1 := rule.MustNewRule("www.baidu.com->www.google.com")
	rules := []rule.Rule {
		rule.MustNewRule("www.baidu.com->www.google.com"),
		rule.MustNewRule("www.baidu1.com->localhost:8080"),
	}

	writeRules(&rules)

	rules1 := GetRules()
	for i, r:= range *rules1 {
		if r.ToString() != rules[i].ToString() {
			log.Fatalf("err when test:%s", rules[i].ToString())
		}
	}
}

func TestWriteTestRule(t *testing.T)  {
	rules := []rule.Rule {
		rule.MustNewRule("www.baidu.com->www.google.com"),
		rule.MustNewRule("www.baidu1.com->localhost:8080"),
		rule.MustNewRule("www.baidu2.com->localhost:8080"),
	}

	writeRules(&rules)
}

func TestRemoveRuleByIndex(t *testing.T) {
	//rule1 := rule.MustNewRule("www.baidu.com->www.google.com")
	rules := []rule.Rule {
		rule.MustNewRule("www.baidu.com->www.google.com"),
		rule.MustNewRule("www.baidu1.com->localhost:8080"),
		rule.MustNewRule("www.baidu2.com->localhost:8080"),
	}

	writeRules(&rules)
	RemoveRuleAndSyncByIndex(0)
	r1 := GetRules()
	if len(*r1) != 2 {
		t.Fatalf("err when remove 0 actual: %d", len(*r1))
	}

	writeRules(&rules)
	RemoveRuleAndSyncByIndex(1)
	r2 := GetRules()
	if len(*r2) != 2 {
		t.Fatalf("err when remove 1 actual: %d", len(*r2))
	}

	writeRules(&rules)
	RemoveRuleAndSyncByIndex(2)
	r3 := GetRules()
	if len(*r3) != 2 {
		t.Fatalf("err when remove 2 actual: %d", len(*r3))
	}

}

func Test(t *testing.T) {
	GetRules()

}

