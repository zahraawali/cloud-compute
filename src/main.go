package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int64
	Title string
}

var products []Product = []Product{
	{ID: 1, Title: "product 1"},
	{ID: 2, Title: "product 2"},
}
var cartItems []Product

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/", handleAllProducts)
	router.GET("/:id", handleProduct)
	router.GET("/cart", handleCart)
	router.GET("/cart/remove/all", handleRemoveAll)
	router.GET("/cart/remove/:id", handleRemoveById)

	s := http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	fmt.Printf("You're running: %s\n", ":8000")
	s.ListenAndServe()
}

func handleAllProducts(c *gin.Context) {
	c.HTML(http.StatusOK, "products.html", gin.H{
		"products":  products,
		"cartCount": strconv.Itoa(len(cartItems)),
	})
}

func handleProduct(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	uid, _ := strconv.ParseInt(id, 8, 64)

	for i := range products {
		if products[i].ID == uid {
			cartItems = append(cartItems, products[i])
			break
		}
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func handleCart(c *gin.Context) {
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"products": cartItems,
	})
}

func handleRemoveAll(c *gin.Context) {
	cartItems = []Product{}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func handleRemoveById(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	uid, _ := strconv.ParseInt(id, 8, 64)
	var newProducts []Product

	for i := range cartItems {
		if cartItems[i].ID != uid {
			newProducts = append(newProducts, cartItems[i])
		}
	}
	cartItems = newProducts

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
