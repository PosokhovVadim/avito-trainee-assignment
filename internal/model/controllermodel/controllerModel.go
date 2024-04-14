package controllermodel

type Banner struct {
	TagID     *[]int64               `json:"tag_ids,omitempty"`
	FeatureID *int64                 `json:"feature_id,omitempty"`
	Content   map[string]interface{} `json:"content,omitempty"`
	IsActive  *bool                  `json:"is_active,omitempty"`
}
