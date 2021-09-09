package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

type Item struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type Cart struct {
	CartId string `json:"cartId"`
	Items  []Item `json:"items"`
}

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
		log.Printf("the cart with id %v created and the product with id %s added.\n", cartId, productName)
	} else {
		// If the item exists, we update its quantity
		if cmd := redis.SIsMember(ctx, cartId, productIdentity); cmd.Val() {
			// perform the operation on quantitiy. if current quantity of the product equals to zero we delete it from the cart
			if currentQuantity := redis.HIncrBy(ctx, productIdentity, "quantity", int64(productQuantity)); currentQuantity.Val() == 0 {
				log.Printf("the product with id %s has been updated\n", productName)
				redis.SRem(ctx, cartId, productIdentity)
				redis.Del(ctx, productIdentity)
			}
		} else {
			redis.HSet(ctx, productIdentity, item)
			redis.SAdd(ctx, cartId, productIdentity)
			log.Printf("the product with name %s added to the cart\n", productName)
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
		return fmt.Errorf("token: %v\n", err)
	}

	keys := redis.SMembers(ctx, cartId).Val()

	for _, productKey := range keys {
		fmt.Println(productKey)
		redis.Del(ctx, productKey)
	}
	redis.Expire(ctx, cartId, 500*time.Millisecond)
	log.Printf("the cart with id %v has been deleted\n", cartId)
	return nil
}

// retrieves the cart
func GetCart(redis *redis.Client, token string) (*Cart, error) {
	ctx := context.Background()
	allProducts := make(map[string][]interface{})

	cartId, err := getUserIdFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("token: %v", err)
	}

	keys := redis.SMembers(ctx, cartId).Val()

	cart := Cart{
		CartId: cartId,
		Items:  nil,
	}
	for _, productKey := range keys {
		item := redis.HGetAll(ctx, productKey).Val()
		log.Println("ITEM BURADA=>", item)
		log.Println("ITEM PRICE=>", item["price"])
		product := Item{
			Id:       item["id"],
			Name:     item["name"],
			Price:    item["price"],
			Quantity: item["quantity"],
		}
		cart.Items = append(cart.Items, product)
		allProducts[cartId] = append(allProducts[cartId], item)
	}

	log.Printf("the items of the card with id %v has been retrieved", cartId)
	return &cart, nil
}
