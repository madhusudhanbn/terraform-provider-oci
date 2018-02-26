// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func InstanceShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readInstanceShapes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readInstanceShapes(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

type InstanceShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListShapesResponse
}

func (s *InstanceShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *InstanceShapesDataSourceCrud) Get() error {
	request := oci_core.ListShapesRequest{}

	if availabilityDomain, ok := s.D.GetOk("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOk("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if imageId, ok := s.D.GetOk("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if limit, ok := s.D.GetOk("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOk("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	response, err := s.Client.ListShapes(context.Background(), request, getRetryOptions(false, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *InstanceShapesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		shape := map[string]interface{}{}

		if r.Shape != nil {
			shape["name"] = *r.Shape
		}

		resources = append(resources, shape)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("shapes", resources); err != nil {
		panic(err)
	}

	return
}
