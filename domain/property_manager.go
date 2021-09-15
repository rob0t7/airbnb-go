/*
Copyright © 2021 Robert Jackiewicz <rob@jackiewicz.ca>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package domain

import (
	"context"

	"github.com/google/uuid"
)

type PropertyManager struct {
	repository PropertyRepository
}

func NewPropertyManager() (*PropertyManager, error) {
	return &PropertyManager{
		repository: &InMemoryPropertyRepository{},
	}, nil
}

func (pm *PropertyManager) ListProperties(ctx context.Context) []Property {
	return pm.repository.ListAll()
}

type CreatePropertyRequest struct {
	Name        string
	Description string
}

func (pm *PropertyManager) RegisterProperty(ctx context.Context, req CreatePropertyRequest) (Property, error) {
	property := Property{Name: req.Name, Description: req.Description}
	property.ID = uuid.NewString()
	pm.repository.CreateProperty(&property)
	return property, nil
}
