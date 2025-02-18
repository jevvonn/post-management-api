package dto

type CreatePostRequest struct {
	Title                  string `json:"title"`
	ContentDescription     string `json:"content_description"`
	Caption                string `json:"caption"`
	Platforms              string `json:"platforms"`
	DesignLink             string `json:"design_link"`
	DeadlineBeforeRevision string `json:"deadline_before_revision"`
	DeadlineAfterRevision  string `json:"deadline_after_revision"`
	UploadDate             string `json:"upload_date"`
	Status                 string `json:"status"`
}
