package main

import "github.com/hashicorp/terraform/helper/schema"

func provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"conditional_list_of_string": resource(schema.TypeList, &schema.Schema{Type: schema.TypeString}),
			"conditional_list_of_list_of_string": resource(schema.TypeList, &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			}),
			"conditional_list_of_map": resource(schema.TypeList, &schema.Schema{Type: schema.TypeMap}),
			"conditional_map":         resource(schema.TypeMap, nil),
		},
	}
}
