/*
 * @Author: liziwei01
 * @Date: 2022-04-17 16:53:02
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-22 20:43:40
 * @Description: file content
 */
package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	diaryModel "gin-idiary-appui/modules/diary/model"
)

func MyDiary(ctx context.Context, pars diaryModel.DiaryListRequestPars) ([]diaryModel.DiaryInfoUnmarshaled, int64, error) {
	// 1 get diary
	diaries, count, err := diaryData.GetDiary(ctx, pars)
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
