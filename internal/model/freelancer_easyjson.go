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

func easyjson92f8d27bDecodeGithubComGoParkMailRu20192ComandusInternalModel(in *jlexer.Lexer, out *FreelancerOutput) {
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
		case "ID":
			out.ID = int64(in.Int64())
		case "AccountId":
			out.AccountId = int64(in.Int64())
		case "Country":
			out.Country = string(in.String())
		case "City":
			out.City = string(in.String())
		case "Address":
			out.Address = string(in.String())
		case "Phone":
			out.Phone = string(in.String())
		case "TagLine":
			out.TagLine = string(in.String())
		case "Overview":
			out.Overview = string(in.String())
		case "ExperienceLevelId":
			out.ExperienceLevelId = int64(in.Int64())
		case "SpecialityId":
			out.SpecialityId = int64(in.Int64())
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
func easyjson92f8d27bEncodeGithubComGoParkMailRu20192ComandusInternalModel(out *jwriter.Writer, in FreelancerOutput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"AccountId\":"
		out.RawString(prefix)
		out.Int64(int64(in.AccountId))
	}
	{
		const prefix string = ",\"Country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"Phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	{
		const prefix string = ",\"TagLine\":"
		out.RawString(prefix)
		out.String(string(in.TagLine))
	}
	{
		const prefix string = ",\"Overview\":"
		out.RawString(prefix)
		out.String(string(in.Overview))
	}
	{
		const prefix string = ",\"ExperienceLevelId\":"
		out.RawString(prefix)
		out.Int64(int64(in.ExperienceLevelId))
	}
	{
		const prefix string = ",\"SpecialityId\":"
		out.RawString(prefix)
		out.Int64(int64(in.SpecialityId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FreelancerOutput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson92f8d27bEncodeGithubComGoParkMailRu20192ComandusInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FreelancerOutput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson92f8d27bEncodeGithubComGoParkMailRu20192ComandusInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FreelancerOutput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson92f8d27bDecodeGithubComGoParkMailRu20192ComandusInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FreelancerOutput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson92f8d27bDecodeGithubComGoParkMailRu20192ComandusInternalModel(l, v)
}
func easyjson92f8d27bDecodeGithubComGoParkMailRu20192ComandusInternalModel1(in *jlexer.Lexer, out *Freelancer) {
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
			out.AccountId = int64(in.Int64())
		case "country":
			out.Country = int64(in.Int64())
		case "city":
			out.City = int64(in.Int64())
		case "address":
			out.Address = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		case "tagline":
			out.TagLine = string(in.String())
		case "overview":
			out.Overview = string(in.String())
		case "experienceLevelId":
			out.ExperienceLevelId = int64(in.Int64())
		case "specialityId":
			out.SpecialityId = int64(in.Int64Str())
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
func easyjson92f8d27bEncodeGithubComGoParkMailRu20192ComandusInternalModel1(out *jwriter.Writer, in Freelancer) {
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
		out.Int64(int64(in.AccountId))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.Int64(int64(in.Country))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.Int64(int64(in.City))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	{
		const prefix string = ",\"tagline\":"
		out.RawString(prefix)
		out.String(string(in.TagLine))
	}
	{
		const prefix string = ",\"overview\":"
		out.RawString(prefix)
		out.String(string(in.Overview))
	}
	{
		const prefix string = ",\"experienceLevelId\":"
		out.RawString(prefix)
		out.Int64(int64(in.ExperienceLevelId))
	}
	{
		const prefix string = ",\"specialityId\":"
		out.RawString(prefix)
		out.Int64Str(int64(in.SpecialityId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Freelancer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson92f8d27bEncodeGithubComGoParkMailRu20192ComandusInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Freelancer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson92f8d27bEncodeGithubComGoParkMailRu20192ComandusInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Freelancer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson92f8d27bDecodeGithubComGoParkMailRu20192ComandusInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Freelancer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson92f8d27bDecodeGithubComGoParkMailRu20192ComandusInternalModel1(l, v)
}