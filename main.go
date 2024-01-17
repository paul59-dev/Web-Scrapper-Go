package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("table#top20", func(table *colly.HTMLElement) {
		table.ForEach("thead", func(_ int, row *colly.HTMLElement) {
			// Initialize a variable to store the concatenated values
			var result string = ""

			// Iterate over each cell (td) in the row
			row.ForEach("th", func(_ int, th *colly.HTMLElement) {
				// Concatenate the text value with a space and a dash
				row.ForEach("th:nth-child(3)", func(_ int, th *colly.HTMLElement) {
					th.Text = ""
				})
				if th.Text == "" {
					result += ""
				} else {
					result += th.Text + " - "
				}

			})

			// Remove the trailing space and dash, then print the result
			fmt.Println(strings.TrimSuffix(result, " - "))
		})

		// Iterate over each row in the table
		table.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			// Initialize a variable to store the concatenated values
			var result string = ""

			// Iterate over each cell (td) in the row
			row.ForEach("td", func(_ int, td *colly.HTMLElement) {
				// Concatenate the text value with a space and a dash
				if td.Text == "" {
					result += ""
				} else {
					result += td.Text + "	- 	"
				}

			})

			// Remove the trailing space and dash, then print the result
			fmt.Println(strings.TrimSuffix(result, ""))
		})
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	err := c.Visit("https://www.tiobe.com/tiobe-index/")
	if err != nil {
		fmt.Println("Erreur lors de la visite de la page :", err)
	}
}
