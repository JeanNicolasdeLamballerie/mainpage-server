package v1routes

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	. "main/server/database"
	types "main/server/types"
	"net/http"
)

func V1Blog() http.Handler {
	v1router := chi.NewRouter()
	v1router.Route("/category", func(r chi.Router) {
		r.Get("/", getCategories)
		v1router.Route("/category/{idx}", func(r chi.Router) {
			r.Use(CategoryCtx)
			r.Get("/", getCategory)
		})

	})

	v1router.Route("/post", func(r chi.Router) {
		r.Get("/", getPosts)
		v1router.Route("/post/{idx}", func(r chi.Router) {
			r.Use(PostCtx)
			r.Get("/", getPost)
		})

	})
	return v1router
}

// ----------------------------------------------------------------------------------------------------------------------//
// ///////////////////////////////////////////////////// Categories ///////////////////////////////////////////////////////
// ----------------------------------------------------------------------------------------------------------------------//

func getCategories(w http.ResponseWriter, r *http.Request) {
	categories := DataBase.Blog.GetCategories() //, ok := ctx.Value("category").(*types.BlogCategory)
	// if !ok {
	// 	http.Error(w, http.StatusText(422), 422)
	// 	return
	// }
	w.Write([]byte(fmt.Sprintf("%s", categories)))
}

func CategoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idx := chi.URLParam(r, "idx")
		// TODO integrate error
		category := DataBase.Blog.GetCategory(idx)
		// if err != nil {
		//   http.Error(w, http.StatusText(404), 404)
		//   return
		// }
		ctx := context.WithValue(r.Context(), "category", category)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value("category").(*types.BlogCategory)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("%s", category)))
}

// -----------------------------------------------------------------------------------------------------------------//
// ///////////////////////////////////////////////////// POSTS ///////////////////////////////////////////////////////
// -----------------------------------------------------------------------------------------------------------------//

func getPosts(w http.ResponseWriter, r *http.Request) {
	categories := DataBase.Blog.GetCategories() //, ok := ctx.Value("category").(*types.BlogCategory)
	// if !ok {
	// 	http.Error(w, http.StatusText(422), 422)
	// 	return
	// }
	w.Write([]byte(fmt.Sprintf("%s", categories)))
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idx := chi.URLParam(r, "idx")
		// TODO integrate error
		category := DataBase.Blog.GetCategory(idx)
		// if err != nil {
		//   http.Error(w, http.StatusText(404), 404)
		//   return
		// }
		ctx := context.WithValue(r.Context(), "category", category)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category, ok := ctx.Value("category").(*types.BlogCategory)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("%s", category)))
}
