package main

import (
	"context"
	"database/sql"
	"strings"

	"github.com/moby/moby/client"
)

func getDockerServices(db *sql.DB) ([]Service, error) {
	ctx := context.Background()

	cli, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, client.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	overrides, err := loadOverride(db)
	if err != nil {
		return nil, err
	}

	var services []Service
	hosts, _ := getNPMProxyHosts()

	for _, c := range containers.Items {
		nom := strings.TrimPrefix(c.Names[0], "/")

		override, exists := overrides[nom]
		if !exists || !override.Enabled {
			continue
		}

		url := override.URL
		if url == "" {
			url = matchNPMUrl(nom, hosts)
		}

		services = append(services, Service{
			Titre:       nom,
			Description: c.Status,
			Categories:  override.Category,
			Url:         url,
			Online:      c.State == "running",
		})
	}
	return services, nil
}

func getContainerNames() ([]string, error) {

	cli, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, client.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var names []string
	for _, c := range containers.Items {
		nom := strings.TrimPrefix(c.Names[0], "/")
		names = append(names, nom)
	}

	return names, nil
}
