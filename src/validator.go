package src

import (
    "github.com/go-playground/validator/v10"
)

func MultiItemSize(fl validator.FieldLevel) bool {
    list,_,_,listOk:=fl.GetStructFieldOKAdvanced2(fl.Parent(), "MenuItemList")
    itemSize,_,_,sizeOk:=fl.GetStructFieldOKAdvanced2(fl.Parent(), "MenuItemSize")
    if (listOk && sizeOk) {
        slice,ok:= list.Interface().([]MenuItemRequest)
        if (ok) {
            return int64(len(slice)) == itemSize.Int()
        }
    }

    return false
}
