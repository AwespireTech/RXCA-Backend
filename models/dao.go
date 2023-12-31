package models

const (
	DAO_STATE_PENDING = iota
	DAO_STATE_APPROVED
	DAO_STATE_DENIED
)

type DAO struct {
	ID                 int    `json:"id" bson:"displayId,omitempty"`
	Address            string `json:"address" bson:"addr,omitempty"`
	Contract           string `json:"network" bson:"contract,omitempty"`
	Name               string `json:"name" bson:"name,omitempty"`
	Description        string `json:"description" bson:"description,omitempty"`
	Framework          string `json:"framework" bson:"framwork,omitempty"`
	MembersUri         string `json:"membersUri" bson:"membersUri,omitempty"`
	ProposalsUri       string `json:"proposalsUri" bson:"proposalsUri,omitempty"`
	IssuersUri         string `json:"issuersUri" bson:"issuersUri,omitempty"`
	ContractsRegUri    string `json:"contractsRegUri" bson:"contractsRegUri,omitempty"`
	ManagerAddress     string `json:"managerAddress" bson:"managerAddress,omitempty"`
	GovernanceDocument string `json:"governanceDocument" bson:"governanceDocument,omitempty"`
	State              int    `json:"state" bson:"state"`
	TokenId            int    `json:"tokenId" bson:"tokenId,omitempty"`
	Creator            string `json:"creator" bson:"creator,omitempty"`
}
type DAOFilter struct {
	Name    interface{} `bson:"name,omitempty"`
	State   interface{} `bson:"state,omitempty"`
	Creator string      `bson:"creator,omitempty"`
}
type DAOAddressFilter struct {
	Address string `bson:"addr,omitempty"`
}

type DAOsResponse struct {
	DAOs  []DAO `json:"daos"`
	Count int   `json:"count"`
}
type DAOVerifyRequest struct {
	Validate bool   `json:"validate"`
	TxHash   string `json:"opHash"`
}
type DAORevokeRequest struct {
	TxHash string `json:"opHash" binding:"required"`
}
type DAOid struct {
	ID int `json:"id" bson:"id,omitempty"`
}
type DAOExploreParams struct {
	Search  string `form:"search"`
	State   int    `form:"state"`
	Limit   int    `form:"limit"`
	Offset  int    `form:"offset"`
	Creator string `form:"creator"`
}
