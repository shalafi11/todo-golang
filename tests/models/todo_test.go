package test

import (
	"testing"
	"todo-golang/models"
)

func TestCreateTodo(t *testing.T) {
	todo1 := models.NewTodo("testTitle", "testBody")
	todo2 := models.NewTodo("testTitle2", "testBody2")
	if todo1.ID != 0 || todo2.ID != 1 {
		t.Error("Invalid id")
	}
}

func TestAddTodoCollection(t *testing.T) {
	tc := models.NewTodoCollection()
	todo := models.NewTodo("testTitle", "testBody")

	if len(tc.TodoCollection) != 0 {
		t.Error("TodoCillection should be empty")
	}

	tc.Add(todo)
	if len(tc.TodoCollection) != 1 {
		t.Error("Todo was not added")
	}
}

func TestRetrieveTodo(t *testing.T) {
	tc := models.NewTodoCollection()
	todo1 := models.NewTodo("testTitle_1", "testBody_1")
	todo2 := models.NewTodo("testTitle_2", "testBody_2")

	tc.Add(todo1)
	tc.Add(todo2)

	result := tc.Retrieve(todo1.ID)
	if result == nil {
		t.Error("Should return: ", todo1, " but return nil")
	}

	result = tc.Retrieve(-10)
	if result != nil {
		t.Error("Should return nil instad of: ", result)
	}
}

func TestUpdateTodo(t *testing.T) {
	tc := models.NewTodoCollection()
	todo1 := models.NewTodo("testTitle_1", "testBody_1")

	tc.Add(todo1)
	result := tc.Update(todo1.ID, "UpdateTitle", "UpdateBody")

	if result == true {
		if tc.TodoCollection[0].Title != "UpdateTitle" && tc.TodoCollection[0].Body != "UpdateBody" {
			t.Error("Update did not update given object")
		}

	} else {
		t.Error("Update Should return true as output")
	}

}

func TestRetrieveAllTodo(t *testing.T) {
	tc := models.NewTodoCollection()
	todo1 := models.NewTodo("testTitle_1", "testBody_1")
	todo2 := models.NewTodo("testTitle_2", "testBody_2")

	result := tc.RetrieveAll()
	if len(result) != 0 {
		t.Error("Should return empty Todo slice")
	}

	tc.Add(todo1)
	result = tc.RetrieveAll()
	if len(result) != 1 {
		t.Error("Should return 1 element Todo slice")
	}

	tc.Add(todo2)
	result = tc.RetrieveAll()
	if len(result) != 2 {
		t.Error("Should return 2 element Todo slice")
	}
}

func TestDeleteTodo(t *testing.T) {
	tc := models.NewTodoCollection()
	todo1 := models.NewTodo("testTitle_1", "testBody_1")
	todo2 := models.NewTodo("testTitle_2", "testBody_2")

	tc.Add(todo1)
	tc.Add(todo2)

	result := tc.Delete(todo1.ID)
	if result != true && len(tc.TodoCollection) != 1 {
		t.Error("Should remove one TodoItem")
	}

	result = tc.Delete(-10)
	if result != false && len(tc.TodoCollection) != 1 {
		t.Error("Should return false and not remove any items from TodoCollection")
	}
}
