package webhook

type IWebhookRepository interface {
	Set(string, string) error
	Get(string) (string, error)
}
