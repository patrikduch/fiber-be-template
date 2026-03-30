package database

import (
	"context"
	"log"
	"strings"

	entRole "fiber-be-template/ent/role"
	"github.com/google/uuid"
)

func Seed() {
	seedRoles()
}

func seedRoles() {
	roles := []string{"CUSTOMER"}

	for _, name := range roles {
		exists, err := EntClient.Role.
			Query().
			Where(entRole.NameEQ(name)).
			Exist(context.Background())
		if err != nil {
			log.Fatalf("Failed to check role %s: %v", name, err)
		}
		if !exists {
			EntClient.Role.
				Create().
				SetName(name).
				SetNormalizedName(strings.ToUpper(name)).
				SetConcurrencyStamp(uuid.New().String()).
				SaveX(context.Background())
			log.Printf("Seeded role: %s", name)
		}
	}
}
