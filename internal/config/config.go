package config

import (
	"text/template"
	channelData "vigilante/internal/channel-data"
	driver "vigilante/internal/db-driver"

	"github.com/alexedwards/scs/v2"
	"github.com/pusher/pusher-http-go"
	"github.com/robfig/cron/v3"
)

type AppConfig struct {
	DB            *driver.DB
	Session       *scs.SessionManager
	InProduction  bool
	MonitorMap    map[int]cron.EntryID
	Scheduler     *cron.Cron
	PreferenceMap map[string]string
	WsClient      pusher.Client
	TemplateCache map[string]*template.Template
	MailQueue     chan channelData.MailJob
	PuserSecret,
	Domain,
	Version,
	Identifier string
}
