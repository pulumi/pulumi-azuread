// Copyright 2026, Pulumi Corporation.

package provider

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/require"
)

func TestMigrateApplicationState(t *testing.T) {
	t.Parallel()

	const guid = "9456d3da-3f88-4c7b-90ed-181ab1ae0473"

	cases := []struct {
		name string
		in   resource.PropertyMap
		want resource.PropertyMap
	}{
		{
			name: "v5 state is fully migrated",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty(guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty("/applications/" + guid),
			},
		},
		{
			name: "already-migrated state is unchanged",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty("/applications/" + guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty("/applications/" + guid),
			},
		},
		{
			name: "existing clientId is preserved",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty("custom"),
				"id":            resource.NewStringProperty(guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty("custom"),
				"id":            resource.NewStringProperty("/applications/" + guid),
			},
		},
		{
			name: "missing applicationId leaves clientId alone",
			in: resource.PropertyMap{
				"id": resource.NewStringProperty(guid),
			},
			want: resource.PropertyMap{
				"id": resource.NewStringProperty("/applications/" + guid),
			},
		},
		{
			name: "empty applicationId is not backfilled",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(""),
				"id":            resource.NewStringProperty(guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(""),
				"id":            resource.NewStringProperty("/applications/" + guid),
			},
		},
		{
			name: "empty clientId is treated as missing and backfilled",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(""),
				"id":            resource.NewStringProperty(guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty("/applications/" + guid),
			},
		},
		{
			name: "missing id field is a no-op for id rewrite",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
			},
		},
		{
			name: "non-string id is left untouched",
			in: resource.PropertyMap{
				"id": resource.NewNumberProperty(42),
			},
			want: resource.PropertyMap{
				"id": resource.NewNumberProperty(42),
			},
		},
		{
			name: "empty map is unchanged",
			in:   resource.PropertyMap{},
			want: resource.PropertyMap{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := migrateApplicationState(context.Background(), tc.in)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestMigrateServicePrincipalState(t *testing.T) {
	t.Parallel()

	const guid = "fae162c3-ade4-49b5-84ba-a7254efa23be"

	cases := []struct {
		name string
		in   resource.PropertyMap
		want resource.PropertyMap
	}{
		{
			name: "v5 state is fully migrated",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty(guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty("/servicePrincipals/" + guid),
			},
		},
		{
			name: "already-migrated state is unchanged",
			in: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty("/servicePrincipals/" + guid),
			},
			want: resource.PropertyMap{
				"applicationId": resource.NewStringProperty(guid),
				"clientId":      resource.NewStringProperty(guid),
				"id":            resource.NewStringProperty("/servicePrincipals/" + guid),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := migrateServicePrincipalState(context.Background(), tc.in)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}
