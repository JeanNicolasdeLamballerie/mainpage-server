package database

import (
	"main/server/data"
	types "main/server/types"
)

var generatedData = mockData.Generate_data_blog()

func getCategories() []types.BlogCategory {

	data := generatedData.Categories
	return data[:]
}

func getAllPosts() []types.Post {
	data := generatedData.Posts
	return data[:]
}
func getPost(idx string) []types.Post {
	allData := generatedData.Posts
	for _, value := range allData {
		if value.Unique_idx == idx {
			return []types.Post{value}
		}
	}
	return []types.Post{}
}
func getCategory(idx string) []types.BlogCategory {
	allData := generatedData.Categories
	for _, value := range allData {
		if value.Unique_idx == idx {
			return []types.BlogCategory{value}
		}
	}
	return []types.BlogCategory{}
}

var DataBase types.DB = types.DB{
	Blog: types.BlogDatabase{
		GetCategories: getCategories,
		GetAllPosts:   getAllPosts,
		GetPost:       getPost,
		GetCategory:   getCategory,
	},
}
