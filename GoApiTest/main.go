package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"` // struct tag like `json..` used for JSON serialization or ORM mapping
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price,omitempty"` // can also omit the json mapping when the field data is not included or null
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {

	if len(albums) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no albums found"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum Album

	err := c.BindJSON(&newAlbum)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/add-album", addAlbum)
	router.Run("localhost:8080")
}

// func main() {

// 	a := Album{ID: 1, Title: "test", Artist: "test artust", Price: 29.03}
// 	data, err := json.Marshal(a) // marshal convert struct u into a JSON byte slice []byte
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	jsonStr := string(data)
// 	fmt.Println("data is : ", jsonStr)

// 	json.Unmarshal([]byte(jsonStr))

// }
