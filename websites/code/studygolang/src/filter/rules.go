// Copyright 2013 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	studygolang@gmail.com

package filter

import (
	"regexp"
)

func Rule(uri string) map[string]map[string]map[string]string {
	if rule, ok := rules[uri]; ok {
		return rule
	}
	for key, rule := range rules {
		reg := regexp.MustCompile(key)
		if reg.MatchString(uri) {
			return rule
		}
	}
	return nil
}

// 定义所有表单验证规则
var rules = map[string]map[string]map[string]map[string]string{
	// 用户注册验证规则
	"/account/register.json": {
		"username": {
			"require": {"error": "用户名不能为空！"},
			"regex":   {"pattern": `^\w*$`, "error": "用户名只能包含大小写字母、数字和下划线"},
			"length":  {"range": "4,20", "error": "用户名长度必须在%d个字符和%d个字符之间"},
		},
		"email": {
			"require": {"error": "邮箱不能为空！"},
			"email":   {"error": "邮箱格式不正确！"},
		},
		"passwd": {
			"require": {"error": "密码不能为空！"},
			"length":  {"range": "6,32", "error": "密码长度必须在%d个字符和%d个字符之间"},
		},
		"pass2": {
			"require": {"error": "确认密码不能为空！"},
			"compare": {"field": "passwd", "rule": "=", "error": "两次密码不一致"},
		},
	},
	// 发新帖
	"/topics/new.json": {
		"nid": {
			"int": {"range": "0,", "error": "请选择节点"},
		},
		"title": {
			"require": {"error": "标题不能为空"},
			"length":  {"range": "3,", "error": "话题标题长度必不能少于%d个字符"},
		},
		"content": {
			"require": {"error": "内容不能为空！"},
			"length":  {"range": "2,", "error": "话题内容长度必不能少于%d个字符"},
		},
	},
	// 发回复
	`/comment/\d+\.json`: {
		"content": {
			"require": {"error": "内容不能为空！"},
			"length":  {"range": "2,", "error": "回复内容长度必不能少于%d个字符"},
		},
	},
	// 发wiki
	"/wiki/new.json": {
		"title": {
			"require": {"error": "标题不能为空"},
			"length":  {"range": "3,", "error": "标题长度必不能少于%d个字符"},
		},
		"uri": {
			"require": {"error": "URL不能为空"},
		},
		"content": {
			"require": {"error": "内容不能为空！"},
			"length":  {"range": "2,", "error": "内容长度必不能少于%d个字符"},
		},
	},
}
