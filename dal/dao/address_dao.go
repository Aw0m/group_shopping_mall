package dao

import (
	"context"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_redis"
	"group_shopping_mall/model/bdm"
)

func SaveShoppingAddress(ctx context.Context, addressList []bdm.ShoppingAddress) error {
	val, err := jsoniter.MarshalToString(addressList)
	if err != nil {
		return errors.Errorf("marshal addressList failed! err %s", err.Error())
	}

	err = client_redis.SetValue(ctx, GetAddressKey(), val, 0)
	if err != nil {
		return errors.Errorf("set addressList to redis failed! err %s", err.Error())
	}
	return nil
}

func GetAddressKey() string {
	key := fmt.Sprintf("[address_list]")
	return key
}
