/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:12:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-22 20:33:56
 * @Description: file content
 */
/*
 * @Author: liziwei01
 * @Date: 2022-04-17 14:12:20
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-17 14:12:20
 * @Description: file content
 */
package info

import (
	"context"
	"fmt"
	infoDao "gin-idiary-appui/modules/user/dao/info"
	infoModel "gin-idiary-appui/modules/user/model/info"
)

func BatchGetUserInfo(ctx context.Context, userIDs []int64) ([]infoModel.UserInfo, error) {
	userInfos := make([]infoModel.UserInfo, 0)

	for _, userID := range userIDs {
		if userID <= 0 {
			return nil, fmt.Errorf("invalid user id: %d", userID)
		}
		userInfo, err := infoDao.GetUserInfoByUserID(ctx, userID)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

func GetUserInfoByUserID(ctx context.Context, userID int64) (infoModel.UserInfo, error) {
	info, err := infoDao.GetUserInfoByUserID(ctx, userID)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	return info, nil
}

func GetUserInfoByEmail(ctx context.Context, email string) (infoModel.UserInfo, error) {
	if email == "" {
		return infoModel.UserInfo{}, fmt.Errorf("invalid email: %s", email)
	}

	info, err := infoDao.GetUserInfoByEmail(ctx, email)
	if err != nil {
		return infoModel.UserInfo{}, err
	}

	if len(info) != 1 {
		return infoModel.UserInfo{}, fmt.Errorf("invalid user info len: %v", len(info))
	}

	return info[0], nil
}
func GetAllUserInfo(ctx context.Context) ([]infoModel.UserInfo, error) {
	infos, err := infoDao.GetAllUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
