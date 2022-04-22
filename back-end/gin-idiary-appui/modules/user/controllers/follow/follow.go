/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:32:48
 * @Description: file content
 */
package follow

import (
	"gin-idiary-appui/library/response"
	followModel "gin-idiary-appui/modules/user/model/follow"
	followService "gin-idiary-appui/modules/user/services/follow"

	"github.com/gin-gonic/gin"
)

func Follow(ctx *gin.Context) {
	inputs, hasError := getFollowPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := followService.Follow(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getFollowPars(ctx *gin.Context) (followModel.FollowPars, bool) {
	var inputs followModel.FollowPars
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	return inputs, false
}
