package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ tfsdk.ResourceType = (*timeRotatingResourceType)(nil)

type timeRotatingResourceType struct{}

func (t timeRotatingResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"day": {
				Description: "Number day of timestamp.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"rotation_days": {
				Description: "Number of days to add to the base timestamp to configure the rotation timestamp. " +
					"When the current time has passed the rotation timestamp, the resource will trigger recreation. " +
					"At least one of the 'rotation_' arguments must be configured.",
				Type:     types.Int64Type,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AtLeastOneOf(path.MatchRoot("rotation_hours"),
						path.MatchRoot("rotation_minutes"),
						path.MatchRoot("rotation_months"),
						path.MatchRoot("rotation_rfc3339"),
						path.MatchRoot("rotation_years")),
					int64validator.AtLeast(1),
				},
				//AtLeastOneOf: []string{
				//	"rotation_days",
				//	"rotation_hours",
				//	"rotation_minutes",
				//	"rotation_months",
				//	"rotation_rfc3339",
				//	"rotation_years",
				//},
				//ValidateFunc: validation.IntAtLeast(1),
			},
			"rotation_hours": {
				Description: "Number of hours to add to the base timestamp to configure the rotation timestamp. " +
					"When the current time has passed the rotation timestamp, the resource will trigger recreation. " +
					"At least one of the 'rotation_' arguments must be configured.",
				Type:     types.Int64Type,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AtLeastOneOf(path.MatchRoot("rotation_days"),
						path.MatchRoot("rotation_minutes"),
						path.MatchRoot("rotation_months"),
						path.MatchRoot("rotation_rfc3339"),
						path.MatchRoot("rotation_years")),
					int64validator.AtLeast(1),
				},
				//AtLeastOneOf: []string{
				//	"rotation_days",
				//	"rotation_hours",
				//	"rotation_minutes",
				//	"rotation_months",
				//	"rotation_rfc3339",
				//	"rotation_years",
				//},
				//ValidateFunc: validation.IntAtLeast(1),
			},
			"rotation_minutes": {
				Description: "Number of minutes to add to the base timestamp to configure the rotation timestamp. " +
					"When the current time has passed the rotation timestamp, the resource will trigger recreation. " +
					"At least one of the 'rotation_' arguments must be configured.",
				Type:     types.Int64Type,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AtLeastOneOf(path.MatchRoot("rotation_days"),
						path.MatchRoot("rotation_hours"),
						path.MatchRoot("rotation_months"),
						path.MatchRoot("rotation_rfc3339"),
						path.MatchRoot("rotation_years")),
					int64validator.AtLeast(1),
				},
				//AtLeastOneOf: []string{
				//	"rotation_days",
				//	"rotation_hours",
				//	"rotation_minutes",
				//	"rotation_months",
				//	"rotation_rfc3339",
				//	"rotation_years",
				//},
				//ValidateFunc: validation.IntAtLeast(1),
			},
			"rotation_months": {
				Description: "Number of months to add to the base timestamp to configure the rotation timestamp. " +
					"When the current time has passed the rotation timestamp, the resource will trigger recreation. " +
					"At least one of the 'rotation_' arguments must be configured.",
				Type:     types.Int64Type,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AtLeastOneOf(path.MatchRoot("rotation_days"),
						path.MatchRoot("rotation_hours"),
						path.MatchRoot("rotation_minutes"),
						path.MatchRoot("rotation_rfc3339"),
						path.MatchRoot("rotation_years")),
					int64validator.AtLeast(1),
				},
				//AtLeastOneOf: []string{
				//	"rotation_days",
				//	"rotation_hours",
				//	"rotation_minutes",
				//	"rotation_months",
				//	"rotation_rfc3339",
				//	"rotation_years",
				//},
				//ValidateFunc: validation.IntAtLeast(1),
			},
			"rotation_rfc3339": {
				Description: "Configure the rotation timestamp with an " +
					"[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format of the offset timestamp. " +
					"When the current time has passed the rotation timestamp, the resource will trigger recreation. " +
					"At least one of the 'rotation_' arguments must be configured.",
				Type:     types.StringType,
				Optional: true,
				Computed: true,
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AtLeastOneOf(path.MatchRoot("rotation_days"),
						path.MatchRoot("rotation_hours"),
						path.MatchRoot("rotation_minutes"),
						path.MatchRoot("rotation_months"),
						path.MatchRoot("rotation_years")),
				},
				//AtLeastOneOf: []string{
				//	"rotation_days",
				//	"rotation_hours",
				//	"rotation_minutes",
				//	"rotation_months",
				//	"rotation_rfc3339",
				//	"rotation_years",
				//},
				//ValidateFunc: validation.IsRFC3339Time,
			},
			"rotation_years": {
				Description: "Number of years to add to the base timestamp to configure the rotation timestamp. " +
					"When the current time has passed the rotation timestamp, the resource will trigger recreation. " +
					"At least one of the 'rotation_' arguments must be configured.",
				Type:     types.Int64Type,
				Optional: true,
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AtLeastOneOf(path.MatchRoot("rotation_days"),
						path.MatchRoot("rotation_hours"),
						path.MatchRoot("rotation_minutes"),
						path.MatchRoot("rotation_months"),
						path.MatchRoot("rotation_rfc3339")),
					int64validator.AtLeast(1),
				},
				//AtLeastOneOf: []string{
				//	"rotation_days",
				//	"rotation_hours",
				//	"rotation_minutes",
				//	"rotation_months",
				//	"rotation_rfc3339",
				//	"rotation_years",
				//},
				//ValidateFunc: validation.IntAtLeast(1),
			},
			"hour": {
				Description: "Number hour of timestamp.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"triggers": {
				Description: "Arbitrary map of values that, when changed, will trigger a new base timestamp value to be saved." +
					" These conditions recreate the resource in addition to other rotation arguments. " +
					"See [the main provider documentation](../index.md) for more information.",
				Type: types.MapType{
					ElemType: types.StringType,
				},
				Optional: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
				//ForceNew: true,
				//Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"minute": {
				Description: "Number minute of timestamp.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"month": {
				Description: "Number month of timestamp.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"rfc3339": {
				Description: "Base timestamp in " +
					"[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format " +
					"(see [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., " +
					"`YYYY-MM-DDTHH:MM:SSZ`). Defaults to the current time.",
				Type:     types.StringType,
				Optional: true,
				Computed: true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					tfsdk.RequiresReplace(),
				},
				//ForceNew:     true,
				//ValidateFunc: validation.IsRFC3339Time,
			},
			"second": {
				Description: "Number second of timestamp.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"unix": {
				Description: "Number of seconds since epoch time, e.g. `1581489373`.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"year": {
				Description: "Number year of timestamp.",
				Type:        types.Int64Type,
				Computed:    true,
			},
			"id": {
				Description: "RFC3339 format of the offset timestamp, e.g. `2020-02-12T06:36:13Z`.",
				Type:        types.StringType,
				Computed:    true,
			},
		},
	}, nil
}

func (t timeRotatingResourceType) NewResource(ctx context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return &timeRotatingResource{}, nil
}

var (
	_ tfsdk.Resource                = (*timeRotatingResource)(nil)
	_ tfsdk.ResourceWithImportState = (*timeRotatingResource)(nil)
)

type timeRotatingResource struct {
}

func (t timeRotatingResource) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	//TODO implement me
	panic("implement me")
}

type timeRotatingModelV0 struct {
	Day             types.Int64  `tfsdk:"day"`
	RotationDays    types.Int64  `tfsdk:"rotation_days"`
	RotationHours   types.Int64  `tfsdk:"rotation_hours"`
	RotationMinutes types.Int64  `tfsdk:"rotation_minutes"`
	RotationMonths  types.Int64  `tfsdk:"rotation_months"`
	RotationRFC3339 types.String `tfsdk:"rotation_rfc3339"`
	RotationYears   types.Int64  `tfsdk:"rotation_years"`
	Hour            types.Int64  `tfsdk:"hour"`
	Triggers        types.Map    `tfsdk:"triggers"`
	Minute          types.Int64  `tfsdk:"minute"`
	Month           types.Int64  `tfsdk:"month"`
	RFC3339         types.String `tfsdk:"rfc3339"`
	Second          types.Int64  `tfsdk:"second"`
	Unix            types.Int64  `tfsdk:"unix"`
	Year            types.Int64  `tfsdk:"year"`
	ID              types.String `tfsdk:"id"`
}

func (t timeRotatingResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan timeRotatingModelV0

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	timestamp := time.Now().UTC()

	if plan.RFC3339.Value != "" {
		var err error

		if timestamp, err = time.Parse(time.RFC3339, plan.RFC3339.Value); err != nil {
			resp.Diagnostics.AddError(
				"Create time rotating error",
				"The rfc3339 timestamp that was supplied could not be parsed as RFC3339.\n\n+"+
					fmt.Sprintf("Original Error: %s", err),
			)
			return
		}
	}

	state, err := setRotationValues(&plan, timestamp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Create time rotating error",
			"The rfc3339 timestamp that was supplied could not be parsed as RFC3339.\n\n+"+
				fmt.Sprintf("Original Error: %s", err),
		)
		return
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (t timeRotatingResource) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {

}

func (t timeRotatingResource) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	//TODO implement me
	panic("implement me")
}

func (t timeRotatingResource) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {

}

func setRotationValues(plan *timeRotatingModelV0, timestamp time.Time) (timeRotatingModelV0, error) {
	formattedTimestamp := timestamp.Format(time.RFC3339)

	var rotationTimestamp time.Time

	if plan.RotationDays.Value != 0 {
		rotationTimestamp = timestamp.AddDate(0, 0, int(plan.RotationDays.Value))
	}

	if plan.RotationHours.Value != 0 {
		hours := time.Duration(plan.RotationHours.Value) * time.Hour
		rotationTimestamp = timestamp.Add(hours)
	}

	if plan.RotationMinutes.Value != 0 {
		minutes := time.Duration(plan.RotationMinutes.Value) * time.Minute
		rotationTimestamp = timestamp.Add(minutes)
	}

	if plan.RotationMonths.Value != 0 {
		rotationTimestamp = timestamp.AddDate(0, int(plan.RotationMonths.Value), 0)
	}

	if plan.RotationRFC3339.Value != "" {
		var err error

		if rotationTimestamp, err = time.Parse(time.RFC3339, plan.RotationRFC3339.Value); err != nil {
			return timeRotatingModelV0{}, err
		}
	}

	if plan.RotationYears.Value != 0 {
		rotationTimestamp = timestamp.AddDate(int(plan.RotationYears.Value), 0, 0)
	}

	formattedRotationTimestamp := rotationTimestamp.Format(time.RFC3339)

	return timeRotatingModelV0{
		RotationRFC3339: types.String{Value: formattedRotationTimestamp},
		Triggers:        plan.Triggers,
		Year:            types.Int64{Value: int64(rotationTimestamp.Year())},
		Month:           types.Int64{Value: int64(rotationTimestamp.Month())},
		Day:             types.Int64{Value: int64(rotationTimestamp.Day())},
		Hour:            types.Int64{Value: int64(rotationTimestamp.Hour())},
		Minute:          types.Int64{Value: int64(rotationTimestamp.Minute())},
		Second:          types.Int64{Value: int64(rotationTimestamp.Second())},
		RotationYears:   plan.RotationYears,
		RotationMonths:  plan.RotationMonths,
		RotationDays:    plan.RotationDays,
		RotationHours:   plan.RotationHours,
		RotationMinutes: plan.RotationMinutes,
		RFC3339:         types.String{Value: formattedTimestamp},
		Unix:            types.Int64{Value: rotationTimestamp.Unix()},
		ID:              types.String{Value: formattedTimestamp},
	}, nil

}
