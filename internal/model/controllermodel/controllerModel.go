package controllermodel

import "time"

type Banner struct {
	ID        int64                  `json:"id"`
	TagID     *[]int64               `json:"tag_ids,omitempty"`
	FeatureID *int64                 `json:"feature_id,omitempty"`
	Content   map[string]interface{} `json:"content,omitempty"`
	IsActive  *bool                  `json:"is_active,omitempty"`
	CreatedAt time.Time              `json:"created_at,omitempty"`
	UpdatedAt time.Time              `json:"updated_at,omitempty"`
}
