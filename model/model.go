package model

type ProblemData struct {
	// IssueId          string        `bson:"id"`
	IssueTitle         string             `json:"issuetitle"`
	IssueType          string             `json:"issuetype"`
	IssueDescription   string             `json:"issuedescription"`
	IssueLevel         string             `json:"issuelevel"`
	IssueLocation      IssueLocation      `json:"issuelocation"`
	IssueProgress      string             `json:"issueprogress"`
	IssueDate          string             `json:"issuedate"`
	IssueMedia         []string           `json:"issuemedia"`
	IssueComments      []Comment          `json:"issuecomments"`
	IssueRaiserDetails IssueRaiserDetails `json:"issueraiserdetails"`
}

type IssueRaiserDetails struct {
	IssueRaiserName         string `json:"issueraisername"`
	IssueRaiserId           string `json:"issueraiserid"`
	IssueRaiserMail         string `json:"issueraisermail"`
	IssueRaiserPhone        string `json:"issueraiserphone"`
	IssueRaiserProfilePhoto string `json:"issueraiserprofilephoto"`
}

type Comment struct {
	Body        string `json:"body"`
	UserName    string `json:"username"`
	CommentType string `json:"commenttype"`
	CommentTime int64  `json:"commenttime"`
}

type UserData struct {
	// UserId           string `bson:"id"`
	UserName         string `json:"username"`
	UserEmail        string `json:"useremail"`
	Gender           string `json:"gender"`
	UserPhone        string `json:"userphone"`
	UserProfilePhoto string `json:"userprofilephoto"`
	UserLocation     string `json:"userlocation"`
	UserAge          string `json:"userage"`
	UserVerified     bool   `json:"userverified"`
	UserIdProof      string `json:"useridproof"`
	UserRole     string `json:"userrole"`
}

type Phone struct {
	Phone string `json:"userphone"`
}

type IssueLocation struct {
	LocalAddress string `json:"localaddress"`
	District     string `json:"district"`
	SubDivision  string `json:"subdivision"`
}
