package main

import (
	"context"

	"github.com/moby/moby/client"
)

func getDockerServices() ([]Service, error) {
	ctx := context.Background()

	cli, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, client.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var services []Service
	for _, c := range containers.Items {
		services = append(services, Service{
			Titre:       c.Names[0],
			Description: c.Status,
			Categories:  "Homelab",
		})
	}
	return services, nil
}
