// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/epilot-dev/terraform-provider-epilot-opportunity/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-opportunity/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net/http"
)

var _ provider.Provider = &EpilotOpportunityProvider{}

type EpilotOpportunityProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// EpilotOpportunityProviderModel describes the provider data model.
type EpilotOpportunityProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	EpilotAuth types.String `tfsdk:"epilot_auth"`
	EpilotOrg  types.String `tfsdk:"epilot_org"`
}

func (p *EpilotOpportunityProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "epilot-opportunity"
	resp.Version = p.version
}

func (p *EpilotOpportunityProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://opportunity.sls.epilot.io)",
				Optional:            true,
				Required:            false,
			},
			"epilot_auth": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"epilot_org": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *EpilotOpportunityProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data EpilotOpportunityProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://opportunity.sls.epilot.io"
	}

	epilotAuth := new(string)
	if !data.EpilotAuth.IsUnknown() && !data.EpilotAuth.IsNull() {
		*epilotAuth = data.EpilotAuth.ValueString()
	} else {
		epilotAuth = nil
	}
	epilotOrg := new(string)
	if !data.EpilotOrg.IsUnknown() && !data.EpilotOrg.IsNull() {
		*epilotOrg = data.EpilotOrg.ValueString()
	} else {
		epilotOrg = nil
	}
	security := shared.Security{
		EpilotAuth: epilotAuth,
		EpilotOrg:  epilotOrg,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(http.DefaultClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *EpilotOpportunityProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOpportunityResource,
	}
}

func (p *EpilotOpportunityProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOpportunityDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &EpilotOpportunityProvider{
			version: version,
		}
	}
}
