/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:16:58
 * @Description: file content
 */
package follow

import (
	"context"
	"gin-idiary-appui/library/mysql"
	"gin-idiary-appui/modules/user/constant"
)

const (
	//用户关注表
	USER_FOLLOWING_TABLE = "tb_user_followings"
)

func GetFollowing(ctx context.Context, userID int64) (string, error) {
	var followings = make([]struct {
		FollowingList string `ddb:"following_list"`
	}, 1)

	client, err := mysql.GetClient(ctx, constant.SERVICE_CONF_DB_IDIARY_USER)
	if err != nil {
		return "", err
	}

	tableName := USER_FOLLOWING_TABLE

	where := map[string]interface{}{
		"user_id": userID,
	}

	columns := []string{"following_list"}

	err = client.Query(ctx, tableName, where, columns, &followings)
	if err != nil {
		if err.Error() == "[scanner]: empty result" {
			return "", nil
		}
		return "", err
	}

	return followings[0].FollowingList, nil
}
