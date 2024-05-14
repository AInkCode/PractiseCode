package model

import "PratiseCode/dao"

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 增
func AddATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

// 删
func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}

// 改
func UpdataATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

// 查
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}
