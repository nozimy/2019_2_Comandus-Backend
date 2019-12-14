// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonAf94a8adDecodeGithubComGoParkMailRu20192ComandusInternalModel(in *jlexer.Lexer, out *HireManager) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "accountId":
			out.AccountID = int64(in.Int64())
		case "location":
			out.Location = string(in.String())
		case "companyId":
			out.CompanyID = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAf94a8adEncodeGithubComGoParkMailRu20192ComandusInternalModel(out *jwriter.Writer, in HireManager) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"accountId\":"
		out.RawString(prefix)
		out.Int64(int64(in.AccountID))
	}
	{
		const prefix string = ",\"location\":"
		out.RawString(prefix)
		out.String(string(in.Location))
	}
	{
		const prefix string = ",\"companyId\":"
		out.RawString(prefix)
		out.Int64(int64(in.CompanyID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HireManager) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAf94a8adEncodeGithubComGoParkMailRu20192ComandusInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HireManager) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAf94a8adEncodeGithubComGoParkMailRu20192ComandusInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HireManager) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAf94a8adDecodeGithubComGoParkMailRu20192ComandusInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HireManager) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAf94a8adDecodeGithubComGoParkMailRu20192ComandusInternalModel(l, v)
}