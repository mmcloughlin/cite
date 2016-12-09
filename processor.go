package cite

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/mvdan/xurls"
)

var ErrUnknownAction = errors.New("cite: unknown action")

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

type Handler func(Resource, []string) ([]string, []string, error)

type Processor struct {
	Handlers         map[string]Handler
	ResourceBuilders []ResourceBuilder
}

func NewProcessor(builders []ResourceBuilder) Processor {
	return Processor{
		Handlers:         make(map[string]Handler),
		ResourceBuilders: builders,
	}
}

func (p Processor) AddHandler(action string, handler Handler) {
	p.Handlers[CanonicalAction(action)] = handler
}

func (p Processor) Process(src Source) (Source, error) {
	blocks := make([]CodeBlock, len(src.Blocks))
	for i, block := range src.Blocks {
		var err error
		blocks[i], err = p.ProcessCodeBlock(block)
		if err != nil {
			return Source{}, nil
		}
	}
	return Source{
		Blocks: blocks,
	}, nil
}

func (p Processor) ProcessCodeBlock(block CodeBlock) (CodeBlock, error) {
	comment, err := p.ProcessCommentBlock(block.CommentBlock)
	if err != nil {
		return CodeBlock{}, err
	}
	return CodeBlock{
		CommentBlock: comment,
		Code:         block.Code,
	}, nil
}

func (p Processor) ProcessCommentBlock(comment CommentBlock) (CommentBlock, error) {
	var err error
	comment.Lines, err = p.ProcessLines(comment.Lines)
	if err != nil {
		return CommentBlock{}, err
	}
	return comment, nil
}

func (p Processor) ProcessLines(lines []string) ([]string, error) {
	var output []string

	for len(lines) > 0 {
		line := lines[0]
		fmt.Println("processing:", line)

		dir, err := ParseDirective(line)
		if err != nil {
			return nil, err
		}

		if dir == nil {
			output = append(output, line)
			lines = lines[1:]
			continue
		}

		fmt.Println("directive:", dir)

		r, err := p.getResource(dir.Citation)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println("resource:", r)

		handler, ok := p.Handlers[dir.Action()]
		if !ok {
			return nil, ErrUnknownAction
		}

		insert, remainder, err := handler(r, lines)
		if err != nil {
			return nil, err
		}

		output = append(output, insert...)
		lines = remainder
	}

	return output, nil
}

func (p Processor) getResource(citation Citation) (Resource, error) {
	for _, builder := range p.ResourceBuilders {
		r, err := builder(citation)
		if r != nil || err != nil {
			return r, err
		}
	}
	return nil, nil
}
