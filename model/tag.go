package model



func InsertTag(tag Tag) error {
	return nil
}

func UpdateTag(tag Tag) error {
	return nil
}

func ListTags(task Task, selector, sort []string, skip, limit int) ([]*Tag, int, error) {
	return nil, 0, nil
}

func LoadTag(id string, selector []string) (*Tag, error) {
	return nil, nil
}

func RemoveTags(uid string, ids []string) error {
	return nil
}

func CheckTagsExist(tags []string) error {
	if len(tags) == 0 {
		return nil
	}

	return nil
}