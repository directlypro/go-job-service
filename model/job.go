package model

type LinkedinJobApiResponse struct {
	Id                                  string         `json:"id"`
	DatePosted                          string         `json:"date_posted"`
	DateCreated                         string         `json:"date_created"`
	Title                               string         `json:"title"`
	Organization                        string         `json:"organization"`
	OrganizationUrl                     string         `json:"organization_url"`
	DateValidthrough                    string         `json:"date_validthrough"`
	LocationsRaw                        []RawLocations `json:"locations_raw"`
	LocationType                        interface{}    `json:"location_type"`
	LocationRequirementsRaw             interface{}    `json:"location_requirements_raw"`
	SalaryRaw                           interface{}    `json:"salary_raw"`
	EmploymentType                      []string       `json:"employment_type"`
	Url                                 string         `json:"url"`
	SourceType                          string         `json:"source_type"`
	Source                              string         `json:"source"`
	SourceDomain                        string         `json:"source_domain"`
	OrganizationLogo                    string         `json:"organization_logo"`
	CitiesDerived                       []string       `json:"cities_derived"`
	RegionsDerived                      []string       `json:"regions_derived"`
	CountriesDerived                    []string       `json:"countries_derived"`
	LocationsDerived                    []string       `json:"locations_derived"`
	TimezonesDerived                    []string       `json:"timezones_derived"`
	LatsDerived                         []float64      `json:"lats_derived"`
	LngsDerived                         []float64      `json:"lngs_derived"`
	RemoteDerived                       bool           `json:"remote_derived"`
	RecruiterName                       string         `json:"recruiter_name"`
	RecruiterTitle                      string         `json:"recruiter_title"`
	RecruiterUrl                        string         `json:"recruiter_url"`
	LinkedinOrgEmployees                int            `json:"linkedin_org_employees"`
	LinkedinOrgUrl                      string         `json:"linkedin_org_url"`
	LinkedinOrgSize                     string         `json:"linkedin_org_size"`
	LinkedinOrgSlogan                   string         `json:"linkedin_org_slogan"`
	LinkedinOrgIndustry                 string         `json:"linkedin_org_industry"`
	LinkedinOrgFollowers                int            `json:"linkedin_org_followers"`
	LinkedinOrgHeadquarters             string         `json:"linkedin_org_headquarters"`
	LinkedinOrgType                     string         `json:"linkedin_org_type"`
	LinkedinOrgFoundeddate              string         `json:"linkedin_org_foundeddate"`
	LinkedinOrgSpecialties              []string       `json:"linkedin_org_specialties"`
	LinkedinOrgLocations                []string       `json:"linkedin_org_locations"`
	LinkedinOrgDescription              string         `json:"linkedin_org_description"`
	LinkedinOrgRecruitmentAgencyDerived bool           `json:"linkedin_org_recruitment_agency_derived"`
	Seniority                           string         `json:"seniority"`
	Directapply                         bool           `json:"directapply"`
	LinkedinOrgSlug                     string         `json:"linkedin_org_slug"`
}

type RawLocations struct {
	Type      string     `json:"@type"`
	Address   RawAddress `json:"address"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
}

type RawAddress struct {
	Type            string      `json:"@type"`
	AddressCountry  string      `json:"addressCountry"`
	AddressLocality string      `json:"addressLocality"`
	AddressRegion   interface{} `json:"addressRegion"`
	StreetAddress   interface{} `json:"streetAddress"`
}
