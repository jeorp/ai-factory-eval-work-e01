package fixture
import "encoding/json"
type Request struct { Text string `json:"text"`; Language string `json:"language,omitempty"` }
type Session struct { Text string `json:"text"`; Language string `json:"language,omitempty"` }
type Call struct { Prompt string; Language string }
var Prompts = map[string]map[string]string{"chat":{"ja":"質問に回答してください","en":"Please answer the question"},"plan":{"ja":"実装を計画してください","en":"Please plan the implementation"},"review":{"ja":"レビューしてください","en":"Please review"},"implement":{"ja":"実装してください","en":"Please implement"},"verify":{"ja":"変更を確認してください","en":"Please verify changes"}}
func Receive(body []byte) Session { var r Request; _ = json.Unmarshal(body, &r); lang := r.Language; if lang == "" { lang = "ja" }; return Session{Text:r.Text, Language:lang} }
func Save(s Session) []byte { b,_:=json.Marshal(Session{Text:s.Text,Language:s.Language});return b }
func Load(b []byte) Session {var s Session; _=json.Unmarshal(b,&s); if s.Language == "" { s.Language = "ja" }; return s}
func Calls(s Session) map[string]Call { out:=map[string]Call{}; lang := s.Language; if lang == "" { lang = "ja" }; for k,p:=range Prompts { promptText := p[lang]; if promptText == "" { promptText = p["ja"] }; out[k]=Call{Prompt:promptText,Language:lang}};return out }
func Echo(s string) string { return s }
