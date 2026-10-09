package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    if err != nil {
        panic("This shouldn't have happened")
    }
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[-=~*+]*>`)
    if err != nil {
        panic("This shouldn't have happened")
    }
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`(?i)".*password.*"`)
    if err != nil {
        panic("This shouldn't have happened")
    }

    var total int
    for _, l := range lines {
        if re.MatchString(l) {
            total++
        }
    }
    return total
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)
    if err != nil {
        panic("This shouldn't have happened")
    }
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User\s+(\w+)`)
    if err != nil {
        panic("This shouldn't have happened")
    }

    var res []string
    for i, l := range lines {
        sl := re.FindStringSubmatch(l)
        if sl == nil {
            res = append(res, l)
        } else {
            res = append(res, "[USR] " + sl[1] + " " + lines[i])
        }
    }
    return res
}
