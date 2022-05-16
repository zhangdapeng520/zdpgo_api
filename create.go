package zdpgo_api

import (
	"os"
)

/*
@Time : 2022/5/16 17:07
@Author : 张大鹏
@File : create.go
@Software: Goland2021.3.1
@Description: create 创建相关
*/

//CreateDirs 调用os.MkdirAll递归创建文件夹
func (a *Api) CreateDirs(filePath string) error {
	if !a.IsExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			a.Log.Error("创建文件夹失败", "error", err)
			return err
		}
		return err
	}
	return nil
}

// IsExists 判断所给路径文件/文件夹是否存在(返回true是存在)
func (a *Api) IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		a.Log.Error("获取文件信息失败", "error", err, "path", path)
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
