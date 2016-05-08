package main

type GoogleResults struct {
    Kind string `json:"kind"`
    URL struct {
        Type string `json:"type"`
        Template string `json:"template"`
    } `json:"url"`
    Queries struct {
        Request []struct {
            Title string `json:"title"`
            TotalResults string `json:"totalResults"`
            SearchTerms string `json:"searchTerms"`
            Count int `json:"count"`
            StartIndex int `json:"startIndex"`
            InputEncoding string `json:"inputEncoding"`
            OutputEncoding string `json:"outputEncoding"`
            Safe string `json:"safe"`
            Cx string `json:"cx"`
        } `json:"request"`
        NextPage []struct {
            Title string `json:"title"`
            TotalResults string `json:"totalResults"`
            SearchTerms string `json:"searchTerms"`
            Count int `json:"count"`
            StartIndex int `json:"startIndex"`
            InputEncoding string `json:"inputEncoding"`
            OutputEncoding string `json:"outputEncoding"`
            Safe string `json:"safe"`
            Cx string `json:"cx"`
        } `json:"nextPage"`
    } `json:"queries"`
    Context struct {
        Title string `json:"title"`
    } `json:"context"`
    SearchInformation struct {
        SearchTime float64 `json:"searchTime"`
        FormattedSearchTime string `json:"formattedSearchTime"`
        TotalResults string `json:"totalResults"`
        FormattedTotalResults string `json:"formattedTotalResults"`
    } `json:"searchInformation"`
    Items []struct {
        Kind string `json:"kind"`
        Title string `json:"title"`
        HTMLTitle string `json:"htmlTitle"`
        Link string `json:"link"`
        DisplayLink string `json:"displayLink"`
        Snippet string `json:"snippet"`
        HTMLSnippet string `json:"htmlSnippet"`
        CacheID string `json:"cacheId"`
        FormattedURL string `json:"formattedUrl"`
        HTMLFormattedURL string `json:"htmlFormattedUrl"`
        Pagemap struct {
            CseThumbnail []struct {
                Width string `json:"width"`
                Height string `json:"height"`
                Src string `json:"src"`
            } `json:"cse_thumbnail"`
            Metatags []struct {
                Viewport string `json:"viewport"`
                OgTitle string `json:"og:title"`
                OgSiteName string `json:"og:site_name"`
                OgDescription string `json:"og:description"`
                OgImage string `json:"og:image"`
                MsapplicationTilecolor string `json:"msapplication-tilecolor"`
                MsapplicationTileimage string `json:"msapplication-tileimage"`
            } `json:"metatags"`
            CseImage []struct {
                Src string `json:"src"`
            } `json:"cse_image"`
        } `json:"pagemap,omitempty"`
        Mime string `json:"mime,omitempty"`
        FileFormat string `json:"fileFormat,omitempty"`
    } `json:"items"`
}