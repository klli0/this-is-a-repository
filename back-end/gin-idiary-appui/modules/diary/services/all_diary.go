/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-22 20:43:27
 * @Description: file content
 */
package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func AllDiary(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfoUnmarshaled, int64, error) {
	// 1 get diary
	diaries, count, err := diaryData.AllDiary(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	// 2 all diary process
	diariesUnmarshaled, err := ProcessDiary(ctx, pars.UserID, diaries)
	if err != nil {
		return nil, 0, err
	}

	return diariesUnmarshaled, count, nil
}
