package acrtoken

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"testing"

	mgmtcontainerregistry "github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2019-06-01-preview/containerregistry"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/golang/mock/gomock"

	"github.com/Azure/ARO-RP/pkg/api"
	mock_containerregistry "github.com/Azure/ARO-RP/pkg/util/mocks/azureclient/mgmt/containerregistry"
	mock_env "github.com/Azure/ARO-RP/pkg/util/mocks/env"
)

func TestEnsureTokenAndPassword(t *testing.T) {
	ctx := context.Background()

	controller := gomock.NewController(t)
	defer controller.Finish()

	env := mock_env.NewMockInterface(controller)
	env.EXPECT().ACRResourceID().AnyTimes().Return("/subscriptions/93aeba23-2f76-4307-be82-02921df010cf/resourceGroups/global/providers/Microsoft.ContainerRegistry/registries/arointsvc")

	tokens := mock_containerregistry.NewMockTokensClient(controller)
	tokens.EXPECT().
		CreateAndWait(ctx, "global", "arointsvc", gomock.Any(), mgmtcontainerregistry.Token{
			TokenProperties: &mgmtcontainerregistry.TokenProperties{
				ScopeMapID: to.StringPtr(env.ACRResourceID() + "/scopeMaps/_repositories_pull"),
				Status:     mgmtcontainerregistry.TokenStatusEnabled,
			},
		}).
		Return(nil)

	registries := mock_containerregistry.NewMockRegistriesClient(controller)
	registries.EXPECT().
		GenerateCredentials(ctx, "global", "arointsvc", gomock.Any()).
		Return(mgmtcontainerregistry.GenerateCredentialsResult{
			Passwords: &[]mgmtcontainerregistry.TokenPassword{
				{
					Value: to.StringPtr("foo"),
				},
			},
		}, nil)

	r, err := azure.ParseResourceID(env.ACRResourceID())
	if err != nil {
		t.Fatal(err)
	}

	m := &manager{
		env: env,
		r:   r,

		registries: registries,
		tokens:     tokens,
	}

	password, err := m.EnsureTokenAndPassword(ctx, &api.RegistryProfile{Username: "token-12345"})
	if err != nil {
		t.Fatal(err)
	}
	if password != "foo" {
		t.Error(password)
	}
}
