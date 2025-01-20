package channelData

import "html/template"

// An entity that holds info for sending an email.
type MailData struct {
	ToName,
	ToAddress,
	FromName,
	FromAddress,
	Subject,
	Template string

	AdditionalTo,
	CC,
	Attachments []string

	Content   template.HTML
	UseHermes bool
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	RowSets   map[string]any
}

// An aggregate that actions will be performed on.
type MailJob struct {
	MailMessage MailData
}
