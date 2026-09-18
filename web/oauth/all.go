package oauth

import (
	_ "github.com/icelee123/komari125/web/oauth/cloudflare"
	_ "github.com/icelee123/komari125/web/oauth/factory"
	_ "github.com/icelee123/komari125/web/oauth/generic"
	_ "github.com/icelee123/komari125/web/oauth/github"
	_ "github.com/icelee123/komari125/web/oauth/qq"
)

func All() {
	//empty function to ensure all OIDC providers are registered
}
