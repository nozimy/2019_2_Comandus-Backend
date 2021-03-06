package responseRepository

import (
	"database/sql"
	user_response "github.com/go-park-mail-ru/2019_2_Comandus/internal/app/proposal"
	"github.com/go-park-mail-ru/2019_2_Comandus/internal/model"
	"github.com/go-park-mail-ru/2019_2_Comandus/monitoring"
	"github.com/prometheus/client_golang/prometheus"
)

type ResponseRepository struct {
	db *sql.DB
}

func NewResponseRepository(db *sql.DB) user_response.Repository {
	return &ResponseRepository{db}
}

func (r *ResponseRepository) Create(response *model.Response) error {
	timer := prometheus.NewTimer(monitoring.DBQueryDuration.With(prometheus.
		Labels{"rep": "response", "method": "create"}))
	defer timer.ObserveDuration()

	return r.db.QueryRow(
		"INSERT INTO responses (freelancerId, jobId, files, date, statusManager, statusFreelancer, paymentAmount,"+
			"coverLetter, timeEstimation) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		response.FreelancerId,
		response.JobId,
		response.Files,
		response.Date,
		response.StatusManager,
		response.StatusFreelancer,
		response.PaymentAmount,
		response.CoverLetter,
		response.TimeEstimation,
	).Scan(&response.ID)
}

func (r *ResponseRepository) Edit(response *model.Response) error {
	timer := prometheus.NewTimer(monitoring.DBQueryDuration.With(prometheus.
		Labels{"rep": "response", "method": "edit"}))
	defer timer.ObserveDuration()

	return r.db.QueryRow(
		"UPDATE responses SET files = $1, statusmanager = $2, statusFreelancer = $3, paymentAmount = $4,"+
			"coverLetter = $5, timeEstimation = $6 WHERE id = $7 "+
			"RETURNING id",
		response.Files,
		response.StatusManager,
		response.StatusFreelancer,
		response.PaymentAmount,
		response.CoverLetter,
		response.TimeEstimation,
		response.ID,
	).Scan(&response.ID)
}

func (r *ResponseRepository) ListForFreelancer(id int64) ([]model.ExtendResponse, error) {
	timer := prometheus.NewTimer(monitoring.DBQueryDuration.With(prometheus.
		Labels{"rep": "response", "method": "listForFreelancer"}))
	defer timer.ObserveDuration()

	var responses []model.ExtendResponse
	rows, err := r.db.Query(
		"SELECT r.id, r.freelancerId, r.jobId, r.files, r.date, r.statusManager, r.statusFreelancer, r.paymentAmount,"+
			"r.coverLetter, r.timeEstimation, U.firstName, U.secondName, J.title "+
			"FROM responses AS r "+
			"INNER JOIN freelancers AS F "+
			"ON R.freelancerid = F.id "+
			"INNER JOIN users AS U "+
			"ON U.accountid = F.accountid "+
			"INNER JOIN jobs AS J "+
			"ON J.id = r.jobid "+
			"WHERE r.freelancerId = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := model.Response{}
		exR := model.ExtendResponse{}
		err := rows.Scan(&r.ID, &r.FreelancerId, &r.JobId, &r.Files, &r.Date, &r.StatusManager,
			&r.StatusFreelancer, &r.PaymentAmount, &r.CoverLetter, &r.TimeEstimation, &exR.FirstName,
			&exR.SecondName, &exR.JobTitle)
		if err != nil {
			return nil, err
		}
		exR.R = &r
		responses = append(responses, exR)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return responses, nil
}

func (r *ResponseRepository) ListForManager(id int64) ([]model.ExtendResponse, error) {
	timer := prometheus.NewTimer(monitoring.DBQueryDuration.With(prometheus.
		Labels{"rep": "response", "method": "listForManager"}))
	defer timer.ObserveDuration()

	var responses []model.ExtendResponse
	rows, err := r.db.Query(
		"SELECT responses.id, responses.freelancerId, responses.jobId, responses.files, responses.date, "+
			"responses.statusManager, responses.statusFreelancer, responses.paymentAmount,"+
			"responses.coverLetter, responses.timeEstimation, U.firstName, U.secondName, J.title "+
			"FROM responses "+
			"INNER JOIN freelancers AS F "+
			"ON responses.freelancerid = F.id "+
			"INNER JOIN users AS U "+
			"ON U.accountid = F.accountid "+
			"INNER JOIN jobs AS J "+
			"ON J.id = responses.jobid "+
			"WHERE J.managerId = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		exR := model.ExtendResponse{}
		r := model.Response{}
		err := rows.Scan(&r.ID, &r.FreelancerId, &r.JobId, &r.Files, &r.Date, &r.StatusManager,
			&r.StatusFreelancer, &r.PaymentAmount, &r.CoverLetter, &r.TimeEstimation,
			&exR.FirstName, &exR.SecondName, &exR.JobTitle)
		if err != nil {
			return nil, err
		}
		exR.R = &r
		responses = append(responses, exR)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return responses, nil
}

func (r *ResponseRepository) Find(id int64) (*model.Response, error) {
	timer := prometheus.NewTimer(monitoring.DBQueryDuration.With(prometheus.
		Labels{"rep": "response", "method": "find"}))
	defer timer.ObserveDuration()

	response := &model.Response{}
	if err := r.db.QueryRow(
		"SELECT id, freelancerId, jobId, files, date, statusManager, statusFreelancer, paymentAmount, "+
			"coverLetter, timeEstimation FROM responses WHERE id = $1",
		id,
	).Scan(
		&response.ID,
		&response.FreelancerId,
		&response.JobId,
		&response.Files,
		&response.Date,
		&response.StatusManager,
		&response.StatusFreelancer,
		&response.PaymentAmount,
		&response.CoverLetter,
		&response.TimeEstimation,
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *ResponseRepository) ListResponsesOnJobID(jobID int64) ([]model.ExtendResponse, error) {
	timer := prometheus.NewTimer(monitoring.DBQueryDuration.With(prometheus.
		Labels{"rep": "response", "method": "listResponsesOnJobID"}))
	defer timer.ObserveDuration()

	var responses []model.ExtendResponse
	rows, err := r.db.Query(
		"SELECT R.id, R.freelancerId, R.jobId, R.files, R.date, R.statusManager, R.statusFreelancer, "+
			"R.paymentAmount, R.coverLetter, R.timeEstimation, U.firstname , U.secondname, J.title "+
			"FROM responses AS R "+
			"INNER JOIN freelancers AS F "+
			"ON R.freelancerid = F.id "+
			"INNER JOIN users AS U "+
			"ON U.accountid = F.accountid "+
			"INNER JOIN jobs AS J "+
			"ON J.id = R.jobid "+
			"WHERE jobid = $1", jobID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := model.Response{}
		exR := model.ExtendResponse{}
		err := rows.Scan(&r.ID, &r.FreelancerId, &r.JobId, &r.Files, &r.Date, &r.StatusManager,
			&r.StatusFreelancer, &r.PaymentAmount, &r.CoverLetter, &r.TimeEstimation, &exR.FirstName,
			&exR.SecondName, &exR.JobTitle)
		if err != nil {
			return nil, err
		}
		exR.R = &r

		responses = append(responses, exR)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return responses, nil
}

func (r *ResponseRepository) CheckForHavingResponse(jobID int64, freelID int64) (bool, error) {
	var count int
	err := r.db.QueryRow(
		"SELECT COUNT(*) "+
			"FROM responses "+
			"WHERE jobid = $1 AND freelancerid = $2 AND statusManager = 'REVIEW' ",
		jobID, freelID).Scan(&count)
	if err != nil {
		return false, err
	}

	if count != 0 {
		return true, nil
	}

	return false, nil
}
