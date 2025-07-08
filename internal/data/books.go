package data

// BookRecommendations contains curated daily book recommendations
var BookRecommendations = []struct {
	Title       string
	Author      string
	Genre       string
	ImageURL    string
	Description string
}{
	{
		"The Seven Husbands of Evelyn Hugo",
		"Taylor Jenkins Reid",
		"Contemporary Fiction",
		"/media/books/evelyn-hugo.jpg",
		"A reclusive Hollywood icon reveals her life's secrets to an unknown journalist.",
	},
	{
		"Where the Crawdads Sing",
		"Delia Owens",
		"Mystery/Nature",
		"/media/books/crawdads-sing.jpg",
		"A mystery set in the natural world of North Carolina's marshlands.",
	},
	{
		"The Midnight Library",
		"Matt Haig",
		"Philosophical Fiction",
		"/media/books/midnight-library.png",
		"Between life and death lies the Midnight Library, where every book is a chance to try another life you could have lived.",
	},
	{
		"Educated",
		"Tara Westover",
		"Memoir",
		"/media/books/educated.jpg",
		"A powerful memoir about education and the struggle for self-invention.",
	},
}
