package models

type SimpleBook struct {
	BookIDEncrypt string
	Href          string
	Pid           string
	GroupID       string
	TrainID       string
	Type          string
}

type DetailedBook struct {
	Status int `json:"status"`
	Data   struct {
		BookId           string `json:"book_id"`
		BookName         string `json:"book_name"`
		BookIdf          string `json:"book_idf"`
		BookType         string `json:"book_type"`
		BookAreaId       string `json:"book_area_id"`
		BookCoverUrl     string `json:"book_cover_url"`
		CompanyId        string `json:"company_id"`
		CompanyIdf       string `json:"company_idf"`
		BookKind         string `json:"book_kind"`
		BookTrialInfo    string `json:"book_trial_info"`
		IsElastic        string `json:"is_elastic"`
		BookFinishedRule string `json:"book_finished_rule"`
	} `json:"data"`
}

type BundleBook struct {
	SimpleBook
	DetailedBook
}
