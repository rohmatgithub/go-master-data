package regional_dto

type SubDistrictRequest struct {
}

type SubDistrictResponse struct {
}

type SubDistrictListResponse struct {
	ID       int64  `json:"id"`
	ParentID int64  `json:"parent_id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}
