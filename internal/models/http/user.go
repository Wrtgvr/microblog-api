package modelshttp

type UserResponse struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=6,max=30"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

type UpdateUserRequest struct {
	Username *string `json:"username,omitempty" binding:"omitempty,alphanum"`
	Password *string `json:"password,omitempty" binding:"omitempty"`
}
