package google

// https://developers.google.com/custom-search/v1/using_rest

// CustomSearchResponse ...
type CustomSearchResponse struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			CX             string `json:"cx"`
			SearchType     string `json:"searchType"`
		} `json:"nextPage"`
		Request []CustomSearchRequest `json:"request"`
	}
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []CustomSearchItem `json:"items"`
	// if error
	Error struct {
		Errors []struct {
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
			Message string `json:"message"`
		} `json:"errors"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type CustomSearchRequest struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	CX             string `json:"cx"`
	SearchType     string `json:"searchType"`
}

// CustomSearchItem ...
type CustomSearchItem struct {
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	HTMLTitle   string `json:"htmlTitle"`
	Link        string `json:"link"`
	DisplayLink string `json:"displayLink"`
	Snippet     string `json:"snippet"`
	HTMLSnippet string `json:"htmlSnippet"`
	Mime        string `json:"mime"`
	FileFormat  string `json:"fileFormat"`
	Image       struct {
		ContextLink     string `json:"contextLink"`
		Height          int    `json:"height"`
		Width           int    `json:"width"`
		ByteSize        int64  `json:"byteSize"`
		ThumbnailLink   string `json:"thumbnailLink"`
		ThumbnailHeight int    `json:"thumbnailHeight"`
		ThumbnailWidth  int    `json:"thumbnailWidth"`
	}
}
