package model

// struct from https://online.moysklad.ru/api/remap/1.2/entity/employee
type Employees struct {
	// Context and Meta not so important for my task now
	Context struct {
		Employee struct {
			Meta struct {
				Href         string `json:"href"`
				MetadataHref string `json:"metadataHref"`
				Type         string `json:"type"`
				MediaType    string `json:"mediaType"`
			} `json:"meta"`
		} `json:"employee"`
	} `json:"context"`
	Meta struct {
		Href      string `json:"href"`
		Type      string `json:"type"`
		MediaType string `json:"mediaType"`
		Size      int    `json:"size"`
		Limit     int    `json:"limit"`
		Offset    int    `json:"offset"`
	} `json:"meta"`

	// Rows is array of employees' data
	Rows `json:"rows"`
}

// Rows is array of employees' data
type Rows []struct {
	Meta struct {
		Href         string `json:"href"`
		MetadataHref string `json:"metadataHref"`
		Type         string `json:"type"`
		MediaType    string `json:"mediaType"`
		UUIDHref     string `json:"uuidHref"`
	} `json:"meta"`
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Owner     struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
			UUIDHref     string `json:"uuidHref"`
		} `json:"meta"`
	} `json:"owner"`
	Shared bool `json:"shared"`
	Group  struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
		} `json:"meta"`
	} `json:"group"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	Created      string `json:"created"`
	UID          string `json:"uid"`
	Email        string `json:"email"`
	LastName     string `json:"lastName"`
	FullName     string `json:"fullName"`
	ShortFio     string `json:"shortFio"`
	Cashiers     []struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
	} `json:"cashiers"`
}

type Employee struct {
	Meta struct {
		Href         string `json:"href"`
		MetadataHref string `json:"metadataHref"`
		Type         string `json:"type"`
		MediaType    string `json:"mediaType"`
	} `json:"meta"`
	ID           string `json:"id"`
	AccountID    string `json:"accountId"`
	Updated      string `json:"updated"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
	Archived     bool   `json:"archived"`
	UID          string `json:"uid"`
	Email        string `json:"email"`
	LastName     string `json:"lastName"`
	FullName     string `json:"fullName"`
	ShortFio     string `json:"shortFio"`
	Cashiers     []struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
	} `json:"cashiers"`
	RetailStore struct {
		Meta struct {
			Href         string `json:"href"`
			MetadataHref string `json:"metadataHref"`
			Type         string `json:"type"`
			MediaType    string `json:"mediaType"`
		} `json:"meta"`
	} `json:"retailStore"`
	Attributes []struct {
		Meta struct {
			Href      string `json:"href"`
			Type      string `json:"type"`
			MediaType string `json:"mediaType"`
		} `json:"meta"`
		ID    string `json:"id"`
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"attributes"`
	Inn      string `json:"inn"`
	Position string `json:"position"`
}
