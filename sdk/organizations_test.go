package meraki

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestHeadLicenseID_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput HeadLicenseID
		shouldFail     bool
	}{
		{`"12345"`, "12345", false},
		{`12345`, "12345", false},
		{`12345.0`, "12345.0", false},
		{`null`, "", false},
		{`true`, "", true},
		{`{}`, "", true},
		// testing some bigger numbers
		{`9192581925819285192581295812958125981`, "9192581925819285192581295812958125981", false},
		{`"9192581925819285192581295812958125981"`, "9192581925819285192581295812958125981", false},
	}

	for _, tt := range tests {
		var id HeadLicenseID
		err := json.Unmarshal([]byte(tt.input), &id)
		if tt.shouldFail && err == nil {
			t.Errorf("expected failure for input %s but got success", tt.input)
		} else if !tt.shouldFail && (err != nil || id != tt.expectedOutput) {
			t.Errorf("unexpected result for input %s: got %v, err %v, expected %v", tt.input, id, err, tt.expectedOutput)
		}
	}
}

type NoopRateLimiter struct{}

func (n *NoopRateLimiter) Take() time.Time {
	return time.Time{}
}

func TestOrganizationsService_GetOrganizationLicenses_HeadLicenseIDHandling(t *testing.T) {
	tests := []struct {
		name                  string
		mockResponseStatus    int
		mockResponseBody      string
		organizationID        string
		queryParams           *GetOrganizationLicensesQueryParams
		expectedHeadLicenseID HeadLicenseID
		expectError           bool
		checkErrorContains    string
	}{
		{
			name:                  "HeadLicenseId as string",
			mockResponseStatus:    http.StatusOK,
			mockResponseBody:      `[{"id": "L_123", "headLicenseId": "9591829581295182591859158195"}]`,
			organizationID:        "org1",
			queryParams:           &GetOrganizationLicensesQueryParams{},
			expectedHeadLicenseID: "9591829581295182591859158195",
			expectError:           false,
		},
		{
			name:                  "HeadLicenseId as number",
			mockResponseStatus:    http.StatusOK,
			mockResponseBody:      `[{"id": "L_456", "headLicenseId": 987659185291859128519258195}]`,
			organizationID:        "org2",
			queryParams:           &GetOrganizationLicensesQueryParams{},
			expectedHeadLicenseID: "987659185291859128519258195",
			expectError:           false,
		},
		{
			name:                  "HeadLicenseId as number with decimal",
			mockResponseStatus:    http.StatusOK,
			mockResponseBody:      `[{"id": "L_789", "headLicenseId": 123.45}]`,
			organizationID:        "org3",
			queryParams:           &GetOrganizationLicensesQueryParams{},
			expectedHeadLicenseID: "123.45",
			expectError:           false,
		},
		{
			name:                  "HeadLicenseId is null",
			mockResponseStatus:    http.StatusOK,
			mockResponseBody:      `[{"id": "L_ABC", "headLicenseId": null}]`,
			organizationID:        "org4",
			queryParams:           &GetOrganizationLicensesQueryParams{},
			expectedHeadLicenseID: "",
			expectError:           false,
		},
		{
			name:                  "HeadLicenseId is absent",
			mockResponseStatus:    http.StatusOK,
			mockResponseBody:      `[{"id": "L_DEF"}]`,
			organizationID:        "org5",
			queryParams:           &GetOrganizationLicensesQueryParams{},
			expectedHeadLicenseID: "",
			expectError:           false,
		},
		{
			name:               "HeadLicenseId boolean",
			mockResponseStatus: http.StatusOK,
			mockResponseBody:   `[{"id": "L_JKL", "headLicenseId": true}]`,
			organizationID:     "org8",
			queryParams:        &GetOrganizationLicensesQueryParams{},
			expectError:        true,
			checkErrorContains: "HeadLicenseID: unsupported type: true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectedPath := fmt.Sprintf("/api/v1/organizations/%s/licenses", tt.organizationID)
				if r.URL.Path != expectedPath {
					t.Errorf("handler expected path %s, got %s", expectedPath, r.URL.Path)
					http.Error(w, "bad path", http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.mockResponseStatus)
				_, _ = w.Write([]byte(tt.mockResponseBody))
			}))
			defer server.Close()

			client := resty.New()
			client.SetBaseURL(server.URL)

			service := &OrganizationsService{
				client:      client,
				ratelimiter: &NoopRateLimiter{},
			}

			licenses, resp, err := service.GetOrganizationLicenses(tt.organizationID, tt.queryParams)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil. Response status: %s, body: %s", resp.Status(), resp.String())
				}
				if tt.checkErrorContains != "" && !strings.Contains(err.Error(), tt.checkErrorContains) {
					t.Errorf("expected error to contain '%s', but got '%s'", tt.checkErrorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Fatalf("did not expect an error but got: %v. Response status: %s, body: %s", err, resp.Status(), resp.String())
				}
				if licenses == nil {
					t.Fatalf("expected licenses but got nil")
				}
				if len(*licenses) == 0 && tt.mockResponseBody != "[]" && tt.mockResponseBody != "" {
					t.Fatalf("expected at least one license item, but got zero")
				}
				if len(*licenses) > 0 {
					item := (*licenses)[0]
					if item.HeadLicenseID != tt.expectedHeadLicenseID {
						t.Errorf("expected HeadLicenseID '%s', got '%s'", tt.expectedHeadLicenseID, item.HeadLicenseID)
					}
				}
			}
		})
	}
}
