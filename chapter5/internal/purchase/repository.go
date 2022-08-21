package purchase

import (
	"context"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"
)

type Repository interface {
	Store(ctx context.Context, purchase Purchase) error
	Ping(ctx context.Context) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")

	return &MongoRepository{
		purchases: purchases,
	}, nil
}

func (mr *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongoP := New(purchase)
	_, err := mr.purchases.InsertOne(ctx, mongoP)
	if err != nil {
		return fmt.Errorf("failed to persist purchase: %w", err)
	}
	return nil
}

type mongoPurchase struct {
	id                 uuid.UUID
	store              store.Store
	productsToPurchase []coffeeco.Product
	total              money.Money
	paymentMeans       payment.Means
	timeOfPurchase     time.Time
	cardToken          *string
}

func New(p Purchase) mongoPurchase {
	return mongoPurchase{
		id:                 p.id,
		store:              p.Store,
		productsToPurchase: p.ProductsToPurchase,
		total:              p.total,
		paymentMeans:       p.PaymentMeans,
		timeOfPurchase:     p.timeOfPurchase,
		cardToken:          p.cardToken,
	}
}

func (m mongoPurchase) ToPurchase() Purchase {
	return Purchase{
		id:                 m.id,
		Store:              m.store,
		ProductsToPurchase: m.productsToPurchase,
		total:              m.total,
		PaymentMeans:       m.paymentMeans,
		timeOfPurchase:     m.timeOfPurchase,
		cardToken:          m.cardToken,
	}
}

func (mr *MongoRepository) Ping(ctx context.Context) error {
	if _, err := mr.purchases.EstimatedDocumentCount(ctx); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}
	return nil
}
