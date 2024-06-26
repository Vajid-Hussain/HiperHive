package requestmodel_auth_svc

type UserSignup struct {
	UserName        string `json:"UserName"  validate:"required"`
	Name            string `json:"Name" validate:"min=1,alphanum"`
	Email           string `json:"Email" validate:"required,email"`
	Password        string `json:"Password" validate:"min=5"`
	ConfirmPassword string `json:"ConfirmPassword"`
	// ConfirmPassword string `json:"ConfirmPassword" validate:"eqfield=Password,required"`
}

type UserLogin struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=5"`
}

type UserProfileStatus struct {
	UserID   string  `json:"-"`
	Status   string  `json:"Status"`
	Duration float32 `json:"Duration"`
}

type UserProfileDescription struct {
	UserID      string `json:"-"`
	Description string `json:"Description"`
}

type UserIDReq struct {
	UserID string `json:"UserID" param:"UserID"`
}

type EmailReq struct {
	Email string `json:"Email" validate:"email,required"`
}

type ForgotPassword struct {
	Password        string `json:"Password" validate:"min=5"`
	ConfirmPassword string `json:"ConfirmPassword"`
	Otp             string `json:"Otp" validate:"len=4"`
}

type DeleteUserProfilePhotoType struct {
	Types string `json:"Type" validate:"required"`
}

// type Pagination struct {
// 	Limit  string `json:"Limit"`
// 	OffSet string `json:"Offset"`
// }

type RefreshToken struct{
	RefreshToken string `json:"RefreshToken"`
}

// type AccessToken struct{
// 	AccessToken string `json:"AccessToken"`
// }