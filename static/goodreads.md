# Goodreads Unofficial API

#### List items:

##### Action: *search*

##### Parametres:
`q`: Search terms

`filter`: Can be 0, 1, 2, 3 represents all, title, author, genre respectively

##### Example: [/search?q=test&filter=0&count=20](/search?q=test&filter=0&count=20)

##### Output Sample:
```    {
[
        "title": "Get a Life, Chloe Brown (The Brown Sisters, #1)",
        "cover": "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1614273529i/43884209._SY75_.jpg",
        "authors": [
            "Talia Hibbert"
        ],
        "pages": 0,
        "published": 0,
        "isbn": 0,
        "rating": 0,
        "link": "https://www.goodreads.com/book/show/43884209-get-a-life-chloe-brown?from_search=true\u0026from_srp=true\u0026qid=pTGlsZdP2O\u0026rank=19"
    },
...
]
```

#### Get single book data:

##### Action: *view*

##### Parametres:
`link`: Goodreads book URL

##### Example: [/view?link=https://www.goodreads.com/book/show/231804.The_Outsiders](/view?link=https://www.goodreads.com/book/show/231804.The_Outsiders)

##### Output Sample:
```    
{
    "title": "The Outsiders",
    "cover": "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1442129426l/231804.jpg",
    "authors": [
        "S.E. Hinton"
    ],
    "pages": 192,
    "published": 1997,
    "isbn": 0,
    "rating": 4.11,
    "link": "https://www.goodreads.com/book/show/231804.The_Outsiders"
}
```
