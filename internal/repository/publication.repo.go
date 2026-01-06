package repository

import (
	"time"
	"web-lab/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicationRepository interface {
	Create(publication *entity.Publication) error
	Delete(publicationID uuid.UUID) error
	Update(publication *entity.Publication) error

	FindByID(publicationID uuid.UUID) (*entity.Publication, error)
	FindByUserID(userID uuid.UUID, isDraft bool) ([]entity.Publication, error)
	FindAll(isDraft bool) ([]entity.Publication, error)

	GetAllCategories() ([]entity.Category, error)

	GetAllFavByUserID(userID uuid.UUID) ([]entity.FavoritePublications, error)
	CheckIsFavorite(userID, publicationID uuid.UUID) bool
	SaveFavorite(userID, publicationID uuid.UUID) error
	RemoveFavorite(userID, publicationID uuid.UUID) error
	// GetAllFavByPubID(userID uuid.UUID) ([]entity.FavoritePublications, error)
}

type publicationRepository struct {
	db *gorm.DB
}

func NewPublicationRepository(db *gorm.DB) PublicationRepository {
	return &publicationRepository{db: db}
}

func (p *publicationRepository) Create(publication *entity.Publication) error {
	return p.db.Create(&publication).Error
}

func (p *publicationRepository) Delete(publicationID uuid.UUID) error {
	return p.db.Delete(&entity.Publication{}, publicationID).Error
}

func (p *publicationRepository) Update(publication *entity.Publication) error {
	response := p.db.Model(&entity.Publication{}).Where("id = ?", publication.ID).Updates(map[string]interface{}{
		"title":       publication.Title,
		"description": publication.Description,
		"updated_at":  time.Now(),
	})
	return response.Error
}

func (p *publicationRepository) FindByID(publicationID uuid.UUID) (*entity.Publication, error) {
	var publication entity.Publication
	err := p.db.
		Preload("User").
		Preload("PublicationCategories").
		Preload("PublicationCategories.Category").
		First(&publication, publicationID).Error
	if err != nil {
		return nil, err
	}
	return &publication, nil
}

func (p *publicationRepository) FindByUserID(userID uuid.UUID, isDraft bool) ([]entity.Publication, error) {
	var publications []entity.Publication
	err := p.db.
		Where("is_draft = ?", isDraft).
		Where("user_id = ?", userID).
		Preload("User").
		Preload("PublicationCategories").
		Preload("PublicationCategories.Category").
		Find(&publications).Error
	if err != nil {
		return nil, err
	}
	return publications, nil
}

func (p *publicationRepository) FindAll(isDraft bool) ([]entity.Publication, error) {
	var publications []entity.Publication
	if err := p.db.
		Where("is_draft = ?", isDraft).
		Preload("User").
		Preload("PublicationCategories").
		Preload("PublicationCategories.Category").
		Find(&publications).Error; err != nil {
		return nil, err
	}
	return publications, nil
}

func (p *publicationRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	err := p.db.Preload("PublicationCategories").Preload("PublicationCategories.Publication").Order("name ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (p *publicationRepository) GetAllFavByUserID(userID uuid.UUID) ([]entity.FavoritePublications, error) {
	var fav []entity.FavoritePublications
	err := p.db.
		Where("user_id = ?", userID).
		Preload("User").
		Preload("Publication").
		Preload("Publication.PublicationCategories").
		Preload("Publication.PublicationCategories.Category").
		Find(&fav).Error
	if err != nil {
		return nil, err
	}
	return fav, nil
}

func (p *publicationRepository) CheckIsFavorite(userID, publicationID uuid.UUID) bool {
	err := p.db.Where("user_id = ? AND publication_id = ?", userID, publicationID).First(&entity.FavoritePublications{}).Error
	if err != nil {
		return false
	}
	return true
}

func (p *publicationRepository) SaveFavorite(userID, publicationID uuid.UUID) error {
	return p.db.Create(&entity.FavoritePublications{
		UserID:        userID,
		PublicationID: publicationID,
	}).Error
}

func (p *publicationRepository) RemoveFavorite(userID, publicationID uuid.UUID) error {
	return p.db.Where("user_id = ? AND publication_id = ?", userID, publicationID).Delete(&entity.FavoritePublications{}).Error
}

//func (p *publicationRepository) GetAllFavByPubID(publicationID uuid.UUID) ([]entity.FavoritePublications, error) {
//	var fav []entity.FavoritePublications
//	err := p.db.Where("publication_id = ?", publicationID).Preload("Publication").Find(&fav).Error
//	if err != nil {
//		return nil, err
//	}
//	return fav, nil
//}
