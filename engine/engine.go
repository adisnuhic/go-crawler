package engine

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

// Post struct that contains main post data
type Post struct {
	ID       string
	Author   string
	PostSlug string
	Title    string
	Subtitle string
	Content  string
}

// GetWebsiteLatestPosts returns latest list of posts with: author, slug, title, subtitle
func GetWebsiteLatestPosts(l string, c chan []Post) {
	posts := []Post{}
	latest, _ := http.Get(l)
	latestBytes, _ := ioutil.ReadAll(latest.Body)
	latestJSON := string(latestBytes)[16:]
	apiPosts := gjson.Get(latestJSON, "payload.references.Post").Map()
	apiAuthor := gjson.Get(latestJSON, "payload.user")
	postAuthor := apiAuthor.Get("username").String()

	for _, ap := range apiPosts {
		post := new(Post)
		post.Author = postAuthor
		post.ID = ap.Get("id").String()
		post.PostSlug = ap.Get("uniqueSlug").String()
		post.Title = ap.Get("title").String()
		post.Subtitle = ap.Get("content.subtitle").String()
		posts = append(posts, *post)
	}

	c <- posts
}

// GetWebsiteFullPost get full posts with content
func GetWebsiteFullPost(post Post, domainRoot string, c chan Post) {
	p, _ := http.Get(domainRoot + post.Author + "/" + post.PostSlug + "?format=json")
	pByptes, _ := ioutil.ReadAll(p.Body)
	pJSON := string(pByptes)[16:]
	post.Content = gjson.Get(pJSON, "payload.value.content.bodyModel.paragraphs.#.text").String()
	c <- post
}
