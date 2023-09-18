package syspolicy

import "tailscale.com/util/winutil"

type windowsHandler struct{}

func (windowsHandler) ReadBool(key string) (bool, error) {
	return winutil.GetPolicyInteger(key, 0) != 0, nil
}

func (windowsHandler) ReadString(key string) (string, error) {
	const defaultValue = "{defaultvalue}"
	v := winutil.GetPolicyString(key, defaultValue)
	if v == defaultValue {
		return "", ErrNoSuchKey
	}
	return v, nil
}
