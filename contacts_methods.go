package provisionclient

import "encoding/json"

type ContactsMethods struct {
	Client *Client
}

func (contacts *ContactsMethods) GetRoles() ([]ContactsRole, error) {
	body, err := contacts.Client.doRequest("GET", "/contacts/roles", nil)
	if err != nil {
		return nil, err
	}

	roles_ret := []ContactsRole{}
	err = json.Unmarshal(body, &roles_ret)
	if err != nil {
		return nil, err
	}

	return roles_ret, nil
}

func (contacts *ContactsMethods) GetContactByID(contactId string) (*ContactsContact, error) {
	body, err := contacts.Client.doRequest("GET", "/contacts/"+contactId, nil)
	if err != nil {
		return nil, err
	}

	resources_ret_json := &ContactsContact{}

	err = json.Unmarshal(body, resources_ret_json)
	if err != nil {
		return nil, err
	}
	return resources_ret_json, nil
}
