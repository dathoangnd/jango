package routers

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const STATIC_DIR = "/static/"
const UPLOAD_DIR = "/upload/"

func InitRouters() {
	r := mux.NewRouter()

	uploadFolder := UPLOAD_DIR[1:]
	if _, err := os.Stat(uploadFolder); os.IsNotExist(err) {
    os.Mkdir(uploadFolder, 0777)
	}
	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("." + STATIC_DIR))))
	r.PathPrefix(UPLOAD_DIR).Handler(http.StripPrefix(UPLOAD_DIR, http.FileServer(http.Dir("." + UPLOAD_DIR))))

	user := r.PathPrefix("/user").Subrouter()
	{
		user.HandleFunc("/single", getUser).Methods("GET")
		user.HandleFunc("/register", register).Methods("POST")
		user.HandleFunc("/login", login).Methods("POST")
		user.HandleFunc("/update", updateUser).Methods("PUT")
		user.HandleFunc("/changepassword", changePassword).Methods("PUT")
	}

	image := r.PathPrefix("/image") .Subrouter()
	{
		image.HandleFunc("/upload", uploadImage).Methods("POST")
	}

	pdf := r.PathPrefix("/pdf") .Subrouter()
	{
		pdf.HandleFunc("/upload", uploadPdf).Methods("POST")
	}
	
	article := r.PathPrefix("/article") .Subrouter()
	{
		article.HandleFunc("/list", getArticles).Methods("GET")
		article.HandleFunc("/single", getArticle).Methods("GET")
		article.HandleFunc("/single", removeArticle).Methods("DELETE")
		article.HandleFunc("/savedraft", saveDraft).Methods("POST", "PUT")
		article.HandleFunc("/publish", publish).Methods("POST", "PUT")
		article.HandleFunc("/comment", getComments).Methods("GET")
		article.HandleFunc("/comment", comment).Methods("POST")
		article.HandleFunc("/comment", removeComment).Methods("DELETE")
		article.HandleFunc("/like", like).Methods("POST")
		article.HandleFunc("/like", unlike).Methods("DELETE")
		article.HandleFunc("/status", changeStatus).Methods("PUT")
	}

	document := r.PathPrefix("/document") .Subrouter()
	{
		document.HandleFunc("/list", getDocuments).Methods("GET")
		document.HandleFunc("/single", getDocument).Methods("GET")
		document.HandleFunc("/single", removeDocument).Methods("DELETE")
		document.HandleFunc("/publish", publishDocument).Methods("POST", "PUT")
		document.HandleFunc("/status", changeDocumentStatus).Methods("PUT")
	}

	voca := r.PathPrefix("/vocabulary") .Subrouter()
	{
		voca.HandleFunc("/list", getVocas).Methods("GET")
		voca.HandleFunc("/like", getLikedVocas).Methods("GET")
		voca.HandleFunc("/like", likeVoca).Methods("POST")
		voca.HandleFunc("/like", unlikeVoca).Methods("DELETE")
	}

	grammar := r.PathPrefix("/grammar") .Subrouter()
	{
		grammar.HandleFunc("/like", likeGrammar).Methods("POST")
		grammar.HandleFunc("/like", unlikeGrammar).Methods("DELETE")
	}

	kanji := r.PathPrefix("/kanji") .Subrouter()
	{
		kanji.HandleFunc("/like", likeKanji).Methods("POST")
		kanji.HandleFunc("/like", unlikeKanji).Methods("DELETE")
	}

	discussion := r.PathPrefix("/discussion") .Subrouter()
	{
		discussion.HandleFunc("/list", getThreads).Methods("GET")
		discussion.HandleFunc("/single", getThread).Methods("GET")
		discussion.HandleFunc("/single", removeThread).Methods("DELETE")
		discussion.HandleFunc("/publish", publishThread).Methods("POST")
		discussion.HandleFunc("/reply", getReplies).Methods("GET")
		discussion.HandleFunc("/reply", reply).Methods("POST")
		discussion.HandleFunc("/reply", removeReply).Methods("DELETE")
		discussion.HandleFunc("/upvotethread", upvoteThread).Methods("POST")
		discussion.HandleFunc("/downvotethread", downvoteThread).Methods("POST")
		discussion.HandleFunc("/upvotereply", upvoteReply).Methods("POST")
		discussion.HandleFunc("/downvotereply", downvoteReply).Methods("POST")
	}

	search := r.PathPrefix("/search") .Subrouter()
	{
		search.HandleFunc("/list", getResults).Methods("GET")
	}
	
	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)
	
	http.ListenAndServe(":3000", handler)
}