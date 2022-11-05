package tools

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

//判断文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}


//判断文件夹是否存在不存存在就创建
func PathExistsCreate(Path string)  {
	_, err := os.Stat(Path)
	res := os.IsNotExist(err)
	if res == true {
		os.MkdirAll(Path, os.ModePerm)
	}
}


//生产随机字符串 不能保证唯一
func CreateRandomString(len int) string  {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0;i < len ;i++  {
		randomInt,_ := rand.Int(rand.Reader,bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func Uuid() string{
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid

}


//创建并且写入文件
func WriteFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	} else {
		_, err = f.Write([]byte(content))
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteByteFile(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return err
		}
	}

	return nil
}


func CopyFile (srcFilePath string,dstFilePath string)(written int64, err error){
	srcFile,err := os.Open(srcFilePath)
	if err != nil{
		fmt.Printf("打开源文件错误，错误信息=%v\n",err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	if PathExists(dstFilePath){
		os.Remove(dstFilePath)
	}
	dstFile,err := os.OpenFile(dstFilePath,os.O_WRONLY | os.O_CREATE,0777)
	if err != nil{
		fmt.Printf("打开目标文件错误，错误信息=%v\n",err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer,reader)
}

//获取程序根目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs("")
	if err != nil {
		return "/"
	}
	return strings.Replace(dir, "\\", "/", -1)
}


//移动文件
func MoveFile(file string, toPath string,fileName string) error {
	_, err := os.Stat(toPath)
	res := os.IsNotExist(err)
	if res == true {
		os.MkdirAll(toPath, os.ModePerm)
	}
	if len(fileName) == 0{
		fileName = path.Base(file)
	}

	var cmd *exec.Cmd
	cmd = exec.Command("mv", file, toPath+"/"+fileName)
	_, err = cmd.Output()
	return err

}


//删除目录
func RemoveAll(toPath string)  {
	_, err := os.Stat(toPath)
	res := os.IsNotExist(err)
	if res != true {
		os.RemoveAll(toPath)
	}
}

//binding type interface 要修改的结构体
//value type interace 有数据的结构体
//fieldValue := reflect.ValueOf(student).FieldByName(fieldName)
func StructAssign(value interface{},binding interface{}, ) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			if bVal.FieldByName(name).Type() == vVal.FieldByName(name).Type() {
				bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
			}
		}
	}
}

//动态赋值
func StructuralSetValue(structural interface{},key string,value interface{})  {
	bVal := reflect.ValueOf(structural).Elem()
	if ok := bVal.FieldByName(key).IsValid(); ok {
		bVal.FieldByName(key).Set(reflect.ValueOf(value))
	}
}

//动态调用函数
func ByNameCall(structural interface{},functionName string, variable []reflect.Value)  {
	_ = reflect.ValueOf(structural).MethodByName(functionName).Call(variable)
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}



func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}


func GetFileList(path string) []string {
	list :=make([]string,0)
	fs,_:= ioutil.ReadDir(path)
	for _,file:=range fs{
		if file.IsDir(){
			lists := GetFileList(path+"/"+file.Name())
			list = append(list,lists...)
		}else{
			list = append(list,path+"/"+file.Name() )
		}
	}
	return  list
}

func ListDataV(list []interface{}, ele interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		if v != ele {
			result = append(result, v)
		}
	}
	return result
}

func ListDataI(list []interface{}, index int) []interface{} {
	result := make([]interface{}, 0)
	for i, v := range list {
		if i != index {
			result = append(result, v)
		}
	}
	return result
}