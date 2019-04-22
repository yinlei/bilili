// Code generated by MockGen. DO NOT EDIT.
// Source: article.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "go-common/app/interface/openplatform/article/model"
	reflect "reflect"
)

// MockArticleRPC is a mock of ArticleRPC interface
type MockArticleRPC struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRPCMockRecorder
}

// MockArticleRPCMockRecorder is the mock recorder for MockArticleRPC
type MockArticleRPCMockRecorder struct {
	mock *MockArticleRPC
}

// NewMockArticleRPC creates a new mock instance
func NewMockArticleRPC(ctrl *gomock.Controller) *MockArticleRPC {
	mock := &MockArticleRPC{ctrl: ctrl}
	mock.recorder = &MockArticleRPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticleRPC) EXPECT() *MockArticleRPCMockRecorder {
	return m.recorder
}

// AddArticle mocks base method
func (m *MockArticleRPC) AddArticle(c context.Context, arg *model.ArgArticle) (int64, error) {
	ret := m.ctrl.Call(m, "AddArticle", c, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddArticle indicates an expected call of AddArticle
func (mr *MockArticleRPCMockRecorder) AddArticle(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddArticle", reflect.TypeOf((*MockArticleRPC)(nil).AddArticle), c, arg)
}

// AddArticleCache mocks base method
func (m *MockArticleRPC) AddArticleCache(c context.Context, arg *model.ArgAid) error {
	ret := m.ctrl.Call(m, "AddArticleCache", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddArticleCache indicates an expected call of AddArticleCache
func (mr *MockArticleRPCMockRecorder) AddArticleCache(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddArticleCache", reflect.TypeOf((*MockArticleRPC)(nil).AddArticleCache), c, arg)
}

// UpdateArticleCache mocks base method
func (m *MockArticleRPC) UpdateArticleCache(c context.Context, arg *model.ArgAidCid) error {
	ret := m.ctrl.Call(m, "UpdateArticleCache", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArticleCache indicates an expected call of UpdateArticleCache
func (mr *MockArticleRPCMockRecorder) UpdateArticleCache(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticleCache", reflect.TypeOf((*MockArticleRPC)(nil).UpdateArticleCache), c, arg)
}

// DelArticleCache mocks base method
func (m *MockArticleRPC) DelArticleCache(c context.Context, arg *model.ArgAidMid) error {
	ret := m.ctrl.Call(m, "DelArticleCache", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelArticleCache indicates an expected call of DelArticleCache
func (mr *MockArticleRPCMockRecorder) DelArticleCache(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelArticleCache", reflect.TypeOf((*MockArticleRPC)(nil).DelArticleCache), c, arg)
}

// UpdateArticle mocks base method
func (m *MockArticleRPC) UpdateArticle(c context.Context, arg *model.ArgArticle) error {
	ret := m.ctrl.Call(m, "UpdateArticle", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArticle indicates an expected call of UpdateArticle
func (mr *MockArticleRPCMockRecorder) UpdateArticle(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockArticleRPC)(nil).UpdateArticle), c, arg)
}

// CreationWithdrawArticle mocks base method
func (m *MockArticleRPC) CreationWithdrawArticle(c context.Context, arg *model.ArgAidMid) error {
	ret := m.ctrl.Call(m, "CreationWithdrawArticle", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreationWithdrawArticle indicates an expected call of CreationWithdrawArticle
func (mr *MockArticleRPCMockRecorder) CreationWithdrawArticle(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreationWithdrawArticle", reflect.TypeOf((*MockArticleRPC)(nil).CreationWithdrawArticle), c, arg)
}

// DelArticle mocks base method
func (m *MockArticleRPC) DelArticle(c context.Context, arg *model.ArgAidMid) error {
	ret := m.ctrl.Call(m, "DelArticle", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelArticle indicates an expected call of DelArticle
func (mr *MockArticleRPCMockRecorder) DelArticle(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelArticle", reflect.TypeOf((*MockArticleRPC)(nil).DelArticle), c, arg)
}

// CreationArticle mocks base method
func (m *MockArticleRPC) CreationArticle(c context.Context, arg *model.ArgAidMid) (*model.Article, error) {
	ret := m.ctrl.Call(m, "CreationArticle", c, arg)
	ret0, _ := ret[0].(*model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreationArticle indicates an expected call of CreationArticle
func (mr *MockArticleRPCMockRecorder) CreationArticle(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreationArticle", reflect.TypeOf((*MockArticleRPC)(nil).CreationArticle), c, arg)
}

// CreationUpperArticles mocks base method
func (m *MockArticleRPC) CreationUpperArticles(c context.Context, arg *model.ArgCreationArts) (*model.CreationArts, error) {
	ret := m.ctrl.Call(m, "CreationUpperArticles", c, arg)
	ret0, _ := ret[0].(*model.CreationArts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreationUpperArticles indicates an expected call of CreationUpperArticles
func (mr *MockArticleRPCMockRecorder) CreationUpperArticles(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreationUpperArticles", reflect.TypeOf((*MockArticleRPC)(nil).CreationUpperArticles), c, arg)
}

// Categories mocks base method
func (m *MockArticleRPC) Categories(c context.Context, arg *model.ArgIP) (*model.Categories, error) {
	ret := m.ctrl.Call(m, "Categories", c, arg)
	ret0, _ := ret[0].(*model.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Categories indicates an expected call of Categories
func (mr *MockArticleRPCMockRecorder) Categories(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Categories", reflect.TypeOf((*MockArticleRPC)(nil).Categories), c, arg)
}

// CategoriesMap mocks base method
func (m *MockArticleRPC) CategoriesMap(c context.Context, arg *model.ArgIP) (map[int64]*model.Category, error) {
	ret := m.ctrl.Call(m, "CategoriesMap", c, arg)
	ret0, _ := ret[0].(map[int64]*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CategoriesMap indicates an expected call of CategoriesMap
func (mr *MockArticleRPCMockRecorder) CategoriesMap(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CategoriesMap", reflect.TypeOf((*MockArticleRPC)(nil).CategoriesMap), c, arg)
}

// SetStat mocks base method
func (m *MockArticleRPC) SetStat(c context.Context, arg *model.ArgStats) error {
	ret := m.ctrl.Call(m, "SetStat", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStat indicates an expected call of SetStat
func (mr *MockArticleRPCMockRecorder) SetStat(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStat", reflect.TypeOf((*MockArticleRPC)(nil).SetStat), c, arg)
}

// UpsArtMetas mocks base method
func (m *MockArticleRPC) UpsArtMetas(c context.Context, arg *model.ArgUpsArts) (map[int64][]*model.Meta, error) {
	ret := m.ctrl.Call(m, "UpsArtMetas", c, arg)
	ret0, _ := ret[0].(map[int64][]*model.Meta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsArtMetas indicates an expected call of UpsArtMetas
func (mr *MockArticleRPCMockRecorder) UpsArtMetas(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsArtMetas", reflect.TypeOf((*MockArticleRPC)(nil).UpsArtMetas), c, arg)
}

// ArticleMetas mocks base method
func (m *MockArticleRPC) ArticleMetas(c context.Context, arg *model.ArgAids) (map[int64]*model.Meta, error) {
	ret := m.ctrl.Call(m, "ArticleMetas", c, arg)
	ret0, _ := ret[0].(map[int64]*model.Meta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticleMetas indicates an expected call of ArticleMetas
func (mr *MockArticleRPCMockRecorder) ArticleMetas(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticleMetas", reflect.TypeOf((*MockArticleRPC)(nil).ArticleMetas), c, arg)
}

// UpdateRecommends mocks base method
func (m *MockArticleRPC) UpdateRecommends(c context.Context) error {
	ret := m.ctrl.Call(m, "UpdateRecommends", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRecommends indicates an expected call of UpdateRecommends
func (mr *MockArticleRPCMockRecorder) UpdateRecommends(c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecommends", reflect.TypeOf((*MockArticleRPC)(nil).UpdateRecommends), c)
}

// Recommends mocks base method
func (m *MockArticleRPC) Recommends(c context.Context, arg *model.ArgRecommends) ([]*model.Meta, error) {
	ret := m.ctrl.Call(m, "Recommends", c, arg)
	ret0, _ := ret[0].([]*model.Meta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recommends indicates an expected call of Recommends
func (mr *MockArticleRPCMockRecorder) Recommends(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recommends", reflect.TypeOf((*MockArticleRPC)(nil).Recommends), c, arg)
}

// UpArtMetas mocks base method
func (m *MockArticleRPC) UpArtMetas(c context.Context, arg *model.ArgUpArts) (*model.UpArtMetas, error) {
	ret := m.ctrl.Call(m, "UpArtMetas", c, arg)
	ret0, _ := ret[0].(*model.UpArtMetas)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpArtMetas indicates an expected call of UpArtMetas
func (mr *MockArticleRPCMockRecorder) UpArtMetas(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpArtMetas", reflect.TypeOf((*MockArticleRPC)(nil).UpArtMetas), c, arg)
}

// AddArtDraft mocks base method
func (m *MockArticleRPC) AddArtDraft(c context.Context, arg *model.ArgArticle) (int64, error) {
	ret := m.ctrl.Call(m, "AddArtDraft", c, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddArtDraft indicates an expected call of AddArtDraft
func (mr *MockArticleRPCMockRecorder) AddArtDraft(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddArtDraft", reflect.TypeOf((*MockArticleRPC)(nil).AddArtDraft), c, arg)
}

// UpdateArtDraft mocks base method
func (m *MockArticleRPC) UpdateArtDraft(c context.Context, arg *model.ArgAidMid) error {
	ret := m.ctrl.Call(m, "UpdateArtDraft", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArtDraft indicates an expected call of UpdateArtDraft
func (mr *MockArticleRPCMockRecorder) UpdateArtDraft(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArtDraft", reflect.TypeOf((*MockArticleRPC)(nil).UpdateArtDraft), c, arg)
}

// DelArtDraft mocks base method
func (m *MockArticleRPC) DelArtDraft(c context.Context, arg *model.ArgAidMid) error {
	ret := m.ctrl.Call(m, "DelArtDraft", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelArtDraft indicates an expected call of DelArtDraft
func (mr *MockArticleRPCMockRecorder) DelArtDraft(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelArtDraft", reflect.TypeOf((*MockArticleRPC)(nil).DelArtDraft), c, arg)
}

// ArtDraft mocks base method
func (m *MockArticleRPC) ArtDraft(c context.Context, arg *model.ArgAidMid) (*model.Draft, error) {
	ret := m.ctrl.Call(m, "ArtDraft", c, arg)
	ret0, _ := ret[0].(*model.Draft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArtDraft indicates an expected call of ArtDraft
func (mr *MockArticleRPCMockRecorder) ArtDraft(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArtDraft", reflect.TypeOf((*MockArticleRPC)(nil).ArtDraft), c, arg)
}

// UpperDrafts mocks base method
func (m *MockArticleRPC) UpperDrafts(c context.Context, arg *model.ArgUpDraft) (*model.Drafts, error) {
	ret := m.ctrl.Call(m, "UpperDrafts", c, arg)
	ret0, _ := ret[0].(*model.Drafts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpperDrafts indicates an expected call of UpperDrafts
func (mr *MockArticleRPCMockRecorder) UpperDrafts(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpperDrafts", reflect.TypeOf((*MockArticleRPC)(nil).UpperDrafts), c, arg)
}

// UpdateNewArts mocks base method
func (m *MockArticleRPC) UpdateNewArts(c context.Context, arg *model.ArgIP) error {
	ret := m.ctrl.Call(m, "UpdateNewArts", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNewArts indicates an expected call of UpdateNewArts
func (mr *MockArticleRPCMockRecorder) UpdateNewArts(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNewArts", reflect.TypeOf((*MockArticleRPC)(nil).UpdateNewArts), c, arg)
}

// ArticleRemainCount mocks base method
func (m *MockArticleRPC) ArticleRemainCount(c context.Context, arg *model.ArgMid) (int, error) {
	ret := m.ctrl.Call(m, "ArticleRemainCount", c, arg)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticleRemainCount indicates an expected call of ArticleRemainCount
func (mr *MockArticleRPCMockRecorder) ArticleRemainCount(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticleRemainCount", reflect.TypeOf((*MockArticleRPC)(nil).ArticleRemainCount), c, arg)
}

// DelRecommendArtCache mocks base method
func (m *MockArticleRPC) DelRecommendArtCache(c context.Context, arg *model.ArgAidCid) error {
	ret := m.ctrl.Call(m, "DelRecommendArtCache", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelRecommendArtCache indicates an expected call of DelRecommendArtCache
func (mr *MockArticleRPCMockRecorder) DelRecommendArtCache(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelRecommendArtCache", reflect.TypeOf((*MockArticleRPC)(nil).DelRecommendArtCache), c, arg)
}

// CheckPrivilege mocks base method
func (m *MockArticleRPC) CheckPrivilege(c context.Context, arg *model.ArgMid) error {
	ret := m.ctrl.Call(m, "CheckPrivilege", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPrivilege indicates an expected call of CheckPrivilege
func (mr *MockArticleRPCMockRecorder) CheckPrivilege(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPrivilege", reflect.TypeOf((*MockArticleRPC)(nil).CheckPrivilege), c, arg)
}

// Favorites mocks base method
func (m *MockArticleRPC) Favorites(c context.Context, arg *model.ArgFav) ([]*model.Favorite, error) {
	ret := m.ctrl.Call(m, "Favorites", c, arg)
	ret0, _ := ret[0].([]*model.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Favorites indicates an expected call of Favorites
func (mr *MockArticleRPCMockRecorder) Favorites(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Favorites", reflect.TypeOf((*MockArticleRPC)(nil).Favorites), c, arg)
}

// UpdateAuthorCache mocks base method
func (m *MockArticleRPC) UpdateAuthorCache(c context.Context, arg *model.ArgAuthor) error {
	ret := m.ctrl.Call(m, "UpdateAuthorCache", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAuthorCache indicates an expected call of UpdateAuthorCache
func (mr *MockArticleRPCMockRecorder) UpdateAuthorCache(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthorCache", reflect.TypeOf((*MockArticleRPC)(nil).UpdateAuthorCache), c, arg)
}

// UpdateSortCache mocks base method
func (m *MockArticleRPC) UpdateSortCache(c context.Context, arg *model.ArgSort) error {
	ret := m.ctrl.Call(m, "UpdateSortCache", c, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSortCache indicates an expected call of UpdateSortCache
func (mr *MockArticleRPCMockRecorder) UpdateSortCache(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSortCache", reflect.TypeOf((*MockArticleRPC)(nil).UpdateSortCache), c, arg)
}

// IsAuthor mocks base method
func (m *MockArticleRPC) IsAuthor(c context.Context, arg *model.ArgMid) (bool, error) {
	ret := m.ctrl.Call(m, "IsAuthor", c, arg)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthor indicates an expected call of IsAuthor
func (mr *MockArticleRPCMockRecorder) IsAuthor(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthor", reflect.TypeOf((*MockArticleRPC)(nil).IsAuthor), c, arg)
}

// NewArticleCount mocks base method
func (m *MockArticleRPC) NewArticleCount(c context.Context, arg *model.ArgNewArt) (int64, error) {
	ret := m.ctrl.Call(m, "NewArticleCount", c, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewArticleCount indicates an expected call of NewArticleCount
func (mr *MockArticleRPCMockRecorder) NewArticleCount(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewArticleCount", reflect.TypeOf((*MockArticleRPC)(nil).NewArticleCount), c, arg)
}

// HadLikesByMid mocks base method
func (m *MockArticleRPC) HadLikesByMid(c context.Context, arg *model.ArgMidAids) (map[int64]int, error) {
	ret := m.ctrl.Call(m, "HadLikesByMid", c, arg)
	ret0, _ := ret[0].(map[int64]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HadLikesByMid indicates an expected call of HadLikesByMid
func (mr *MockArticleRPCMockRecorder) HadLikesByMid(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HadLikesByMid", reflect.TypeOf((*MockArticleRPC)(nil).HadLikesByMid), c, arg)
}

// UpMoreArts mocks base method
func (m *MockArticleRPC) UpMoreArts(c context.Context, arg *model.ArgAid) ([]*model.Meta, error) {
	ret := m.ctrl.Call(m, "UpMoreArts", c, arg)
	ret0, _ := ret[0].([]*model.Meta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpMoreArts indicates an expected call of UpMoreArts
func (mr *MockArticleRPCMockRecorder) UpMoreArts(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpMoreArts", reflect.TypeOf((*MockArticleRPC)(nil).UpMoreArts), c, arg)
}

// CreationUpStat mocks base method
func (m *MockArticleRPC) CreationUpStat(c context.Context, arg *model.ArgMid) (model.UpStat, error) {
	ret := m.ctrl.Call(m, "CreationUpStat", c, arg)
	ret0, _ := ret[0].(model.UpStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreationUpStat indicates an expected call of CreationUpStat
func (mr *MockArticleRPCMockRecorder) CreationUpStat(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreationUpStat", reflect.TypeOf((*MockArticleRPC)(nil).CreationUpStat), c, arg)
}

// CreationUpThirtyDayStat mocks base method
func (m *MockArticleRPC) CreationUpThirtyDayStat(c context.Context, arg *model.ArgMid) ([]*model.ThirtyDayArticle, error) {
	ret := m.ctrl.Call(m, "CreationUpThirtyDayStat", c, arg)
	ret0, _ := ret[0].([]*model.ThirtyDayArticle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreationUpThirtyDayStat indicates an expected call of CreationUpThirtyDayStat
func (mr *MockArticleRPCMockRecorder) CreationUpThirtyDayStat(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreationUpThirtyDayStat", reflect.TypeOf((*MockArticleRPC)(nil).CreationUpThirtyDayStat), c, arg)
}

// UpLists mocks base method
func (m *MockArticleRPC) UpLists(c context.Context, arg *model.ArgMid) (model.UpLists, error) {
	ret := m.ctrl.Call(m, "UpLists", c, arg)
	ret0, _ := ret[0].(model.UpLists)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpLists indicates an expected call of UpLists
func (mr *MockArticleRPCMockRecorder) UpLists(c, arg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpLists", reflect.TypeOf((*MockArticleRPC)(nil).UpLists), c, arg)
}
