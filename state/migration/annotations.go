// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package migration

import (
	"github.com/juju/schema"
)

// Instead of copy / pasting the Annotations, SetAnnotations, and the import
// three lines into every entity that has annotations, we provide a helper
// struct used in composition. Each entity still needs to define the
// annotations map to be serialized. For the two locations where entities are
// created (the new<entity> function, and the import functions), the pointer
// to the annotations needs to be set. This allows the accessor and setter
// methods to work. This type is composed without a name so the methods get
// promoted so they satisfy the HasAnnotations interface.
type hasAnnotations map[string]string

// Annotations implements HasAnnotations.
func (a hasAnnotations) Annotations() map[string]string {
	return a
}

// SetAnnotations implements HasAnnotations.
func (a *hasAnnotations) SetAnnotations(annotations map[string]string) {
	*a = annotations
}

func (a *hasAnnotations) importAnnotations(valid map[string]interface{}) {
	if asInterfaces, ok := valid["annotations"]; ok {
		annotations := make(map[string]string)
		// The schema will return a string map as map[string]interface{}.
		// It will make sure that the interface values are strings, but doesn't
		// return them as strings. So we need to do that here.
		for key, value := range asInterfaces.(map[string]interface{}) {
			annotations[key] = value.(string)
		}
		a.SetAnnotations(annotations)
	}
}

func addAnnotationSchema(fields schema.Fields, defaults schema.Defaults) {
	fields["annotations"] = schema.StringMap(schema.String())
	defaults["annotations"] = schema.Omit
}
