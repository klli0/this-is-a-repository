/*
 * @Author: liziwei01
 * @Date: 2022-04-16 18:24:49
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:17:09
 * @Description: file content
 */
package follow

import (
	"context"
	"encoding/json"
	"gin-idiary-appui/library/logit"
	followDao "gin-idiary-appui/modules/user/dao/follow"
)

func Followings(ctx context.Context, userID int64) ([]int64, error) {
	var followingList = make([]int64, 0)
	followings, err := followDao.GetFollowing(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(followings) == 0 {
		return followingList, nil
	}

	err = json.Unmarshal([]byte(followings), &followingList)
	if err != nil {
		logit.Logger.Error("json.Unmarshal error: #+v", err)
		return nil, err
	}

	return followingList, nil
}
