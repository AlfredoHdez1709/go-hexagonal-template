package repository

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go-hexagonal-template/internal/core/domain"
	"go-hexagonal-template/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	tableName string
	database  *mongo.Database
}

func (u UsersRepository) InsertUser(ctx context.Context, user domain.User) error {
	collection := u.database.Collection(u.tableName)

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UsersRepository) GetUsers(ctx context.Context) ([]domain.User, error) {

	collection := u.database.Collection(u.tableName)

	curs, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(curs *mongo.Cursor, ctx context.Context) {
		err := curs.Close(ctx)
		if err != nil {
			log.Fatalf("failed to close cursor, %v", err)
		}
	}(curs, nil)
	var users []domain.User
	err = curs.All(nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil

}

func NewUsersRepository(database *mongo.Database) ports.IUserRepository {
	return UsersRepository{
		tableName: "users",
		database:  database,
	}
}
