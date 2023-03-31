package mirai

// https://github.com/project-mirai/mirai-api-http/blob/master/docs/api/API.md#%E7%8A%B6%E6%80%81%E7%A0%81

type Code int

const (
	CodeOK             Code = 0
	CodeWrongVerifyKey Code = 1
	CodeBotNotExists   Code = 2
	CodeInvalidSession Code = 3
	CodeUnauthorized   Code = 4
	CodeTargetNotFound Code = 5
	CodeFileNotFound   Code = 6
	CodeNoPermission   Code = 10
	CodeBotMuted       Code = 20
	CodeMessageTooLong Code = 30
	CodeInvalidArgs    Code = 400
)

func (c Code) Equal(code Code) bool {
	return c == code
}

func (c Code) String() string {
	switch c {
	case CodeOK:
		return "OK"
	case CodeWrongVerifyKey:
		return "Wrong Verify Key"
	case CodeBotNotExists:
		return "Bot Not Exists"
	case CodeInvalidSession:
		return "Invalid Session"
	case CodeUnauthorized:
		return "Unauthorized"
	case CodeTargetNotFound:
		return "Target Not Found"
	case CodeFileNotFound:
		return "File Not Found"
	case CodeNoPermission:
		return "No Permission"
	case CodeBotMuted:
		return "Bot Muted"
	case CodeMessageTooLong:
		return "Message Too Long"
	case CodeInvalidArgs:
		return "Invalid Args"
	default:
		return "Unknown Code"
	}
}
