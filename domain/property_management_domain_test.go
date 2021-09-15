package domain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createNewPropertyManager(t *testing.T) *PropertyManager {
	propertyManager, err := NewPropertyManager()
	require.NoError(t, err)
	return propertyManager
}

func TestListingProperties(t *testing.T) {
	propertyManager := createNewPropertyManager(t)
	ctx := context.Background()

	properties := propertyManager.ListProperties(ctx)

	assert.Len(t, properties, 0)
}

func TestAddingNewProperty(t *testing.T) {
	propertyManager := createNewPropertyManager(t)
	ctx := context.Background()

	attrs := CreatePropertyRequest{
		Name:        "Seabreeze",
		Description: "A beautiful lakefront property located on Lake Erie.",
	}
	property, err := propertyManager.RegisterProperty(ctx, attrs)
	require.NoError(t, err)
	assert.Equal(t, "Seabreeze", property.Name)

	properties := propertyManager.ListProperties(ctx)
	assert.Len(t, properties, 1)
	assert.Equal(t, []Property{property}, properties)

}
