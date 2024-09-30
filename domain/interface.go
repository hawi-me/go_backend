package domain

import (
	"context"
)

type DiseaseRepositoryInterface interface {
	GetDiseaseByName(ctx context.Context, name string) (*Disease, error)
	InsertOne(ctx context.Context, disease *Disease) (string, error)
	GetAllDiseases(ctx context.Context,page int) ([]Disease, error)
}

type NutrientRepositoryInterface interface {
	GetNutrientByName(ctx context.Context, name string) (*Nutrient, error)
	InsertOne(ctx context.Context, nutrient *Nutrient) (string, error)
	GetAllNutrients(ctx context.Context) ([]Nutrient, error)
}

type HerbRepositoryInterface interface {
	GetHerbByName(ctx context.Context, name string) (*Herb, error)
	InsertOne(ctx context.Context, herb *Herb) (string, error)
	GetAllHerbs(ctx context.Context) ([]Herb, error)
}

type EducationalUseCaseInterface interface {
	GetDiseaseByName(ctx context.Context, name string) (*Disease, error)
	InsertOneDisease(ctx context.Context, disease *Disease) (string, error)
	GetAllDiseases(ctx context.Context,page int) ([]Disease, error)

	InsertOneHerb(ctx context.Context, herb *Herb) (string, error)
	GetHerbByName(ctx context.Context, name string) (*Herb, error)
	GetAllHerbs(ctx context.Context) ([]Herb, error)

	// Nutrient-related methods
	InsertOneNutrient(ctx context.Context, nutrient *Nutrient) (string, error)
	GetNutrientByName(ctx context.Context, name string) (*Nutrient, error)
	GetAllNutrients(ctx context.Context) ([]Nutrient, error)
}
