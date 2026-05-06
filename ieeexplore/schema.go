package ieeexplore

type SearchResult struct {
	TotalRecords             int           `json:"totalRecords"`
	RecordsPerPage           int           `json:"recordsPerPage"`
	StartRecord              int           `json:"startRecord"`
	EndRecord                int           `json:"endRecord"`
	TotalPages               int           `json:"totalPages"`
	SearchType               string        `json:"searchType"`
	PromoApplied             bool          `json:"promoApplied"`
	SubscribedContentApplied bool          `json:"subscribedContentApplied"`
	ShowStandardDictionary   bool          `json:"showStandardDictionary"`
	Records                  []SearchEntry `json:"records"`
	Facets                   []Facet       `json:"facets,omitempty"`
	BreadCrumbs              []BreadCrumb  `json:"breadCrumbs,omitempty"`
}

type SearchEntry struct {
	Abstract                string            `json:"abstract,omitempty"`
	AccessType              AccessType        `json:"accessType"`
	ArticleContentType      string            `json:"articleContentType,omitempty"`
	ArticleNumber           string            `json:"articleNumber,omitempty"`
	ArticleTitle            string            `json:"articleTitle,omitempty"`
	Authors                 []Author          `json:"authors,omitempty"`
	CitationCount           int               `json:"citationCount"`
	CitationsLink           string            `json:"citationsLink,omitempty"`
	ContentType             string            `json:"contentType,omitempty"`
	Course                  bool              `json:"course"`
	DisplayContentType      string            `json:"displayContentType,omitempty"`
	DisplayPublicationTitle string            `json:"displayPublicationTitle,omitempty"`
	DocIdentifier           string            `json:"docIdentifier,omitempty"`
	DocumentLink            string            `json:"documentLink,omitempty"`
	DOI                     string            `json:"doi,omitempty"`
	DownloadCount           int               `json:"downloadCount"`
	EndPage                 string            `json:"endPage,omitempty"`
	Ephemera                bool              `json:"ephemera"`
	HandleProduct           bool              `json:"handleProduct"`
	HighlightedTitle        string            `json:"highlightedTitle,omitempty"`
	HTMLLink                string            `json:"htmlLink,omitempty"`
	IsBook                  bool              `json:"isBook"`
	IsBookWithoutChapters   bool              `json:"isBookWithoutChapters"`
	IsConference            bool              `json:"isConference"`
	IsEarlyAccess           bool              `json:"isEarlyAccess"`
	IsImmersiveArticle      bool              `json:"isImmersiveArticle"`
	IsJournal               bool              `json:"isJournal"`
	IsJournalAndMagazine    bool              `json:"isJournalAndMagazine"`
	IsMagazine              bool              `json:"isMagazine"`
	IsNumber                string            `json:"isNumber,omitempty"`
	IsOnlineOnly            bool              `json:"isOnlineOnly"`
	IsStandard              bool              `json:"isStandard"`
	Issue                   string            `json:"issue,omitempty"`
	MajorTopic              string            `json:"majorTopic,omitempty"`
	PatentCitationCount     int               `json:"patentCitationCount"`
	PDFLink                 string            `json:"pdfLink,omitempty"`
	PDFSize                 string            `json:"pdfSize,omitempty"`
	PublicationDate         string            `json:"publicationDate,omitempty"`
	PublicationLink         string            `json:"publicationLink,omitempty"`
	PublicationNumber       string            `json:"publicationNumber,omitempty"`
	PublicationTitle        string            `json:"publicationTitle,omitempty"`
	PublicationYear         string            `json:"publicationYear,omitempty"`
	Publisher               string            `json:"publisher,omitempty"`
	Redline                 bool              `json:"redline"`
	RightsLink              string            `json:"rightsLink,omitempty"`
	RightslinkFlag          bool              `json:"rightslinkFlag"`
	ShowAlgorithm           bool              `json:"showAlgorithm"`
	ShowCheckbox            bool              `json:"showCheckbox"`
	ShowDataset             bool              `json:"showDataset"`
	ShowHTML                bool              `json:"showHtml"`
	ShowVideo               bool              `json:"showVideo"`
	StartPage               string            `json:"startPage,omitempty"`
	SupplementGroup         []SupplementGroup `json:"supplementGroup,omitempty"`
	VJ                      bool              `json:"vj"`
	Volume                  string            `json:"volume,omitempty"`
	URL                     string            `json:"url,omitempty"`
	PDFURL                  string            `json:"pdfUrl,omitempty"`
	PublicationURL          string            `json:"publicationUrl,omitempty"`
}

type AccessType struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type Author struct {
	PreferredName           string `json:"preferredName,omitempty"`
	NormalizedName          string `json:"normalizedName,omitempty"`
	FirstName               string `json:"firstName,omitempty"`
	LastName                string `json:"lastName,omitempty"`
	SearchablePreferredName string `json:"searchablePreferredName,omitempty"`
	ID                      int64  `json:"id,omitempty"`
}

type SupplementGroup struct {
	Repository string `json:"repository,omitempty"`
	Type       string `json:"type,omitempty"`
	Badge      string `json:"badge,omitempty"`
}

type Facet struct {
	ID         string       `json:"id,omitempty"`
	Name       string       `json:"name,omitempty"`
	NumRecords int          `json:"numRecords"`
	Active     string       `json:"active,omitempty"`
	Children   []FacetValue `json:"children,omitempty"`
}

type FacetValue struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	NumRecords int    `json:"numRecords"`
	Active     string `json:"active,omitempty"`
}

type BreadCrumb struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}
