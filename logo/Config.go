package logo

var (
	url      string
	password string

	securityHint string
)

func LoadConfig(urlParm string, passwordParm string) error {
	url = urlParm + "/AJAX"
	password = passwordParm
	return generateSecurityHint()
}
