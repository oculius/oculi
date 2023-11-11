package authz

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type (
	Permission struct {
		Resource string
		Action   string
	}

	Permissions []Permission
)

const PermissionSeparator = ":"

func (p *Permission) UnmarshalJSON(b []byte) error {
	unquote, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	if strings.Count(unquote, PermissionSeparator) != 1 {
		return errors.New("invalid permission format, expected: 'action" + PermissionSeparator + "resource'")
	}
	split := strings.SplitN(unquote, PermissionSeparator, 2)
	p.Action = strings.TrimSpace(strings.ToLower(split[0]))
	p.Resource = strings.TrimSpace(strings.ToLower(split[1]))
	return nil
}

func (p *Permission) MarshalJSON() ([]byte, error) {
	if strings.Contains(p.Resource, PermissionSeparator) || strings.Contains(p.Action, PermissionSeparator) {
		return nil, errors.New("resource/action contains invalid character '" + PermissionSeparator + "'")
	}
	return []byte(strconv.Quote(
		fmt.Sprintf("%s%s%s",
			strings.ToLower(p.Action),
			PermissionSeparator,
			strings.ToLower(p.Resource),
		)),
	), nil
}

func NewPermissions(matrix [][]string) Permissions {
	if len(matrix) == 0 {
		return nil
	}
	result := make([]Permission, len(matrix))
	for i, each := range matrix {
		if len(each) < 2 {
			continue
		}
		result[i] = Permission{
			Resource: strings.Replace(each[0], ResourcePrefix, "", 1),
			Action:   strings.Replace(each[1], ActionPrefix, "", 1),
		}
	}
	return result
}
