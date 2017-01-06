package cite

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// Github URL parameters.
const (
	GithubScheme = "https"
	GithubHost   = "github.com"
)

// Types of errors from parsing a github citation.
var (
	ErrGithubMalformedFragment = errors.New("cite: malformed github line range fragment")
	ErrGithubWrongScheme       = errors.New("cite: incorrect scheme for github url")
	ErrGithubShortPath         = errors.New("cite: github url path is too short")
	ErrGithubNotBlob           = errors.New("cite: expected blob github url")
	ErrGithubMissingFragment   = errors.New("cite: missing fragment in github url")
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

var fragmentRegexp = regexp.MustCompile(`^L(\d+)(-L(\d+))?$`)

func parseLineRangeFragment(fragment string) (LineRange, error) {
	match := fragmentRegexp.FindStringSubmatch(fragment)
	if match == nil {
		return LineRange{}, ErrGithubMalformedFragment
	}

	startStr := match[1]
	endStr := match[3]

	start, _ := strconv.Atoi(startStr)

	if endStr == "" {
		return NewSingleLine(start), nil
	}

	end, _ := strconv.Atoi(endStr)

	return NewLineRange(start, end)
}

// GithubResource represents a snippet on github such as:
//
//	https://github.com/mmcloughlin/geohash/blob/master/LICENSE#L3-L10
type GithubResource struct {
	User       string
	Repository string
	GitRef     string
	Path       string
	LineRange  LineRange
}

// BuildGithubResourceFromCitation constructs a GithubResource from a
// reference to it.
func BuildGithubResourceFromCitation(c Citation) (Resource, error) {
	u := c.URL

	if u.Host != GithubHost {
		return nil, nil
	}

	if u.Scheme != GithubScheme {
		return nil, ErrGithubWrongScheme
	}

	parts := strings.Split(u.Path, "/")
	if len(parts) < 6 {
		return nil, ErrGithubShortPath
	}

	if parts[3] != "blob" {
		return nil, ErrGithubNotBlob
	}

	if u.Fragment == "" {
		return nil, ErrGithubMissingFragment
	}

	lines, err := parseLineRangeFragment(u.Fragment)
	if err != nil {
		return nil, err
	}

	return GithubResource{
		User:       parts[1],
		Repository: parts[2],
		GitRef:     parts[4],
		Path:       path.Join(parts[5:]...),
		LineRange:  lines,
	}, nil
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

// URL returns the URL of the full (raw) file.
func (r GithubResource) URL() *url.URL {
	u := r.url("raw")
	u.Fragment = ""
	return u
}

// Cite returns a reference to the resource.
func (r GithubResource) Cite() Citation {
	return Citation{
		URL:   r.url("blob"),
		Extra: "",
	}
}

// Lines returns the selection of lines in the snippet.
func (r GithubResource) Lines() LineSelection {
	return r.LineRange
}
