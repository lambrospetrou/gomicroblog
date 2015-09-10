package post

import (
	"github.com/russross/blackfriday"
	"html/template"
	"time"
)

type id_t string

type BPost struct {
	Id                 id_t          `json:"id"`
	Title              string        `json:"title"`
	Author             string        `json:"author"`
	DateCreated        time.Time     `json:"date_created"`
	UrlFriendlyLink    string        `json:"url_friendly_link"`
	ContentMarkdown    string        `json:"content_markdown"`
	DateEditedMarkdown time.Time     `json:"date_edited_markdown"`
	ContentHtml        string        `json:"content_html"`
	DateCompiledHtml   time.Time     `json:"date_edited_html"`
	BodyHtml           template.HTML `json:"-"`
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
	return a[i].DateEditedMarkdown.Unix() >= a[j].DateEditedMarkdown.Unix()
}

func (p *BPost) IdStr() string {
	//return strconv.Itoa(p.Id)
	return string(p.Id)
}

func (p *BPost) FormattedEditedTime() string {
	return p.DateEditedMarkdown.Format("January 02, 2006 | Monday -- 15:04PM")
}

func (p *BPost) FormattedCreatedTime() string {
	return p.DateCreated.Format("January 02, 2006 | Monday")
}

func (p *BPost) HTML5CreatedTime() string {
	return p.DateCreated.Format("2006-01-02")
}

func (p *BPost) PrepareSave() {
	// update the HTML content
	p.DateEditedMarkdown = time.Now()
	p.ContentHtml = string(blackfriday.MarkdownCommon([]byte(p.ContentMarkdown)))
	p.DateCompiledHtml = p.DateEditedMarkdown

	// set the ID the same as the URL friendly link and it will be updated
	// by the storager used if necessary
	p.Id = id_t(p.UrlFriendlyLink)
}

// New creates a new blog post returns it empty setting its creation date to time.Noe
func New() *BPost {
	return &BPost{DateCreated: time.Now()}
}
