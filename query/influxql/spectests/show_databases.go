package spectests

import (
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/stdlib/universe"
)

func init() {
	RegisterFixture(
		NewFixture(
			`SHOW DATABASES`,
			&flux.Spec{
				Operations: []*flux.Operation{
					{
						ID:   "databases0",
						Spec: &influxdb.DatabasesOpSpec{},
					},
					{
						ID: "rename0",
						Spec: &universe.RenameOpSpec{
							Columns: map[string]string{
								"databaseName": "name",
							},
						},
					},
					{
						ID: "extractcol0",
						Spec: &universe.KeepOpSpec{
							Columns: []string{
								"name",
							},
						},
					},
					{
						ID: "yield0",
						Spec: &universe.YieldOpSpec{
							Name: "0",
						},
					},
				},
				Edges: []flux.Edge{
					{Parent: "databases0", Child: "rename0"},
					{Parent: "rename0", Child: "extractcol0"},
					{Parent: "extractcol0", Child: "yield0"},
				},
				Now: Now(),
			},
		),
	)
}
