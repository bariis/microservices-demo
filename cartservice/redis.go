package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"log"
	"strings"
	"time"
)

// gets and returns the required part of the token
func getUserIdFromToken(token string) (string, error) {
	// remove the `bearer` part from the token
	encodedToken := strings.Split(token, " ")[1]
	parsedToken, _ := jwt.Parse(encodedToken, nil)
	if parsedToken == nil {
		return "", errors.New("token is not valid")
	}

	claims, _ := parsedToken.Claims.(jwt.MapClaims)

	return claims["user_id"].(string), nil
}

// add or updates the cart
func AddItem(redis *redis.Client, productId string, productQuantity int, token string) error {
	ctx := context.Background()

	// user_id will be stored as cartId like key in redis cache
	cartId, err := getUserIdFromToken(token)
	if err != nil {
		return fmt.Errorf("token: %v", err)
	}

	item := map[string]interface{}{
		productId: productQuantity,
	}

	cmd := redis.Exists(ctx, cartId)
	if cmd.Val() == 0 {
		redis.HSet(ctx, cartId, item)
	} else {
		// If the item exists, we update its quantity
		if cmd := redis.HExists(ctx, cartId, productId); cmd.Val() {
			redis.HIncrBy(ctx, cartId, productId, int64(productQuantity))

			// if the quantity of the item equals to zero, we delete this product from cache
			if cmd := redis.HGet(ctx, cartId, productId); cmd.Val() == "0" {
				redis.HDel(ctx, cartId, productId)
			}
		} else {
			redis.HSet(ctx, cartId, item)
		}
	}

	log.Printf("the product with id %v has been updated\n", productId)
	// ttl for the current user's shopping cart
	redis.Expire(ctx, cartId, 3*time.Hour)

	return nil
}

// delete the whole cart
func EmptyCart(redis *redis.Client, token string) error {
	ctx := context.Background()

	cartId, err := getUserIdFromToken(token)
	if err != nil {
		return fmt.Errorf("token: %v", err)
	}

	redis.Del(ctx, cartId)
	log.Printf("the cart with id %v has been deleted\n", cartId)
	return nil
}

// retrieves the cart
func GetCart(redis *redis.Client, token string) (map[string]string, error) {
	ctx := context.Background()

	cartId, err := getUserIdFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("token: %v", err)
	}

	cmd := redis.HGetAll(ctx, cartId).Val()
	log.Printf("the items of the card with id %v has been retrieved", cartId)

	return cmd, nil
}
