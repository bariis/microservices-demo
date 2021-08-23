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
func AddItem(redis *redis.Client, productName, productId string, productQuantity int, productPrice float64, token string) error {
	ctx := context.Background()
	// user_id will be stored as cartId like key in redis cache
	cartId, err := getUserIdFromToken(token)
	if err != nil {
		return fmt.Errorf("token: %v", err)
	}

	item := map[string]interface{}{
		"id":       productId,
		"name":     productName,
		"quantity": productQuantity,
		"price":    productPrice,
	}

	productIdentity := cartId + ":" + productId
	cmd := redis.Exists(ctx, cartId)
	if cmd.Val() == 0 {
		redis.HSet(ctx, productIdentity, item)
		redis.SAdd(ctx, cartId, cartId+":"+productId)
		log.Printf("the cart with id %v created and the product with id %v added.\n", productId)
	} else {
		// If the item exists, we update its quantity
		if cmd := redis.SIsMember(ctx, cartId, productIdentity); cmd.Val() {
			// perform the operation on quantitiy. if current quantity of the product equals to zero we delete it from the cart
			if currentQuantity := redis.HIncrBy(ctx, productIdentity, "quantity", int64(productQuantity)); currentQuantity.Val() == 0 {
				log.Printf("the product with id %v has been updated\n", productId)
				redis.SRem(ctx, cartId, productIdentity)
				redis.Del(ctx, productIdentity)
			}
		} else {
			redis.SAdd(ctx, cartId, productIdentity)
		}
	}

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
