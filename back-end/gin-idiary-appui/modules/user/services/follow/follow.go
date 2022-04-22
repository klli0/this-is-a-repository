/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:36:31
 * @Description: file content
 */
package follow

import (
	"context"
	followData "gin-idiary-appui/modules/user/data/follow"
	followModel "gin-idiary-appui/modules/user/model/follow"
)

func Follow(ctx context.Context, pars followModel.FollowPars) error {
	if !pars.ShouldUnFollow {
		return followData.Follow(ctx, pars.UserID, pars.FollowingID)
	}
	return followData.UnFollow(ctx, pars.UserID, pars.FollowingID)
}
