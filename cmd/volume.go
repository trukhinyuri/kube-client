package cmd

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/cherry"
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	resourceVolumeRootPath   = "/volume"
	resourceVolumePath       = resourceVolumeRootPath + "/{volume}"
	resourceVolumeNamePath   = resourceVolumePath + "/name"
	resourceVolumeAccessPath = resourceVolumePath + "/access"
)

// DeleteVolume -- deletes Volume with provided volume name
func (client *Client) DeleteVolume(volumeName string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		Delete(client.ResourceAddr + resourceVolumePath)
	return mapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted)
}

// GetVolume -- return User Volume by name,
// consumes optional userID param
func (client *Client) GetVolume(volumeName string, userID *string) (model.Volume, error) {
	req := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetResult(model.Volume{})
	if userID != nil {
		req.SetQueryParam("user-id", *userID)
	}
	resp, err := req.Get(client.ResourceAddr + resourceVolumePath)
	if err = mapErrors(resp, err, http.StatusOK); err != nil {
		return model.Volume{}, err
	}
	return *resp.Result().(*model.Volume), nil
}

// GetVolumeList -- get list of volumes,
// consumes optional user ID and filter parameters.
// Returns new_access_level as access if user role = user.
// Should have filters: not deleted, limited, not limited, owner, not owner.
func (client *Client) GetVolumeList(userID, filter *string) ([]model.Volume, error) {
	req := client.Request.
		SetResult([]model.Volume{}).
		SetError(cherry.Err{})
	if userID != nil {
		req.SetQueryParam("user-id", *userID)
	}
	if filter != nil {
		req.SetQueryParam("user-id", *filter)
	}
	resp, err := req.Get(client.ResourceAddr + resourceVolumeRootPath)
	if err = mapErrors(resp, err, http.StatusOK); err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.Volume), nil
}

//RenameVolume -- change volume name
func (client *Client) RenameVolume(volumeName, newName string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetBody(model.ResourceUpdateName{Label: newName}).
		Put(client.ResourceAddr + resourceVolumeNamePath)
	return mapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted)
}

// SetAccess -- sets User Volume access
func (client *Client) SetAccess(volumeName string, accessData model.ResourceUpdateUserAccess) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetBody(accessData).
		Post(client.ResourceAddr + resourceVolumeAccessPath)
	return mapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted)
}

// DeleteAccess -- deletes user Volume access
func (client *Client) DeleteAccess(volumeName, username string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetBody(model.ResourceUpdateUserAccess{
			Username: username,
		}).
		Delete(client.ResourceAddr + resourceVolumeAccessPath)
	return mapErrors(resp, err,
		http.StatusOK,
		http.StatusAccepted)
}