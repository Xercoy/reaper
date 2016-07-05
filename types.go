package ghostimport

type DB_Data_Posts struct {
	AuthorID        int         `json:"author_id"`
	CreatedAt       int         `json:"created_at"`
	CreatedBy       int         `json:"created_by"`
	Featured        int         `json:"featured"`
	HTML            string      `json:"html"`
	ID              int         `json:"id"`
	Image           interface{} `json:"image"`
	Language        string      `json:"language"`
	Markdown        string      `json:"markdown"`
	MetaDescription interface{} `json:"meta_description"`
	MetaTitle       interface{} `json:"meta_title"`
	Mobiledoc       interface{} `json:"mobiledoc"`
	Page            int         `json:"page"`
	PublishedAt     int         `json:"published_at"`
	PublishedBy     int         `json:"published_by"`
	Slug            string      `json:"slug"`
	Status          string      `json:"status"`
	Title           string      `json:"title"`
	UpdatedAt       int         `json:"updated_at"`
	UpdatedBy       int         `json:"updated_by"`
	UUID            string      `json:"uuid"`
	Visibility      string      `json:"visibility"`
}

type DB_Data struct {
	AppFields   []interface{} `json:"app_fields"`
	AppSettings []interface{} `json:"app_settings"`
	Apps        []interface{} `json:"apps"`
	Permissions []struct {
		ActionType string      `json:"action_type"`
		CreatedAt  int         `json:"created_at"`
		CreatedBy  int         `json:"created_by"`
		ID         int         `json:"id"`
		Name       string      `json:"name"`
		ObjectID   interface{} `json:"object_id"`
		ObjectType string      `json:"object_type"`
		UpdatedAt  int         `json:"updated_at"`
		UpdatedBy  int         `json:"updated_by"`
		UUID       string      `json:"uuid"`
	} `json:"permissions"`
	PermissionsApps  []interface{} `json:"permissions_apps"`
	PermissionsRoles []struct {
		ID           int `json:"id"`
		PermissionID int `json:"permission_id"`
		RoleID       int `json:"role_id"`
	} `json:"permissions_roles"`
	PermissionsUsers []interface{}   `json:"permissions_users"`
	Posts            []DB_Data_Posts `json:"posts"`
	PostsTags        []struct {
		ID        int `json:"id"`
		PostID    int `json:"post_id"`
		SortOrder int `json:"sort_order"`
		TagID     int `json:"tag_id"`
	} `json:"posts_tags"`
	Roles []struct {
		CreatedAt   int    `json:"created_at"`
		CreatedBy   int    `json:"created_by"`
		Description string `json:"description"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		UpdatedAt   int    `json:"updated_at"`
		UpdatedBy   int    `json:"updated_by"`
		UUID        string `json:"uuid"`
	} `json:"roles"`
	RolesUsers []struct {
		ID     int `json:"id"`
		RoleID int `json:"role_id"`
		UserID int `json:"user_id"`
	} `json:"roles_users"`
	Settings []struct {
		CreatedAt int    `json:"created_at"`
		CreatedBy int    `json:"created_by"`
		ID        int    `json:"id"`
		Key       string `json:"key"`
		Type      string `json:"type"`
		UpdatedAt int    `json:"updated_at"`
		UpdatedBy int    `json:"updated_by"`
		UUID      string `json:"uuid"`
		Value     string `json:"value"`
	} `json:"settings"`
	Subscribers []interface{} `json:"subscribers"`
	Tags        []struct {
		CreatedAt       int         `json:"created_at"`
		CreatedBy       int         `json:"created_by"`
		Description     interface{} `json:"description"`
		ID              int         `json:"id"`
		Image           interface{} `json:"image"`
		MetaDescription interface{} `json:"meta_description"`
		MetaTitle       interface{} `json:"meta_title"`
		Name            string      `json:"name"`
		ParentID        interface{} `json:"parent_id"`
		Slug            string      `json:"slug"`
		UpdatedAt       int         `json:"updated_at"`
		UpdatedBy       int         `json:"updated_by"`
		UUID            string      `json:"uuid"`
		Visibility      string      `json:"visibility"`
	} `json:"tags"`
	Users []struct {
		Accessibility   interface{} `json:"accessibility"`
		Bio             string      `json:"bio"`
		Cover           interface{} `json:"cover"`
		CreatedAt       int         `json:"created_at"`
		CreatedBy       int         `json:"created_by"`
		Email           string      `json:"email"`
		Facebook        interface{} `json:"facebook"`
		ID              int         `json:"id"`
		Image           string      `json:"image"`
		Language        string      `json:"language"`
		LastLogin       int         `json:"last_login"`
		Location        string      `json:"location"`
		MetaDescription interface{} `json:"meta_description"`
		MetaTitle       interface{} `json:"meta_title"`
		Name            string      `json:"name"`
		Password        string      `json:"password"`
		Slug            string      `json:"slug"`
		Status          string      `json:"status"`
		Tour            interface{} `json:"tour"`
		Twitter         interface{} `json:"twitter"`
		UpdatedAt       int         `json:"updated_at"`
		UpdatedBy       int         `json:"updated_by"`
		UUID            string      `json:"uuid"`
		Visibility      string      `json:"visibility"`
		Website         string      `json:"website"`
	} `json:"users"`
}

type DB_Meta struct {
	ExportedOn int    `json:"exported_on"`
	Version    string `json:"version"`
}

type DB struct {
	Data DB_Data `json:"data"`
	Meta DB_Meta `json:"meta"`
}

type GhostImport struct {
	Db []DB `json:"db"`
}
