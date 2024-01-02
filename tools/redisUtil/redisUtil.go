package redisUtil

import (
	"context"
	"errors"
	"time"
	"wscmakebygo.com/global/constant"
	"wscmakebygo.com/global/redisConn"
)

var ctx = context.Background()

func GetData(key string) (string, error) {
	val, err := redisConn.GetRedis().Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, err
}

func SetData(key string, value string) error {
	err := redisConn.
		GetRedis().
		SetEX(ctx, key, value, time.Second*constant.ATTENDEE_LOGIN_TIMEOUT).
		Err()
	if err != nil {
		return err
	}
	return nil
}

func Exist(key string) (bool, error) {
	isExist, err := redisConn.
		GetRedis().
		Exists(ctx, key).
		Result()
	if err != nil {
		return false, err
	}
	return isExist > 0, nil
}

func RemoveKey(key string) error {
	isExist, err := Exist(key)
	if err != nil {
		return err
	} else if !isExist {
		return errors.New(key + " is no Exist")
	}
	err = redisConn.
		GetRedis().
		Del(ctx, key).
		Err()
	return err
}
