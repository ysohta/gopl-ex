package main

type ReplyCode int

const (
	ReplyCodeOkay                = 200
	ReplyCodeNameSystem          = 215
	ReplyCodeCloeseConnection    = 221
	ReplyCodeUserLoggedIn        = 230
	ReplyCodeFileActionComplete  = 250
	ReplyCodePathNameCreated     = 257
	ReplyCodeNeedAccount         = 332
	ReplyCodeConnectionClosed    = 426
	ReplyCodeCommandUnrecognized = 500
	ReplyCodeParameterError      = 501
	ReplyCodeNotImplemented      = 502
	ReplyCodeFileUnavailable     = 550
)
