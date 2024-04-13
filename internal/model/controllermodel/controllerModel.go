package controllermodel

type Banner struct {
	TagID     *[]int                  `json:"tag_ids,omitempty"`
	FeatureID *int                    `json:"feature_id,omitempty"`
	Content   *map[string]interface{} `json:"content,omitempty"`
	IsActive  *bool                   `json:"is_active,omitempty"`
}
