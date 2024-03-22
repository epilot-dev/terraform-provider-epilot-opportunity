// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-opportunity/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-opportunity/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *OpportunityResourceModel) ToSharedOpportunityCreate() *shared.OpportunityCreate {
	var address *shared.BaseRelationRef
	if r.Address != nil {
		var dollarRelationRef []shared.DollarRelationRef = nil
		for _, dollarRelationRefItem := range r.Address.DollarRelationRef {
			id := new(string)
			if !dollarRelationRefItem.ID.IsUnknown() && !dollarRelationRefItem.ID.IsNull() {
				*id = dollarRelationRefItem.ID.ValueString()
			} else {
				id = nil
			}
			var tags []string = nil
			for _, tagsItem := range dollarRelationRefItem.Tags {
				tags = append(tags, tagsItem.ValueString())
			}
			entityID := new(string)
			if !dollarRelationRefItem.EntityID.IsUnknown() && !dollarRelationRefItem.EntityID.IsNull() {
				*entityID = dollarRelationRefItem.EntityID.ValueString()
			} else {
				entityID = nil
			}
			path := new(string)
			if !dollarRelationRefItem.Path.IsUnknown() && !dollarRelationRefItem.Path.IsNull() {
				*path = dollarRelationRefItem.Path.ValueString()
			} else {
				path = nil
			}
			dollarRelationRef = append(dollarRelationRef, shared.DollarRelationRef{
				ID:       id,
				Tags:     tags,
				EntityID: entityID,
				Path:     path,
			})
		}
		address = &shared.BaseRelationRef{
			DollarRelationRef: dollarRelationRef,
		}
	}
	var billingAddress *shared.BaseRelationRef
	if r.BillingAddress != nil {
		var dollarRelationRef1 []shared.DollarRelationRef = nil
		for _, dollarRelationRefItem1 := range r.BillingAddress.DollarRelationRef {
			id1 := new(string)
			if !dollarRelationRefItem1.ID.IsUnknown() && !dollarRelationRefItem1.ID.IsNull() {
				*id1 = dollarRelationRefItem1.ID.ValueString()
			} else {
				id1 = nil
			}
			var tags1 []string = nil
			for _, tagsItem1 := range dollarRelationRefItem1.Tags {
				tags1 = append(tags1, tagsItem1.ValueString())
			}
			entityId1 := new(string)
			if !dollarRelationRefItem1.EntityID.IsUnknown() && !dollarRelationRefItem1.EntityID.IsNull() {
				*entityId1 = dollarRelationRefItem1.EntityID.ValueString()
			} else {
				entityId1 = nil
			}
			path1 := new(string)
			if !dollarRelationRefItem1.Path.IsUnknown() && !dollarRelationRefItem1.Path.IsNull() {
				*path1 = dollarRelationRefItem1.Path.ValueString()
			} else {
				path1 = nil
			}
			dollarRelationRef1 = append(dollarRelationRef1, shared.DollarRelationRef{
				ID:       id1,
				Tags:     tags1,
				EntityID: entityId1,
				Path:     path1,
			})
		}
		billingAddress = &shared.BaseRelationRef{
			DollarRelationRef: dollarRelationRef1,
		}
	}
	currentTask := new(string)
	if !r.CurrentTask.IsUnknown() && !r.CurrentTask.IsNull() {
		*currentTask = r.CurrentTask.ValueString()
	} else {
		currentTask = nil
	}
	var customer *shared.BaseRelation
	if r.Customer != nil {
		var dollarRelation []shared.DollarRelation = nil
		for _, dollarRelationItem := range r.Customer.DollarRelation {
			var tags2 []string = nil
			for _, tagsItem2 := range dollarRelationItem.Tags {
				tags2 = append(tags2, tagsItem2.ValueString())
			}
			entityId2 := new(string)
			if !dollarRelationItem.EntityID.IsUnknown() && !dollarRelationItem.EntityID.IsNull() {
				*entityId2 = dollarRelationItem.EntityID.ValueString()
			} else {
				entityId2 = nil
			}
			dollarRelation = append(dollarRelation, shared.DollarRelation{
				Tags:     tags2,
				EntityID: entityId2,
			})
		}
		customer = &shared.BaseRelation{
			DollarRelation: dollarRelation,
		}
	}
	var dates []shared.OpportunityCreateDates = nil
	for _, datesItem := range r.Dates {
		var tags3 []string = nil
		for _, tagsItem3 := range datesItem.Tags {
			tags3 = append(tags3, tagsItem3.ValueString())
		}
		dates1 := new(string)
		if !datesItem.Dates.IsUnknown() && !datesItem.Dates.IsNull() {
			*dates1 = datesItem.Dates.ValueString()
		} else {
			dates1 = nil
		}
		value := new(time.Time)
		if !datesItem.Value.IsUnknown() && !datesItem.Value.IsNull() {
			*value, _ = time.Parse(time.RFC3339Nano, datesItem.Value.ValueString())
		} else {
			value = nil
		}
		dates = append(dates, shared.OpportunityCreateDates{
			Tags:  tags3,
			Dates: dates1,
			Value: value,
		})
	}
	var deliveryAddress *shared.BaseRelationRef
	if r.DeliveryAddress != nil {
		var dollarRelationRef2 []shared.DollarRelationRef = nil
		for _, dollarRelationRefItem2 := range r.DeliveryAddress.DollarRelationRef {
			id2 := new(string)
			if !dollarRelationRefItem2.ID.IsUnknown() && !dollarRelationRefItem2.ID.IsNull() {
				*id2 = dollarRelationRefItem2.ID.ValueString()
			} else {
				id2 = nil
			}
			var tags4 []string = nil
			for _, tagsItem4 := range dollarRelationRefItem2.Tags {
				tags4 = append(tags4, tagsItem4.ValueString())
			}
			entityId3 := new(string)
			if !dollarRelationRefItem2.EntityID.IsUnknown() && !dollarRelationRefItem2.EntityID.IsNull() {
				*entityId3 = dollarRelationRefItem2.EntityID.ValueString()
			} else {
				entityId3 = nil
			}
			path2 := new(string)
			if !dollarRelationRefItem2.Path.IsUnknown() && !dollarRelationRefItem2.Path.IsNull() {
				*path2 = dollarRelationRefItem2.Path.ValueString()
			} else {
				path2 = nil
			}
			dollarRelationRef2 = append(dollarRelationRef2, shared.DollarRelationRef{
				ID:       id2,
				Tags:     tags4,
				EntityID: entityId3,
				Path:     path2,
			})
		}
		deliveryAddress = &shared.BaseRelationRef{
			DollarRelationRef: dollarRelationRef2,
		}
	}
	var items *shared.BaseRelation
	if r.Items != nil {
		var dollarRelation1 []shared.DollarRelation = nil
		for _, dollarRelationItem1 := range r.Items.DollarRelation {
			var tags5 []string = nil
			for _, tagsItem5 := range dollarRelationItem1.Tags {
				tags5 = append(tags5, tagsItem5.ValueString())
			}
			entityId4 := new(string)
			if !dollarRelationItem1.EntityID.IsUnknown() && !dollarRelationItem1.EntityID.IsNull() {
				*entityId4 = dollarRelationItem1.EntityID.ValueString()
			} else {
				entityId4 = nil
			}
			dollarRelation1 = append(dollarRelation1, shared.DollarRelation{
				Tags:     tags5,
				EntityID: entityId4,
			})
		}
		items = &shared.BaseRelation{
			DollarRelation: dollarRelation1,
		}
	}
	opportunityTitle := r.OpportunityTitle.ValueString()
	var payment *shared.BaseRelationRef
	if r.Payment != nil {
		var dollarRelationRef3 []shared.DollarRelationRef = nil
		for _, dollarRelationRefItem3 := range r.Payment.DollarRelationRef {
			id3 := new(string)
			if !dollarRelationRefItem3.ID.IsUnknown() && !dollarRelationRefItem3.ID.IsNull() {
				*id3 = dollarRelationRefItem3.ID.ValueString()
			} else {
				id3 = nil
			}
			var tags6 []string = nil
			for _, tagsItem6 := range dollarRelationRefItem3.Tags {
				tags6 = append(tags6, tagsItem6.ValueString())
			}
			entityId5 := new(string)
			if !dollarRelationRefItem3.EntityID.IsUnknown() && !dollarRelationRefItem3.EntityID.IsNull() {
				*entityId5 = dollarRelationRefItem3.EntityID.ValueString()
			} else {
				entityId5 = nil
			}
			path3 := new(string)
			if !dollarRelationRefItem3.Path.IsUnknown() && !dollarRelationRefItem3.Path.IsNull() {
				*path3 = dollarRelationRefItem3.Path.ValueString()
			} else {
				path3 = nil
			}
			dollarRelationRef3 = append(dollarRelationRef3, shared.DollarRelationRef{
				ID:       id3,
				Tags:     tags6,
				EntityID: entityId5,
				Path:     path3,
			})
		}
		payment = &shared.BaseRelationRef{
			DollarRelationRef: dollarRelationRef3,
		}
	}
	var source *shared.OpportunityCreateSource
	if r.Source != nil {
		href := new(string)
		if !r.Source.Href.IsUnknown() && !r.Source.Href.IsNull() {
			*href = r.Source.Href.ValueString()
		} else {
			href = nil
		}
		title := new(string)
		if !r.Source.Title.IsUnknown() && !r.Source.Title.IsNull() {
			*title = r.Source.Title.ValueString()
		} else {
			title = nil
		}
		source = &shared.OpportunityCreateSource{
			Href:  href,
			Title: title,
		}
	}
	sourceType := new(string)
	if !r.SourceType.IsUnknown() && !r.SourceType.IsNull() {
		*sourceType = r.SourceType.ValueString()
	} else {
		sourceType = nil
	}
	status := new(string)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = r.Status.ValueString()
	} else {
		status = nil
	}
	out := shared.OpportunityCreate{
		Address:          address,
		BillingAddress:   billingAddress,
		CurrentTask:      currentTask,
		Customer:         customer,
		Dates:            dates,
		DeliveryAddress:  deliveryAddress,
		Items:            items,
		OpportunityTitle: opportunityTitle,
		Payment:          payment,
		Source:           source,
		SourceType:       sourceType,
		Status:           status,
	}
	return &out
}

func (r *OpportunityResourceModel) RefreshFromSharedOpportunity(resp *shared.Opportunity) {
	if resp != nil {
		if resp.ACL.AdditionalProperties == nil {
			r.ACL.AdditionalProperties = types.StringNull()
		} else {
			additionalPropertiesResult, _ := json.Marshal(resp.ACL.AdditionalProperties)
			r.ACL.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
		}
		r.ACL.Delete = []types.String{}
		for _, v := range resp.ACL.Delete {
			r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
		}
		r.ACL.Edit = []types.String{}
		for _, v := range resp.ACL.Edit {
			r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
		}
		r.ACL.View = []types.String{}
		for _, v := range resp.ACL.View {
			r.ACL.View = append(r.ACL.View, types.StringValue(v))
		}
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.ID = types.StringValue(resp.ID)
		r.Org = types.StringValue(resp.Org)
		if len(r.Owners) > len(resp.Owners) {
			r.Owners = r.Owners[:len(resp.Owners)]
		}
		for ownersCount, ownersItem := range resp.Owners {
			var owners1 tfTypes.BaseEntityOwner
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
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.Title = types.StringValue(resp.Title)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		if resp.Address == nil {
			r.Address = nil
		} else {
			r.Address = &tfTypes.BaseRelationRef{}
			if len(r.Address.DollarRelationRef) > len(resp.Address.DollarRelationRef) {
				r.Address.DollarRelationRef = r.Address.DollarRelationRef[:len(resp.Address.DollarRelationRef)]
			}
			for dollarRelationRefCount, dollarRelationRefItem := range resp.Address.DollarRelationRef {
				var dollarRelationRef1 tfTypes.DollarRelationRef
				dollarRelationRef1.ID = types.StringPointerValue(dollarRelationRefItem.ID)
				dollarRelationRef1.Tags = []types.String{}
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
			r.BillingAddress = &tfTypes.BaseRelationRef{}
			if len(r.BillingAddress.DollarRelationRef) > len(resp.BillingAddress.DollarRelationRef) {
				r.BillingAddress.DollarRelationRef = r.BillingAddress.DollarRelationRef[:len(resp.BillingAddress.DollarRelationRef)]
			}
			for dollarRelationRefCount1, dollarRelationRefItem1 := range resp.BillingAddress.DollarRelationRef {
				var dollarRelationRef3 tfTypes.DollarRelationRef
				dollarRelationRef3.ID = types.StringPointerValue(dollarRelationRefItem1.ID)
				dollarRelationRef3.Tags = []types.String{}
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
			r.Customer = &tfTypes.BaseRelation{}
			if len(r.Customer.DollarRelation) > len(resp.Customer.DollarRelation) {
				r.Customer.DollarRelation = r.Customer.DollarRelation[:len(resp.Customer.DollarRelation)]
			}
			for dollarRelationCount, dollarRelationItem := range resp.Customer.DollarRelation {
				var dollarRelation1 tfTypes.DollarRelation
				dollarRelation1.Tags = []types.String{}
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
			var dates1 tfTypes.Dates
			dates1.Tags = []types.String{}
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
			r.DeliveryAddress = &tfTypes.BaseRelationRef{}
			if len(r.DeliveryAddress.DollarRelationRef) > len(resp.DeliveryAddress.DollarRelationRef) {
				r.DeliveryAddress.DollarRelationRef = r.DeliveryAddress.DollarRelationRef[:len(resp.DeliveryAddress.DollarRelationRef)]
			}
			for dollarRelationRefCount2, dollarRelationRefItem2 := range resp.DeliveryAddress.DollarRelationRef {
				var dollarRelationRef5 tfTypes.DollarRelationRef
				dollarRelationRef5.ID = types.StringPointerValue(dollarRelationRefItem2.ID)
				dollarRelationRef5.Tags = []types.String{}
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
		if resp.Items == nil {
			r.Items = nil
		} else {
			r.Items = &tfTypes.BaseRelation{}
			if len(r.Items.DollarRelation) > len(resp.Items.DollarRelation) {
				r.Items.DollarRelation = r.Items.DollarRelation[:len(resp.Items.DollarRelation)]
			}
			for dollarRelationCount1, dollarRelationItem1 := range resp.Items.DollarRelation {
				var dollarRelation3 tfTypes.DollarRelation
				dollarRelation3.Tags = []types.String{}
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
			r.Payment = &tfTypes.BaseRelationRef{}
			if len(r.Payment.DollarRelationRef) > len(resp.Payment.DollarRelationRef) {
				r.Payment.DollarRelationRef = r.Payment.DollarRelationRef[:len(resp.Payment.DollarRelationRef)]
			}
			for dollarRelationRefCount3, dollarRelationRefItem3 := range resp.Payment.DollarRelationRef {
				var dollarRelationRef7 tfTypes.DollarRelationRef
				dollarRelationRef7.ID = types.StringPointerValue(dollarRelationRefItem3.ID)
				dollarRelationRef7.Tags = []types.String{}
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
			r.Source = &tfTypes.OpportunityCreateSource{}
			r.Source.Href = types.StringPointerValue(resp.Source.Href)
			r.Source.Title = types.StringPointerValue(resp.Source.Title)
		}
		r.SourceType = types.StringPointerValue(resp.SourceType)
		r.Status = types.StringPointerValue(resp.Status)
	}
}
