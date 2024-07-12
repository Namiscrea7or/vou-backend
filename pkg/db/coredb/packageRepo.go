package coredb

import (
	"context"
	"log"
	"time"

	"vou/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PackagesRepo struct {
	*mongo.Collection
}

func NewPackagesRepo() *PackagesRepo {
	packagesCollection := db.GetPackageCollection()
	_, err := packagesCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "user_id", Value: 1}},
			Options: options.Index().SetUnique(false),
		},
	})
	if err != nil {
		log.Println("failed to create indexes for packages collection")
		return nil
	}

	return &PackagesRepo{packagesCollection}
}

func (r *PackagesRepo) GetPackageByID(id primitive.ObjectID) (Package, error) {
	var pkg Package
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.FindOne(ctx, bson.M{"_id": id}).Decode(&pkg)

	return pkg, err
}

func (r *PackagesRepo) CreateNewPackage(pkg Package) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.InsertOne(ctx, pkg)
}

func (r *PackagesRepo) GetPackagesByUserID(userID string) ([]Package, error) {
	var packages []Package
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var pkg Package
		if err = cursor.Decode(&pkg); err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	return packages, cursor.Err()
}

func (r *PackagesRepo) AddVoucherToPackage(pkgID primitive.ObjectID, voucherID string) error {
	filter := bson.M{"_id": pkgID}
	update := bson.M{
		"$addToSet": bson.M{"vouchers": voucherID},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.UpdateOne(ctx, filter, update)
	return err
}

func (r *PackagesRepo) RemoveVoucherFromPackage(pkgID primitive.ObjectID, voucherID string) error {
	filter := bson.M{"_id": pkgID}
	update := bson.M{
		"$pull": bson.M{"vouchers": voucherID},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.UpdateOne(ctx, filter, update)
	return err
}