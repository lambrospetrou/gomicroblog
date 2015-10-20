# MicroBlog
Micro blog service powered by Go - blog.lambrospetrou.com will be based on this!

I will use *Go* [templates](http://golang.org/pkg/text/template/ "Golang templates") for the generation of the web pages.

## Directory structure

The blog articles will be in the *_posts* directory.

### Blog Article

Each post article will have its own directory/folder and inside that folder reside all images, files, markdown code and html for the article.

For example, assuming a blog post with URL firendly name *how-to-make-a-blog* we have the following folder structure:

Folder name: *how-to-make-a-blog*
And inside that folder we have:
- data/
- how-to-make-a-blog.md

## Site generation

In order to generate the website we run the *site generator* which reads from the *published* folder all the articles and based on the html templates located in the directory *_layouts* creates the static website and puts everything in the directory *_site*.

The *post.html* is used to render single posts and the *index.html* is used to render the index page for the articles.

### *_sites* directory

The generated output directory has the following structure.

**_sites**
- *articles/:article-url/:post-data:* 
- *s/* contains static files to be imported
	* _css/_ the css files
	* _libs/_ any javascript libraries
- *index.html* the home index of the website (contains all the blog articles titles and dates, links to CV and bio/work)

## Configuration

Configuration and user-custom variables will be added in a later version of the site generator.








