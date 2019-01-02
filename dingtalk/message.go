package main

const (
	MESSAGETYPE_TEXT       = "text"
	MESSAGETYPE_LINK       = "link"
	MESSAGETYPE_MARKDOWN   = "markdown"
	MESSAGETYPE_ACTIONCARD = "actionCard"
	MESSAGETYPE_FEEDCARD   = "feedCard"
)

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type text struct {
	Content string `json:"content"`
}

type TextMsg struct {
	MsgType string `json:"msgtype"`
	Text    text   `json:"text"`
	At      At     `json:"at"`
}

type link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"`
}

type LinkMsg struct {
	MsgType string `json:"msgtype"`
	Link    link   `json:"link"`
}

func NewTextMsg(content string) TextMsg {
	return TextMsg{
		MsgType: MESSAGETYPE_TEXT,
		Text: text{
			Content: content,
		},
	}
}

func NewLinkMsg(title, text, messageUrl string) LinkMsg {
	return LinkMsg{
		MsgType: MESSAGETYPE_LINK,
		Link: link{
			Title:      title,
			Text:       text,
			MessageUrl: messageUrl,
		},
	}
}

type markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type MarkdownMsg struct {
	MsgType  string   `json:"msgtype"`
	MarkDown markdown `json:"markdown"`
	At       At       `json:"at"`
}

func NewMarkdownMsg(title, text string) MarkdownMsg {
	return MarkdownMsg{
		MsgType: MESSAGETYPE_MARKDOWN,
		MarkDown: markdown{
			Title: title,
			Text:  text,
		},
	}
}

type Button struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type actionCard struct {
	Title          string   `json:"title"`
	Text           string   `json:"text"`
	HideAvatar     string   `json:"hideAvatar"`
	BtnOrientation string   `json:"btnOrientation"`
	SingleTitle    string   `json:"singleTitle"`
	SingleURL      string   `json:"singleURL"`
	Btns           []Button `json:"btns"`
}

type ActionCardMsg struct {
	MsgType    string     `json:"msgtype"`
	ActionCard actionCard `json:"actionCard"`
}

func NewWholeActionCardMsg(title, text, singleTitle, singleURL string) ActionCardMsg {
	return ActionCardMsg{
		MsgType: MESSAGETYPE_ACTIONCARD,
		ActionCard: actionCard{
			Title:       title,
			Text:        text,
			SingleTitle: singleTitle,
			SingleURL:   singleURL,
		},
	}
}

func NewIndependentCardMsg(title, text string, btns []Button) ActionCardMsg {
	return ActionCardMsg{
		MsgType: MESSAGETYPE_ACTIONCARD,
		ActionCard: actionCard{
			Title: title,
			Text:  text,
			Btns:  btns,
		},
	}
}

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

type FeedCard struct {
	Links []FeedCardLink `json:"links"`
}

type FeedCardMsg struct {
	MsgType  string   `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
}

func NewFeedCardMsg(links []FeedCardLink) FeedCardMsg {
	return FeedCardMsg{
		MsgType: MESSAGETYPE_FEEDCARD,
		FeedCard: FeedCard{
			Links: links,
		},
	}
}
