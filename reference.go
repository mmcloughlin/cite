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

type LineRange struct {
	Start int
	End   int
}

func (l LineRange) String() string {
	if l.Start == l.End {
		return fmt.Sprintf("L%d", l.Start)
	}
	return fmt.Sprintf("L%d-L%d", l.Start, l.End)
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

func (r Reference) URL() *url.URL {
	path := path.Join(
		r.User,
		r.Repository,
		"blob",
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

func (r Reference) String() string {
	return r.URL().String()
}
