package consts

//app properties consts
const (
	AppName     = "Notifier"
	AppUsage    = "A simple tool for notification"
	AppHelpName = "Notifier"
	AppVersion  = "1.2.0"
	AppAuthor   = "ZHU YUE"
)

//limitation parameters
const (
	//maximum subject length for email(Bytes)
	MAX_EMAIL_SUBJECT_LEN = 256
)

//config files
const (
	NotifyrcFile string = ".notifyrc"
	DefaultsFile string = ".notifdef"
)

//Notifiers name
const (
	EmailNotifier string = "smtpemailnotifier"
	SlackNotifier string = "slackNotifier"
)

//ERR refers to error code(0~255), equals to uint8
type ERR uint8

//common error codes
const (
	ERR_MAX = 255

	//UNIX/LINUX general error code
	NIL                 ERR = 0
	SUCCESS             ERR = 0
	GENERAL_ERR         ERR = 1
	MISS_USE            ERR = 2
	CMD_CANNOT_EXE      ERR = 126
	CMD_NOT_FOUND       ERR = 127
	INVALID_ARG_TO_EXIT ERR = 128
	CTRLC_TERMINATE     ERR = 130

	//file-reading error code
	NOTIFRC_PARSE_ERR ERR = 55 //eror occurs while parsing notifier config file(P)
	DFLTS_PARSE_ERR   ERR = 56 //error occurs while pasing default config file(P)

	//smtpemail error code
	SMTPM_NOTGT         ERR = 10 //no target email address
	SMTPM_INVAL         ERR = 11 //smtp notif not valid(Not an exact error)
	SMTPM_SVR_CONN_ERR  ERR = 12 //lose internet connection or get refused by remote host.stop and check network, host and port(P)
	SMTPM_CLT_BLD_ERR   ERR = 13 //error occurs while building a client, stop and check network, host and port(P)
	SMTPM_AUTH_ERR      ERR = 14 //error occurs while authenticating mail account and password, stop and check your account and network(P)
	SMTPM_SENDER_ERR    ERR = 15 //error occurs while adding sender account, stop and check your account(P)
	SMTPM_RCVR_ERR      ERR = 16 //error occurs while adding receivers' accounts, stop and check those account IDs(P)
	SMTPM_CLT_IO_ERR    ERR = 17 //... while initializing or close a iowriter for email client(P)
	SMTPM_CLT_DATA_ERR  ERR = 18 //... while trying to write message in the client(P)
	SMTPM_CLT_CLOSE_ERR ERR = 19 //... while closing an smtpemail client(P)

	//slack error code
	SLK_NOTGT        ERR = 28 //no target slack ids
	SLK_INVAL        ERR = 29 //slack notif not valid(Not an exact error)
	SLK_TOKEN_INVAL  ERR = 30 //slack token not invalid(P)
	SLK_CHL_ERR      ERR = 31 //token is right, just got stuck in posting to one target user(or channel)(P)
	SLK_SVR_CONN_ERR ERR = 32 //got stuck because of the network, or be refused by slack host.(T)

	//slack webhook error code
	REQ_FAIL        ERR = 39 //no network connection
	INVALID_PAYLOAD ERR = 40 /*HTTP 400 Bad Request the data sent in your request cannot be understood as presented.
	  verify your content body matches your content type and is structurally valid.*/
	USER_NOT_FOUND ERR = 42 //HTTP 400 bad Request. the user used in your request does not actually exist.
	ACTION_FORBID  ERR = 43 //HTTP 403 Forbidden. the team associated with your request has some kind of restriction on the webhook posting in this context.
	CHL_NOT_FOUND  ERR = 44 //HTTP 404 Not Found. the channel associated with your request does not exist.
	CHL_ARCHIVED   ERR = 41 //HTTP 410 Gone. the channel has been archived and doesn't accept further messages, even from your incoming webhook.
	ROLLUP_ERROR   ERR = 50 //HTTP 500 Server Error. something strange and unusual happened that was likely not your fault at all.

)
