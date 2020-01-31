package config

import (
	"github.com/jingleWang/easy-proxy/src/rule"
)

func GetRules() *[]rule.Rule {
	return getRulesFromConfig()
}

func GetRulesStrs() []string {
	return conf.GetStringSlice("rules")
}

func RemoveRuleAndSyncByIndex(index int) {
	writeRules(removeRuleByIndex(index))
}

func RemoveAllRule() {
	writeRules(&[]rule.Rule{})
}

func AddRule(rule rule.Rule) {
	newRule := *getRulesFromConfig()
	newRule = append(newRule, rule)
	writeRules(&newRule)
}

func removeRuleByIndex(index int) *[]rule.Rule {
	var newRule []rule.Rule
	rules := getRulesFromConfig()

	if len(*rules) == 0 {
		return &[]rule.Rule{}
	}

	if index == 0 {
		if len(*rules) == 1 {
			newRule = []rule.Rule{}
		} else {
			newRule = (*rules)[1:]
		}
	} else if index == len(*rules)-1 {
		newRule = (*rules)[:index]
	} else if index > 0 && index < len(*rules)-1 {
		newRule = (*rules)[:index+copy((*rules)[index:], (*rules)[index+1:])]
	} else {
		newRule = *rules
	}
	return &newRule
}

func getRulesFromConfig() *[]rule.Rule {
	ruleStr := GetRulesStrs()
	rules := make([]rule.Rule, len(ruleStr))
	for i, str := range ruleStr {
		rules[i] = rule.MustNewRule(str)
	}

	return &rules
}

func writeRules(rules *[]rule.Rule) {
	if rules == nil {
		return
	}

	emptyCn := 0
	ruleStr := make([]string, len(*rules))
	for i, r := range *rules {
		if pr, ok := interface{}(r).(*rule.DefaultRule); ok {
			ruleStr[i-emptyCn] = pr.ToString()
		} else {
			ruleStr[len(*rules)-1-emptyCn] = ""
			emptyCn++
		}
	}
	writeRulesStrs(ruleStr[0 : len(*rules)-emptyCn])
}

func writeRulesStrs(rules []string) {
	conf.Set("rules", rules)
	conf.WriteConfig()
}


