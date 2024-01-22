package service

import (
	"io"
	"log"
	"net/http"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/pet/storage"
	"sort"
)

type PetService struct {
	storage storage.PetStorager
}

func NewPetService(storage storage.PetStorager) *PetService {
	return &PetService{storage: storage}
}

func (p *PetService) Create(in CreateIn) CreateOut {
	pet := p.storage.Create(in.Pet)

	return CreateOut{Pet: pet}
}

func (p *PetService) Update(in UpdateIn) UpdateOut {
	err := p.storage.Update(in.Pet)

	return UpdateOut{Error: err}
}

func (p *PetService) Delete(petID int) error {
	return p.storage.Delete(petID)
}

func (p *PetService) GetByID(in GetByIDIn) GetByIDOut {
	pet, err := p.storage.GetByID(in.PetID)

	return GetByIDOut{Pet: pet, Error: err}
}

func (p *PetService) GetList() GetListOut {
	pets := p.storage.GetList()

	sort.Slice(pets, func(i, j int) bool {
		return i < j
	})

	return GetListOut{List: pets}
}

func (p *PetService) UploadImage(in UploadIn) UploadOut {
	fileName := in.Handler.Filename

	err := p.storage.UploadImage(fileName, in.PetID)

	out := UploadOut{
		FileName: fileName,
		Error:    err,
	}

	file, err := io.ReadAll(in.File)
	if err != nil {
		log.Println("ошибка при чтении файла для определения его типа:", err.Error())
		out.FileType = "unknown"
		return out
	}

	fileType := http.DetectContentType(file)

	out.FileType = fileType

	return out
}

func (p *PetService) FindByStatus(values []string) []models.Pet {
	return p.storage.FindByStatus(values)
}

func (p *PetService) UpdateByID(in UpdateByIDIn) error {
	return p.storage.UpdateByID(in.PetID, in.Name, in.Status)
}
