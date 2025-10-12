package service

import (
	albumModel "jackson.com/goApiDb/internal/albums/models"
	albumRepo "jackson.com/goApiDb/internal/albums/repository"
)

type AlbumService struct {
	albumRepo *albumRepo.AlbumRepository
}

func NewAlbumService(r *albumRepo.AlbumRepository) *AlbumService {
	return &AlbumService{albumRepo: r}
}

func (s *AlbumService) GetAlbums() ([]albumModel.Album, error) {
	return s.albumRepo.GetAlbums()
}

func (s *AlbumService) InsertAlbum(a albumModel.Album) (int, error) {
	return s.albumRepo.InsertAlbum(a)
}

func (s *AlbumService) UpdateAlbum(a albumModel.Album) error {
	return s.albumRepo.UpdateAlbum(a)
}

func (s *AlbumService) DeleteAlbum(id int) error {
	return s.albumRepo.DeleteAlbum(id)
}
