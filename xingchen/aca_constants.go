package xingchen

type Version int

const (
	V1 Version = iota + 1
	V2
)

const (
	V1_API_NAME = "v1"
	V2_API_NAME = "v2"
)

const (
	GATEWAY_HEADER_SERVICE_ROUTER = "x-fag-servicename"
	GATEWAY_HEADER_APPCODE        = "x-fag-appcode"
	APP_CODE                      = "aca"
)

const (
	DEFAULT_V1_BASE_PATH = "https://xingchen.aliyun.com"
	DEFAULT_V2_BASE_PATH = "https://nlp.aliyuncs.com"
)

var CHAT_APIS = map[string]struct{}{
	"/v2/api/chat/send":      {},
	"/v2/api/groupchat/send": {},
}

var V2_PATH_ROUTE_MAP = map[string]string{
	"/v2/api/chat/send":                       "aca-chat-send",
	"/v2/api/chat/generate":                   "aca-chat-regenerate",
	"/v2/api/chat/reminder":                   "aca-chat-reminder",
	"/v2/api/chat/stop":                       "aca-chat-stop",
	"/v2/api/groupchat/send":                  "aca-groupchat-send",
	"/v2/api/chat/message/histories":          "aca-message-history",
	"/v2/api/chat/rating":                     "aca-message-rating",
	"/v2/api/chat/reset":                      "aca-chat-reset",
	"/v2/api/character/create":                "aca-character-create",
	"/v2/api/character/update":                "aca-character-update",
	"/v2/api/character/createOrUpdateVersion": "aca-character-version-mgmt",
	"/v2/api/character/newversion/recommend":  "aca-character-version-recommend",
	"/v2/api/character/details":               "aca-character-details",
	"/v2/api/character/delete":                "aca-character-delete",
	"/v2/api/character/search":                "aca-character-search",
	"/v2/api/character/versions":              "aca-character-versions",
	"/v2/api/common/file/asyn/upload":         "aca-doc-converter-submit",
	"/v2/api/common/file/asyn/download":       "aca-doc-converter-result",
	"/v2/api/chat/polling/image":              "aca-polling-image",
	"/v2/api/extract/summary":                 "aca-extract-memory-summary",
	"/v2/api/extract/kv":                      "aca-extract-memory-kv",
	"/v2/api/character/auto/desc":             "aca-char-auto-desc",
}
