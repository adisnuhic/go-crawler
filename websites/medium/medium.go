package medium

// AuthorsLinks returns list of links from medium.com website
func AuthorsLinks() ([]string, string) {
	domainRoot := "https://medium.com/"
	links := []string{
		"https://medium.com/@LanceUlanoff/latest?format=json",

		"https://medium.com/@lexikon1/latest?format=json",
		"https://medium.com/@BonoU2/latest?format=json",
		"https://medium.com/@katiethurmes/latest?format=json",
	}
	return links, domainRoot
}
