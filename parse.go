package cite

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

var commentExpr = `^(\s*)//(.*)$`
var commentRegexp = regexp.MustCompile(commentExpr)

type CommentBlock struct {
	Leader string
	Lines  []string
}

func (c CommentBlock) String() string {
	out := ""
	for _, line := range c.Lines {
		out += fmt.Sprintf("%s//%s\n", c.Leader, line)
	}
	return out
}

type CodeBlock struct {
	CommentBlock CommentBlock
	Code         []string
}

func (b CodeBlock) String() string {
	out := b.CommentBlock.String()
	for _, line := range b.Code {
		out += line + "\n"
	}
	return out
}

type Source struct {
	Blocks []CodeBlock
}

func (s Source) String() string {
	out := ""
	for _, b := range s.Blocks {
		out += b.String()
	}
	return out
}

func ParseCode(r io.Reader) Source {
	src := Source{}
	block := CodeBlock{}
	inComment := true

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		match := commentRegexp.FindStringSubmatch(line)
		isComment := match != nil

		switch {
		case inComment && isComment:
			block.CommentBlock.Lines = append(block.CommentBlock.Lines, match[2])
		case inComment && !isComment:
			block.Code = append(block.Code, line)
		case !inComment && isComment:
			src.Blocks = append(src.Blocks, block)
			block = CodeBlock{
				CommentBlock: CommentBlock{
					Leader: match[1],
					Lines:  []string{match[2]},
				},
			}
		case !inComment && !isComment:
			block.Code = append(block.Code, line)
		}

		inComment = isComment
	}

	src.Blocks = append(src.Blocks, block)

	return src
}
