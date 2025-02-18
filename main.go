package main

import (
	"fmt"
	"kattudden/wiz-controller/config"
	"kattudden/wiz-controller/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		// Erstelle einen Slice, der alle Gruppen mitsamt den Bulb-Status enthält.
		var groups []GroupStatus

		// Für jede Gruppe in den Konfigurationsdaten: Status der Bulbs abfragen.
		for groupName, group := range config.Groups {
			var bulbs []BulbStatus

			for _, bulp := range group.Bulps {
				// Abfrage des Status, hier wird angenommen, dass controller.GetStatus(ip, port) einen Status zurückgibt
				status, err := controller.GetStatus(bulp.IP, bulp.Port)
				bulb := BulbStatus{
					IP:   bulp.IP,
					Port: bulp.Port,
					Name: bulp.Name,
				}

				if err != nil {
					bulb.Message = fmt.Sprintf("Fehler: %v", err)
				} else {
					// Beispiel: Wir nutzen status.Result.State um zu bestimmen, ob das Licht an ist.
					bulb.IsOn = status.Result.State
				}

				bulbs = append(bulbs, bulb)
			}

			groupStatus := GroupStatus{
				Name:  groupName,
				Bulbs: bulbs,
			}
			groups = append(groups, groupStatus)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"groups": groups,
		})
	})

	// static files
	r.StaticFile("/styles.css", "static/styles.css")
	r.StaticFile("/favicon.ico", "static/favicon.ico")
	r.StaticFile("/light-off.png", "images/light-off.png")
	r.StaticFile("/light-on.png", "images/light-on.png")
	r.StaticFile("/room.png", "images/room.png")

	r.Run()
}
