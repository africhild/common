package identity

import (
	"fmt"
	"log"
	"os"

	"github.com/africhild/common/pkg/request"
	"github.com/africhild/common/pkg/util"
)

func Authenticate(method string, url string, headers map[string]string) (bool, util.Object, util.Object, error) {
	baseUrl := os.Getenv("IDENTITY_BASE_URL")
	payload := util.Object{
		"method":  method,
		"url":     url,
		"headers": headers,
	}
	response, err := request.Post(fmt.Sprintf("%v/authenticate", baseUrl), payload, nil, headers, 0)
	if err != nil {
		log.Println(err)
		return false, nil, nil, err
	}

	if !response.Status {
		return false, nil, nil, nil
	}

	status, _ := response.Data.Data["status"].(bool)
	user, _ := response.Data.Data["user"].(util.Object)
	workspace, _ := response.Data.Data["workspace"].(util.Object)

	return status, user, workspace, nil
}

func Authorize(userId, workspaceId, permissionType string, permissions []string) (bool, error) {
	baseUrl := os.Getenv("IDENTITY_BASE_URL")
	payload := util.Object{
		"userId":      userId,
		"workspaceId": workspaceId,
		"permissions": permissions,
		"type":        permissionType,
	}
	response, err := request.Post(fmt.Sprintf("%v/authorize", baseUrl), payload, nil, nil, 0)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if !response.Status {
		return false, nil
	}

	status := response.Data.Data["status"].(bool)

	return status, nil
}

func GetWorkspace(sourceId string) (string, string, error) {
	baseUrl := os.Getenv("IDENTITY_BASE_URL")
	url := fmt.Sprintf("%v/workspace/%v/sourceId", baseUrl, sourceId)
	response, err := request.Get(url, nil, nil, 0)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	if !response.Status {
		return "", "", err
	}

	userId := response.Data.Data["userId"].(string)
	workspaceId := response.Data.Data["workspaceId"].(string)

	return userId, workspaceId, nil
}
