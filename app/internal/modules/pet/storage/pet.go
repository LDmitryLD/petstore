package storage

import (
	"fmt"
	"projects/LDmitryLD/petstore/app/helpers"
	"projects/LDmitryLD/petstore/app/internal/models"
	"sync"
)

type PetStorage struct {
	data               []*models.Pet
	primaryKeyIDx      map[int]*models.Pet
	autoIncrementCount int
	sync.Mutex
}

func NewPetStorage() *PetStorage {
	return &PetStorage{
		data:               make([]*models.Pet, 0, 13),
		primaryKeyIDx:      make(map[int]*models.Pet, 13),
		autoIncrementCount: 1,
	}
}

func (p *PetStorage) Create(pet models.Pet) models.Pet {
	p.Lock()
	defer p.Unlock()

	pet.ID = p.autoIncrementCount
	p.primaryKeyIDx[pet.ID] = &pet
	p.autoIncrementCount++
	p.data = append(p.data, &pet)

	return pet
}

func (p *PetStorage) Update(pet models.Pet) error {
	p.Lock()
	defer p.Unlock()

	if _, ok := p.primaryKeyIDx[pet.ID]; !ok {
		return fmt.Errorf("pet with id %d not found", pet.ID)
	}

	helpers.UpdatePet(&p.data, pet)

	p.primaryKeyIDx[pet.ID] = &pet

	return nil
}

func (p *PetStorage) Delete(petID int) error {
	p.Lock()
	defer p.Unlock()

	if _, ok := p.primaryKeyIDx[petID]; !ok {
		return fmt.Errorf("pet witj id %d not found", petID)
	}

	helpers.DeletePet(&p.data, petID)

	delete(p.primaryKeyIDx, petID)

	return nil
}

func (p *PetStorage) GetByID(petID int) (models.Pet, error) {
	p.Lock()
	defer p.Unlock()

	if v, ok := p.primaryKeyIDx[petID]; ok {
		return *v, nil
	}

	return models.Pet{}, fmt.Errorf("not found")
}

func (p *PetStorage) GetList() []models.Pet {
	p.Lock()
	defer p.Unlock()

	if len(p.data) == 0 {
		return []models.Pet{}
	}

	list := make([]models.Pet, len(p.data))

	for i, ptr := range p.data {
		list[i] = *ptr
	}
	return list
}

func (p *PetStorage) UploadImage(fileName string, id int) error {
	p.Lock()
	defer p.Unlock()

	p.primaryKeyIDx[id].PhotoUrls = append(p.primaryKeyIDx[id].PhotoUrls, fileName)

	for i, pet := range p.data {
		if pet.ID == id {
			p.data[i].PhotoUrls = append(p.data[i].PhotoUrls, fileName)
		}
	}

	return nil
}

func (p *PetStorage) FindByStatus(values []string) []models.Pet {
	p.Lock()
	defer p.Unlock()

	var list []models.Pet
	for _, pet := range p.data {
		for _, status := range values {
			if pet.Status == status {
				list = append(list, *pet)
			}
		}
	}

	return list
}

func (p *PetStorage) UpdateByID(petID int, name, status string) error {
	p.Lock()
	defer p.Unlock()

	if _, ok := p.primaryKeyIDx[petID]; !ok {
		return fmt.Errorf("pet not found")
	}

	if name != "" {
		p.primaryKeyIDx[petID].Name = name
	}

	if status != "" {
		p.primaryKeyIDx[petID].Status = status
	}

	helpers.UpdatePetByID(&p.data, petID, name, status)

	return nil
}
