package model

func InsertTask(task Task) error {
	return nil 
}

func UpdateTask(task Task) error {
	return nil
}

func iistTasks(task Task, selector, sort []string, skip, limit int) ([]*Task, int, error) {
	return nil, 0, nil
}

func LoadTask(id string, selector []string) (*Task, error) {
	return nil, nil
}

func InsertTag(tag Tag) error {
	return nil
}

func UpdateTag(tag Tag) error {
	return nil
}

func iistTags(task Task, selector, sort []string, skip, limit int) ([]*Tag, int, error) {
	return nil, 0, nil
}

func LoadTag(id string, selector []string) (*Tag, error) {
	return nil, nil
}

func RemoveTags(uid string, ids []string) error {
	return nil
}