package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/jobutterfly/olives/consts"
	"github.com/jobutterfly/olives/sqlc"
)

func TestGetPost(t *testing.T) {
	if err := Start(); err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	postId := int32(9)
	post, err := Th.q.GetPost(context.Background(), postId)
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	noRowMsg, err := json.Marshal(consts.ErrorMessage{
		Msg: sql.ErrNoRows.Error(),
	})
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	notAnInt, err := json.Marshal(consts.ErrorMessage{
		Msg: consts.PathNotAnInteger.Error(),
	})
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	testCases := []GetTestCase{
		{
			Name:         "successful get post request",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/"+strconv.Itoa(int(postId)), nil),
			ExpectedRes:  jsonPost,
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "failed request for non existing post",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/"+strconv.Itoa(1000000), nil),
			ExpectedRes:  noRowMsg,
			ExpectedCode: http.StatusNotFound,
		},
		{
			Name:         "failed request for wrong path",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/banana", nil),
			ExpectedRes:  notAnInt,
			ExpectedCode: http.StatusBadRequest,
		},
	}

	TestGet(t, testCases, Th.GetPost)
}

func TestGetSubolivePosts(t *testing.T) {
	suboliveId := int32(2)
	firstPosts, err := Th.q.GetSubolivePosts(context.Background(), sqlc.GetSubolivePostsParams{
		Offset: 0,
		SuboliveID: suboliveId,
	})
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	firstJsonPosts, err := json.Marshal(firstPosts)
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	secondPosts, err := Th.q.GetSubolivePosts(context.Background(), sqlc.GetSubolivePostsParams{
		Offset: 10,
		SuboliveID: suboliveId,
	})
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	secondJsonPosts, err := json.Marshal(secondPosts)
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	noRowMsg, err := json.Marshal(consts.ErrorMessage{
		Msg: sql.ErrNoRows.Error(),
	})

	notAnInt, err := json.Marshal(consts.ErrorMessage{
		Msg: consts.PathNotAnInteger.Error(),
	})
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	pageNotAnInt, err := json.Marshal(consts.ErrorMessage{
		Msg: consts.PageNotAnInteger.Error(),
	})
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	suboliveNonExistant, err := json.Marshal(consts.ErrorMessage{
		Msg: consts.SuboliveNonExistant.Error(),
	})
	if err != nil {
		t.Errorf("expected no errors, got %v", err)
		return
	}

	testCases := []GetTestCase{
		{
			Name:         "successful get posts request",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/subolive/"+strconv.Itoa(int(suboliveId))+"?page=0", nil),
			ExpectedRes:  firstJsonPosts,
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "successful get posts request second page",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/subolive/"+strconv.Itoa(int(suboliveId))+"?page=1", nil),
			ExpectedRes:  secondJsonPosts,
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "unsuccessful get posts request subolive id not number",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/subolive/banana?page=1", nil),
			ExpectedRes:  notAnInt,
			ExpectedCode: http.StatusBadRequest,
		},
		{
			Name:         "unsuccessful get posts request subolive id not exist",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/subolive/"+strconv.Itoa(int(10000))+"?page=1", nil),
			ExpectedRes:  suboliveNonExistant,
			ExpectedCode: http.StatusNotFound,
		},
		{
			Name:         "unsuccessful get posts request not existant page",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/subolive/"+strconv.Itoa(int(suboliveId))+"?page=1000", nil),
			ExpectedRes:  noRowMsg,
			ExpectedCode: http.StatusNotFound,
		},
		{
			Name:         "unsuccessful get posts request page not int",
			Req:          httptest.NewRequest(http.MethodGet, "/posts/subolive/"+strconv.Itoa(int(suboliveId))+"?page=banana", nil),
			ExpectedRes:  pageNotAnInt,
			ExpectedCode: http.StatusInternalServerError,
		},
	}

	TestGet(t, testCases, Th.GetSubolivePosts)
}


















