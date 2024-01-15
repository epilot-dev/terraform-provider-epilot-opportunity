// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-opportunity/internal/sdk/pkg/utils"
	"time"
)

type Dates struct {
	Tags  []string   `json:"_tags,omitempty"`
	Dates *string    `json:"dates,omitempty"`
	Value *time.Time `json:"value,omitempty"`
}

func (d Dates) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Dates) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Dates) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Dates) GetDates() *string {
	if o == nil {
		return nil
	}
	return o.Dates
}

func (o *Dates) GetValue() *time.Time {
	if o == nil {
		return nil
	}
	return o.Value
}

type Source struct {
	Href  *string `default:"null" json:"href"`
	Title *string `default:"manual" json:"title"`
}

func (s Source) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Source) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Source) GetHref() *string {
	if o == nil {
		return nil
	}
	return o.Href
}

func (o *Source) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type Opportunity struct {
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL       BaseEntityACL `json:"_acl"`
	CreatedAt time.Time     `json:"_created_at"`
	ID        string        `json:"_id"`
	// Organization Id the entity belongs to
	Org              string            `json:"_org"`
	Owners           []BaseEntityOwner `json:"_owners"`
	Schema           string            `json:"_schema"`
	Tags             []string          `json:"_tags"`
	Title            string            `json:"_title"`
	UpdatedAt        time.Time         `json:"_updated_at"`
	Address          *BaseRelationRef  `json:"address,omitempty"`
	BillingAddress   *BaseRelationRef  `json:"billing_address,omitempty"`
	CurrentTask      *string           `default:"open" json:"current_task"`
	Customer         *BaseRelation     `json:"customer,omitempty"`
	Dates            []Dates           `json:"dates,omitempty"`
	DeliveryAddress  *BaseRelationRef  `json:"delivery_address,omitempty"`
	Items            *BaseRelation     `json:"items,omitempty"`
	OpportunityTitle string            `json:"opportunity_title"`
	Payment          *BaseRelationRef  `json:"payment,omitempty"`
	Source           *Source           `json:"source,omitempty"`
	SourceType       *string           `default:"manual" json:"source_type"`
	Status           *string           `default:"open" json:"status"`
}

func (o Opportunity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *Opportunity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Opportunity) GetACL() BaseEntityACL {
	if o == nil {
		return BaseEntityACL{}
	}
	return o.ACL
}

func (o *Opportunity) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Opportunity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Opportunity) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *Opportunity) GetOwners() []BaseEntityOwner {
	if o == nil {
		return []BaseEntityOwner{}
	}
	return o.Owners
}

func (o *Opportunity) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *Opportunity) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Opportunity) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Opportunity) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Opportunity) GetAddress() *BaseRelationRef {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Opportunity) GetBillingAddress() *BaseRelationRef {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *Opportunity) GetCurrentTask() *string {
	if o == nil {
		return nil
	}
	return o.CurrentTask
}

func (o *Opportunity) GetCustomer() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Customer
}

func (o *Opportunity) GetDates() []Dates {
	if o == nil {
		return nil
	}
	return o.Dates
}

func (o *Opportunity) GetDeliveryAddress() *BaseRelationRef {
	if o == nil {
		return nil
	}
	return o.DeliveryAddress
}

func (o *Opportunity) GetItems() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *Opportunity) GetOpportunityTitle() string {
	if o == nil {
		return ""
	}
	return o.OpportunityTitle
}

func (o *Opportunity) GetPayment() *BaseRelationRef {
	if o == nil {
		return nil
	}
	return o.Payment
}

func (o *Opportunity) GetSource() *Source {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *Opportunity) GetSourceType() *string {
	if o == nil {
		return nil
	}
	return o.SourceType
}

func (o *Opportunity) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}