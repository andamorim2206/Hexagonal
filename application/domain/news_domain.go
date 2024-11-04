package domain

type NewsReqDomain struct {
	Subject string
	From    string
}

type NewsDomain struct {
	Status       string
	TotalResults string
	Articles     []Articles
}

type Articles struct {
	Source      string
	Id          string
	Name        string
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}
