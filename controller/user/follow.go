package user

import (
	"fmt"
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/user"
	"gateway/sql"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// FollowersController 获取粉丝
func FollowersController(c *gin.Context) {
	var form user.FollowersGetReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	if form.Offset < 0 || form.Number <= 0 {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	ca := cache.Follow{}
	// 获取粉丝总数
	followersCounter, err := ca.CountFollowers(c.GetUint(middleware.KeyUID))
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}
	// 分页获取用户uid
	uidStringList, err := ca.GetFollowers(c.GetUint(middleware.KeyUID), form.Offset, form.Number)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	db := sql.UserInfoTable{}
	// 按uid获取用户信息
	data, err := db.QueryInByUIDs(uidStringList)
	if err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	ret := &user.FollowersGetResModel{
		Total: followersCounter,
		Users: data,
	}
	errno.Perfect(c, ret)
}

// FollowingController 获取关注
func FollowingController(c *gin.Context) {
	var form user.FollowingGetReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	if form.Offset < 0 || form.Number <= 0 {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	ca := cache.Follow{}
	// 获取粉丝总数
	followersCounter, err := ca.CountFollowing(c.GetUint(middleware.KeyUID))
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}
	// 分页获取用户uid和时间
	res, err := ca.GetFollowing(c.GetUint(middleware.KeyUID), form.Offset, form.Number)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	var uidStringList []string
	for _, z := range res {
		uidStringList = append(uidStringList, fmt.Sprintf("%d", z.Member.(uint)))
	}

	db := sql.UserInfoTable{}
	// 按uid获取用户信息
	data, err := db.QueryInByUIDs(uidStringList)
	if err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	ret := &user.FollowingGetResModel{
		Total: followersCounter,
		Users: data,
	}
	errno.Perfect(c, ret)
}

// FollowItController 关注用户
func FollowItController(c *gin.Context) {
	var form user.FollowingActReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	if !(form.Type == 1 || form.Type == 0) {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	var err error
	uid := c.GetUint(middleware.KeyUID)
	ca := cache.Follow{}
	if form.Type == 0 {
		// 关注
		err = ca.Follow(uid, form.UID)
	} else {
		err = ca.UnFollow(uid, form.UID)
	}

	if err == cache.AddItAlreadyErr {
		errno.New(c, errno.TypeUnknownMistake, nil, "已经关注过了")
		return
	} else if err == cache.DelItAlreadyErr {
		errno.New(c, errno.TypeUnknownMistake, nil, "已经取关过了")
		return
	} else if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	errno.Perfect(c, nil)
}
