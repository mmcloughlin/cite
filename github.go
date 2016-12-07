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

func lineRangeFragment(l LineRange) string {
	switch {
	case l.start < l.end:
		return fmt.Sprintf("L%d-L%d", l.start, l.end)
	case l.start == l.end:
		return fmt.Sprintf("L%d", l.start)
	default:
		panic(ErrBadLineRange)
	}
}

// Reference represents a github reference such as:
//
//	https://github.com/mmcloughlin/geohash/blob/master/LICENSE#L3-L10
type GithubResource struct {
	User       string
	Repository string
	GitRef     string
	Path       string
	LineRange  LineRange
}

func (r GithubResource) url(what string) *url.URL {
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
		Fragment: lineRangeFragment(r.LineRange),
	}
}

func (r GithubResource) URL() *url.URL {
	return r.url("raw")
}

func (r GithubResource) Cite() Citation {
	return Citation{
		URL:   r.url("blob"),
		Extra: "",
	}
}

func (r GithubResource) Lines() LineSelection {
	return r.LineRange
}
