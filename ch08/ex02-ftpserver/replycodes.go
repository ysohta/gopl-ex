package main

type ReplyCode int

const (
	ReplyCodeFileStatusOkay           = 150
	ReplyCodeOkay                     = 200
	ReplyCodeNameSystem               = 215
	ReplyCodeCloeseConnection         = 221
	ReplyCodeCloseDataConnection      = 226
	ReplyCodeEnteringPasv             = 227
	ReplyCodeEnteringEpsv             = 229
	ReplyCodeUserLoggedIn             = 230
	ReplyCodeFileActionComplete       = 250
	ReplyCodePathNameCreated          = 257
	ReplyCodeNeedAccount              = 332
	ReplyCodeFailedOpenDataConnection = 425
	ReplyCodeConnectionClosed         = 426
	ReplyCodeCommandUnrecognized      = 500
	ReplyCodeParameterError           = 501
	ReplyCodeNotImplemented           = 502
	ReplyCodeFileUnavailable          = 550
)
