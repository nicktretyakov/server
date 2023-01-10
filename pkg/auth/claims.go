package auth

import (
	"strconv"

	"github.com/pascaldekloe/jwt"

	"be/internal/model"
)

func claimsFromProfileToken(rawClaims jwt.Claims) profileClaims {
	phone, _ := rawClaims.Set["phone"].(string)
	email, _ := rawClaims.Set["email"].(string)
	employeeFirstName, _ := rawClaims.Set["employee_first_name"].(string)
	employeeMiddleName, _ := rawClaims.Set["employee_patronymic_name"].(string)
	employeeLastName, _ := rawClaims.Set["employee_last_name"].(string)
	employeeAvatar, _ := rawClaims.Set["employee_photo_small_file_name"].(string)
	employeeEmail, _ := rawClaims.Set["employee_email"].(string)
	employeePhone, _ := rawClaims.Set["employee_phone_number_mobile_from_portal"].(string)
	employeePosition, _ := rawClaims.Set["employee_job_title_fact"].(string)

	var portalCode uint64

	if portalCodeStr, ok := rawClaims.Set["portal_code"].(string); ok {
		portalCode, _ = strconv.ParseUint(portalCodeStr, 10, 64)
	}

	return profileClaims{
		profileID:          rawClaims.Subject,
		phone:              phone,
		email:              email,
		employeePortalCode: portalCode,
		employeeFirstName:  employeeFirstName,
		employeeMiddleName: employeeMiddleName,
		employeeLastName:   employeeLastName,
		employeeAvatar:     employeeAvatar,
		employeeEmail:      employeeEmail,
		employeePhone:      employeePhone,
		employeePosition:   employeePosition,
	}
}

type profileClaims struct {
	profileID          string
	phone              string
	email              string
	employeePortalCode uint64
	employeeFirstName  string
	employeeMiddleName string
	employeeLastName   string
	employeeAvatar     string
	employeeEmail      string
	employeePhone      string
	employeePosition   string
}

func (c *profileClaims) asUser() model.User {
	return model.User{
		ProfileID: c.profileID,
		Email:     c.email,
		Phone:     c.phone,
		Role:      model.Regular,
		Employee: model.Employee{
			FirstName:  &c.employeeFirstName,
			MiddleName: &c.employeeMiddleName,
			LastName:   &c.employeeLastName,
			Avatar:     &c.employeeAvatar,
			Email:      &c.employeeEmail,
			Phone:      &c.employeePhone,
			Position:   &c.employeePosition,
			PortalCode: c.employeePortalCode,
		},
	}
}
