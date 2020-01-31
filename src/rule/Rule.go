package rule

import (
	"errors"
	"regexp"
	"strconv"
)

const rulePattern string = `(?U)^(.*)->(.*)(:([0-9]+))?$`

var compiledPattern = regexp.MustCompile(rulePattern)

type Rule interface {
	IsMatch(fqna string) bool
	ToString() string
}

//Default implement of Rule
type DefaultRule struct {
	rule          string
	Pattern       string
	Host          string
	Port          int
	parsedPattern *regexp.Regexp
}

func NewRule(rule string) (Rule, error) {
	if pattern, host, port, err := parseRule(rule); err != nil {
		return &DefaultRule{}, err
	} else {
		if compiledPattern, err := regexp.Compile(pattern); err != nil {
			return &DefaultRule{}, errors.New(pattern + "is not a valid regular expression")
		} else {
			return &DefaultRule{
				rule:          rule,
				Pattern:       pattern,
				Host:          host,
				Port:          port,
				parsedPattern: compiledPattern,
			}, nil
		}
	}
}

func MustNewRule(rule string) Rule {
	if r, err := NewRule(rule); err != nil {
		return &EmptyRule{}
	} else {
		return r
	}
}

//return the given fqna is matched for rule
func (rule *DefaultRule) IsMatch(fqna string) bool {
	return rule.parsedPattern.MatchString(fqna)
}

func (rule *DefaultRule) ToString() string {
	return rule.rule
}

//empty implement of Rule
type EmptyRule struct {
}

//return the given fqna is matched for rule
func (rule *EmptyRule) IsMatch(fqna string) bool {
	return false
}

func (rule *EmptyRule) ToString() string {
	return ""
}

func CompiledPattern() *regexp.Regexp {
	return compiledPattern
}

//parse rule to match pattern, proxy host and port
func parseRule(rule string) (p string, h string, port int, err error) {
	rs := CompiledPattern().FindStringSubmatch(rule)

	if len(rs) != 5 {
		err = errors.New("there is grammar error in given rule")
	}

	p = rs[1]
	h = rs[2]
	if rs[4] == "" {
		port = 0
	} else {
		port, err = strconv.Atoi(rs[4])
	}
	err = nil
	return
}
