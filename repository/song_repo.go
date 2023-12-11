package repository

import (
	"encoding/json"
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository/repo"
	"sayamphoo/microservice/utility"
)

type SongRepo struct {
}

var repoMusic = repo.Database{
	Index: entity.EntitySongRequestEntityName,
}

func (m *SongRepo) Sava(wp *wrapper.SongRequestWrapper) (string, error) {
	return repoMusic.RepoSave(wp)
}

func (m *SongRepo) GetSongQueue() (*[]entity.SongRequestEntity, error) {
	data, err := repoMusic.RepoFindByWordRangeLte("timestamp", utility.GetTimeNow())
	if err != nil {
		return nil, err
	}
	
	return entity.HitsToEntitySong(*data), nil

}
func (m *SongRepo) SetSongState(id string) *bool {
	json := json.RawMessage(`{"state" : "Already played"}`)
	_, err := repoMusic.RepoUpdating(id, json)

	if err != nil {
		return nil
	}

	result := true
	return &result
}