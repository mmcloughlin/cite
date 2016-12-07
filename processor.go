package cite

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/mvdan/xurls"
)

var directiveRegex *regexp.Regexp
var subexpIdx = map[string]int{}

func init() {
	urlRegexp, _ := xurls.StrictMatchingScheme("https?")
	directiveExpr := `(?P<action>\w+):?\s+(?P<url>` + urlRegexp.String() + `)(\s+\((?P<extra>.+)\))?`
	directiveRegex = regexp.MustCompile(directiveExpr)

	for i, name := range directiveRegex.SubexpNames() {
		subexpIdx[name] = i
	}
}

type Citation struct {
	URL   *url.URL
	Extra string
}

func (c Citation) String() string {
	if c.Extra == "" {
		return c.URL.String()
	}
	return fmt.Sprintf("%v (%s)", c.URL, c.Extra)
}

type Directive struct {
	ActionRaw string
	Citation  Citation
}

func ParseDirective(line string) (*Directive, error) {
	match := directiveRegex.FindStringSubmatch(line)
	if match == nil {
		return nil, nil
	}

	action := match[subexpIdx["action"]]
	urlStr := match[subexpIdx["url"]]
	extra := match[subexpIdx["extra"]]

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	return &Directive{
		ActionRaw: action,
		Citation: Citation{
			URL:   u,
			Extra: extra,
		},
	}, nil
}

func CanonicalAction(s string) string {
	return strings.ToLower(s)
}

func (d Directive) Action() string {
	return CanonicalAction(d.ActionRaw)
}

func (d Directive) String() string {
	return fmt.Sprintf("%s: %v", d.ActionRaw, d.Citation)
}

//type Handler interface {
//	Handle(Directive, []string) ([]string, []string, error)
//}
//
//type Processor struct {
//	Handlers map[string]Handler
//}
//
//func (p Processor) Process(comment []string) ([]string, error) {
//	for _, line := range comment {
//		dir, err := ParseDirective(line)
//		if err != nil {
//			return nil, err
//		}
//
//	}
//}
