// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"tutorial.ent.go/app/internal/db/ent/pet"
	"tutorial.ent.go/app/internal/db/ent/schema"
	"tutorial.ent.go/app/internal/db/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescID is the schema descriptor for id field.
	petDescID := petFields[0].Descriptor()
	// pet.IDValidator is a validator for the "id" field. It is called by the builders before save.
	pet.IDValidator = func() func(string) error {
		validators := petDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
