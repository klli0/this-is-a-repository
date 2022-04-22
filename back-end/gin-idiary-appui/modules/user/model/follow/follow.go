/*
 * @Author: liziwei01
 * @Date: 2022-04-12 14:24:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-23 00:35:07
 * @Description: file content
 */
package model

type FollowPars struct {
	UserID         int64 `form:"user_id" json:"user_id" binding:"required"`
	FollowingID    int64 `form:"following_id" json:"following_id" binding:"required"`
	ShouldUnFollow bool  `form:"should_unfollow" json:"should_unfollow"`
}

type FollowingPars struct {
	UserID int64 `form:"user_id" json:"user_id" binding:"required"`
}

type FollowerPars struct {
	UserID int64 `form:"user_id" json:"user_id" binding:"required"`
}
