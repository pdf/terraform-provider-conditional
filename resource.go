package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/satori/go.uuid"
)

func resource(kind schema.ValueType, elem *schema.Schema) *schema.Resource {
	if elem == nil {
		elem = &schema.Schema{}
	}

	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"if": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "set to the result of a boolean condition: on `true`, `result` will contain the field `then`; on `false`, `result` will contain the fields `else` (or empty if `else` is omitted)",
			},
			"then": &schema.Schema{
				Type:        kind,
				Required:    true,
				ForceNew:    true,
				Elem:        elem,
				Description: "the value to return as `result` when `if` is `true`",
			},
			"else": &schema.Schema{
				Type:        kind,
				Optional:    true,
				ForceNew:    true,
				Elem:        elem,
				Description: "the value to return as `result` when `if` is `false`",
			},
			"result": &schema.Schema{
				Type:     kind,
				Optional: true,
				Computed: true,
				Elem:     elem,
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	if d.Get("if").(bool) {
		d.Set("result", d.Get("then"))
	} else {
		d.Set("result", d.Get("else"))
	}
	d.SetId(uuid.NewV4().String())
	return nil
}

func resourceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
