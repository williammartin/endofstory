package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type app struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type client struct {
	graph *graphql.Client
	token string
}

func (c client) AllApps() []app {
	req := graphql.NewRequest(`query {
      allApps(condition: {deleted: false}) {
        nodes{
          name
          uuid
        }
      }
    }`)

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	type allAppsData struct {
		AllApps struct {
			Nodes []app `json:"nodes"`
		} `json:"allApps"`
	}

	var appsData allAppsData
	must(c.graph.Run(context.Background(), req, &appsData))

	return appsData.AllApps.Nodes
}

func (c client) DestroyApps(apps []app) {
	for _, app := range apps {
		fmt.Printf("Destroying %s...\n", app.Name)

		req := graphql.NewRequest(`
      mutation ($data: UpdateAppByUuidInput!) {
        updateAppByUuid(input: $data) {
          app {
            uuid
          }
        }
      }`)

		type destroyAppsData struct {
			UUID     string      `json:"uuid"`
			AppPatch interface{} `json:"appPatch"`
		}

		data := destroyAppsData{
			UUID: app.UUID,
			AppPatch: struct {
				Deleted bool `json:"deleted"`
			}{
				Deleted: true,
			}}

		req.Var("data", data)

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

		must(c.graph.Run(context.Background(), req, nil))
	}
}
