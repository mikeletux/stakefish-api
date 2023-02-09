package controller

import regexp2 "regexp"

const ipRegExp string = `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`

func checkIfStringIsIP(ip string) bool {
	regexp := regexp2.MustCompile(ipRegExp)
	return regexp.MatchString(ip)
}
