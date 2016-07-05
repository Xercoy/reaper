package ghostimport

func (gi *GhostImport) GetPosts() ([]DB_Data_Posts, error) {
	DB := gi.Db[0]

	Data := DB.Data

	Posts := Data.Posts

	return Posts, nil
}
