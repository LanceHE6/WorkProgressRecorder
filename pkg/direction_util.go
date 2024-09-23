package pkg

import (
	"WorkProgressRecord/internal/model"
)

// IsInSlice
//
//	@Description: 判断val是否在slice中
//	@param slice []any
//	@param val any
//	@return bool true:在  false:不在
func IsInSlice(slice []interface{}, val interface{}) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// IsContainGoal
//
//	@Description: 判断用户方向中是否包括方向
//	@param slice []any
//	@param val any
//	@return bool true:在  false:不在
func IsContainGoal(direction model.Direction, directionType model.DirectionType) bool {
	for _, item := range direction {
		if item == directionType {
			return true
		}
	}
	return false
}

// RemoveGoal
//
//	@Description: 从用户方向中移除方向
//	@param direction []model.DirectionType 用户方向
//	@param directionType model.DirectionType 要移除的方向
//	@return model.Direction 移除后的用户方向
func RemoveGoal(direction model.Direction, directionType model.DirectionType) model.Direction {
	for i, item := range direction {
		if item == directionType {
			direction = append(direction[:i], direction[i+1:]...)
			break
		}
	}
	return direction
}
