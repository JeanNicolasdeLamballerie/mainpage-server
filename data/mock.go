package mockData

import (
	types "main/server/types"
	"math/rand"
	"strconv"
	"time"
)

// Mock Init values
const MOCK_CATEGORY_NUMBER = 6

// MUST be superior or equal to categ numbers
const MOCK_POST_NUMBER = 12

type BlogDataMock struct {
	Categories [MOCK_CATEGORY_NUMBER]types.BlogCategory
	Posts      [MOCK_POST_NUMBER]types.Post
}

// Generators

func Generate_data_blog() BlogDataMock {
	timenow := time.Now()
	new_categories := [MOCK_CATEGORY_NUMBER]types.BlogCategory{}
	new_posts := [MOCK_POST_NUMBER]types.Post{}

	// Categories
	for index := range new_categories {
		index_str := strconv.Itoa(index)
		// use original loop to implement some blogs
		new_posts[index] = types.Post{
			Unique_idx:  "postidx_" + index_str,
			Name:        "Blog post " + index_str,
			Date:        timenow,
			Description: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
			Content:     "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
		}
		new_categories[index] = types.BlogCategory{
			Unique_idx: `idx_` + index_str,
			Name:       index_str + "Some Blog",
			Latest: struct {
				Unique_idx  string
				Name        string
				Date        time.Time
				Description string
			}{Unique_idx: "postidx_" + index_str,
				Name:        "Blog Post " + index_str,
				Date:        timenow,
				Description: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."},
		}

	}

	for index := MOCK_CATEGORY_NUMBER - 1; index < (MOCK_POST_NUMBER - 1); index++ {
		index_str := strconv.Itoa(index)
		timethen := time.Now().AddDate(-rand.Intn(4), -rand.Intn(4), -rand.Intn(20))
		new_posts[index] = types.Post{
			Unique_idx:  "postidx_" + index_str,
			Name:        "Blog post " + index_str,
			Date:        timethen,
			Description: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
			Content:     "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
		}
	}
	return BlogDataMock{Categories: new_categories, Posts: new_posts}
}
