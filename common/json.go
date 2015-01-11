package common

type UserJson struct {
	Username string   `json:"username"`
	Orgs     []string `bson:"orgs" omitempty`
	Trees    []string `bson:"trees" omitempty`
}

type MissionJson struct {
	Mission string `json:"mission"`
	TreeId  string `json:"treeId,omitempty"`
}

type OrganizationJson struct {
	Organization string `json:"organization"`
	UserId       string `json:"userId" bson:"_id,omitempty"`
	UserName     string `json:"username"`
	UserRole     string `json:"role"`
}

type TreeJson struct {
	TreeName  string `json:"treeName"`
	Timeframe string `json:"timeframe"`
	UserName  string `json:"username"`
	UserId    string `json:"userId" bson:"_id,omitempty"`
	UserRole  string `json:"role"`
}

type MembersJson struct {
	UpdateTree bool     `json:"updateTree,omitempty"`
	TreeName   string   `json:"treeName,omitempty"`
	TreeId     string   `json:"treeId,omitempty"`
	Members    []Member `json:"members"`
}

type MembersJsonDelete struct {
	UpdateTree bool     `json:"updateTree,omitempty"`
	TreeId     string   `json:"treeId,omitempty"`
	Members    []string `json:"members"`
}

type ObjectiveJson struct {
	Id        string   `json:"id"`
	TreeId    string   `json:"treeId,omitempty"`
	Name      string   `json:"name"`
	Body      string   `json:"body"`
	Completed bool     `json:"completed"`
	Members   []Member `json:"members"`
}

type KeyResultJson struct {
	Id        string   `json:"id"`
	TreeId    string   `json:"treeId,omitempty"`
	Name      string   `json:"name"`
	Body      string   `json:"body"`
	Completed bool     `json:"completed"`
	Members   []Member `json:"members"`
	Priority  string   `json:"priority"`
}
