# gomicroblog
Micro blog service powered by Go - blog.lambrospetrou.com will be based on this!

## Work-flow

The blog will have three folders related to blog articles as follows:
1. *published* contains all the articles ready for publishing
2. *pending* contains all the finished articles that are yet to be moved to *published* folder
3. *drafts* contains all unfinished articles that are in progress and should not be published when the generator creates the website

### Blog Article

Each post article has its own directory/folder and inside that folder reside all images, files and markdown code for the article.

For example, assuming a blog post with URL firendly name *how-to-make-a-blog* we have the following folder structure:

Folder name: *how-to-make-a-blog*
And inside that folder we have:
- imgs/
- files/
- how-to-make-a-blog.md

As you can see the directory of the post contains all the information that makeup that article, including all images, files referenced and the markup code for the article text.

The post directory will exist in **only one** of the three post stages described above depending on the post status (*draft*, *pending*, *published*).

## Site generation






