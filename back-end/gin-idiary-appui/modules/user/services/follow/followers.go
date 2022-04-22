/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:27:51
 * @Description: file content
 */
package follow

import (
	"context"
	uploadData "gin-idiary-appui/modules/upload/data"
	followData "gin-idiary-appui/modules/user/data/follow"
	infoData "gin-idiary-appui/modules/user/data/info"
	followModel "gin-idiary-appui/modules/user/model/follow"
	infoModel "gin-idiary-appui/modules/user/model/info"
)

func Followers(ctx context.Context, pars followModel.FollowerPars) ([]infoModel.UserInfoNonsensitive, error) {
	userInfos := make([]infoModel.UserInfo, 0)

	// 1 获取关注列表
	followers, err := followData.Followers(ctx, pars.UserID)
	if err != nil {
		return nil, err
	}

	// 2 获取用户信息
	for _, uid := range followers {
		userInfo, err := infoData.GetUserInfoByUserID(ctx, uid)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}

	// 3 profile url
	for k, v := range userInfos {
		userInfos[k].Profile, err = uploadData.GetImageURL(ctx, v.Profile)
		if err != nil {
			return nil, err
		}
	}

	// 4 去除敏感信息
	userInfosNonsensitive := make([]infoModel.UserInfoNonsensitive, 0)
	for _, v := range userInfos {
		userInfoNonsensitive := infoModel.UserInfoNonsensitive{
			UserID:   v.UserID,
			Email:    v.Email,
			Nickname: v.Nickname,
			Profile:  v.Profile,
		}
		userInfosNonsensitive = append(userInfosNonsensitive, userInfoNonsensitive)
	}

	return userInfosNonsensitive, nil
}
