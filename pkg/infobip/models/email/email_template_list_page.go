package email

// EmailTemplateListPage defines model for EmailTemplateListPage.
type EmailTemplateListPage struct {
	Paging  PageInfo                `json:"paging"`
	Results []EmailTemplateListItem `json:"results"`
}
