package actions

import (
	"fmt"
	"github.com/jingleWang/easy-proxy/src/config"
	"github.com/jingleWang/easy-proxy/src/rule"
	"github.com/urfave/cli/v2"
	"strconv"
)

//list rules
func ListRules(c *cli.Context) error {
	printRules()
	return nil
}

//add rules
func AddRules(c *cli.Context) error {
	if c.NArg() >= 2 {
		ruleStr := c.Args().Get(0) + "->" + c.Args().Get(1)
		fmt.Printf("add rule: %s\n", ruleStr)
		if rule, err := rule.NewRule(ruleStr); err != nil {
			fmt.Println(err.Error())
		} else {
			config.AddRule(rule)
			fmt.Println("current rules:")
			printRules()
		}
	} else {
		fmt.Printf("add rules requires 2 string values\n")
	}
	return nil
}

//remove rule
func RemoveRule(c *cli.Context) error {
	if ok := removeAllRuleIfNeeded(c); ok {
		return nil
	}

	removeGivenRule(c)
	return nil
}

//remove all rule if flag all config.
func removeAllRuleIfNeeded(c *cli.Context) bool {
	if c.Bool("all") {
		config.RemoveAllRule()
		fmt.Println("remove all rules")
		return true
	}
	return false
}

func removeGivenRule(c *cli.Context) {
	if c.NArg() > 0 {
		if i, err := strconv.Atoi(c.Args().Get(0)); err != nil {
			fmt.Printf("the argument of remove rules must an int value\n")
		} else {
			config.RemoveRuleAndSyncByIndex(i)
			fmt.Printf("remove success\n")
			fmt.Printf("current rules:\n")
			printRules()
		}
	} else {
		fmt.Printf("remove rules requires an int value\n")
	}
}

func printRules() {
	for i, v := range config.GetRulesStrs() {
		fmt.Printf("(%d). %s\n", i, v)
	}
}
