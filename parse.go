package cite

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

var (
	commentExpr   = `^(\s*)//(.*)$`
	commentRegexp = regexp.MustCompile(commentExpr)
)

// CommentBlock represents a Golang block comment with "//" syntax. All lines
// of the comment have the same bytes preceeding the "//", which we call the
// leader.
type CommentBlock struct {
	Leader string
	Lines  []string
}

// String converts the CommentBlock into source code.
func (c CommentBlock) String() string {
	out := ""
	for _, line := range c.Lines {
		out += fmt.Sprintf("%s//%s\n", c.Leader, line)
	}
	return out
}

// CodeBlock represents arbitrary lines of Golang source code preceeded by a
// CommentBlock.
type CodeBlock struct {
	CommentBlock CommentBlock
	Code         []string
}

// String converts the CodeBlock into source code.
func (b CodeBlock) String() string {
	out := b.CommentBlock.String()
	for _, line := range b.Code {
		out += line + "\n"
	}
	return out
}

// Source represents a Golang source file broken up into CodeBlock objects, by
// splitting on comment blocks.
type Source struct {
	Blocks []CodeBlock
}

// String converts Source into source code.
func (s Source) String() string {
	out := ""
	for _, b := range s.Blocks {
		out += b.String()
	}
	return out
}

// ParseCode parses the code from the given io.Reader into a Source objects by
// splitting on comment blocks.
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
