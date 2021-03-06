package main

var metro = map[string][]string{
	// Yokohama Municipal Subway Route Blue Line
	"あざみ野":    {"中川"},
	"中川":      {"センター北"},
	"センター北":   {"センター南"},
	"センター南":   {"仲町台", "都筑ふれあいの丘"},
	"仲町台":     {"新羽"},
	"新羽":      {"北新横浜"},
	"北新横浜":    {"新横浜"},
	"新横浜":     {"岸根公園"},
	"岸根公園":    {"片倉町"},
	"片倉町":     {"三ツ沢上町"},
	"三ツ沢上町":   {"三ツ沢下町"},
	"三ツ沢下町":   {"横浜"},
	"横浜":      {"高島町"},
	"高島町":     {"桜木町"},
	"桜木町":     {"関内"},
	"関内":      {"伊勢佐木長者町"},
	"伊勢佐木長者町": {"阪東橋"},
	"阪東橋":     {"吉野町"},
	"吉野町":     {"蒔田"},
	"蒔田":      {"弘明寺"},
	"弘明寺":     {"上大岡"},
	"上大岡":     {"港南中央"},
	"港南中央":    {"上永谷"},
	"上永谷":     {"下永谷"},
	"下永谷":     {"舞岡"},
	"舞岡":      {"戸塚"},
	"戸塚":      {"踊場"},
	"踊場":      {"中田"},
	"中田":      {"立場"},
	"立場":      {"下飯田"},
	"下飯田":     {"湘南台"},

	// Yokohama Municipal Subway Route Green Line
	"川和町":      {"中山"},
	"都筑ふれあいの丘": {"川和町"},
	"北山田":      {"センター北"},
	"東山田":      {"北山田"},
	"高田":       {"東山田"},
	"日吉本町":     {"高田"},
	"日吉":       {"日吉本町"},
}
