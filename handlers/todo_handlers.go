package todo_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aboobakersiddiqr63/go-crud-postgresql/helper"
	todo "github.com/aboobakersiddiqr63/go-crud-postgresql/models"
	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	helper.GetCommonHeaders(w)

	response := getAllTasks()
	json.NewEncoder(w).Encode(response)
}

func Createtask(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "POST")

	var task todo.ToDoList
	json.NewDecoder(r.Body).Decode(&task)

	response := createtask(task)
	json.NewEncoder(w).Encode(response)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "POST")

	params := mux.Vars(r)

	response := taskComplete(params["id"])
	json.NewEncoder(w).Encode(response)
}

func UndoTaskStatus(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "PUT")

	params := mux.Vars(r)

	response := undoTaskStatus(params["id"])
	json.NewEncoder(w).Encode(response)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "DELETE")

	params := mux.Vars(r)

	response := deleteTask(params["id"])
	json.NewEncoder(w).Encode(response)
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "DELETE")

	response := deleteAllTask()
	json.NewEncoder(w).Encode(response)
}

func getAllTasks() []todo.ToDoList {
	var todos []todo.ToDoList

	err := helper.DB.Table("todos").Find(&todos).Error
	helper.HandleException(err, "getAllTasks")

	return todos
}

func createtask(task todo.ToDoList) string {
	err := helper.DB.Table("todos").Create(&task).Error
	helper.HandleException(err, "createtask")
	response := fmt.Sprintln("The task has been added")
	return response
}

func taskComplete(id string) string {
	var todo todo.ToDoList
	err := helper.DB.Table("todos").First(&todo, id).Error
	helper.HandleException(err, "Todo Not found")

	todo.Status = true

	saveErr := helper.DB.Table("todos").Save(&todo).Error
	helper.HandleException(saveErr, "Failed To Save")

	response := fmt.Sprintln("Task is marked as completed")
	return response
}

func undoTaskStatus(id string) string {
	var todo todo.ToDoList
	err := helper.DB.Table("todos").First(&todo, id).Error
	helper.HandleException(err, "Todo Not found")

	todo.Status = false

	saveErr := helper.DB.Table("todos").Save(&todo).Error
	helper.HandleException(saveErr, "Failed To Save")

	response := fmt.Sprintln("Task is marked as uncomplete")
	return response
}

func deleteTask(id string) string {
	var todo todo.ToDoList
	err := helper.DB.Table("todos").Delete(&todo, id).Error
	helper.HandleException(err, "deleteTask")

	response := fmt.Sprintln("Task is deleted")
	return response
}

func deleteAllTask() string {
	err := helper.DB.Table("todos").Where("status = ?", false).Delete(&todo.ToDoList{}).Error
	helper.HandleException(err, "All Tasks are deleted")

	response := fmt.Sprintln("All the tasks removed")
	return response
}
