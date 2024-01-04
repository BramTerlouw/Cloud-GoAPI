package auth

import (
	"ExerciseMicroservice/internal/auth"
	"ExerciseMicroservice/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var adminToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDQzNzg0MDksImlhdCI6MTcwNDM3ODEwOSwianRpIjoiNzRiMjdlOTQtMzI2Ny00YjRlLWFhZDUtZWZiNzliZTA4NjBmIiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiNmMxY2U0NDgtNjcwZi00N2IyLTgzZjctNGQ3NzFiMDE3NzViIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6ImM5YmZjYjkxLTk0OTItNGM1ZS1iN2NjLWRiNzhiNzU0ZjdjYyIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZmlsdGVyX2V4ZXJjaXNlX2RpZmZpY3VsdHkiLCJmaWx0ZXJfc2Nob29sX25hbWUiLCJmaWx0ZXJfZXhlcmNpc2VfbW9kdWxlX2lkIiwiZmlsdGVyX21vZHVsZV9jYXRlZ29yeSIsImRlbGV0ZV9tb2R1bGVfYWxsIiwiY3JlYXRlX2V4ZXJjaXNlIiwiZ2V0X3NjaG9vbCIsImZpbHRlcl9zY2hvb2xfbG9jYXRpb24iLCJmaWx0ZXJfbW9kdWxlX2RpZmZpY3VsdHkiLCJnZXRfbW9kdWxlIiwiZ2V0X21vZHVsZXMiLCJmaWx0ZXJfc2Nob29sX3NvZnREZWxldGUiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJmaWx0ZXJfbW9kdWxlX21hZGVfYnkiLCJmaWx0ZXJfZXhlcmNpc2VfcXVlc3Rpb25fdHlwZV9pZCIsInVwZGF0ZV9jbGFzc19hbGwiLCJnZXRfY2xhc3MiLCJnZXRfc2Nob29sc19hbGwiLCJmaWx0ZXJfY2xhc3Nfc29mdERlbGV0ZSIsImdldF9jbGFzc2VzX2FsbCIsInVwZGF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX21hZGVfYnkiLCJnZXRfZXhlcmNpc2VzX2FsbCIsImZpbHRlcl9jbGFzc19tYWRlX2J5IiwiZmlsdGVyX21vZHVsZV9zb2Z0RGVsZXRlIiwiZ2V0X2V4ZXJjaXNlcyIsImdldF9jbGFzc2VzIiwiZGVsZXRlX21vZHVsZSIsImdldF9zY2hvb2xzIiwiZGVsZXRlX2V4ZXJjaXNlIiwidXBkYXRlX2V4ZXJjaXNlIiwiZ2V0X2V4ZXJjaXNlIiwiZmlsdGVyX2V4ZXJjaXNlX25hbWUiLCJkZWxldGVfZXhlcmNpc2VfYWxsIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJjcmVhdGVfbW9kdWxlIiwidXBkYXRlX2V4ZXJjaXNlX2FsbCIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJmaWx0ZXJfZXhlcmNpc2Vfc29mdGRlbGV0ZWQiLCJnZXRfbW9kdWxlc19hbGwiLCJmaWx0ZXJfZXhlcmNpc2VfY2xhc3NfaWQiLCJmaWx0ZXJfbW9kdWxlX3NjaG9vbF9pZCIsImZpbHRlcl9jbGFzc19uYW1lIiwiZmlsdGVyX3NjaG9vbF9oYXNfb3BlbmFpX2FjY2VzcyIsInVwZGF0ZV9tb2R1bGUiLCJmaWx0ZXJfbW9kdWxlX25hbWUiLCJmaWx0ZXJfbW9kdWxlX21hZGVfYnlfbmFtZSIsImZpbHRlcl9leGVyY2lzZV9tYWRlX2J5IiwiZGVsZXRlX3NjaG9vbF9hbGwiLCJ1cGRhdGVfY2xhc3MiLCJmaWx0ZXJfbW9kdWxlX3ByaXZhdGUiLCJkZWxldGVfY2xhc3NfYWxsIl19LCJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiJjOWJmY2I5MS05NDkyLTRjNWUtYjdjYy1kYjc4Yjc1NGY3Y2MiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5hbWUiOiJjaGFkIGFkbWluIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW5AYWRtaW4uY29tIiwiZ2l2ZW5fbmFtZSI6ImNoYWQiLCJmYW1pbHlfbmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBhZG1pbi5jb20ifQ.cqVrVDahp_gFrxF-ol6Ji31MRmOxpRkSeKeeym142luWGT6bzL-Gz4bqqz3omH-f240sagWOsjL-hD6cxFwonJxhKgqewvONCuYn10oat3jeWQYY2GOfC-WVOouzZjOBvH4JM4hftADsHPWhCWgj5kCjOAxatlzUbhzKCqL6jTx4LE85sUzZG6Gh7V26kwamkoJKwTNpAckRMA4qXj1Z-BciA2KlQaSsacWNFyEwQFrFDzVR2PHTc9IPBj2o8Xz0qPPS7hZecHk1r2A2l0RGYsqqzpbhHT96pSjEdOx4j97aoeMK34RlzCRssS5zYzeMiNAdecNYfyCVvQ5ZbqioNg"
var teacherToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDQzNzg0NjEsImlhdCI6MTcwNDM3ODE2MSwianRpIjoiZDJkNWE4ZWItYWEyZS00ZTI5LWI5NTctNjlkMDQyMWUwNTE0IiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiNTk3OGU2YmEtZDE5OS00MjZkLWE2NDMtM2Y3YjM1MDliMGQ1IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6IjE0MWQwZDE1LWVlZWQtNDA4Yi1hM2M2LWRlMjY5MmU3ZjVhYSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZmlsdGVyX2V4ZXJjaXNlX2RpZmZpY3VsdHkiLCJmaWx0ZXJfZXhlcmNpc2VfbW9kdWxlX2lkIiwiZ2V0X2V4ZXJjaXNlcyIsImdldF9jbGFzc2VzIiwiZGVsZXRlX21vZHVsZSIsImRlbGV0ZV9leGVyY2lzZSIsImdldF9leGVyY2lzZSIsInVwZGF0ZV9leGVyY2lzZSIsImNyZWF0ZV9leGVyY2lzZSIsImZpbHRlcl9leGVyY2lzZV9uYW1lIiwiZGVsZXRlX2NsYXNzIiwiY3JlYXRlX21vZHVsZSIsImdldF9tb2R1bGUiLCJnZXRfbW9kdWxlcyIsImNyZWF0ZV9jbGFzcyIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJmaWx0ZXJfZXhlcmNpc2VfY2xhc3NfaWQiLCJmaWx0ZXJfbW9kdWxlX3NjaG9vbF9pZCIsImZpbHRlcl9jbGFzc19uYW1lIiwiZmlsdGVyX2V4ZXJjaXNlX3F1ZXN0aW9uX3R5cGVfaWQiLCJ1cGRhdGVfbW9kdWxlIiwiZmlsdGVyX21vZHVsZV9tYWRlX2J5X25hbWUiLCJnZXRfY2xhc3MiLCJ1cGRhdGVfY2xhc3MiXX0sImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoiZW1haWwgcHJvZmlsZSIsInNpZCI6IjE0MWQwZDE1LWVlZWQtNDA4Yi1hM2M2LWRlMjY5MmU3ZjVhYSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwibmFtZSI6IkJyYW0gVGVybG91dyIsInByZWZlcnJlZF91c2VybmFtZSI6ImJyYW1AdGVhY2hlci5jb20iLCJnaXZlbl9uYW1lIjoiQnJhbSIsImZhbWlseV9uYW1lIjoiVGVybG91dyIsImVtYWlsIjoiYnJhbUB0ZWFjaGVyLmNvbSJ9.Ym1c2EMv-t1rtVEel7C609ac0qzD5lritwnFh2hnxMrNUto2wz4iFVbfQB0dBcXyLItHTbzSi6_OjOTivKqhNIsdNVvnGT_tdFBrbDshuJFaHoH7HSlfi_tCSViIHhYjbPTDnGLJDx6j6nq-DumoOA65t3qwn2inFx6wm44ZOA5NvD-v2npSYVQv_uf03LopuabXMWl3DH5RcngDQB9DK1pTE8jc_eBfYuNloAGEySq_Pn0xnUoLAjAtGDp5z3DmsMQ44-1wRSWLCR2n02Kk2otD_VnDOWxexO6vrTitowT_KCcYm1RE27cIbKBh5fo8uzZsf6KHDvZARJkwAHnWlw"
var studentToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDQzNzg1MzIsImlhdCI6MTcwNDM3ODIzMiwianRpIjoiOTVmMDVhYjItNmM4Yy00YTFjLWEwYjEtNDlhM2M3ZWYyMTc3IiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiM2U4MDVlOTctZmM3Ni00MzI0LWExOTktNDYzZjYwZTQzYjQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6IjVhMGJlMTY5LWQyNzctNDgwNC1iNjlmLWQzMTNjZDMxOTIwYiIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImdldF9leGVyY2lzZSIsImdldF9tb2R1bGUiLCJnZXRfbW9kdWxlcyIsImZpbHRlcl9tb2R1bGVfbWFkZV9ieV9uYW1lIiwiZ2V0X2NsYXNzIiwiZ2V0X2V4ZXJjaXNlcyIsImZpbHRlcl9leGVyY2lzZV9jbGFzc19pZCIsImdldF9jbGFzc2VzIiwiZmlsdGVyX21vZHVsZV9zY2hvb2xfaWQiXX0sImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoiZW1haWwgcHJvZmlsZSIsInNpZCI6IjVhMGJlMTY5LWQyNzctNDgwNC1iNjlmLWQzMTNjZDMxOTIwYiIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwibmFtZSI6Ik1lcmxpam4gQnVzY2giLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJtZXJsaWpuQHN0dWRlbnQuY29tIiwiZ2l2ZW5fbmFtZSI6Ik1lcmxpam4iLCJmYW1pbHlfbmFtZSI6IkJ1c2NoIiwiZW1haWwiOiJtZXJsaWpuQHN0dWRlbnQuY29tIn0.InwXY-fwpicRpoij63WHz3arsIfdQ-pyn-EPCARUWbYHSrJ1FVkw1Kb2Ea3-2zwg56K4hveidofESHBTQEIcGmpC0qRgdOyNGnPgE6Wt6xYkhdhGYlvYbvWg9F3dMmDtaTq_PBASYGrz7Cr38IZEI6IaTo4zzLpAWPSwkKpJZyohbGnq_bPWQOYCbjUVYS63INwqGCha_lO97WW1h5iONEtX4XjWH43MAIp9sH-TXrdpnFa2tVwNwLl8qz_bunoWAZb1l_Fde_VTt2plGJQWVDF3vnGN3kRjB84GpPSiyPFBK8Jj9f7mmkOi9VO-56yQRfmyScoq2-qcd_SQmO23OQ"

var adminUUID = "6c1ce448-670f-47b2-83f7-4d771b01775b"
var teacherUUID = "5978e6ba-d199-426d-a643-3f7b3509b0d5"
var studentUUID = "3a3bd756-6353-4e29-8aba-5b3531bdb9ed"

func TestCreateExerciseWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	uuid, err := exercisePolicy.CreateExercise(adminToken)

	assert.Nil(t, err)
	assert.Equal(t, adminUUID, uuid)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateExerciseWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	uuid, err := exercisePolicy.CreateExercise(teacherToken)

	assert.Nil(t, err)
	assert.Equal(t, teacherUUID, uuid)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateExerciseWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	uuid, err := exercisePolicy.CreateExercise(studentToken)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid permissions for this action", err.Error())
	assert.Empty(t, uuid)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUpdateExerciseWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On(
			"GetExerciseByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	_, err := exercisePolicy.UpdateExercise(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUpdateExerciseWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockRepo.
		On(
			"GetExerciseByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := exercisePolicy.UpdateExercise(teacherToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUpdateExerciseWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockRepo.
		On(
			"GetExerciseByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := exercisePolicy.UpdateExercise(studentToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid permissions for this action", err.Error())
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestDeleteExerciseWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On(
			"GetExerciseByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	_, _, err := exercisePolicy.DeleteExercise(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestDeleteExerciseWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockRepo.
		On(
			"GetExerciseByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, _, err := exercisePolicy.DeleteExercise(teacherToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestDeleteExerciseWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockRepo.
		On(
			"GetExerciseByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, _, err := exercisePolicy.DeleteExercise(studentToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid permissions for this action", err.Error())
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetExerciseWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On("GetExerciseByID", "3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	exercise, err := exercisePolicy.GetExercise(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	assert.NotNil(t, exercise)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetExerciseWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On("GetExerciseByID", "3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	exercise, err := exercisePolicy.GetExercise(teacherToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	assert.NotNil(t, exercise)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetExerciseWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On("GetExerciseByID", "3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockExercise, nil)

	exercise, err := exercisePolicy.GetExercise(studentToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	assert.NotNil(t, exercise)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestListExercisesWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	result, err := exercisePolicy.ListExercises(adminToken)

	assert.Nil(t, err)
	assert.True(t, result)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestListExercisesWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	result, err := exercisePolicy.ListExercises(teacherToken)

	assert.Nil(t, err)
	assert.False(t, result)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestListExercisesWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := exercisePolicy.ListExercises(studentToken)

	assert.Nil(t, err)
	assert.False(t, false)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestHasPermissionsWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	// Define the expected roles for the admin token
	expectedRoles := []string{
		"filter_exercise_difficulty",
		"filter_exercise_module_id",
		"create_exercise",
		"filter_exercise_question_type_id",
		"get_exercises_all",
		"get_exercises",
		"delete_exercise",
		"update_exercise",
		"get_exercise",
		"filter_exercise_name",
		"delete_exercise_all",
		"update_exercise_all",
		"filter_exercise_softdeleted",
		"filter_exercise_class_id",
		"filter_exercise_made_by",
	}

	//loop through the expected roles and check if the admin token has them
	for _, role := range expectedRoles {
		hasPermissions := exercisePolicy.HasPermissions(adminToken, role)
		assert.True(t, hasPermissions, "Admin token should have all exercise-related roles")
	}

	mockToken.AssertExpectations(t)
}

func TestHasPermissionsWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	// Define the expected roles for the admin token
	expectedRoles := []string{
		"filter_exercise_difficulty",
		"filter_exercise_module_id",
		"create_exercise",
		"filter_exercise_question_type_id",
		"get_exercises",
		"delete_exercise",
		"update_exercise",
		"get_exercise",
		"filter_exercise_name",
		"filter_exercise_class_id",
	}

	//loop through the expected roles and check if the admin token has them
	for _, role := range expectedRoles {
		hasPermissions := exercisePolicy.HasPermissions(teacherToken, role)
		assert.True(t, hasPermissions, "Teacher token should have all expected roles")
	}

	mockToken.AssertExpectations(t)
}

func TestHasPermissionsWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockExerciseRepository)
	mockToken := new(mocks.MockToken)

	exercisePolicy := auth.ExercisePolicy{
		Token:              mockToken,
		ExerciseRepository: mockRepo,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	// Define the expected roles for the admin token
	expectedRoles := []string{
		"get_exercises",
		"get_exercise",
		"filter_exercise_class_id",
	}

	//loop through the expected roles and check if the admin token has them
	for _, role := range expectedRoles {
		hasPermissions := exercisePolicy.HasPermissions(studentToken, role)
		assert.True(t, hasPermissions, "Student token should have all expected roles")
	}

	mockToken.AssertExpectations(t)
}
