/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:41:16
 * @Description: file content
 */
package follow

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-idiary-appui/library/logit"
	"gin-idiary-appui/library/utils"
	followDao "gin-idiary-appui/modules/user/dao/follow"

	"github.com/gogf/gf/util/gconv"
)

func Follow(ctx context.Context, userID, followingID int64) error {
	var followingList = make([]int64, 0)
	var followerList = make([]int64, 0)

	// 1 拿到关注列表
	followings, err := followDao.GetFollowing(ctx, userID)
	if err != nil {
		return err
	}

	// 2 添加关注
	if followings != "" {
		err = json.Unmarshal([]byte(followings), &followingList)
		if err != nil {
			logit.Logger.Error("json.Unmarshal error: #+v", err)
			return err
		}
	}
	if utils.Slice.In(followingID, followingList) == true {
		return fmt.Errorf("已经关注过了")
	}
	followingList = append(followingList, followingID)
	followingsBytes, err := json.Marshal(followingList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 3 更新关注列表
	err = followDao.ModifyFollowing(ctx, userID, gconv.String(followingsBytes))
	if err != nil {
		return err
	}

	// 4 拿到被关注者的粉丝列表
	followers, err := followDao.GetFollower(ctx, followingID)
	if err != nil {
		return err
	}

	// 5 添加粉丝
	if followers != "" {
		err = json.Unmarshal([]byte(followers), &followerList)
		if err != nil {
			logit.Logger.Error("json.Unmarshal error: #+v", err)
			return err
		}
	}

	followerList = append(followerList, userID)
	followersBytes, err := json.Marshal(followerList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 6 更新粉丝列表
	err = followDao.ModifyFollower(ctx, followingID, gconv.String(followersBytes))
	if err != nil {
		return err
	}

	return nil
}

func UnFollow(ctx context.Context, userID, followingID int64) error {
	var followingList = make([]int64, 0)
	var followerList = make([]int64, 0)

	// 1 拿到关注列表
	followings, err := followDao.GetFollowing(ctx, userID)
	if err != nil {
		return err
	}

	// 2 取消关注
	if followings != "" {
		err = json.Unmarshal([]byte(followings), &followingList)
		if err != nil {
			logit.Logger.Error("json.Unmarshal error: #+v", err)
			return err
		}
	}
	if utils.Slice.In(followingID, followingList) == false {
		return fmt.Errorf("已经取关过了")
	}
	utils.Slice.Remove(&followingList, followingID)
	followingsBytes, err := json.Marshal(followingList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 3 更新关注列表
	err = followDao.ModifyFollowing(ctx, userID, gconv.String(followingsBytes))
	if err != nil {
		return err
	}

	// 4 拿到被关注者的粉丝列表
	followers, err := followDao.GetFollower(ctx, followingID)
	if err != nil {
		return err
	}

	// 5 取消粉丝
	if followers != "" {
		err = json.Unmarshal([]byte(followers), &followerList)
		if err != nil {
			logit.Logger.Error("json.Unmarshal error: #+v", err)
			return err
		}
	}
	utils.Slice.Remove(&followerList, userID)
	followersBytes, err := json.Marshal(followerList)
	if err != nil {
		logit.Logger.Error("json.Marshal error: #+v", err)
		return err
	}

	// 6 更新粉丝列表
	err = followDao.ModifyFollower(ctx, followingID, gconv.String(followersBytes))
	if err != nil {
		return err
	}

	return nil
}
