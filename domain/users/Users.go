package domain

import (
	"time"
)

type User struct {
	Email       string    `json:"email"`
	IdNumber    string    `json:"idNumber"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type UserAddress struct {
	UserId        string `json:"userId"`
	AddressTypeId string `json:"addressTypeId"`
	Address       string `json:"address"`
	PostalCode    string `json:"postalCode"`
}

type UserApplication struct {
	UserId        string    `json:"userId"`
	ApplicationId string    `json:"applicationId"`
	DateTime      time.Time `json:"dateTime"`
}

type UserApplicationCourse struct {
	UserId        string `json:"userId"`
	ApplicationId string `json:"applicationId"`
	CourseId      string `json:"courseId"`
}

type UserApplicationInstitution struct {
	UserId        string `json:"userId"`
	ApplicationId string `json:"applicationId"`
	InstitutionId string `json:"institutionId"`
}

type UserCommunication struct {
	CommunicationId string `json:"communicationId"`
	Description     string `json:"description"`
}

type UserContacts struct {
	UserId        string `json:"userId"`
	ContactTypeId string `json:"contactTypeId"`
	Contact       string `json:"contact"`
}

type UserCourentLocation struct {
	UserId   string `json:"userId"`
	Province string `json:"province"`
	District string `json:"district"`
	TownCode string `json:"townCode"`
}

type UserDemographics struct {
	UserId   string `json:"userId"`
	TitleId  string `json:"titleId"`
	GenderId string `json:"genderId"`
	RaceId   string `json:"raceId"`
}

type UserDocument struct {
	UserId     string `json:"userId"`
	DocumentId string `json:"documentId"`
}

type UserPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRelative struct {
	UserId       string `json:"userId"`
	Name         string `json:"name"`
	Cellphone    string `json:"cellphone"`
	Email        string `json:"email"`
	Relationship string `json:"relationship"`
}

type UserResults struct {
	UserResultsId string `json:"userResultsId"`
	Description   string `json:"description"`
}

type UserRole struct {
	UserId string `json:"userId"`
	RoleId string `json:"roleId"`
}

type UserTertiaryCourse struct {
	UserId        string `json:"userId"`
	ApplicationId string `json:"applicationId"`
	CourseId      string `json:"courseId"`
}

type UserTertiaryInstitution struct {
	UserId        string  `json:"userId"`
	ApplicationId string  `json:"applicationId"`
	InstitutionId string  `json:"institutionId"`
	DebtAmount    float64 `json:"debtAmount"`
}

type UserTertiarySubject struct {
	UserId        string  `json:"userId"`
	ApplicationId string  `json:"applicationId"`
	SubjectId     string  `json:"subjectId"`
	SubjectMark   float64 `json:"subjectMark"`
}

type UserTown struct {
	UserId     string `json:"userId"`
	LocationId string `json:"locationId"`
}

type UserMatricInstitution struct {
	UserId        string `json:"userId"`
	InstitutionId string `json:"institutionId"`
}

type UserMatricSubject struct {
	UserId      string  `json:"userId"`
	SubjectId   string  `json:"subjectId"`
	SubjectMark float64 `json:"subjectMark"`
}
