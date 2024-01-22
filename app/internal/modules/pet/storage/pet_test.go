package storage

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"reflect"
	"sync"
	"testing"
)

func TestPetStorage_GetByID(t *testing.T) {
	type fields struct {
		data               []*models.Pet
		primaryKeyIDx      map[int]*models.Pet
		autoIncrementCount int
		Mutex              sync.Mutex
	}
	type args struct {
		petID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Pet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PetStorage{
				data:               tt.fields.data,
				primaryKeyIDx:      tt.fields.primaryKeyIDx,
				autoIncrementCount: tt.fields.autoIncrementCount,
				Mutex:              tt.fields.Mutex,
			}
			got, err := p.GetByID(tt.args.petID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PetStorage.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PetStorage.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPetStorage_Create(t *testing.T) {
	type fields struct {
		data               []*models.Pet
		primaryKeyIDx      map[int]*models.Pet
		autoIncrementCount int
		Mutex              sync.Mutex
	}
	type args struct {
		pet models.Pet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.Pet
	}{
		{
			name: "first test",
			args: args{
				pet: models.Pet{
					ID: 0,
					Category: models.Category{
						ID:   0,
						Name: "test category",
					},
					Name: "Alma",
				},
			},
			fields: fields{
				data:               make([]*models.Pet, 0, 13),
				primaryKeyIDx:      make(map[int]*models.Pet, 13),
				autoIncrementCount: 0,
				Mutex:              sync.Mutex{},
			},
			want: models.Pet{
				ID: 0,
				Category: models.Category{
					ID:   0,
					Name: "test category",
				},
				Name: "Alma",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PetStorage{
				data:               tt.fields.data,
				primaryKeyIDx:      tt.fields.primaryKeyIDx,
				autoIncrementCount: tt.fields.autoIncrementCount,
				Mutex:              tt.fields.Mutex,
			}
			if got := p.Create(tt.args.pet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PetStorage.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
