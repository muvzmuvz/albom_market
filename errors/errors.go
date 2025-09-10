package errors

import (
	"errors"
	"gin/db"
	"gin/sturct"
	"strings"
)

// Проверка нового альбома при создании
func ValidateAlbums(newAlbum sturct.Album) (sturct.Album, error) {
	newAlbum.Title = strings.TrimSpace(newAlbum.Title)
	newAlbum.Artist = strings.TrimSpace(newAlbum.Artist)
	newAlbum.ImagePath = strings.TrimSpace(newAlbum.ImagePath)

	if newAlbum.ID == "" {
		return sturct.Album{}, errors.New("айди не может быть пустым")
	}
	if newAlbum.Title == "" {
		return sturct.Album{}, errors.New("заголовок не может быть пустым")
	}
	if newAlbum.Artist == "" {
		return sturct.Album{}, errors.New("артист не может быть пустым")
	}
	if newAlbum.ImagePath == "" {
		return sturct.Album{}, errors.New("путь к изображению не может быть пустым")
	}

	var count int64
	db.DB.Model(&sturct.Album{}).Where("title = ?", newAlbum.Title).Count(&count)
	if count > 0 {
		return sturct.Album{}, errors.New("альбом с таким заголовком уже существует")
	}

	return newAlbum, nil
}

// Проверка обновления альбома
func ValidateUpdateAlbum(album sturct.Album) error {
	album.Title = strings.TrimSpace(album.Title)
	album.Artist = strings.TrimSpace(album.Artist)
	album.ImagePath = strings.TrimSpace(album.ImagePath)

	if album.Title == "" {
		return errors.New("заголовок не может быть пустым")
	}
	if album.Artist == "" {
		return errors.New("артист не может быть пустым")
	}
	if album.ImagePath == "" {
		return errors.New("путь к изображению не может быть пустым")
	}

	// Проверка уникальности title для других альбомов
	var count int64
	db.DB.Model(&sturct.Album{}).
		Where("title = ? AND id != ?", album.Title, album.ID).
		Count(&count)
	if count > 0 {
		return errors.New("альбом с таким заголовком уже существует")
	}

	return nil
}

// Проверка пользователя
func ValidateUser(newUser sturct.User) (sturct.User, error) {
	newUser.Username = strings.TrimSpace(newUser.Username)
	newUser.Password = strings.TrimSpace(newUser.Password)

	if newUser.Username == "" {
		return sturct.User{}, errors.New("username не может быть пустым")
	}
	if newUser.Password == "" {
		return sturct.User{}, errors.New("password не может быть пустым")
	}

	var count int64
	db.DB.Model(&sturct.User{}).Where("username = ?", newUser.Username).Count(&count)
	if count > 0 {
		return sturct.User{}, errors.New("пользователь с таким username уже существует")
	}

	return newUser, nil
}
