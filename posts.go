package ghostimport

func (gi *GhostImport) GetPosts() []DB_Data_Posts {
	DB := gi.Db[0]

	Data := DB.Data

	posts := Data.Posts

	return posts
}

func (gi *GhostImport) GetPublishedPosts() []DB_Data_Posts {
	var publishedPosts []DB_Data_Posts

	DB := gi.Db[0]
	Data := DB.Data
	posts := Data.Posts

	for _, p := range posts {
		if p.Status != "draft" {
			publishedPosts = append(publishedPosts, p)
		}
	}

	return publishedPosts
}
