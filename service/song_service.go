package service

import (
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository"
	"sayamphoo/microservice/utility"
)

type SongService struct{}

var repoSong *repository.SongRepo

func init() {
	repoSong = &repository.SongRepo{}
}

func (m *SongService) Request(wp *wrapper.SongRequestWrapper) int {
	userName, _ := repoMember.FindById(wp.UserID)
	wp.Timestamp = utility.GetTimeNow()
	wp.NameUser = (*userName).Name
	wp.State = "q"
	_, err := repoSong.Sava(wp) //บันทึกรายการเพลงที่ขอ
	if err != nil {
		return -1
	}

	return -1
}

func (m *SongService) GetSong() *[]entity.SongRequestEntity {
	entityList, err := repoSong.GetSongQueue()

	if err != nil {
		return nil
	}

	return entityList
}
