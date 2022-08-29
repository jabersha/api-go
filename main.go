package main

import(
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
)

type exemplo struct {
	ID	string		
	Item string		
	Completed bool	
}

var exemplos = []exemplo{
	{ID: "1", Item: "Exemplo1", Completed: false},
	{ID: "2", Item: "Exemplo2", Completed: false},
	{ID: "3", Item: "Exemplo3", Completed: true},
}

func getExamplos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, exemplos)
}

func postExample(context *gin.Context){
	var newExemplo exemplo

	if err := context.BindJSON(&newExemplo); err != nil {
		return
	}

	exemplos = append(exemplos, newExemplo)
	context.IndentedJSON(http.StatusCreated, newExemplo)
}

func getExemploByID(id string) (*exemplo, error) {
	for i, t := range exemplos {
		if t.ID == id {
			return &exemplos[i], nil
		}
	}

	return nil, errors.New("Exemplo não existe!!")
}

func getExamplo(context *gin.Context){
	id := context.Param("id")
	exemplo, err := getExemploByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, errors.New("Error 404"))
		return
	}

	context.IndentedJSON(http.StatusOK, exemplo)
}

func main(){
	router := gin.Default()
	router.GET("/exemplo", getExamplos)
	router.GET("/exemplo/:id", getExamplo)
	router.POST("/exemplo", postExample)
	router.Run("localhost:9090")

}