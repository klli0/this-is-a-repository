/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:01:05
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-22 20:33:20
 * @Description: file content
 */
package comment

import (
	"context"
	commentDao "gin-idiary-appui/modules/diary/dao/comment"
	commentModel "gin-idiary-appui/modules/diary/model/comment"
	infoDao "gin-idiary-appui/modules/user/dao/info"
)

func GetCommentWithUserInfoByDiaryID(ctx context.Context, diaryID int64) ([]commentModel.CommentInfo, int64, error) {
	comments, count, err := GetCommentByDiaryID(ctx, diaryID)
	if err != nil {
		return nil, 0, err
	}

	for k, v := range comments {
		userInfo, err := infoDao.GetUserInfoByUserID(ctx, v.UserID)
		if err != nil {
			return nil, 0, err
		}
		comments[k].Nickname = userInfo.Nickname
		comments[k].Profile = userInfo.Profile
	}

	return comments, count, err
}

func GetCommentByDiaryID(ctx context.Context, diaryID int64) ([]commentModel.CommentInfo, int64, error) {
	comments, err := commentDao.GetCommentByDiaryID(ctx, diaryID)
	if err != nil {
		return nil, 0, err
	}

	count, err := commentDao.GetCommentByDiaryIDCount(ctx, diaryID)
	if err != nil {
		return nil, 0, err
	}

	return comments, count, err
}
