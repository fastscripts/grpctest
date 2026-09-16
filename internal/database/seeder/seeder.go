package seeder

import (
	"fmt"

	"gorm.io/gorm"
)

type Seeder struct {
	Name string
	Func func(db *gorm.DB) error
}

var Seeders = []Seeder{
	{
		Name: "SeedUsers",
		Func: SeedUsers,
	},
}

type Opts struct {
	DB      *gorm.DB
	Seeders []Seeder
}

func RunAllSeeders(opts Opts) error {

	if len(opts.Seeders) == 0 {
		opts.Seeders = Seeders
	}

	for _, seeder := range opts.Seeders {
		fmt.Printf("Running seeder: %s\n", seeder.Name)
		err := seeder.Func(opts.DB)
		if err != nil {
			fmt.Printf("Seeder %s Failed: %v\n", seeder.Name, err)
			return err
		}
		fmt.Printf("Seeder %s Completed\n", seeder.Name)
	}
	return nil
}
