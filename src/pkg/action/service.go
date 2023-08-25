// Copyright 2022 The ILLA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package action

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/illacloud/builder-backend/internal/idconvertor"
	"github.com/illacloud/builder-backend/internal/repository"
	"github.com/illacloud/builder-backend/internal/util/resourcelist"

	"go.uber.org/zap"
)

const (
	FIELD_VIRTUAL_RESOURCE = "virtualResource"
)

type ActionService interface {
	IsPublicAction(teamID int, actionID int) bool
	CreateAction(action ActionDto) (*ActionDtoForExport, error)
	DeleteAction(teamID int, id int) error
	UpdateAction(action ActionDto) (*ActionDtoForExport, error)
	UpdatePublic(teamID int, appID int, userID int, actionConfig *repository.ActionConfig) error
	GetAction(teamID int, id int) (*ActionDtoForExport, error)
	FindActionsByAppVersion(teamID int, app, version int) ([]*ActionDtoForExport, error)
	RunAction(teamID int, action ActionDto, actionRuntimeInfo *repository.ActionRuntimeInfo) (interface{}, error)
	ValidateActionOptions(actionType string, options map[string]interface{}) error
}

type ActionDto struct {
	ID          int                      `json:"actionId"`
	UID         uuid.UUID                `json:"uid"`
	TeamID      int                      `json:"teamID"`
	App         int                      `json:"-"`
	Version     int                      `json:"-"`
	Resource    int                      `json:"resourceId,omitempty"`
	DisplayName string                   `json:"displayName" validate:"required"`
	Type        string                   `json:"actionType" validate:"required"`
	Template    map[string]interface{}   `json:"content" validate:"required"`
	Transformer map[string]interface{}   `json:"transformer" validate:"required"`
	TriggerMode string                   `json:"triggerMode" validate:"oneof=manually automate"`
	Config      *repository.ActionConfig `json:"config"`
	CreatedAt   time.Time                `json:"createdAt,omitempty"`
	CreatedBy   int                      `json:"createdBy,omitempty"`
	UpdatedAt   time.Time                `json:"updatedAt,omitempty"`
	UpdatedBy   int                      `json:"updatedBy,omitempty"`
}

type ActionDtoForExport struct {
	ID                string                   `json:"actionId"`
	UID               uuid.UUID                `json:"uid"`
	TeamID            string                   `json:"teamID"`
	App               string                   `json:"-"`
	Version           int                      `json:"-"`
	Resource          string                   `json:"resourceId,omitempty"`
	DisplayName       string                   `json:"displayName" validate:"required"`
	Type              string                   `json:"actionType" validate:"required"`
	IsVirtualResource bool                     `json:"isVirtualResource"`
	Template          map[string]interface{}   `json:"content" validate:"required"`
	Transformer       map[string]interface{}   `json:"transformer" validate:"required"`
	TriggerMode       string                   `json:"triggerMode" validate:"oneof=manually automate"`
	Config            *repository.ActionConfig `json:"config"`
	CreatedAt         time.Time                `json:"createdAt,omitempty"`
	CreatedBy         string                   `json:"createdBy,omitempty"`
	UpdatedAt         time.Time                `json:"updatedAt,omitempty"`
	UpdatedBy         string                   `json:"updatedBy,omitempty"`
}

func NewActionDtoForExport(a *ActionDto) *ActionDtoForExport {
	return &ActionDtoForExport{
		ID:                idconvertor.ConvertIntToString(a.ID),
		UID:               a.UID,
		TeamID:            idconvertor.ConvertIntToString(a.TeamID),
		App:               idconvertor.ConvertIntToString(a.App),
		Version:           a.Version,
		Resource:          idconvertor.ConvertIntToString(a.Resource),
		DisplayName:       a.DisplayName,
		Type:              a.Type,
		IsVirtualResource: resourcelist.IsVirtualResource(a.Type),
		Template:          a.Template,
		Transformer:       a.Transformer,
		TriggerMode:       a.TriggerMode,
		Config:            a.Config,
		CreatedAt:         a.CreatedAt,
		CreatedBy:         idconvertor.ConvertIntToString(a.CreatedBy),
		UpdatedAt:         a.UpdatedAt,
		UpdatedBy:         idconvertor.ConvertIntToString(a.UpdatedBy),
	}
}

func (resp *ActionDtoForExport) ExportActionDto() ActionDto {
	actionDto := ActionDto{
		UID:         resp.UID,
		DisplayName: resp.DisplayName,
		Type:        resp.Type,
		Template:    resp.Template,
		Transformer: resp.Transformer,
		TriggerMode: resp.TriggerMode,
		Config:      resp.Config,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
	}
	// fill converted fields
	if resp.App != "" {
		actionDto.App = idconvertor.ConvertStringToInt(resp.App)
	}
	if resp.ID != "" {
		actionDto.ID = idconvertor.ConvertStringToInt(resp.ID)
	}
	if resp.TeamID != "" {
		actionDto.TeamID = idconvertor.ConvertStringToInt(resp.TeamID)
	}
	if resp.Resource != "" {
		actionDto.Resource = idconvertor.ConvertStringToInt(resp.Resource)
	}
	if resp.CreatedBy != "" {
		actionDto.CreatedBy = idconvertor.ConvertStringToInt(resp.CreatedBy)
	}
	if resp.UpdatedBy != "" {
		actionDto.UpdatedBy = idconvertor.ConvertStringToInt(resp.UpdatedBy)
	}
	if resp.Config == nil {
		resp.Config = repository.NewActionConfig()
	}
	return actionDto
}

func (resp *ActionDtoForExport) ExportID() string {
	return resp.ID
}

func (resp *ActionDtoForExport) ExportResourceID() string {
	return resp.Resource
}

func (resp *ActionDtoForExport) ExportResourceIDInInt() int {
	return idconvertor.ConvertStringToInt(resp.Resource)
}

func (resp *ActionDtoForExport) ExportTeamID() string {
	return resp.TeamID
}

func (resp *ActionDtoForExport) ExportType() string {
	return resp.Type
}

func (resp *ActionDtoForExport) ExportCreatedAt() time.Time {
	return resp.CreatedAt
}

func (resp *ActionDtoForExport) ExportUpdatedAt() time.Time {
	return resp.UpdatedAt
}

func (resp *ActionDtoForExport) ExportCreatedBy() string {
	return resp.CreatedBy
}

func (resp *ActionDtoForExport) ExportUpdatedBy() string {
	return resp.UpdatedBy
}

func (resp *ActionDtoForExport) ExportForFeedback() interface{} {
	return resp
}

func (resp *ActionDtoForExport) AppendAdditionalTemplate(field string, value interface{}) {
	resp.Template[field] = value
}

func (resp *ActionDtoForExport) AppendVirtualResourceToTemplate(value interface{}) {
	resp.Template[FIELD_VIRTUAL_RESOURCE] = value
}

func (a *ActionDto) InitUID() {
	a.UID = uuid.New()
}

func (a *ActionDto) SetTeamID(teamID int) {
	a.TeamID = teamID
}

func (a *ActionDto) SetPublicStatus(isPublic bool) {
	if a.Config == nil {
		a.Config = repository.NewActionConfig()
	}
	if isPublic {
		a.Config.SetPublic()
	} else {
		a.Config.SetPrivate()
	}
}

func (a *ActionDto) ExportActionType() string {
	return a.Type
}

type ActionServiceImpl struct {
	logger             *zap.SugaredLogger
	appRepository      repository.AppRepository
	actionRepository   repository.ActionRepository
	resourceRepository repository.ResourceRepository
}

func NewActionServiceImpl(logger *zap.SugaredLogger, appRepository repository.AppRepository,
	actionRepository repository.ActionRepository, resourceRepository repository.ResourceRepository) *ActionServiceImpl {
	return &ActionServiceImpl{
		logger:             logger,
		appRepository:      appRepository,
		actionRepository:   actionRepository,
		resourceRepository: resourceRepository,
	}
}

func (impl *ActionServiceImpl) IsPublicAction(teamID int, actionID int) bool {
	action, err := impl.actionRepository.RetrieveActionByIDAndTeamID(actionID, teamID)
	if err != nil {
		return false
	}
	return action.IsPublic()
}

func (impl *ActionServiceImpl) CreateAction(action ActionDto) (*ActionDtoForExport, error) {
	// validate app
	if appDto, err := impl.appRepository.RetrieveAppByIDAndTeamID(action.App, action.TeamID); err != nil || appDto.ID != action.App {
		return nil, errors.New("app not found")
	}
	// validate resource
	if !resourcelist.IsVirtualResource(action.Type) {
		if rscDto, err := impl.resourceRepository.RetrieveByID(action.TeamID, action.Resource); err != nil || rscDto.ID != action.Resource {
			return nil, errors.New("resource not found")
		}
	}

	id, err := impl.actionRepository.Create(&repository.Action{
		ID:          action.ID,
		UID:         action.UID,
		TeamID:      action.TeamID,
		App:         action.App,
		Version:     action.Version,
		Resource:    action.Resource,
		Name:        action.DisplayName,
		Type:        resourcelist.GetResourceNameMappedID(action.Type),
		TriggerMode: action.TriggerMode,
		Transformer: action.Transformer,
		Template:    action.Template,
		Config:      action.Config.ExportToJSONString(),
		CreatedAt:   action.CreatedAt,
		CreatedBy:   action.CreatedBy,
		UpdatedAt:   action.UpdatedAt,
		UpdatedBy:   action.UpdatedBy,
	})
	if err != nil {
		return nil, err
	}
	action.ID = id

	// update app `updatedAt` field
	_ = impl.appRepository.UpdateUpdatedAt(&repository.App{
		ID:        action.App,
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: action.UpdatedBy,
	})

	return NewActionDtoForExport(&action), nil
}

func (impl *ActionServiceImpl) DeleteAction(teamID int, actionID int) error {
	action, _ := impl.actionRepository.RetrieveByID(teamID, actionID)

	if err := impl.actionRepository.Delete(teamID, actionID); err != nil {
		return err
	}

	// update app `updatedAt` field
	_ = impl.appRepository.UpdateUpdatedAt(&repository.App{
		ID:        action.App,
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: action.UpdatedBy,
	})

	return nil
}

func (impl *ActionServiceImpl) UpdateAction(action ActionDto) (*ActionDtoForExport, error) {
	// validate app
	if appDto, err := impl.appRepository.RetrieveAppByIDAndTeamID(action.App, action.TeamID); err != nil || appDto.ID != action.App {
		return nil, errors.New("app not found")
	}
	// validate resource
	if !resourcelist.IsVirtualResource(action.Type) {
		if rscDto, err := impl.resourceRepository.RetrieveByID(action.TeamID, action.Resource); err != nil || rscDto.ID != action.Resource {
			return nil, errors.New("resource not found")
		}
	}

	if err := impl.actionRepository.Update(&repository.Action{
		ID:          action.ID,
		UID:         action.UID,
		TeamID:      action.TeamID,
		Resource:    action.Resource,
		Name:        action.DisplayName,
		Type:        resourcelist.GetResourceNameMappedID(action.Type),
		TriggerMode: action.TriggerMode,
		Transformer: action.Transformer,
		Template:    action.Template,
		Config:      action.Config.ExportToJSONString(),
		UpdatedAt:   action.UpdatedAt,
		UpdatedBy:   action.UpdatedBy,
	}); err != nil {
		return nil, err
	}

	// update app `updatedAt` field
	_ = impl.appRepository.UpdateUpdatedAt(&repository.App{
		ID:        action.App,
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: action.UpdatedBy,
	})

	return NewActionDtoForExport(&action), nil
}

func (impl *ActionServiceImpl) UpdatePublic(teamID int, appID int, userID int, actionConfig *repository.ActionConfig) error {
	return impl.actionRepository.UpdatePublicByTeamIDAndAppIDAndUserID(teamID, appID, userID, actionConfig)
}

func (impl *ActionServiceImpl) GetAction(teamID int, actionID int) (*ActionDtoForExport, error) {
	res, err := impl.actionRepository.RetrieveByID(teamID, actionID)
	if err != nil {
		return nil, err
	}
	action := ActionDto{
		ID:          res.ID,
		UID:         res.UID,
		TeamID:      res.TeamID,
		Resource:    res.Resource,
		DisplayName: res.Name,
		Type:        resourcelist.GetResourceIDMappedType(res.Type),
		TriggerMode: res.TriggerMode,
		Transformer: res.Transformer,
		Template:    res.Template,
		Config:      res.ExportConfig(),
		CreatedBy:   res.CreatedBy,
		CreatedAt:   res.CreatedAt,
		UpdatedBy:   res.UpdatedBy,
		UpdatedAt:   res.UpdatedAt,
	}
	return NewActionDtoForExport(&action), nil

}

func (impl *ActionServiceImpl) FindActionsByAppVersion(teamID int, appID int, version int) ([]*ActionDtoForExport, error) {
	res, err := impl.actionRepository.RetrieveActionsByAppVersion(teamID, appID, version)
	if err != nil {
		return nil, err
	}

	actionDtoForExportSlice := make([]*ActionDtoForExport, 0, len(res))
	for _, value := range res {
		actionDto := ActionDto{
			ID:          value.ID,
			UID:         value.UID,
			TeamID:      value.TeamID,
			Resource:    value.Resource,
			DisplayName: value.Name,
			Type:        resourcelist.GetResourceIDMappedType(value.Type),
			TriggerMode: value.TriggerMode,
			Transformer: value.Transformer,
			Template:    value.Template,
			Config:      value.ExportConfig(),
			CreatedBy:   value.CreatedBy,
			CreatedAt:   value.CreatedAt,
			UpdatedBy:   value.UpdatedBy,
			UpdatedAt:   value.UpdatedAt,
		}
		actionDtoForExportSlice = append(actionDtoForExportSlice, NewActionDtoForExport(&actionDto))
	}
	return actionDtoForExportSlice, nil
}

func (impl *ActionServiceImpl) RunAction(teamID int, action ActionDto, actionRuntimeInfo *repository.ActionRuntimeInfo) (interface{}, error) {
	if action.Resource == 0 {
		return nil, errors.New("no resource")
	}
	actionFactory := Factory{Type: action.Type}
	actionAssemblyLine := actionFactory.Build()
	if actionAssemblyLine == nil {
		return nil, errors.New("invalid ActionType:: unsupported type")
	}

	rsc := &repository.Resource{}
	if !resourcelist.IsVirtualResource(action.ExportActionType()) {
		// process non-virtual resource
		var errInRetrieveResource error
		rsc, errInRetrieveResource = impl.resourceRepository.RetrieveByID(teamID, action.Resource)
		if rsc.ID == 0 {
			return nil, errors.New("resource not found")
		}
		if errInRetrieveResource != nil {
			return nil, errInRetrieveResource
		}
		if _, err := actionAssemblyLine.ValidateResourceOptions(rsc.Options); err != nil {
			return nil, errors.New("invalid resource content")
		}
	} else {
		// process virtual resource runtime info
		action.Template = actionRuntimeInfo.AppendToActionTemplate(action.Template)
	}
	if _, err := actionAssemblyLine.ValidateActionOptions(action.Template); err != nil {
		return nil, errors.New("invalid action content")
	}
	res, err := actionAssemblyLine.Run(rsc.Options, action.Template)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (impl *ActionServiceImpl) ValidateActionOptions(actionType string, options map[string]interface{}) error {
	if actionType == resourcelist.TYPE_TRANSFORMER {
		return nil
	}
	actionFactory := Factory{Type: actionType}
	actionAssemblyLine := actionFactory.Build()
	if actionAssemblyLine == nil {
		return errors.New("invalid ActionType:: unsupported type")
	}
	if _, err := actionAssemblyLine.ValidateActionOptions(options); err != nil {
		return errors.New("invalid action content")
	}
	return nil
}