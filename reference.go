package cite

import (
	"fmt"
	"net/url"
	"path"
)

const (
	GithubScheme = "https"
	GithubHost   = "github.com"
)

// LineRange
type LineRange struct {
	start int
	end   int
}

func NewSingleLine(n int) LineRange {
	return NewLineRange(n, n)
}

func NewLineRange(s, e int) LineRange {
	return LineRange{
		start: s,
		end:   e,
	}
}

func (l LineRange) Start() int { return l.start }
func (l LineRange) End() int   { return l.end }

func (l LineRange) String() string {
	switch {
	case l.start < l.end:
		return fmt.Sprintf("L%d-L%d", l.start, l.end)
	case l.start == l.end:
		return fmt.Sprintf("L%d", l.start)
	default:
		return ""
	}
}

// Reference represents a github reference such as:
//
//	https://github.com/mmcloughlin/geohash/blob/master/LICENSE#L3-L10
type Reference struct {
	User       string
	Repository string
	GitRef     string
	Path       string
	Lines      LineRange
}

func (r Reference) url(what string) *url.URL {
	path := path.Join(
		r.User,
		r.Repository,
		what,
		r.GitRef,
		r.Path,
	)
	return &url.URL{
		Scheme:   GithubScheme,
		Host:     GithubHost,
		Path:     path,
		Fragment: r.Lines.String(),
	}
}

func (r Reference) BlobURL() *url.URL {
	return r.url("blob")
}

func (r Reference) RawFileURL() *url.URL {
	u := r.url("raw")
	u.Fragment = ""
	return u
}
