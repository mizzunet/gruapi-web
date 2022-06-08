package gruapi

import (
	"fmt"
	// "net/http"
	// "time"
	// "net/url"
	"strconv"
	"strings"

	// "github.com/PuerkitoBio/goquery"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Book struct {
	TITLE     string   `json:"title"`
	COVER     string   `json:"cover"`
	AUTHORS   []string `json:"authors"`
	PAGES     int      `json:"pages"`
	PUBLISHED int      `json:"published"`
	ISBN      int      `json:"isbn"`
	RATING    float64  `json:"rating"`
	LINK      string   `json:"link"`
	// LANGAUGE     string
	// RATING    struct {
	// RATINGS string `json:"ratings_count"`
	// REVIEWS string `json:"review_count"`
	// }
}

func View(URL string) Book {
	var book Book

	c := colly.NewCollector(
		colly.CacheDir("cache"),
	)
	// t := time.Date(2029, time.November, 10, 23, 0, 0, 0, time.UTC)

	// cookie := []*http.Cookie{
	// 	{
	// 		Name:    "mobvious.device_type",
	// 		Value:   "mobile",
	// 		Domain:  "www.goodreads.com",
	// 		Path:    "/",
	// 		Expires: t,
	// 	},
	// }
	// c.SetCookies("www.goodreads.com", cookie)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// Scrape books data
	c.OnHTML("#topcol", func(e *colly.HTMLElement) {

		book.TITLE = strings.Replace(strings.TrimSpace(e.DOM.Find("#bookTitle").Text()), " pages", "", 1)
		e.DOM.Find("#bookAuthors").Find("span[itemprop='name']").Each(func(i int, s *goquery.Selection) {
			book.AUTHORS = append(book.AUTHORS, s.Text())
		})
		book.COVER = e.ChildAttr("#coverImage", "src")
		// book.DESCRIPTIONS = desc
		book.PAGES, _ = strconv.Atoi(strings.Replace(strings.TrimSpace(e.DOM.Find("span[itemprop='numberOfPages']").Text()), " pages", "", 1))
		book.ISBN, _ = strconv.Atoi(strings.Replace(strings.Split(strings.TrimSpace(e.DOM.Find("#bookDataBox > div:nth-child(2) > div.infoBoxRowItem").Text()), " ")[0], "\n", "", 1))
		book.PUBLISHED, _ = strconv.Atoi(strings.Replace(strings.Split(strings.TrimSpace(e.DOM.Find("#details > div:nth-child(2)").Text()), " ")[10], "\n", "", 1))
		book.RATING, _ = strconv.ParseFloat(strings.TrimSpace(e.DOM.Find("div[itemprop='aggregateRating']").Find("span[itemprop='ratingValue']").Text()), 2)
		// link, _ := e.DOM.Find("a[title]").Attr("href")
		book.LINK = URL
	})
	c.Visit(URL)

	return book
}

func Search(q string, filter int, count int) []Book {
	var Books []Book
	var URL string

	q = strings.ReplaceAll(q, " ", "+")
	URL = baseURL(&filter) + q

	c := colly.NewCollector(
		colly.CacheDir("cache"),
		colly.MaxDepth(2),
		colly.Async(),
	)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// c.OnHTML("body", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.DOM.Html())
	// })
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 20})

	// Scrape books data

	// Visit book pages from results
	if count == 20 {
		c.OnHTML("tr[itemtype='http://schema.org/Book']", func(e *colly.HTMLElement) {
			var book Book
			// if data == 0 {
			link, _ := e.DOM.Find("a[title]").Attr("href")

			book.TITLE = strings.TrimSpace(e.ChildText("span[role='heading']"))
			e.DOM.Find("[itemprop='author']").Find("span[itemprop='name']").Each(func(i int, s *goquery.Selection) {
				book.AUTHORS = append(book.AUTHORS, s.Text())
			})
			book.COVER = e.ChildAttr("img[itemprop='image']", "src")
			book.LINK = e.Request.AbsoluteURL(link)

			Books = append(Books, book)
			// } else {
			// 	// Get Book Links
			// 	link, _ := e.DOM.Find("a[title]").Attr("href")
			// 	// Visit book page
			// 	c.Visit(e.Request.AbsoluteURL(link))
			// }

		})
	} else {
		c.OnHTML("table > tbody", func(e *colly.HTMLElement) {
			// for i < 10 {
			e.DOM.Find(".bookTitle").Each(func(i int, s *goquery.Selection) {
				if i+1 <= count {
					link, _ := s.Attr("href")
					e.Request.Visit(link)
				}
			})
		})
	}
	c.Visit(URL)
	c.Wait()
	return Books
}

// c.OnHTML("#topcol", func(e *colly.HTMLElement) {

// 	fmt.Println("djflakdjlfjl")
// 	var book Book
// 	book.TITLE = strings.Replace(strings.TrimSpace(e.DOM.Find("#bookTitle").Text()), " pages", "", 1)
// 	e.DOM.Find("#bookAuthors").Find("span[itemprop='name']").Each(func(i int, s *goquery.Selection) {
// 		book.AUTHORS = append(book.AUTHORS, s.Text())
// 	})
// 	book.COVER = e.ChildAttr("#coverImage", "src")
// 	book.PAGES, _ = strconv.Atoi(strings.Replace(strings.TrimSpace(e.DOM.Find("span[itemprop='numberOfPages']").Text()), " pages", "", 1))
// 	book.ISBN, _ = strconv.Atoi(strings.Replace(strings.Split(strings.TrimSpace(e.DOM.Find("#bookDataBox > div:nth-child(2) > div.infoBoxRowItem").Text()), " ")[0], "\n", "", 1))
// 	book.PUBLISHED, _ = strconv.Atoi(strings.Replace(strings.Split(strings.TrimSpace(e.DOM.Find("#details > div:nth-child(2)").Text()), " ")[10], "\n", "", 1))
// 	book.RATING, _ = strconv.ParseFloat(strings.TrimSpace(e.DOM.Find("div[itemprop='aggregateRating']").Find("span[itemprop='ratingValue']").Text()), 2)
// 	// link, _ := e.DOM.Find("a[title]").Attr("href")
// 	book.LINK = URL

// 	Books = append(Books, book)
// })

func baseURL(filter *int) string {
	switch *filter {
	case 0: //"all"
		return "https://www.goodreads.com/search?search_type=books&query="
	case 1: // "title"
		return "https://www.goodreads.com/search?search_type=books&search%5Bfield%5D=title&q="
	case 2: //"author":
		return "https://www.goodreads.com/search?search_type=books&search%5Bfield%5D=author&q="
	case 3: //"genre":
		return "https://www.goodreads.com/search?search_type=books&search%5Bfield%5D=genre&q="
	default:
		return "https://www.goodreads.com/search?search_type=books&query="
	}
}

// func isValidUrl(toTest string) bool {
// 	_, err := url.ParseRequestURI(toTest)
// 	if err != nil {
// 		return false
// 	}

// 	u, err := url.Parse(toTest)
// 	if err != nil || u.Scheme == "" || u.Host == "" {
// 		return false
// 	}

// 	return true
// }
