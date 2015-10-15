package post

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type BPost struct {
	Title           string        `json:"title"`
	Author          string        `json:"author"`
	UrlPermalink    string        `json:"url_permalink"`
	DateCreated     time.Time     `json:"date_created"`
	DateEdited      time.Time     `json:"date_edited"`
	ContentMarkdown string        `json:"content_markdown"`
	ContentHtml     template.HTML `json:"content_html"`
}

// ByDate implements sort.Interface for []BPost based on
// the DateCreated field.
type ByDate []*BPost

func (a ByDate) Len() int      { return len(a) }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool {
	if a[i].DateCreated.Unix() > a[j].DateCreated.Unix() {
		return true
	}
	if a[i].DateCreated.Unix() < a[j].DateCreated.Unix() {
		return false
	}
	return a[i].DateEdited.Unix() >= a[j].DateEdited.Unix()
}

func (p *BPost) FormattedEditedTime() string {
	return p.DateEdited.Format("January 02, 2006 | Monday -- 15:04PM")
}

func (p *BPost) FormattedCreatedTime() string {
	return p.DateCreated.Format("January 02, 2006 | Monday")
}

func (p *BPost) HTML5CreatedTime() string {
	return p.DateCreated.Format("2006-01-02")
}

func FromJson(b []byte) *BPost {
	bp := &BPost{}
	if err := json.Unmarshal(b, bp); err != nil {
		return nil
	}
	return bp
}

// FromFile reads a post folder (that follows our special structure) and creates a new
// post structure with fields filled from the file loaded.
func FromMarkdown(pathname string) (*BPost, error) {
	bp := &BPost{}

	markdown, err := ioutil.ReadFile(pathname)
	if err != nil {
		return nil, err
	}
	bytesRead, err := parseFrontMatter(bp, markdown)

	bp.ContentMarkdown = string(markdown[bytesRead:])
	bp.ContentHtml = template.HTML(string(blackfriday.MarkdownCommon(markdown[bytesRead:])))

	return bp, nil
}

func parseFrontMatter(bp *BPost, markdown []byte) (int, error) {
	bRead := 0
	scanner := bufio.NewScanner(bytes.NewReader(markdown))
	if scanner == nil {
		return -1, errors.New("Could not create reader from the markdown bytes: " + string(markdown))
	}
	lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		bRead += len(line)
		lines++
		if line == "---" {
			if lines > 1 {
				return bRead, nil
			}
		} else if lines == 1 {
			// no front-matter is defined - should start from the first line
			return 0, nil
		}
		segments := strings.Split(line, ":")
		switch segments[0] {
		case "title":
			bp.Title = strings.Trim(segments[1], " ")
			break
		default:
			log.Println(errors.New("Wrong property defined in the front-matter: " + line))
		}
	}
	return bRead, nil
}
