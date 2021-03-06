package user

import (
	"fmt"
	"seuxw/embrice/entity/user"
)

// CreateUserDB 用户创建 DB 操作
func (db *Database) CreateUserDB(user *user.User) error {
	var (
		insertSQL           string
		insertID            int64
		selectCheckSQL      string
		selectCheckWhereStr string
		err                 error
		count               int
	)

	selectCheckSQL = `
	select
		count(1) as count
	from
		sd_user
	where
		%s and deleted = 0
	`

	insertSQL = `
	insert into sd_user (
		card_id, user_uuid, qq_id, wechat_id, stu_no, real_name,
		nick_name, gender, user_type, pwd, session, mobile
	) values (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)
	`

	// 查询待插入的记录是否已经存在 避免重复插入
	if user.QQID != 0 {
		selectCheckWhereStr = "qq_id = ?"
		insertID = user.QQID
	} else if user.CardID != 0 {
		selectCheckWhereStr = "card_id = ?"
		insertID = user.CardID
	}

	err = db.Get(&count, fmt.Sprintf(selectCheckSQL, selectCheckWhereStr), insertID)
	if err != nil {
		err = fmt.Errorf("数据库预查询错误 err:%s", err)
		goto END
	}

	// 插入操作
	if count == 0 {
		_ = db.MustExec(
			insertSQL, user.CardID, user.UserUUID, user.QQID, user.WeChatID,
			user.StuNo, user.RealName, user.NickName, user.Gender,
			user.UserType, user.Pwd, user.Session, user.Mobile)
	} else {
		err = fmt.Errorf("用户 %d 已经存在！", insertID)
	}

END:
	return err
}

// GetUserByUUIDDB
func (db *Database) GetUserByUUIDDB(uuid string) (*user.GetUserByUUIDResp, error) {
	var (
		selectSQL string
		err       error
		rtnData   user.GetUserByUUIDResp
	)

	selectSQL = `
	select
		card_id, qq_id, wechat_id, stu_no, real_name, nick_name,
		gender, user_type, identity, class, dept_name, major_name,
		grade, nick_name, vip, vip_level, rmk_name, hometown,
		address, birthday
	from
		v_insensitive_userinfo
	where
		user_uuid = ?
	`

	if err = db.Get(&rtnData, selectSQL, uuid); err != nil {
		goto END
	}

END:
	return &rtnData, err
}
