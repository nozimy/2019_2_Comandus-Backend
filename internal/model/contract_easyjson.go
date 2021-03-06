// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	company_grpc "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/company/delivery/grpc/company_grpc"
	freelancer_grpc "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/freelancer/delivery/grpc/freelancer_grpc"
	job_grpc "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/job/delivery/grpc/job_grpc"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

func easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalModel(in *jlexer.Lexer, out *ContractOutput) {
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
		case "Company":
			easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppCompanyDeliveryGrpcCompanyGrpc(in, &out.Company)
		case "Freelancer":
			easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc(in, &out.Freelancer)
		case "Job":
			easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppJobDeliveryGrpcJobGrpc(in, &out.Job)
		case "Contract":
			(out.Contract).UnmarshalEasyJSON(in)
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
func easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalModel(out *jwriter.Writer, in ContractOutput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Company\":"
		out.RawString(prefix[1:])
		easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppCompanyDeliveryGrpcCompanyGrpc(out, in.Company)
	}
	{
		const prefix string = ",\"Freelancer\":"
		out.RawString(prefix)
		easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc(out, in.Freelancer)
	}
	{
		const prefix string = ",\"Job\":"
		out.RawString(prefix)
		easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppJobDeliveryGrpcJobGrpc(out, in.Job)
	}
	{
		const prefix string = ",\"Contract\":"
		out.RawString(prefix)
		(in.Contract).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContractOutput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContractOutput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContractOutput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContractOutput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalModel(l, v)
}
func easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppJobDeliveryGrpcJobGrpc(in *jlexer.Lexer, out *job_grpc.Job) {
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
		case "HireManagerId":
			out.HireManagerId = int64(in.Int64())
		case "Title":
			out.Title = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Files":
			out.Files = string(in.String())
		case "SpecialityId":
			out.SpecialityId = int64(in.Int64())
		case "ExperienceLevelId":
			out.ExperienceLevelId = int64(in.Int64())
		case "PaymentAmount":
			out.PaymentAmount = float32(in.Float32())
		case "Country":
			out.Country = int64(in.Int64())
		case "City":
			out.City = int64(in.Int64())
		case "JobTypeId":
			out.JobTypeId = int64(in.Int64())
		case "Date":
			if in.IsNull() {
				in.Skip()
				out.Date = nil
			} else {
				if out.Date == nil {
					out.Date = new(timestamp.Timestamp)
				}
				easyjson42340b0aDecodeGithubComGolangProtobufPtypesTimestamp(in, out.Date)
			}
		case "Status":
			out.Status = string(in.String())
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
func easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppJobDeliveryGrpcJobGrpc(out *jwriter.Writer, in job_grpc.Job) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"ID\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	if in.HireManagerId != 0 {
		const prefix string = ",\"HireManagerId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.HireManagerId))
	}
	if in.Title != "" {
		const prefix string = ",\"Title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Description != "" {
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Files != "" {
		const prefix string = ",\"Files\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Files))
	}
	if in.SpecialityId != 0 {
		const prefix string = ",\"SpecialityId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SpecialityId))
	}
	if in.ExperienceLevelId != 0 {
		const prefix string = ",\"ExperienceLevelId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ExperienceLevelId))
	}
	if in.PaymentAmount != 0 {
		const prefix string = ",\"PaymentAmount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.PaymentAmount))
	}
	if in.Country != 0 {
		const prefix string = ",\"Country\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Country))
	}
	if in.City != 0 {
		const prefix string = ",\"City\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.City))
	}
	if in.JobTypeId != 0 {
		const prefix string = ",\"JobTypeId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JobTypeId))
	}
	if in.Date != nil {
		const prefix string = ",\"Date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson42340b0aEncodeGithubComGolangProtobufPtypesTimestamp(out, *in.Date)
	}
	if in.Status != "" {
		const prefix string = ",\"Status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	out.RawByte('}')
}
func easyjson42340b0aDecodeGithubComGolangProtobufPtypesTimestamp(in *jlexer.Lexer, out *timestamp.Timestamp) {
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
		case "seconds":
			out.Seconds = int64(in.Int64())
		case "nanos":
			out.Nanos = int32(in.Int32())
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
func easyjson42340b0aEncodeGithubComGolangProtobufPtypesTimestamp(out *jwriter.Writer, in timestamp.Timestamp) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Seconds != 0 {
		const prefix string = ",\"seconds\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Seconds))
	}
	if in.Nanos != 0 {
		const prefix string = ",\"nanos\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Nanos))
	}
	out.RawByte('}')
}
func easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc(in *jlexer.Lexer, out *freelancer_grpc.ExtendedFreelancer) {
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
		case "Fr":
			if in.IsNull() {
				in.Skip()
				out.Fr = nil
			} else {
				if out.Fr == nil {
					out.Fr = new(freelancer_grpc.Freelancer)
				}
				easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc1(in, out.Fr)
			}
		case "FirstName":
			out.FirstName = string(in.String())
		case "SecondName":
			out.SecondName = string(in.String())
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
func easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc(out *jwriter.Writer, in freelancer_grpc.ExtendedFreelancer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Fr != nil {
		const prefix string = ",\"Fr\":"
		first = false
		out.RawString(prefix[1:])
		easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc1(out, *in.Fr)
	}
	if in.FirstName != "" {
		const prefix string = ",\"FirstName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FirstName))
	}
	if in.SecondName != "" {
		const prefix string = ",\"SecondName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SecondName))
	}
	out.RawByte('}')
}
func easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc1(in *jlexer.Lexer, out *freelancer_grpc.Freelancer) {
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
			out.Country = int64(in.Int64())
		case "City":
			out.City = int64(in.Int64())
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
func easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppFreelancerDeliveryGrpcFreelancerGrpc1(out *jwriter.Writer, in freelancer_grpc.Freelancer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"ID\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	if in.AccountId != 0 {
		const prefix string = ",\"AccountId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AccountId))
	}
	if in.Country != 0 {
		const prefix string = ",\"Country\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Country))
	}
	if in.City != 0 {
		const prefix string = ",\"City\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.City))
	}
	if in.Address != "" {
		const prefix string = ",\"Address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	if in.Phone != "" {
		const prefix string = ",\"Phone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Phone))
	}
	if in.TagLine != "" {
		const prefix string = ",\"TagLine\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TagLine))
	}
	if in.Overview != "" {
		const prefix string = ",\"Overview\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Overview))
	}
	if in.ExperienceLevelId != 0 {
		const prefix string = ",\"ExperienceLevelId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ExperienceLevelId))
	}
	if in.SpecialityId != 0 {
		const prefix string = ",\"SpecialityId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SpecialityId))
	}
	out.RawByte('}')
}
func easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalAppCompanyDeliveryGrpcCompanyGrpc(in *jlexer.Lexer, out *company_grpc.CompanyOutput) {
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
		case "CompanyName":
			out.CompanyName = string(in.String())
		case "Site":
			out.Site = string(in.String())
		case "TagLine":
			out.TagLine = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Country":
			out.Country = string(in.String())
		case "City":
			out.City = string(in.String())
		case "Address":
			out.Address = string(in.String())
		case "Phone":
			out.Phone = string(in.String())
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
func easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalAppCompanyDeliveryGrpcCompanyGrpc(out *jwriter.Writer, in company_grpc.CompanyOutput) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"ID\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	if in.CompanyName != "" {
		const prefix string = ",\"CompanyName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CompanyName))
	}
	if in.Site != "" {
		const prefix string = ",\"Site\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Site))
	}
	if in.TagLine != "" {
		const prefix string = ",\"TagLine\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TagLine))
	}
	if in.Description != "" {
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Country != "" {
		const prefix string = ",\"Country\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Country))
	}
	if in.City != "" {
		const prefix string = ",\"City\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.City))
	}
	if in.Address != "" {
		const prefix string = ",\"Address\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Address))
	}
	if in.Phone != "" {
		const prefix string = ",\"Phone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Phone))
	}
	out.RawByte('}')
}
func easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalModel1(in *jlexer.Lexer, out *Contract) {
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
		case "responseId":
			out.ResponseID = int64(in.Int64())
		case "companyId":
			out.CompanyID = int64(in.Int64())
		case "freelancerId":
			out.FreelancerID = int64(in.Int64())
		case "startTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
			}
		case "endTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EndTime).UnmarshalJSON(data))
			}
		case "status":
			out.Status = string(in.String())
		case "statusFreelancerWork":
			out.StatusFreelancerWork = string(in.String())
		case "freelancerGrade":
			out.FreelancerGrade = int(in.Int())
		case "freelancerComment":
			out.FreelancerComment = string(in.String())
		case "clientGrade":
			out.ClientGrade = int(in.Int())
		case "clientComment":
			out.ClientComment = string(in.String())
		case "paymentAmount":
			out.PaymentAmount = float32(in.Float32())
		case "timeEstimation":
			out.TimeEstimation = int(in.Int())
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
func easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalModel1(out *jwriter.Writer, in Contract) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"responseId\":"
		out.RawString(prefix)
		out.Int64(int64(in.ResponseID))
	}
	{
		const prefix string = ",\"companyId\":"
		out.RawString(prefix)
		out.Int64(int64(in.CompanyID))
	}
	{
		const prefix string = ",\"freelancerId\":"
		out.RawString(prefix)
		out.Int64(int64(in.FreelancerID))
	}
	{
		const prefix string = ",\"startTime\":"
		out.RawString(prefix)
		out.Raw((in.StartTime).MarshalJSON())
	}
	{
		const prefix string = ",\"endTime\":"
		out.RawString(prefix)
		out.Raw((in.EndTime).MarshalJSON())
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"statusFreelancerWork\":"
		out.RawString(prefix)
		out.String(string(in.StatusFreelancerWork))
	}
	{
		const prefix string = ",\"freelancerGrade\":"
		out.RawString(prefix)
		out.Int(int(in.FreelancerGrade))
	}
	{
		const prefix string = ",\"freelancerComment\":"
		out.RawString(prefix)
		out.String(string(in.FreelancerComment))
	}
	{
		const prefix string = ",\"clientGrade\":"
		out.RawString(prefix)
		out.Int(int(in.ClientGrade))
	}
	{
		const prefix string = ",\"clientComment\":"
		out.RawString(prefix)
		out.String(string(in.ClientComment))
	}
	{
		const prefix string = ",\"paymentAmount\":"
		out.RawString(prefix)
		out.Float32(float32(in.PaymentAmount))
	}
	{
		const prefix string = ",\"timeEstimation\":"
		out.RawString(prefix)
		out.Int(int(in.TimeEstimation))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contract) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contract) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson42340b0aEncodeGithubComGoParkMailRu20192ComandusInternalModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contract) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contract) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson42340b0aDecodeGithubComGoParkMailRu20192ComandusInternalModel1(l, v)
}
