package socialmeta

import (
	"log"

	"github.com/caddyserver/caddy/v2"
	templates "github.com/caddyserver/caddy/v2/modules/caddyhttp/templates"
)

func init() {
	err := caddy.RegisterModule(SocialMeta{})
	if err != nil {
		log.Fatal(err)
	}

}

type SocialMeta struct {
}

// CaddyModule returns the Caddy module information.
func (SocialMeta) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "template.function.social_meta",
		New: func() caddy.Module { return new(SocialMeta) },
	}
}

func (s *SocialMeta) Provision(ctx caddy.Context) error {
	templates.TemplateFuncs["socialMeta"] = s.GetMeta
	return nil
}

func (s *SocialMeta) GetMeta(ctx templates.CustomContext) string {
	return ""
}
