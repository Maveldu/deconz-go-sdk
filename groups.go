package sdk

import (
	"fmt"
)

func (c Client) GetGroups() (map[string]Group, error) {
	result := map[string]Group{}
	err := c.Get("/groups", nil, &result)
	if err  != nil {
		return nil, err
	}
	return result, nil
}

func (c Client) GetGroup(id string) (Group, error) {
	result := Group{}
	err := c.Get(fmt.Sprintf("/groups/%s", id), nil, &result)
	if err  != nil {
		return Group{}, err
	}
	return result, nil
}

func (c Client) GetGroupScenes(groupID string) (map[string]GroupSceneDetailed, error) {
	result := map[string]GroupSceneDetailed{}
	err := c.Get(fmt.Sprintf("/groups/%s/scenes", groupID), nil, &result)
	if err  != nil {
		return nil, err
	}
	return result, nil
}

func (c Client) GetGroupScene(groupID, sceneID string) (Scene, error) {
	result := Scene{}
	err := c.Get(fmt.Sprintf("/groups/%s/scenes/%s", groupID, sceneID), nil, &result)
	if err  != nil {
		return Scene{}, err
	}
	return result, nil
}
