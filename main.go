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
	 * Create channel to communicate between Go routines for basic posts
	 */
	c1 := make(chan []engine.Post)

	/*
	 * Get basic posts data from the Web
	 */
	fmt.Println("Please wait - Fetching data...")
	links, _ := medium.AuthorsLinks() // Get medium authors list

	for _, l := range links {
		go engine.GetWebsiteLatestPosts(l, c1) // for each link
	}

	basicPosts := []engine.Post{}
	for i := 0; i < len(links); i++ {
		basicPosts = append(basicPosts, <-c1...)
	}

	/*
	 * Get Full posts data from the Web
	 */
	c2 := make(chan engine.Post)
	for _, fp := range basicPosts {
		go engine.GetWebsiteFullPost(fp, "https://medium.com/@", c2)
	}

	fullPosts := []engine.Post{}
	for j := 0; j < len(basicPosts); j++ {
		fullPosts = append(fullPosts, <-c2)
	}

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
