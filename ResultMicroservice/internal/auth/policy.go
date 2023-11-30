package auth

import (
	"ResultMicroservice/graph/model"
	"ResultMicroservice/internal/repository"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type IResultPolicy interface {
	CreateResult(bearerToken string) error
	UpdateResult(bearerToken string, id string) (*model.Result, error)
	DeleteResult(bearerToken string, id string) error
	GetResultByExercise(bearerToken string, exerciseID string) error
	GetResultsByClass(bearerToken string, classID string) error
	GetResultsByUser(bearerToken string, userID string) error
	GetResultByID(bearerToken string, id string) error
	DeleteResultByClass(bearerToken string, classID string) (*string, error)
}

type ResultPolicy struct {
	Token            IToken
	ResultRepository repository.IResultRepository
}

func NewResultPolicy(collection *mongo.Collection) IResultPolicy {
	token := NewToken()

	return &ResultPolicy{
		Token:            token,
		ResultRepository: repository.NewResultRepository(collection),
	}
}

func (p *ResultPolicy) CreateResult(bearerToken string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "create_result") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

func (p *ResultPolicy) UpdateResult(bearerToken string, id string) (*model.Result, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	result, err := p.ResultRepository.GetResultByID(id)
	if err != nil {
		return nil, errors.New("invalid permissions for this action")
	}

	if p.hasRole(roles, "update_result") && result.UserID == uuid {
		return result, nil
	}

	if p.hasRole(roles, "update_result_all") {
		return result, nil
	}

	return nil, errors.New("invalid permissions for this action")
}

func (p *ResultPolicy) DeleteResult(bearerToken string, id string) error {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	result, err := p.ResultRepository.GetResultByID(id)
	if err != nil {
		return errors.New("invalid permissions for this action")
	}

	if p.hasRole(roles, "delete_result") && result.UserID == uuid {
		return nil
	}

	if p.hasRole(roles, "delete_result_all") {
		return nil
	}

	return errors.New("invalid permissions for this action")
}

func (p *ResultPolicy) GetResultByExercise(bearerToken string, exerciseID string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "get_result_by_exercise") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

func (p *ResultPolicy) GetResultsByClass(bearerToken string, classID string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "get_results_by_class") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

func (p *ResultPolicy) GetResultsByUser(bearerToken string, userID string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "get_results_by_user") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

func (p *ResultPolicy) GetResultByID(bearerToken string, id string) error {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return err2
	}

	if !p.hasRole(roles, "get_result_by_id") {
		return errors.New("invalid permissions for this action")
	}

	return nil
}

func (p *ResultPolicy) DeleteResultByClass(bearerToken string, classID string) (*string, error) {
	uuid, roles, err := p.getSubAndRoles(bearerToken)
	if err != nil {
		return nil, err
	}

	if !p.hasRole(roles, "delete_result_by_class") {
		return nil, errors.New("invalid permissions for this action")
	}

	return &uuid, nil
}

func (p *ResultPolicy) getSubAndRoles(bearerToken string) (string, []interface{}, error) {
	token, err := p.Token.IntrospectToken(bearerToken)
	if err != nil || token == false {
		return "", nil, errors.New("invalid token")
	}

	decodeToken, err := p.Token.DecodeToken(bearerToken)
	if err != nil {
		return "", nil, err
	}

	sub, _ := decodeToken["sub"].(string)

	resourceAccess, ok := decodeToken["resource_access"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New("invalid token")
	}

	userManagementClient, ok := resourceAccess["user-management-client"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New("invalid token")
	}

	roles, ok := userManagementClient["roles"].([]interface{})
	if !ok {
		return "", nil, errors.New("invalid token")
	}
	return sub, roles, nil
}

func (p *ResultPolicy) hasRole(roles []interface{}, targetRole string) bool {
	for _, role := range roles {
		if role == targetRole {
			return true
		}
	}
	return false
}
