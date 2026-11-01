package fixture
import "encoding/json"
type Request struct { Text string `json:"text"`; Language string `json:"language,omitempty"` }
type Session struct { Text string `json:"text"`; Language string `json:"language,omitempty"` }
type Call struct { Prompt string; Language string }
var PromptsJA = map[string]string{"chat":"質問に回答してください", "plan":"実装を計画してください", "review":"レビューしてください", "implement":"実装してください", "verify":"変更を確認してください"}
var PromptsEN = map[string]string{"chat":"Please answer the question", "plan":"Please plan the implementation", "review":"Please review", "implement":"Please implement", "verify":"Please verify the changes"}
func Receive(body []byte) Session { var r Request; _ = json.Unmarshal(body, &r); lang := r.Language; if lang == "" { lang = "ja" }; return Session{Text:r.Text, Language:lang} }
func Save(s Session) []byte { b,_:=json.Marshal(s);return b }
func Load(b []byte) Session {var s Session; _=json.Unmarshal(b,&s); if s.Language == "" { s.Language = "ja" }; return s}
func Calls(s Session) map[string]Call { 
    lang := s.Language
    if lang == "" { lang = "ja" }
    prompts := PromptsJA
    if lang == "en" { prompts = PromptsEN }
    out:=map[string]Call{}; 
    for k,p:=range prompts {out[k]=Call{Prompt:p,Language:lang}};
    return out 
}
func Echo(s string) string { return s }
