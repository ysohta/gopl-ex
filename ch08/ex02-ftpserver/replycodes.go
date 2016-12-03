package main

type ReplyCode int

const (
	ReplyCodeOkay                = 200
	ReplyCodeUserLoggedIn        = 230
	ReplyCodeNeedAccount         = 332
	ReplyCodeNameSystem          = 215
	ReplyCodeCommandUnrecognized = 500
	ReplyCodeNotImplemented      = 502
	ReplyCodeCloeseConnection    = 221
	ReplyCodePathNameCreated     = 257
	ReplyCodeFileActionComplete  = 250
	ReplyCodeParameterError      = 501
	ReplyCodeFileUnavailable     = 550
)
