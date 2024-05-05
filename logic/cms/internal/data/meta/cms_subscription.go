// Code generated by "daox.gen"; DO NOT EDIT.
package meta

import (
    "github.com/fengjx/daox/sqlbuilder"
    "github.com/fengjx/daox/sqlbuilder/ql"


)



// CmsSubscriptionM 用户订阅主题表
type CmsSubscriptionM struct {
    ID string
    UID string
    Topic string
}

func (m CmsSubscriptionM) TableName() string {
    return "cms_subscription"
}

func (m CmsSubscriptionM) IsAutoIncrement() bool {
    return false
}

func (m CmsSubscriptionM) PrimaryKey() string {
    return "id"
}

func (m CmsSubscriptionM) Columns() []string {
	return []string{
        "id",
        "uid",
        "topic",
    }
}

var CmsSubscriptionMeta = CmsSubscriptionM{
    ID: "id",
    UID: "uid",
    Topic: "topic",
}




func (m CmsSubscriptionM) IdIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).In(args...)
}

func (m CmsSubscriptionM) IdNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.ID).NotIn(args...)
}

func (m CmsSubscriptionM) IdEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).EQ(val)
}

func (m CmsSubscriptionM) IdNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotEQ(val)
}

func (m CmsSubscriptionM) IdLT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LT(val)
}

func (m CmsSubscriptionM) IdLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).LTEQ(val)
}

func (m CmsSubscriptionM) IdGT(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GT(val)
}

func (m CmsSubscriptionM) IdGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).GTEQ(val)
}

func (m CmsSubscriptionM) IdLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).Like(val)
}

func (m CmsSubscriptionM) IdNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.ID).NotLike(val)
}

func (m CmsSubscriptionM) IdDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.ID)
}

func (m CmsSubscriptionM) IdAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.ID)
}



func (m CmsSubscriptionM) UidIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.UID).In(args...)
}

func (m CmsSubscriptionM) UidNotIn(vals ...int64) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.UID).NotIn(args...)
}

func (m CmsSubscriptionM) UidEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).EQ(val)
}

func (m CmsSubscriptionM) UidNotEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).NotEQ(val)
}

func (m CmsSubscriptionM) UidLT(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).LT(val)
}

func (m CmsSubscriptionM) UidLTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).LTEQ(val)
}

func (m CmsSubscriptionM) UidGT(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).GT(val)
}

func (m CmsSubscriptionM) UidGTEQ(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).GTEQ(val)
}

func (m CmsSubscriptionM) UidLike(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).Like(val)
}

func (m CmsSubscriptionM) UidNotLike(val int64) sqlbuilder.Column {
	return ql.Col(m.UID).NotLike(val)
}

func (m CmsSubscriptionM) UidDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.UID)
}

func (m CmsSubscriptionM) UidAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.UID)
}



func (m CmsSubscriptionM) TopicIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Topic).In(args...)
}

func (m CmsSubscriptionM) TopicNotIn(vals ...string) sqlbuilder.Column {
	var args []any
    for _, val := range vals {
        args = append(args, val)
    }
    return ql.Col(m.Topic).NotIn(args...)
}

func (m CmsSubscriptionM) TopicEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).EQ(val)
}

func (m CmsSubscriptionM) TopicNotEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).NotEQ(val)
}

func (m CmsSubscriptionM) TopicLT(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).LT(val)
}

func (m CmsSubscriptionM) TopicLTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).LTEQ(val)
}

func (m CmsSubscriptionM) TopicGT(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).GT(val)
}

func (m CmsSubscriptionM) TopicGTEQ(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).GTEQ(val)
}

func (m CmsSubscriptionM) TopicLike(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).Like(val)
}

func (m CmsSubscriptionM) TopicNotLike(val string) sqlbuilder.Column {
	return ql.Col(m.Topic).NotLike(val)
}

func (m CmsSubscriptionM) TopicDesc() sqlbuilder.OrderBy {
	return ql.Desc(m.Topic)
}

func (m CmsSubscriptionM) TopicAsc() sqlbuilder.OrderBy {
	return ql.Asc(m.Topic)
}
