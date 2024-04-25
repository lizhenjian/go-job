package util

import (
	"errors"
	"fmt"
	"go-jobs/configs/constants"
	"reflect"
	"sync/atomic"
	"unsafe"
)

// 调用结构体下方法
func CallMethod(instance interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	// 获取方法的反射值
	methodValue := reflect.ValueOf(instance).MethodByName(methodName)

	// 检查方法是否存在
	if !methodValue.IsValid() {
		return nil, fmt.Errorf("method not found: %s", methodName)
	}

	// 如果方法是通过指针接收者定义的，则获取指针类型的反射值
	if methodValue.Type().IsVariadic() {
		ptrValue := reflect.ValueOf(instance)
		methodValue = ptrValue.MethodByName(methodName)
	}

	// 构造方法的参数列表
	methodArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		methodArgs[i] = reflect.ValueOf(arg)
	}

	// 调用方法并获取返回值
	resultValues := methodValue.Call(methodArgs)

	// 如果方法没有返回值，则返回 nil
	if len(resultValues) == 0 {
		return nil, nil
	}

	// 将返回值转为 interface{} 类型并返回
	results := make([]interface{}, len(resultValues))
	for i, v := range resultValues {
		results[i] = v.Interface()
	}
	return results, nil
}

// 注册全局变量
// var importedFunctions = make(map[string]map[string]interface{})
// 注册
func RegisterFunction(packageName, funcName string, function interface{}) {
	var importedFunctionsTemp = make(map[string]map[string]interface{})
	if importedFunctionsTemp[packageName] == nil {
		importedFunctionsTemp[packageName] = make(map[string]interface{})
	}
	importedFunctionsTemp[packageName][funcName] = function
	atomic.StorePointer(&constants.ImportedFunctions, unsafe.Pointer(&importedFunctionsTemp))
}

// 调用包下面的方法
func Eval(packageName string, funcName string, args ...interface{}) ([]interface{}, error) {
	importedFunctions2 := *(*map[string]map[string]interface{})(atomic.LoadPointer(&constants.ImportedFunctions))
	// Get the function's reflect.Value using the package path and function name
	funcValue := reflect.ValueOf(importedFunctions2[packageName][funcName])
	// Check if the function exists
	if !funcValue.IsValid() {
		return nil, errors.New("Failed to find function: " + funcName)
	}

	// Construct the function's argument list
	funcArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		funcArgs[i] = reflect.ValueOf(arg)
	}

	// Call the function and get the return values
	resultValues := funcValue.Call(funcArgs)

	// Convert the return values to interface{} and save them in the result slice
	results := make([]interface{}, len(resultValues))
	for i, result := range resultValues {
		results[i] = result.Interface()
	}

	return results, nil
}
