package profile

import (
	"github.com/google/uuid"

	"be/internal/lib"
	"be/internal/model"
)

type Employee struct {
	PhoneNumberMobileFromPortal   *string    `json:"phone_number_mobile_from_portal"`
	PhoneNumberInternalFromPortal *string    `json:"phone_number_internal_from_portal"`
	ActiveDirectoryLogin          *string    `json:"active_directory_login"`
	Email                         *string    `json:"email"`
	JobTitleAdm                   *string    `json:"job_title_adm"`
	JobTitleBuh                   *string    `json:"job_title_buh"`
	JobTitleFact                  *string    `json:"job_title_fact"`
	LastName                      *string    `json:"last_name"`
	FirstName                     *string    `json:"first_name"`
	PatronymicName                *string    `json:"patronymic_name"`
	PortalCode                    *string    `json:"portal_code"`
	FIOLeadFunc                   *string    `json:"fio_lead_func"`
	PortalCodeLeadFunc            *string    `json:"portal_code_lead_func"`
	FIOLeadADM                    *string    `json:"fio_lead_adm"`
	PortalCodeLeadADM             *string    `json:"portal_code_lead_adm"`
	SBE                           *string    `json:"sbe"`
	Region                        *string    `json:"region"`
	Place                         *string    `json:"place"`
	PhotoFileName                 *string    `json:"photo_file_name"`
	PhotoSmallFileName            *string    `json:"photo_small_file_name"`
	BirthDate                     *string    `json:"birth_date"`
	Gender                        *string    `json:"gender"`
	Position                      *string    `json:"position"`
	FrcDtcCode                    *string    `json:"frc_dtc_code"`
	FrcName                       *string    `json:"frc_name"`
	CountryIsoCde                 *string    `json:"country_iso_cde"`
	CountryName                   *string    `json:"country_name"`
	City                          *string    `json:"city"`
	FiasCode                      *string    `json:"fias_code"`
	LeadEks                       *string    `json:"lead_eks"`
	AlternateEmployeeEks          *string    `json:"alternate_employee_eks"`
	AlternateEmployeeDate         *string    `json:"alternate_employee_date"`
	JobTitleAdmID                 *string    `json:"job_title_adm_id"`
	JobTitleBuhID                 *string    `json:"job_title_buh_id"`
	JobTitleFactID                *string    `json:"job_title_fact_id"`
	SbeID                         *string    `json:"sbe_id"`
	RegionID                      *string    `json:"region_id"`
	PlaceID                       *string    `json:"place_id"`
	LeadFio                       *string    `json:"lead_fio"`
	AlternateEmployeeFio          *string    `json:"alternate_employee_fio"`
	BookingIt                      *string    `json:"booking_it"`
	Organization                  *string    `json:"organization"`
	SubdivisionPays               *string    `json:"subdivision_pays"`
	Ekk                           *string    `json:"ekk"`
	EkkName                       *string    `json:"ekk_name"`
	ProfileID                     *uuid.UUID `json:"profile_id"`
}

func (e Employee) Cast() model.Employee {
	employee := model.Employee{
		PortalCode: lib.MustParseUint64(e.PortalCode),
	}

	if e.FirstName != nil {
		employee.FirstName = e.FirstName
	}

	if e.PatronymicName != nil {
		employee.MiddleName = e.PatronymicName
	}

	if e.LastName != nil {
		employee.LastName = e.LastName
	}

	if e.PhotoSmallFileName != nil {
		employee.Avatar = e.PhotoSmallFileName
	}

	if e.Email != nil {
		employee.Email = e.Email
	}

	if e.PhoneNumberMobileFromPortal != nil {
		employee.Phone = e.PhoneNumberMobileFromPortal
	}

	if e.JobTitleFact != nil {
		employee.Position = e.JobTitleFact
	}

	return employee
}

type EmployeeList []Employee

func (e EmployeeList) Cast() []model.Employee {
	res := make([]model.Employee, 0, len(e))

	for _, employee := range e {
		res = append(res, employee.Cast())
	}

	return res
}

type employeeListResponse struct {
	Data []Employee `json:"data"`
}
