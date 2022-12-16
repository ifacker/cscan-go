package regex2

import "github.com/dlclark/regexp2"

// 正则表达式 全局查找 支持最新规则 （传入 re 参数和将要筛选的内容）
func Regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

// 正则表达式 简单使用 （传入将要筛选的内容和正则表达式）
func Regexp2SimpleUse(body, regex string) ([]string, error) {
	re, err := regexp2.Compile(regex, 0)
	if err != nil {
		return nil, err
	}
	regexResult := Regexp2FindAllString(re, body)
	return regexResult, nil
}

// 正则表达式 替换
func Regexp2SimpleReplace(src, dest, regex string) (string, error) {
	re, err := regexp2.Compile(regex, 0)
	if err != nil {
		return "", err
	}
	result, err := re.Replace(src, dest, -1, -1)
	if err != nil {
		return "", err
	}
	return result, nil
}
