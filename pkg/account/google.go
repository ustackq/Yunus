package account

// GetAccessRule ...
func (ag *Account) GetAccessRule() RuleAction {
	ruleAction := make(RuleAction)
	ruleAction["ruleType"] = Black
	ruleAction["actions"] = []string{"registry"}
	return ruleAction
}
