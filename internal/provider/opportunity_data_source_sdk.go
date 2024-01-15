// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-opportunity/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *OpportunityDataSourceModel) RefreshFromSharedOpportunity(resp *shared.Opportunity) {
	if resp.ACL.AdditionalProperties == nil {
		r.ACL.AdditionalProperties = types.StringNull()
	} else {
		additionalPropertiesResult, _ := json.Marshal(resp.ACL.AdditionalProperties)
		r.ACL.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
	}
	r.ACL.Delete = nil
	for _, v := range resp.ACL.Delete {
		r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
	}
	r.ACL.Edit = nil
	for _, v := range resp.ACL.Edit {
		r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
	}
	r.ACL.View = nil
	for _, v := range resp.ACL.View {
		r.ACL.View = append(r.ACL.View, types.StringValue(v))
	}
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	r.Org = types.StringValue(resp.Org)
	if len(r.Owners) > len(resp.Owners) {
		r.Owners = r.Owners[:len(resp.Owners)]
	}
	for ownersCount, ownersItem := range resp.Owners {
		var owners1 BaseEntityOwner
		owners1.OrgID = types.StringValue(ownersItem.OrgID)
		owners1.UserID = types.StringPointerValue(ownersItem.UserID)
		if ownersCount+1 > len(r.Owners) {
			r.Owners = append(r.Owners, owners1)
		} else {
			r.Owners[ownersCount].OrgID = owners1.OrgID
			r.Owners[ownersCount].UserID = owners1.UserID
		}
	}
	r.Schema = types.StringValue(resp.Schema)
	r.Tags = nil
	for _, v := range resp.Tags {
		r.Tags = append(r.Tags, types.StringValue(v))
	}
	r.Title = types.StringValue(resp.Title)
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	if resp.Address == nil {
		r.Address = nil
	} else {
		r.Address = &BaseRelationRef{}
		if len(r.Address.DollarRelationRef) > len(resp.Address.DollarRelationRef) {
			r.Address.DollarRelationRef = r.Address.DollarRelationRef[:len(resp.Address.DollarRelationRef)]
		}
		for dollarRelationRefCount, dollarRelationRefItem := range resp.Address.DollarRelationRef {
			var dollarRelationRef1 DollarRelationRef
			dollarRelationRef1.ID = types.StringPointerValue(dollarRelationRefItem.ID)
			dollarRelationRef1.Tags = nil
			for _, v := range dollarRelationRefItem.Tags {
				dollarRelationRef1.Tags = append(dollarRelationRef1.Tags, types.StringValue(v))
			}
			dollarRelationRef1.EntityID = types.StringPointerValue(dollarRelationRefItem.EntityID)
			dollarRelationRef1.Path = types.StringPointerValue(dollarRelationRefItem.Path)
			if dollarRelationRefCount+1 > len(r.Address.DollarRelationRef) {
				r.Address.DollarRelationRef = append(r.Address.DollarRelationRef, dollarRelationRef1)
			} else {
				r.Address.DollarRelationRef[dollarRelationRefCount].ID = dollarRelationRef1.ID
				r.Address.DollarRelationRef[dollarRelationRefCount].Tags = dollarRelationRef1.Tags
				r.Address.DollarRelationRef[dollarRelationRefCount].EntityID = dollarRelationRef1.EntityID
				r.Address.DollarRelationRef[dollarRelationRefCount].Path = dollarRelationRef1.Path
			}
		}
	}
	if resp.BillingAddress == nil {
		r.BillingAddress = nil
	} else {
		r.BillingAddress = &BaseRelationRef{}
		if len(r.BillingAddress.DollarRelationRef) > len(resp.BillingAddress.DollarRelationRef) {
			r.BillingAddress.DollarRelationRef = r.BillingAddress.DollarRelationRef[:len(resp.BillingAddress.DollarRelationRef)]
		}
		for dollarRelationRefCount1, dollarRelationRefItem1 := range resp.BillingAddress.DollarRelationRef {
			var dollarRelationRef3 DollarRelationRef
			dollarRelationRef3.ID = types.StringPointerValue(dollarRelationRefItem1.ID)
			dollarRelationRef3.Tags = nil
			for _, v := range dollarRelationRefItem1.Tags {
				dollarRelationRef3.Tags = append(dollarRelationRef3.Tags, types.StringValue(v))
			}
			dollarRelationRef3.EntityID = types.StringPointerValue(dollarRelationRefItem1.EntityID)
			dollarRelationRef3.Path = types.StringPointerValue(dollarRelationRefItem1.Path)
			if dollarRelationRefCount1+1 > len(r.BillingAddress.DollarRelationRef) {
				r.BillingAddress.DollarRelationRef = append(r.BillingAddress.DollarRelationRef, dollarRelationRef3)
			} else {
				r.BillingAddress.DollarRelationRef[dollarRelationRefCount1].ID = dollarRelationRef3.ID
				r.BillingAddress.DollarRelationRef[dollarRelationRefCount1].Tags = dollarRelationRef3.Tags
				r.BillingAddress.DollarRelationRef[dollarRelationRefCount1].EntityID = dollarRelationRef3.EntityID
				r.BillingAddress.DollarRelationRef[dollarRelationRefCount1].Path = dollarRelationRef3.Path
			}
		}
	}
	r.CurrentTask = types.StringPointerValue(resp.CurrentTask)
	if resp.Customer == nil {
		r.Customer = nil
	} else {
		r.Customer = &BaseRelation{}
		if len(r.Customer.DollarRelation) > len(resp.Customer.DollarRelation) {
			r.Customer.DollarRelation = r.Customer.DollarRelation[:len(resp.Customer.DollarRelation)]
		}
		for dollarRelationCount, dollarRelationItem := range resp.Customer.DollarRelation {
			var dollarRelation1 DollarRelation
			dollarRelation1.Tags = nil
			for _, v := range dollarRelationItem.Tags {
				dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
			}
			dollarRelation1.EntityID = types.StringPointerValue(dollarRelationItem.EntityID)
			if dollarRelationCount+1 > len(r.Customer.DollarRelation) {
				r.Customer.DollarRelation = append(r.Customer.DollarRelation, dollarRelation1)
			} else {
				r.Customer.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
				r.Customer.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
			}
		}
	}
	if len(r.Dates) > len(resp.Dates) {
		r.Dates = r.Dates[:len(resp.Dates)]
	}
	for datesCount, datesItem := range resp.Dates {
		var dates1 Dates
		dates1.Tags = nil
		for _, v := range datesItem.Tags {
			dates1.Tags = append(dates1.Tags, types.StringValue(v))
		}
		dates1.Dates = types.StringPointerValue(datesItem.Dates)
		if datesItem.Value != nil {
			dates1.Value = types.StringValue(datesItem.Value.Format(time.RFC3339Nano))
		} else {
			dates1.Value = types.StringNull()
		}
		if datesCount+1 > len(r.Dates) {
			r.Dates = append(r.Dates, dates1)
		} else {
			r.Dates[datesCount].Tags = dates1.Tags
			r.Dates[datesCount].Dates = dates1.Dates
			r.Dates[datesCount].Value = dates1.Value
		}
	}
	if resp.DeliveryAddress == nil {
		r.DeliveryAddress = nil
	} else {
		r.DeliveryAddress = &BaseRelationRef{}
		if len(r.DeliveryAddress.DollarRelationRef) > len(resp.DeliveryAddress.DollarRelationRef) {
			r.DeliveryAddress.DollarRelationRef = r.DeliveryAddress.DollarRelationRef[:len(resp.DeliveryAddress.DollarRelationRef)]
		}
		for dollarRelationRefCount2, dollarRelationRefItem2 := range resp.DeliveryAddress.DollarRelationRef {
			var dollarRelationRef5 DollarRelationRef
			dollarRelationRef5.ID = types.StringPointerValue(dollarRelationRefItem2.ID)
			dollarRelationRef5.Tags = nil
			for _, v := range dollarRelationRefItem2.Tags {
				dollarRelationRef5.Tags = append(dollarRelationRef5.Tags, types.StringValue(v))
			}
			dollarRelationRef5.EntityID = types.StringPointerValue(dollarRelationRefItem2.EntityID)
			dollarRelationRef5.Path = types.StringPointerValue(dollarRelationRefItem2.Path)
			if dollarRelationRefCount2+1 > len(r.DeliveryAddress.DollarRelationRef) {
				r.DeliveryAddress.DollarRelationRef = append(r.DeliveryAddress.DollarRelationRef, dollarRelationRef5)
			} else {
				r.DeliveryAddress.DollarRelationRef[dollarRelationRefCount2].ID = dollarRelationRef5.ID
				r.DeliveryAddress.DollarRelationRef[dollarRelationRefCount2].Tags = dollarRelationRef5.Tags
				r.DeliveryAddress.DollarRelationRef[dollarRelationRefCount2].EntityID = dollarRelationRef5.EntityID
				r.DeliveryAddress.DollarRelationRef[dollarRelationRefCount2].Path = dollarRelationRef5.Path
			}
		}
	}
	r.ID = types.StringValue(resp.ID)
	if resp.Items == nil {
		r.Items = nil
	} else {
		r.Items = &BaseRelation{}
		if len(r.Items.DollarRelation) > len(resp.Items.DollarRelation) {
			r.Items.DollarRelation = r.Items.DollarRelation[:len(resp.Items.DollarRelation)]
		}
		for dollarRelationCount1, dollarRelationItem1 := range resp.Items.DollarRelation {
			var dollarRelation3 DollarRelation
			dollarRelation3.Tags = nil
			for _, v := range dollarRelationItem1.Tags {
				dollarRelation3.Tags = append(dollarRelation3.Tags, types.StringValue(v))
			}
			dollarRelation3.EntityID = types.StringPointerValue(dollarRelationItem1.EntityID)
			if dollarRelationCount1+1 > len(r.Items.DollarRelation) {
				r.Items.DollarRelation = append(r.Items.DollarRelation, dollarRelation3)
			} else {
				r.Items.DollarRelation[dollarRelationCount1].Tags = dollarRelation3.Tags
				r.Items.DollarRelation[dollarRelationCount1].EntityID = dollarRelation3.EntityID
			}
		}
	}
	r.OpportunityTitle = types.StringValue(resp.OpportunityTitle)
	if resp.Payment == nil {
		r.Payment = nil
	} else {
		r.Payment = &BaseRelationRef{}
		if len(r.Payment.DollarRelationRef) > len(resp.Payment.DollarRelationRef) {
			r.Payment.DollarRelationRef = r.Payment.DollarRelationRef[:len(resp.Payment.DollarRelationRef)]
		}
		for dollarRelationRefCount3, dollarRelationRefItem3 := range resp.Payment.DollarRelationRef {
			var dollarRelationRef7 DollarRelationRef
			dollarRelationRef7.ID = types.StringPointerValue(dollarRelationRefItem3.ID)
			dollarRelationRef7.Tags = nil
			for _, v := range dollarRelationRefItem3.Tags {
				dollarRelationRef7.Tags = append(dollarRelationRef7.Tags, types.StringValue(v))
			}
			dollarRelationRef7.EntityID = types.StringPointerValue(dollarRelationRefItem3.EntityID)
			dollarRelationRef7.Path = types.StringPointerValue(dollarRelationRefItem3.Path)
			if dollarRelationRefCount3+1 > len(r.Payment.DollarRelationRef) {
				r.Payment.DollarRelationRef = append(r.Payment.DollarRelationRef, dollarRelationRef7)
			} else {
				r.Payment.DollarRelationRef[dollarRelationRefCount3].ID = dollarRelationRef7.ID
				r.Payment.DollarRelationRef[dollarRelationRefCount3].Tags = dollarRelationRef7.Tags
				r.Payment.DollarRelationRef[dollarRelationRefCount3].EntityID = dollarRelationRef7.EntityID
				r.Payment.DollarRelationRef[dollarRelationRefCount3].Path = dollarRelationRef7.Path
			}
		}
	}
	if resp.Source == nil {
		r.Source = nil
	} else {
		r.Source = &OpportunityCreateSource{}
		r.Source.Href = types.StringPointerValue(resp.Source.Href)
		r.Source.Title = types.StringPointerValue(resp.Source.Title)
	}
	r.SourceType = types.StringPointerValue(resp.SourceType)
	r.Status = types.StringPointerValue(resp.Status)
}