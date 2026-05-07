package ieeexplore

type SearchResult struct {
	TotalRecords             int                      `json:"totalRecords"`
	RecordsPerPage           int                      `json:"recordsPerPage"`
	StartRecord              int                      `json:"startRecord"`
	EndRecord                int                      `json:"endRecord"`
	TotalPages               int                      `json:"totalPages"`
	SearchType               string                   `json:"searchType"`
	PromoApplied             bool                     `json:"promoApplied"`
	SubscribedContentApplied bool                     `json:"subscribedContentApplied"`
	ShowStandardDictionary   bool                     `json:"showStandardDictionary"`
	Records                  []SearchResultRecord     `json:"records"`
	Facets                   []SearchResultFacet      `json:"facets,omitempty"`
	BreadCrumbs              []SearchResultBreadCrumb `json:"breadCrumbs,omitempty"`
}

type SearchResultRecord struct {
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

type SearchResultFacet struct {
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

type SearchResultBreadCrumb struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type XPLGlobal struct {
	UserInfo struct {
		CustomerNameRaw               string   `json:"customerNameRaw"`
		InstitutionName               string   `json:"institutionName"`
		AuthType                      []string `json:"authType"`
		Institute                     bool     `json:"institute"`
		Member                        bool     `json:"member"`
		Individual                    bool     `json:"individual"`
		Guest                         bool     `json:"guest"`
		SubscribedContent             bool     `json:"subscribedContent"`
		FileCabinetContent            bool     `json:"fileCabinetContent"`
		FileCabinetUser               bool     `json:"fileCabinetUser"`
		DtrDeleteCookieOnLogout       bool     `json:"dtrDeleteCookieOnLogout"`
		InstitutionalFileCabinetUser  bool     `json:"institutionalFileCabinetUser"`
		Products                      string   `json:"products"`
		InstType                      string   `json:"instType"`
		UserIds                       []int    `json:"userIds"`
		DtrCookieExpirationDays       int      `json:"dtrCookieExpirationDays"`
		ShowPatentCitations           bool     `json:"showPatentCitations"`
		ShowGet802Link                bool     `json:"showGet802Link"`
		OpenURLImgLoc                 string   `json:"openUrlImgLoc"`
		OpenURLLink                   string   `json:"openUrlLink"`
		ShowOpenURLLink               bool     `json:"showOpenUrlLink"`
		Tracked                       bool     `json:"tracked"`
		RingGoldID                    string   `json:"ringGoldId"`
		InstitutionUserID             int      `json:"institutionUserId"`
		DelegatedAdmin                bool     `json:"delegatedAdmin"`
		Desktop                       bool     `json:"desktop"`
		IsInstitutionDashboardEnabled bool     `json:"isInstitutionDashboardEnabled"`
		IsInstitutionProfileEnabled   bool     `json:"isInstitutionProfileEnabled"`
		IsRoamingEnabled              bool     `json:"isRoamingEnabled"`
		IsDelegatedAdmin              bool     `json:"isDelegatedAdmin"`
		IsMdl                         bool     `json:"isMdl"`
		IsCwg                         bool     `json:"isCwg"`
		IsIel                         bool     `json:"isIel"`
		IsReadAndPublish              bool     `json:"isReadAndPublish"`
		IsAcademic                    bool     `json:"isAcademic"`
	} `json:"userInfo"`
	Authors []struct {
		Name        string   `json:"name"`
		Affiliation []string `json:"affiliation"`
		FirstName   string   `json:"firstName"`
		LastName    string   `json:"lastName"`
		ID          string   `json:"id"`
	} `json:"authors"`
	Isbn []struct {
		Format   string `json:"format"`
		Value    string `json:"value"`
		IsbnType string `json:"isbnType"`
	} `json:"isbn"`
	Issn []struct {
		Format string `json:"format"`
		Value  string `json:"value"`
	} `json:"issn"`
	ArticleNumber string `json:"articleNumber"`
	MediaPath     string `json:"mediaPath"`
	DBTime        string `json:"dbTime"`
	Metrics       struct {
		WosCitationCount    int `json:"wosCitationCount"`
		CitationCountPaper  int `json:"citationCountPaper"`
		CitationCountPatent int `json:"citationCountPatent"`
		TotalDownloads      int `json:"totalDownloads"`
	} `json:"metrics"`
	PurchaseOptions struct {
		ShowOtherFormatPricingTab                     bool   `json:"showOtherFormatPricingTab"`
		ShowPdfFormatPricingTab                       bool   `json:"showPdfFormatPricingTab"`
		PdfPricingInfoAvailable                       bool   `json:"pdfPricingInfoAvailable"`
		OtherPricingInfoAvailable                     bool   `json:"otherPricingInfoAvailable"`
		MandatoryBundle                               bool   `json:"mandatoryBundle"`
		OptionalBundle                                bool   `json:"optionalBundle"`
		DisplayTextWhenOtherFormatPricingNotAvailable string `json:"displayTextWhenOtherFormatPricingNotAvailable"`
		PdfPricingInfo                                []struct {
			MemberPrice    string `json:"memberPrice"`
			NonMemberPrice string `json:"nonMemberPrice"`
			PartNumber     string `json:"partNumber"`
			Type           string `json:"type"`
		} `json:"pdfPricingInfo"`
	} `json:"purchaseOptions"`
	GetProgramTermsAccepted bool `json:"getProgramTermsAccepted"`
	Sections                struct {
		Abstract       string `json:"abstract"`
		Authors        string `json:"authors"`
		Disclaimer     string `json:"disclaimer"`
		Figures        string `json:"figures"`
		Multimedia     string `json:"multimedia"`
		References     string `json:"references"`
		Citedby        string `json:"citedby"`
		Keywords       string `json:"keywords"`
		Definitions    string `json:"definitions"`
		Algorithm      string `json:"algorithm"`
		Dataset        string `json:"dataset"`
		Cadmore        string `json:"cadmore"`
		Footnotes      string `json:"footnotes"`
		RelatedContent string `json:"relatedContent"`
		Metrics        string `json:"metrics"`
	} `json:"sections"`
	Abstract                    string `json:"abstract"`
	PublicationDate             string `json:"publicationDate"`
	PublicationYear             string `json:"publicationYear"`
	PdfPath                     string `json:"pdfPath"`
	PublicationTitle            string `json:"publicationTitle"`
	PdfURL                      string `json:"pdfUrl"`
	FormulaStrippedArticleTitle string `json:"formulaStrippedArticleTitle"`
	DisplayPublicationTitle     string `json:"displayPublicationTitle"`
	StartPage                   string `json:"startPage"`
	EndPage                     string `json:"endPage"`
	Keywords                    []struct {
		Type string   `json:"type"`
		Kwd  []string `json:"kwd"`
	} `json:"keywords"`
	PubLink                  string `json:"pubLink"`
	IssueLink                string `json:"issueLink"`
	AuthorNames              string `json:"authorNames"`
	ArticleCopyRight         string `json:"articleCopyRight"`
	RightsLink               string `json:"rightsLink"`
	DisplayPublicationDate   string `json:"displayPublicationDate"`
	DateOfInsertion          string `json:"dateOfInsertion"`
	AllowComments            bool   `json:"allowComments"`
	DOILink                  string `json:"doiLink"`
	IsGetAddressInfoCaptured bool   `json:"isGetAddressInfoCaptured"`
	IsMarketingOptIn         bool   `json:"isMarketingOptIn"`
	PubTopics                []struct {
		Name string `json:"name"`
	} `json:"pubTopics"`
	Publisher             string `json:"publisher"`
	DisplayDocTitle       string `json:"displayDocTitle"`
	IsFreeDocument        bool   `json:"isFreeDocument"`
	IsJournal             bool   `json:"isJournal"`
	IsBook                bool   `json:"isBook"`
	IsBookWithoutChapters bool   `json:"isBookWithoutChapters"`
	IsTutorial            bool   `json:"isTutorial"`
	IsDynamicHTML         bool   `json:"isDynamicHtml"`
	IsSpringer            bool   `json:"isSpringer"`
	IsStandard            bool   `json:"isStandard"`
	IsChapter             bool   `json:"isChapter"`
	IsPromo               bool   `json:"isPromo"`
	IsEarlyAccess         bool   `json:"isEarlyAccess"`
	IsOpenAccess          bool   `json:"isOpenAccess"`
	IsEphemera            bool   `json:"isEphemera"`
	IsConference          bool   `json:"isConference"`
	IsOnlineOnly          bool   `json:"isOnlineOnly"`
	XploreDocumentType    string `json:"xploreDocumentType"`
	IsProduct             bool   `json:"isProduct"`
	IsLatestStandard      bool   `json:"isLatestStandard"`
	HTMLLink              string `json:"htmlLink"`
	IsTranslation         bool   `json:"isTranslation"`
	IsGiveaway            bool   `json:"isGiveaway"`
	PersistentLink        string `json:"persistentLink"`
	IsDegruyter           bool   `json:"isDegruyter"`
	IsCustomDenial        bool   `json:"isCustomDenial"`
	IsOUP                 bool   `json:"isOUP"`
	IsSMPTE               bool   `json:"isSMPTE"`
	IsSAE                 bool   `json:"isSAE"`
	IsNow                 bool   `json:"isNow"`
	HTMLAbstractLink      string `json:"htmlAbstractLink"`
	HasStandardVersions   bool   `json:"hasStandardVersions"`
	OpenAccessFlag        string `json:"openAccessFlag"`
	InsertDate            string `json:"insertDate"`
	EphemeraFlag          string `json:"ephemeraFlag"`
	Title                 string `json:"title"`
	ConfLoc               string `json:"confLoc"`
	HTMLFlag              string `json:"html_flag"`
	MlHTMLFlag            string `json:"ml_html_flag"`
	SourcePdf             string `json:"sourcePdf"`
	MlTime                string `json:"mlTime"`
	XplorePubID           string `json:"xplore-pub-id"`
	IsNumber              string `json:"isNumber"`
	RightsLinkFlag        string `json:"rightsLinkFlag"`
	ContentType           string `json:"contentType"`
	PublicationNumber     string `json:"publicationNumber"`
	CitationCount         string `json:"citationCount"`
	XploreIssue           string `json:"xplore-issue"`
	ArticleID             string `json:"articleId"`
	ContentTypeDisplay    string `json:"contentTypeDisplay"`
	ReferenceCount        int    `json:"referenceCount"`
	SubType               string `json:"subType"`
	Value                 string `json:"_value"`
	Lastupdate            string `json:"lastupdate"`
	DOI                   string `json:"doi"`
}
