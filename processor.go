package cite

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/mvdan/xurls"
)

// ErrUnknownAction is returned when an action is encountered for which no
// handlers are registered.
var ErrUnknownAction = errors.New("cite: unknown action")

// ErrUnknownResource is returned when the Processor is unable to map a
// Citation to the Resource it references.
var ErrUnknownResource = errors.New("cite: unknown resource type")

var directiveRegex *regexp.Regexp
var subexpIdx = map[string]int{}

func init() {
	urlRegexp, _ := xurls.StrictMatchingScheme("https?")
	directiveExpr := `^\s*(?P<action>\w+):\s+(?P<url>` + urlRegexp.String() + `)(\s+\((?P<extra>.+)\))?`
	directiveRegex = regexp.MustCompile(directiveExpr)

	for i, name := range directiveRegex.SubexpNames() {
		subexpIdx[name] = i
	}
}

// Directive represents a line in godoc to be processed by cite. This is an
// action together with a citation.
type Directive struct {
	ActionRaw string
	Citation  Citation
}

// ParseDirective tests the given line for a Directive and returns it if
// found.
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

// CanonicalAction returns the canonical casing of an action string. This
// allows for arbitrary casing in the godoc.
func CanonicalAction(s string) string {
	return strings.ToLower(s)
}

// Action returns the canonical representation of the action in the directive.
func (d Directive) Action() string {
	return CanonicalAction(d.ActionRaw)
}

func (d Directive) String() string {
	return fmt.Sprintf("%s: %v", d.ActionRaw, d.Citation)
}

// Handler does something in response to a directive. It takes the referenced
// Resource and all remaining lines in the comment block. It returns lines to
// be inserted, lines remaining to be processed, and possibly an error.
type Handler func(Resource, []string) ([]string, []string, error)

// Processor modifies source code by performing some configured actions on
// citations found in comment blocks.
type Processor struct {
	Handlers         map[string]Handler
	ResourceBuilders []ResourceBuilder
}

// NewProcessor constructs a new Processor. It takes a list of
// ResourceBuilder, which defines the types of references the Processor
// understands.
func NewProcessor(builders []ResourceBuilder) Processor {
	return Processor{
		Handlers:         make(map[string]Handler),
		ResourceBuilders: builders,
	}
}

// AddHandler registers a handler for the given action type. Note casing of
// the action string does not matter.
func (p Processor) AddHandler(action string, handler Handler) {
	p.Handlers[CanonicalAction(action)] = handler
}

// Process transforms source code using registered handlers.
func (p Processor) Process(src Source) (Source, error) {
	blocks := make([]CodeBlock, len(src.Blocks))
	for i, block := range src.Blocks {
		var err error
		blocks[i], err = p.processCodeBlock(block)
		if err != nil {
			return Source{}, err
		}
	}
	return Source{
		Blocks: blocks,
	}, nil
}

func (p Processor) processCodeBlock(block CodeBlock) (CodeBlock, error) {
	comment, err := p.processCommentBlock(block.CommentBlock)
	if err != nil {
		return CodeBlock{}, err
	}
	return CodeBlock{
		CommentBlock: comment,
		Code:         block.Code,
	}, nil
}

func (p Processor) processCommentBlock(comment CommentBlock) (CommentBlock, error) {
	var err error
	comment.Lines, err = p.processLines(comment.Lines)
	if err != nil {
		return CommentBlock{}, err
	}
	return comment, nil
}

func (p Processor) processLines(lines []string) ([]string, error) {
	var output []string

	for len(lines) > 0 {
		line := lines[0]

		dir, err := ParseDirective(line)
		if err != nil {
			return nil, err
		}

		if dir == nil {
			output = append(output, line)
			lines = lines[1:]
			continue
		}

		r, err := p.getResource(dir.Citation)
		if err != nil {
			return nil, err
		}

		if r == nil {
			return nil, ErrUnknownResource
		}

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
