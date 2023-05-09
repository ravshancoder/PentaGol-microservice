package postgres

import (
	"log"
	"time"

	p "github.com/PentaGol/liga_service/genproto/liga"
)

func (r *Repo) CreateGame(game *p.GameRequest) (*p.GameResponse, error) {
	var res p.GameResponse
	err := r.db.QueryRow(`
		insert into 
			games(
				time,
				condtion,
				first_team_id,
				second_team_id,
				result_first_team,
				result_second_team,
				first_team_point,
				second_team_point,
				liga_id
			)
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		returning 
			id, 
			time,
			condtion,
			first_team_id,
			second_team_id,
			result_first_team,
			result_second_team,
			first_team_point,
			second_team_point,
			liga_id, 
			created_at, 
			updated_at`, game.Time, game.Condtion, game.FirstTeamId, game.SecondTeamId, game.ResultFirstTeam, game.ResultSecondTeam, game.FirstTeamPoint, game.SecondTeamPoint, game.LigaId).
		Scan(
			&res.Id,
			&res.Time,
			&res.Condtion,
			&res.FirstTeamId,
			&res.SecondTeamId,
			&res.ResultFirstTeam,
			&res.ResultSecondTeam,
			&res.FirstTeamPoint,
			&res.SecondTeamPoint,
			&res.LigaId,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to create liga")
		return &p.GameResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetGameById(game *p.IdRequest) (*p.GameResponse, error) {
	res := p.GameResponse{}
	err := r.db.QueryRow(`
		select 
			id, 
			time,
			condtion,
			first_team_id,
			second_team_id,
			result_first_team,
			result_second_team,
			first_team_point,
			second_team_point,
			liga_id, 
			created_at, 
			updated_at
		from 
			games 
		where 
			id = $1`, game.Id).
		Scan(
			&res.Id,
			&res.Time,
			&res.Condtion,
			&res.FirstTeamId,
			&res.SecondTeamId,
			&res.ResultFirstTeam,
			&res.ResultSecondTeam,
			&res.FirstTeamPoint,
			&res.SecondTeamPoint,
			&res.LigaId,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to get game")
		return &p.GameResponse{}, err
	}

	return &res, nil
}

func (r *Repo) GetAllGames(req *p.AllGameRequest) (*p.Games, error) {
	res := p.Games{}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Query(`
		select 
			id, 
			time,
			condtion,
			first_team_id,
			second_team_id,
			result_first_team,
			result_second_team,
			first_team_point,
			second_team_point,
			liga_id, 
			created_at, 
			updated_at 
		from 
			games 
		limit $1 offset $2`, req.Limit, offset,
	)
	if err != nil {
		log.Println("failed to get game")
		return &p.Games{}, err
	}

	for rows.Next() {
		temp := p.GameResponse{}
		err = rows.Scan(
			&temp.Id,
			&temp.Time,
			&temp.Condtion,
			&temp.FirstTeamId,
			&temp.SecondTeamId,
			&temp.ResultFirstTeam,
			&temp.ResultSecondTeam,
			&temp.FirstTeamPoint,
			&temp.SecondTeamPoint,
			&temp.LigaId,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			log.Println("failed to scanning game")
			return &p.Games{}, err
		}

		res.Games = append(res.Games, &temp)
	}

	return &res, nil
}

func (r *Repo) DeleteGame(id *p.IdRequest) (*p.GameResponse, error) {
	game := p.GameResponse{}
	err := r.db.QueryRow(`
		update 
			games 
		set 
			deleted_at = $1 
		where 
			id = $2
		returning 
			id, 
			time,
			condtion,
			first_team_id,
			second_team_id,
			result_first_team,
			result_second_team,
			first_team_point,
			second_team_point,
			liga_id, 
			created_at, 
			updated_at `, time.Now(), id.Id).
		Scan(
			&game.Id,
			&game.Time,
			&game.Condtion,
			&game.FirstTeamId,
			&game.SecondTeamId,
			&game.ResultFirstTeam,
			&game.ResultSecondTeam,
			&game.FirstTeamPoint,
			&game.SecondTeamPoint,
			&game.LigaId,
			&game.CreatedAt,
			&game.UpdatedAt,
		)

	if err != nil {
		log.Println("failed to delete game")
		return &p.GameResponse{}, err
	}

	return &game, nil
}
