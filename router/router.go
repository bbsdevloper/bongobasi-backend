package router

import (
	"net/http"

	authhandler "github.com/Prosecutor1x/citizen-connect-frontend/handlers/auth_handler"
	"github.com/Prosecutor1x/citizen-connect-frontend/handlers/issue_handler"
	"github.com/Prosecutor1x/citizen-connect-frontend/handlers/media_handler"
	"github.com/Prosecutor1x/citizen-connect-frontend/handlers/protected_handler"
	"github.com/Prosecutor1x/citizen-connect-frontend/handlers/user_handler"
	"github.com/Prosecutor1x/citizen-connect-frontend/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// issue routes
	router.HandleFunc("/api/addNewIssue", issue_handler.CreateIssueHandler).Methods("POST")
	router.HandleFunc("/api/fetchIssue", issue_handler.FetchAllIssueHandler).Methods("GET")
	router.HandleFunc("/api/deleteIssue/{id}", issue_handler.DeleteIssueHandler).Methods("DELETE")
	router.HandleFunc("/api/fetchIssue/{id}", issue_handler.FetchSingleIssueHandler).Methods("GET")
	router.HandleFunc("/api/fetchIssues/{id}", issue_handler.FetchAllUserIssueHandler).Methods("GET")
	router.HandleFunc("/api/updateIssue/{id}", issue_handler.UpdateIssueHandler).Methods("PUT")

	// auth routes
	router.HandleFunc("/api/sendOtp", authhandler.SendOtp).Methods("POST")
	router.HandleFunc("/api/verifyOtp", authhandler.VerifyOtp).Methods("POST")

	// user routes
	router.HandleFunc("/api/createUser", user_handler.CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/checkUser", user_handler.CheckUserExist).Methods("POST")
	router.HandleFunc("/api/getUser/{id}", user_handler.GetUser).Methods("GET")
	// router.Handle("/api/getUser/{id}", middleware.AuthMiddleware(http.HandlerFunc(user_handler.GetUser))).Methods("GET")

	// media routes
	router.HandleFunc("/api/uploadMedia", media_handler.UploadMediaToS3).Methods("POST")

	// protected routes
	router.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(protected_handler.ProtectedHandler))).Methods("GET")

	return router
}
