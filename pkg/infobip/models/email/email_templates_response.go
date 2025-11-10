package email

// EmailTemplate struct for EmailTemplate
type EmailTemplatesResponse struct {
	Paging *Paging
	// List of domains that belong to the account.
	Results []EmailTemplateResponse
}
