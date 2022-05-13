package user

import (
	"Firstgin/global"
	"Firstgin/models/ent"
	"Firstgin/models/ent/predicate"
	"Firstgin/models/ent/user"
	"Firstgin/utils"
	"context"
	"errors"
	"strconv"
)

func Service(ctx context.Context) *ServiceUser {
	art := &ServiceUser{}
	art.db = global.Db.User
	art.ctx = ctx
	return art
}

type ServiceUser struct {
	db  *ent.UserClient
	ctx context.Context
}

func (m *ServiceUser) Create(data *Date) (resp Resp, err error) {
	password, err := utils.EncryptPassword(data.Password)
	if err != nil {
		return resp, nil
	}
	db := m.db.Create().SetAge(data.Age).SetName(data.Username).SetPassword(password)
	u, err := db.Save(m.ctx)
	if err != nil {
		return resp, err
	}
	resp.User = u
	return resp, nil
}

func (m *ServiceUser) Delete(id string) (err error) {
	userObj, err := m.FindById(id)
	if err != nil {
		return errors.New("user is not find")
	}
	err = m.db.DeleteOne(userObj).Exec(m.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *ServiceUser) FindById(id string) (resp *ent.User, err error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return resp, errors.New("user is not find")
	}
	resp, err = m.db.Query().Where(user.IDEQ(intId)).First(m.ctx)
	if err != nil {
		return resp, errors.New("user is not find")
	}
	return resp, err
}

func (m *ServiceUser) FindByUsername(name string) (resp *ent.User, err error) {
	resp, err = m.db.Query().Where(user.NameEQ(name)).First(m.ctx)
	if err != nil {
		return resp, errors.New("user is not find")
	}
	return resp, err
}

func (m *ServiceUser) Update(id string, data *Date) (resp Resp, err error) {
	userObj, err := m.FindById(id)
	if err != nil {
		return resp, errors.New("user is not find")
	}
	password, err := utils.EncryptPassword(data.Password)
	if err != nil {
		return resp, nil
	}
	u, err := userObj.Update().SetAge(data.Age).SetName(data.Username).SetPassword(password).Save(m.ctx)
	if err != nil {
		return resp, err
	}
	resp.User = u
	return resp, nil
}

func (m *ServiceUser) Retrieve(data *DateListParams) (resp RespList, err error) {
	// 设置分页默认值
	utils.PipePagerFn(data)
	//查询条件数组
	var whereArr []predicate.User
	if data.Username != "" {
		whereArr = append(whereArr, user.NameContains(data.Username))
	}
	if data.Age != 0 {
		whereArr = append(whereArr, user.AgeEQ(data.Age))
	}
	// 查询
	db := m.db.Query().Where(whereArr...)
	// 获取总条数
	total, err := db.Count(m.ctx)
	if err != nil {
		return resp, err
	}
	resp.Total = total
	// 自动分页
	utils.PipeLimitFn(db, data)
	resp.Data, err = db.All(m.ctx)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
