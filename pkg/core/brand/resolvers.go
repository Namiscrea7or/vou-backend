package brand

import (
	"fmt"
	"log"
	"vou/pkg/db/coredb"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandResolver struct {
	BrandRepo *coredb.BrandRepo
	UsersRepo *coredb.UsersRepo
}

func NewBrandResolver() *BrandResolver {
	return &BrandResolver{
		BrandRepo: coredb.NewBrandRepo(),
	}
}

func (r *BrandResolver) CreateBrand(params graphql.ResolveParams) (interface{}, error) {
	name, ok := params.Args["name"].(string)
	if !ok {
		fmt.Errorf("Don't find name")
	}

	industry, ok := params.Args["industry"].(string)
	if !ok {
		fmt.Errorf("Don't find industry")
	}

	address, ok := params.Args["address"].(string)
	if !ok {
		fmt.Errorf("Don't find address")
	}

	latitude, ok := params.Args["latitude"].(float64)
	if !ok {
		fmt.Errorf("Don't find latitude")
	}

	longitude, ok := params.Args["longitude"].(float64)
	if !ok {
		fmt.Errorf("Don't find longitude")
	}

	status, ok := params.Args["status"].(bool)
	if !ok {
		fmt.Errorf("Don't find status")
	}

	creatorId, ok := params.Args["creatorId"].(string)
	if !ok {
		fmt.Errorf("Don't find creatorId")
	}

	brand := coredb.Brand{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Industry: industry,
		Address:  address,
		Location: coredb.Gps{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Status:    status,
		CreatorId: creatorId,
	}

	_, err := r.BrandRepo.CreateBrand(brand)
	if err != nil {
		log.Printf("failed to create brand: %v\n", err)
		return false, err
	}

	return true, nil
}

func (r *BrandResolver) GetBrandByID(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, nil
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	brand, err := r.BrandRepo.GetBrandByID(objectID)
	if err != nil {
		return nil, fmt.Errorf("brand not found: %v", err)
	}

	return brand, nil

}

func (r *BrandResolver) GetAllBrands(params graphql.ResolveParams) (interface{}, error) {
	brands, err := r.BrandRepo.GetAllBrands()
	if err != nil {
		log.Printf("failed to find Brands: %v\n", err)
		return nil, err
	}

	return brands, nil
}
