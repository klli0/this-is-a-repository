/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-22 20:43:34
 * @Description: file content
 */
package services

import (
	"context"
	diaryData "gin-idiary-appui/modules/diary/data"
	diaryModel "gin-idiary-appui/modules/diary/model"
	followData "gin-idiary-appui/modules/user/data/follow"
)

func FriendDiary(ctx context.Context, pars diaryModel.FriendDiaryListRequestPars) ([]diaryModel.DiaryInfoUnmarshaled, int64, error) {
	// 1 获取关注列表
	followingList, err := followData.Followings(ctx, pars.UserID)
	if err != nil {
		return nil, 0, err
	}

	// 2 获取粉丝列表
	followerList, err := followData.Followers(ctx, pars.UserID)
	if err != nil {
		return nil, 0, err
	}

	// 3 计算 朋友：关注列表 与 粉丝列表重合的部分
	if len(followerList) == 0 || len(followingList) == 0 {
		return nil, 0, nil
	}

	friendIDs, err := diaryData.CalcFriend(ctx, followingList, followerList)
	if err != nil {
		return nil, 0, err
	}

	// 4 获取朋友的日记列表
	diaries, count, err := diaryData.FriendDiary(ctx, pars, friendIDs)
	if err != nil {
		return nil, 0, err
	}

	// 5 all diary process
	diariesUnmarshaled, err := ProcessDiary(ctx, pars.UserID, diaries)
	if err != nil {
		return nil, 0, err
	}

	return diariesUnmarshaled, count, nil

}
