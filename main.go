package main

import (
	"fmt"
	"kattudden/wiz-controller/config"
	"kattudden/wiz-controller/controller"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	for groupName, group := range config.Groups {
		fmt.Println("Gruppe:", groupName)

		for _, bulp := range group.Bulps {
			fmt.Printf("  Bulp: ip: %s, port: %s\n", bulp.IP, bulp.Port)
			status, err := controller.GetStatus(bulp.IP, bulp.Port)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("status:", status.Result.State)
		}
	}
}
