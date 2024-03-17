package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/ar4mirez/terraform-provider-ct/ct"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ct.Provider,
	})
}
