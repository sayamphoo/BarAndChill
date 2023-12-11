package repository

import (
	"errors"
	"sayamphoo/microservice/models/entity"
	"sayamphoo/microservice/models/wrapper"
	"sayamphoo/microservice/repository/repo"
)

type MemberRepo struct {
	repoMember repo.Database
}

func NewMemberRepo() *MemberRepo {
	return &MemberRepo{
		repoMember: repo.Database{
			Index: entity.EntityMemberName,
		},
	}
}

func (mr *MemberRepo) Save(model *wrapper.RegisterWrapper) (*entity.MemberEntity, error) {
	id, err := mr.repoMember.RepoSave(model)
	if err != nil {
		return nil, err
	}

	member := model.ToMemberEntity(id)
	return member, nil
}

func (mr *MemberRepo) FindByUsername(username string) (*entity.MemberEntity, error) {
	rawData, err := mr.repoMember.RepoFindByWord("username.keyword", username)
	if err != nil {
		return nil, err
	}

	if rawData.Total.Value == 0 {
		return nil, errors.New("User Not Found")
	}

	entity, _, err := entity.HitsToMemberEntity(*rawData)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (mr *MemberRepo) FindById(id string) (*entity.MemberEntity, error) {
	rawData, err := mr.repoMember.RepoFindByID(id)
	if err != nil {
		return nil, err
	}

	if rawData.Total.Value == 0 {
		return nil, errors.New("User Not Found")
	}

	entity, _, err := entity.HitsToMemberEntity(*rawData)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (mr *MemberRepo) Update(id string, wp interface{}) error {
	_, err := mr.repoMember.RepoUpdating(id, wp)
	return err
}

func (mr *MemberRepo) GetAllUsers() (*[]entity.MemberEntity, error) {
	result, err := mr.repoMember.RepoGetIndex()
	if err != nil {
		return nil, err
	}

	_, entityList, err := entity.HitsToMemberEntity(*result)
	if err != nil {
		return nil, err
	}

	return &entityList, nil
}
