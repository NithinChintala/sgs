# Scuffed Google Scholar
Scuffed Google Scholar (SGS) is an attempt to create a small version
of an application that indexes academic journals

# Motivation
This was my CS3200 Final project. Also I use Google Scholar a lot and it
has a simple UI and seemed like a cool way to practice making relational
databases.

# Why Go?
I just wanted to get better with the language, so I decided to use it
to make this. Also Go is a pretty cool language. Though I have to say 
not having generics in CRUD based applications feels AWFUL.

# Install
To build the server just run `go build`. To populate the databse run `go run scripts/populate.go`.

# Random
I intially wanted to scrape BibTeX files from Google Scholar to populate
my databse, but quickly learned that this is explicitily made incredibly difficult.
There is a way, but it essentially gets flagged and banned immediately.

So, instead of actual academic papers I just made random data in `scripts/populate.go`
