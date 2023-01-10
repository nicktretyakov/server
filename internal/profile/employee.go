package profile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/imroc/req"
	"github.com/pkg/errors"

	"be/internal/lib"
)

const employeeQuerySearchURL = "/api/v1/employees"

func (s service) FindEmployees(ctx context.Context, query string, limit, offset uint64) (EmployeeList, error) {
	params := req.QueryParam{
		"q":      query,
		"limit":  limit,
		"offset": offset,
	}

	resp, err := s.r.Get(s.url(employeeQuerySearchURL), ctx, s.headers(), params)
	if err != nil {
		s.logger.Err(err).Msg("profile error")
		return nil, err
	}

	if statusCode := resp.Response().StatusCode; statusCode != http.StatusOK { //nolint:bodyclose
		s.logger.Err(err).
			Int("status_code", statusCode).
			Str("response", resp.String()).
			Msg("profile error")

		return nil, errors.New("profile error")
	}

	var employeesResponse employeeListResponse

	if err = resp.ToJSON(&employeesResponse); err != nil {
		return nil, err
	}

	employeesResponse.Data = checkPortalCode(employeesResponse.Data)

	return employeesResponse.Data, nil
}

func (s service) FindEmployeeByPortalCode(ctx context.Context, portalCode uint64) (*Employee, error) {
	url := s.portalCodeURL(portalCode)

	resp, err := s.get(ctx, url)
	if err != nil {
		return nil, err
	}

	rawResponse := resp.Response()
	defer rawResponse.Body.Close()

	if respCode := rawResponse.StatusCode; respCode != http.StatusOK {
		errMsg, _ := resp.ToString()
		s.logger.Error().
			Uint64("portal_code", portalCode).
			Int("profile_status_code", respCode).
			Msgf("profile error, resp=%s", errMsg)

		return nil, errors.New("profile error")
	}

	var employee Employee
	if err = resp.ToJSON(&employee); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s service) portalCodeURL(portalCode uint64) string {
	return fmt.Sprintf("/api/v1/employees/portal_code/%d", portalCode)
}

func checkPortalCode(employees []Employee) []Employee {
	employeeWithPortalCode := make([]Employee, 0, len(employees))

	for _, employee := range employees {
		if lib.MustParseUint64(employee.PortalCode) != 0 {
			employeeWithPortalCode = append(employeeWithPortalCode, employee)
		}
	}

	return employeeWithPortalCode
}
