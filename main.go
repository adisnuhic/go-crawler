package main

import (
	"fmt"
	"go_crawler/db"
	"go_crawler/engine"
	"go_crawler/websites/medium"
)

func main() {

	/*
	 * Setup & Get instance of ArrangoDB
	 */
	instance := db.SetupArrangoDB()

	/*
	 * Get data from Web
	 */
	fmt.Println("Please wait - Fetching data...")
	links, _ := medium.AuthorsLinks()                                           // Get medium authors list
	posts := engine.GetWebsiteLatestPosts(links)                                // Get basic post data (postid,author,title,subtitle,slug), maybe it can be useful for listing basic data
	fullPosts := engine.GetWebsiteLatesFullPosts(posts, "https://medium.com/@") // Get full posts data (postid, author, title, subtitle, slug, content)

	/*
	 * Make document pattern & fill it with data & call func to create Document for each of them
	 */
	fmt.Println("Please wait - Inserting data...")
	doc := make(map[string]interface{})
	for _, fp := range fullPosts {
		doc["ID"] = fp.ID
		doc["Author"] = fp.Author
		doc["Title"] = fp.Title
		doc["Subtitle"] = fp.Subtitle
		doc["Slug"] = fp.PostSlug
		doc["Content"] = fp.Content
		db.CreateDocument(doc, instance) // create document in DB
	}

	fmt.Println("...FINISHED")

}
