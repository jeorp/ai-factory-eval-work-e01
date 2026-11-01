package fixture
import "encoding/json"
type Request struct { Text string `json:"text"`; Language string `json:"language,omitempty"` }
type Session struct { Text string `json:"text"`; Language string `json:"language,omitempty"` }
type Call struct { Prompt string; Language string }
var Prompts = map[string]string{"chat":"Please answer the question", "plan":"Please plan the implementation", "review":"Please review the changes", "implement":"Please implement the changes", "verify":"Please verify the changes"}
func Receive(body []byte) Session { var r Request; _ = json.Unmarshal(body, &r); lang := r.Language; if lang == "" { lang = "ja" }; return Session{Text:r.Text, Language:lang} }
func Save(s Session) []byte { b,_:=json.Marshal(Session{Text:s.Text, Language:s.Language});return b }
func Load(b []byte) Session {var s Session; _=json.Unmarshal(b,&s); if s.Language == "" { s.Language = "ja" }; return s}
func Calls(s Session) map[string]Call { out:=map[string]Call{}; for k,p:=range Prompts {out[k]=Call{Prompt:p,Language:s.Language}};return out }
func Echo(s string) string { return s }
