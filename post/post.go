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
	"path/filepath"
	"strings"
	"time"
)

type BPost struct {
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	Author          string        `json:"author"`
	UrlPermalink    string        `json:"url"`
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

func (p *BPost) FormattedEditedTimeLong() string {
	return p.DateEdited.Format("January 02, 2006 | Monday -- 15:04PM")
}

func (p *BPost) FormattedCreatedTimeLong() string {
	return p.DateCreated.Format("January 02, 2006 | Monday -- 15:04PM")
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

// FromMarkdown reads a post folder (that follows our special structure) and creates a new
// post structure with fields filled from the file loaded.
func FromMarkdown(pathname string) (*BPost, error) {
	bp := &BPost{}

	var err error

	// extract the date from the filename
	dateStr := filepath.Base(pathname)[:10]
	bp.DateCreated, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	// assign the Url to the current pathname but might be overwritten
	bp.UrlPermalink = filepath.Base(pathname)[11:]
	bp.UrlPermalink = bp.UrlPermalink[:len(bp.UrlPermalink)-3]

	markdown, err := ioutil.ReadFile(pathname)
	if err != nil {
		return nil, err
	}
	bytesRead, err := parseFrontMatter(bp, markdown)

	remBytes := markdown[bytesRead:]
	bp.ContentMarkdown = string(remBytes)
	bp.ContentHtml = template.HTML(string(blackfriday.MarkdownCommon(remBytes)))

	return bp, nil
}

func parseFrontMatter(bp *BPost, markdown []byte) (int, error) {
	bRead := 0
	scanner := bufio.NewScanner(bytes.NewReader(markdown))
	if scanner == nil {
		return -1, errors.New("Could not create reader from the markdown bytes: " + string(markdown))
	}

	foundDateEdited := false

	lines := 0
	for scanner.Scan() {
		bytesScanned := scanner.Bytes()
		bRead += len(bytesScanned)
		line := string(bytesScanned)
		lines++
		if line == "---" {
			if lines > 1 {
				// We found the second line of the matter so stop the reading
				break
			} else {
				// the first '---'
				continue
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
		case "url":
			bp.UrlPermalink = strings.Trim(segments[1], " ")
			break
		case "description":
			bp.Description = strings.Trim(segments[1], " ")
			break
		case "date-edited":
			// Ignore invalid dates.
			bp.DateEdited, _ = time.Parse("2006-01-02", strings.Trim(segments[1], " "))
			foundDateEdited = true
			break
		default:
			log.Println(errors.New("Wrong property defined in the front-matter: " + line))
		}
	}
	// add the new line bytes since they are discarded by the scanner
	bRead += len([]byte("\n")) * lines
	// Set the date edited to those created.
	if !foundDateEdited {
		bp.DateEdited = bp.DateCreated
	}
	return bRead, nil
}
