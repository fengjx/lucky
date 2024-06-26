// Code generated by "daox.gen"; DO NOT EDIT.
package meta

import (
    "github.com/fengjx/daox/sqlbuilder"
    "github.com/fengjx/daox/sqlbuilder/ql"


    "time"

)



// SysUserM 用户信息表
type SysUserM struct {
    ID string
    Username string
    Pwd string
    Salt string
    Email string
    Nickname string
    Avatar string
    Phone string
    Status string
    Remark string
    Utime string
    Ctime string
}

func (m SysUserM) TableName() string {
    return "sys_user"
}

func (m SysUserM) IsAutoIncrement() bool {
    return true
}

func (m SysUserM) PrimaryKey() string {
    return "id"
}

func (m SysUserM) Columns() []string {
	return []string{
        "id",
        "username",
        "pwd",
        "salt",
        "email",
        "nickname",
        "avatar",
        "phone",
        "status",
        "remark",
        "utime",
        "ctime",
    }
}

var SysUserMeta = SysUserM{
    ID: "id",
    Username: "username",
    Pwd: "pwd",
    Salt: "salt",
    Email: "email",
    Nickname: "nickname",
    Avatar: "avatar",
    Phone: "phone",
    Status: "status",
    Remark: "remark",
    Utime: "utime",
    Ctime: "ctime",
}




func (m SysUserM) IdIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).In(args...)
}

func (m SysUserM) IdNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).NotIn(args...)
}

func (m SysUserM) IdEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).EQ(val)
}

func (m SysUserM) IdNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotEQ(val)
}

func (m SysUserM) IdLT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LT(val)
}

func (m SysUserM) IdLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LTEQ(val)
}

func (m SysUserM) IdGT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GT(val)
}

func (m SysUserM) IdGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GTEQ(val)
}

func (m SysUserM) IdLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).Like(val)
}

func (m SysUserM) IdNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotLike(val)
}

func (m SysUserM) IdDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.ID)
}

func (m SysUserM) IdAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.ID)
}



func (m SysUserM) UsernameIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Username).In(args...)
}

func (m SysUserM) UsernameNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Username).NotIn(args...)
}

func (m SysUserM) UsernameEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).EQ(val)
}

func (m SysUserM) UsernameNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).NotEQ(val)
}

func (m SysUserM) UsernameLT(val string) sqlbuilder.Column {
	return ql.Col(m.Username).LT(val)
}

func (m SysUserM) UsernameLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).LTEQ(val)
}

func (m SysUserM) UsernameGT(val string) sqlbuilder.Column {
	return ql.Col(m.Username).GT(val)
}

func (m SysUserM) UsernameGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Username).GTEQ(val)
}

func (m SysUserM) UsernameLike(val string) sqlbuilder.Column {
	return ql.Col(m.Username).Like(val)
}

func (m SysUserM) UsernameNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Username).NotLike(val)
}

func (m SysUserM) UsernameDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Username)
}

func (m SysUserM) UsernameAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Username)
}



func (m SysUserM) PwdIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Pwd).In(args...)
}

func (m SysUserM) PwdNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Pwd).NotIn(args...)
}

func (m SysUserM) PwdEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).EQ(val)
}

func (m SysUserM) PwdNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).NotEQ(val)
}

func (m SysUserM) PwdLT(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).LT(val)
}

func (m SysUserM) PwdLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).LTEQ(val)
}

func (m SysUserM) PwdGT(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).GT(val)
}

func (m SysUserM) PwdGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).GTEQ(val)
}

func (m SysUserM) PwdLike(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).Like(val)
}

func (m SysUserM) PwdNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Pwd).NotLike(val)
}

func (m SysUserM) PwdDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Pwd)
}

func (m SysUserM) PwdAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Pwd)
}



func (m SysUserM) SaltIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Salt).In(args...)
}

func (m SysUserM) SaltNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Salt).NotIn(args...)
}

func (m SysUserM) SaltEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).EQ(val)
}

func (m SysUserM) SaltNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).NotEQ(val)
}

func (m SysUserM) SaltLT(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).LT(val)
}

func (m SysUserM) SaltLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).LTEQ(val)
}

func (m SysUserM) SaltGT(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).GT(val)
}

func (m SysUserM) SaltGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).GTEQ(val)
}

func (m SysUserM) SaltLike(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).Like(val)
}

func (m SysUserM) SaltNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Salt).NotLike(val)
}

func (m SysUserM) SaltDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Salt)
}

func (m SysUserM) SaltAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Salt)
}



func (m SysUserM) EmailIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Email).In(args...)
}

func (m SysUserM) EmailNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Email).NotIn(args...)
}

func (m SysUserM) EmailEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Email).EQ(val)
}

func (m SysUserM) EmailNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Email).NotEQ(val)
}

func (m SysUserM) EmailLT(val string) sqlbuilder.Column {
	return ql.Col(m.Email).LT(val)
}

func (m SysUserM) EmailLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Email).LTEQ(val)
}

func (m SysUserM) EmailGT(val string) sqlbuilder.Column {
	return ql.Col(m.Email).GT(val)
}

func (m SysUserM) EmailGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Email).GTEQ(val)
}

func (m SysUserM) EmailLike(val string) sqlbuilder.Column {
	return ql.Col(m.Email).Like(val)
}

func (m SysUserM) EmailNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Email).NotLike(val)
}

func (m SysUserM) EmailDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Email)
}

func (m SysUserM) EmailAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Email)
}



func (m SysUserM) NicknameIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Nickname).In(args...)
}

func (m SysUserM) NicknameNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Nickname).NotIn(args...)
}

func (m SysUserM) NicknameEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).EQ(val)
}

func (m SysUserM) NicknameNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).NotEQ(val)
}

func (m SysUserM) NicknameLT(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).LT(val)
}

func (m SysUserM) NicknameLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).LTEQ(val)
}

func (m SysUserM) NicknameGT(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).GT(val)
}

func (m SysUserM) NicknameGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).GTEQ(val)
}

func (m SysUserM) NicknameLike(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).Like(val)
}

func (m SysUserM) NicknameNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Nickname).NotLike(val)
}

func (m SysUserM) NicknameDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Nickname)
}

func (m SysUserM) NicknameAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Nickname)
}



func (m SysUserM) AvatarIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Avatar).In(args...)
}

func (m SysUserM) AvatarNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Avatar).NotIn(args...)
}

func (m SysUserM) AvatarEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).EQ(val)
}

func (m SysUserM) AvatarNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).NotEQ(val)
}

func (m SysUserM) AvatarLT(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).LT(val)
}

func (m SysUserM) AvatarLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).LTEQ(val)
}

func (m SysUserM) AvatarGT(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).GT(val)
}

func (m SysUserM) AvatarGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).GTEQ(val)
}

func (m SysUserM) AvatarLike(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).Like(val)
}

func (m SysUserM) AvatarNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Avatar).NotLike(val)
}

func (m SysUserM) AvatarDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Avatar)
}

func (m SysUserM) AvatarAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Avatar)
}



func (m SysUserM) PhoneIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Phone).In(args...)
}

func (m SysUserM) PhoneNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Phone).NotIn(args...)
}

func (m SysUserM) PhoneEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).EQ(val)
}

func (m SysUserM) PhoneNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).NotEQ(val)
}

func (m SysUserM) PhoneLT(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).LT(val)
}

func (m SysUserM) PhoneLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).LTEQ(val)
}

func (m SysUserM) PhoneGT(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).GT(val)
}

func (m SysUserM) PhoneGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).GTEQ(val)
}

func (m SysUserM) PhoneLike(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).Like(val)
}

func (m SysUserM) PhoneNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Phone).NotLike(val)
}

func (m SysUserM) PhoneDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Phone)
}

func (m SysUserM) PhoneAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Phone)
}



func (m SysUserM) StatusIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Status).In(args...)
}

func (m SysUserM) StatusNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Status).NotIn(args...)
}

func (m SysUserM) StatusEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).EQ(val)
}

func (m SysUserM) StatusNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).NotEQ(val)
}

func (m SysUserM) StatusLT(val string) sqlbuilder.Column {
	return ql.Col(m.Status).LT(val)
}

func (m SysUserM) StatusLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).LTEQ(val)
}

func (m SysUserM) StatusGT(val string) sqlbuilder.Column {
	return ql.Col(m.Status).GT(val)
}

func (m SysUserM) StatusGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Status).GTEQ(val)
}

func (m SysUserM) StatusLike(val string) sqlbuilder.Column {
	return ql.Col(m.Status).Like(val)
}

func (m SysUserM) StatusNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Status).NotLike(val)
}

func (m SysUserM) StatusDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Status)
}

func (m SysUserM) StatusAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Status)
}



func (m SysUserM) RemarkIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Remark).In(args...)
}

func (m SysUserM) RemarkNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Remark).NotIn(args...)
}

func (m SysUserM) RemarkEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).EQ(val)
}

func (m SysUserM) RemarkNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).NotEQ(val)
}

func (m SysUserM) RemarkLT(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).LT(val)
}

func (m SysUserM) RemarkLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).LTEQ(val)
}

func (m SysUserM) RemarkGT(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).GT(val)
}

func (m SysUserM) RemarkGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).GTEQ(val)
}

func (m SysUserM) RemarkLike(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).Like(val)
}

func (m SysUserM) RemarkNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Remark).NotLike(val)
}

func (m SysUserM) RemarkDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Remark)
}

func (m SysUserM) RemarkAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Remark)
}



func (m SysUserM) UtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).In(args...)
}

func (m SysUserM) UtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Utime).NotIn(args...)
}

func (m SysUserM) UtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).EQ(val)
}

func (m SysUserM) UtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotEQ(val)
}

func (m SysUserM) UtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LT(val)
}

func (m SysUserM) UtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).LTEQ(val)
}

func (m SysUserM) UtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GT(val)
}

func (m SysUserM) UtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).GTEQ(val)
}

func (m SysUserM) UtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).Like(val)
}

func (m SysUserM) UtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Utime).NotLike(val)
}

func (m SysUserM) UtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Utime)
}

func (m SysUserM) UtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Utime)
}



func (m SysUserM) CtimeIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).In(args...)
}

func (m SysUserM) CtimeNotIn(vals ...time.Time) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Ctime).NotIn(args...)
}

func (m SysUserM) CtimeEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).EQ(val)
}

func (m SysUserM) CtimeNotEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotEQ(val)
}

func (m SysUserM) CtimeLT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LT(val)
}

func (m SysUserM) CtimeLTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).LTEQ(val)
}

func (m SysUserM) CtimeGT(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GT(val)
}

func (m SysUserM) CtimeGTEQ(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).GTEQ(val)
}

func (m SysUserM) CtimeLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).Like(val)
}

func (m SysUserM) CtimeNotLike(val time.Time) sqlbuilder.Column {
	return ql.Col(m.Ctime).NotLike(val)
}

func (m SysUserM) CtimeDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Ctime)
}

func (m SysUserM) CtimeAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Ctime)
}
