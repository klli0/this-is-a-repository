/*
 * @Author: liziwei01
 * @Date: 2022-04-16 18:24:49
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:17:06
 * @Description: file content
 */
package follow

import (
	"context"
	"encoding/json"
	"gin-idiary-appui/library/logit"
	followDao "gin-idiary-appui/modules/user/dao/follow"
)

func Followers(ctx context.Context, userID int64) ([]int64, error) {
	var followerList = make([]int64, 0)
	followers, err := followDao.GetFollower(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(followers) == 0 {
		return followerList, nil
	}

	err = json.Unmarshal([]byte(followers), &followerList)
	if err != nil {
		logit.Logger.Error("json.Unmarshal error: #+v", err)
		return nil, err
	}

	return followerList, nil
}
